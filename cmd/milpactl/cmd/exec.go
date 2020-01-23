package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/wsstream"
	"github.com/kr/pty"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/net/context"
)

var (
	execPodName      string
	execUnitName     string
	execInteractive  bool
	execTTY          bool
	execUsageStr     = "exec POD_NAME -- command"
	oldTerminalState *terminal.State
)

const (
	wsTTYControlChan = 4
)

func restoreOldTerminalState() {
	if oldTerminalState != nil {
		_ = terminal.Restore(int(os.Stdin.Fd()), oldTerminalState)
	}
}

func restoreTerminalAndExit(returnCode int, format string, v ...interface{}) {
	restoreOldTerminalState()
	if format != "" {
		msg := fmt.Sprintf(format, v...)
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		fmt.Fprint(os.Stderr, msg)
	}
	os.Exit(returnCode)
}

func exec(cmd *cobra.Command, args []string) {
	argsLenAtDash := cmd.ArgsLenAtDash()
	if len(args) == 0 || argsLenAtDash == 0 {
		fatal("A pod name is required: " + execUsageStr)
	}
	execPodName = args[0]

	command := args[1:]
	if len(command) == 0 {
		fatal("A command is required: " + execUsageStr)
	}

	params := api.ExecParams{
		PodName:     execPodName,
		UnitName:    execUnitName,
		Command:     command,
		Interactive: execInteractive,
		TTY:         execTTY,
	}

	client, conn, err := getMilpaClient(cmd.InheritedFlags(), false)
	dieIfError(err, "Failed to create milpa client")
	defer conn.Close()

	stream, err := client.Exec(context.Background())
	dieIfError(err, "Failed to setup exec streaming client")

	b, err := json.Marshal(params)
	dieIfError(err, "Error serializing exec parameters")
	paramMsg := &clientapi.StreamMsg{Data: b}
	err = stream.Send(paramMsg)
	dieIfError(err, "Error sending initial exec parameters")

	// Send tty resize messages to the other side
	if execTTY {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGWINCH)
		go func() {
			for range ch {
				s, err := pty.GetsizeFull(os.Stdin)
				if err != nil {
					log.Printf("error resizing pty: %s", err)
					continue
				}
				b, err := json.Marshal(s)
				if err != nil {
					log.Printf("error resizing pty: %s", err)
					continue
				}
				f := wsstream.PackMessage(wsTTYControlChan, b)
				controlMsg := &clientapi.StreamMsg{Data: f}
				stream.Send(controlMsg)
			}
		}()
		ch <- syscall.SIGWINCH // Initial resize.
	}

	if execInteractive {
		go func() {
			defer stream.CloseSend()
			if execTTY {
				// put local terminal in raw mode so that we are basically a pass-through for the remote terminal
				oldTerminalState, _ = terminal.MakeRaw(int(os.Stdin.Fd()))
				defer restoreOldTerminalState()
			}
			var b []byte = make([]byte, 1)
			for {
				n, err := os.Stdin.Read(b)
				if err != nil {
					restoreTerminalAndExit(1, "Error reading stdin: %s", err)
				}
				if n == 0 {
					continue
				}
				f := wsstream.PackMessage(wsstream.StdinChan, []byte(b))
				sm := &clientapi.StreamMsg{Data: f}
				if err := stream.Send(sm); err != nil {
					return
				}
			}
		}()
	}
	// Write to local stdout and stderr, if we get an exit code,
	// exit with that code
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			restoreTerminalAndExit(1, "Error in grpc receive %s", err)
		}
		c, msg, err := wsstream.UnpackMessage(resp.Data)
		if err != nil {
			restoreTerminalAndExit(1, "error unserializing websocket data: %s", err)
		}
		if len(msg) > 0 {
			if c == wsstream.StdoutChan {
				fmt.Fprint(os.Stdout, string(msg))
			} else if c == wsstream.StderrChan {
				fmt.Fprint(os.Stderr, string(msg))
			} else if c == wsstream.ExitCodeChan {
				i, err := strconv.Atoi(string(msg))
				if err != nil {
					restoreTerminalAndExit(1, "Invalid exit code: %s", msg)
				}
				restoreTerminalAndExit(i, "")
			}
		}
	}
}

func ExecCommand() *cobra.Command {
	var execCmd = &cobra.Command{
		Use:   "exec",
		Short: "Execute a command in the context unit",
		Long:  `Execute a command in the context unit`,
		Example: `# Get output from running 'date' from pod my-pod, using the first unit by default
milpactl exec my-pod date

# Get output from running 'date' in the rubyserver unit from pod my-pod
milpactl exec my-pod -u rubyserver date

# Switch to raw terminal mode, sends stdin to 'bash' in rubyserver from pod my-pod
# and sends stdout/stderr from 'bash' back to the client
milpactl exec my-pod -u rubyserver -i -t -- bash -il`,
		Run: func(cmd *cobra.Command, args []string) {
			exec(cmd, args)
		},
	}

	execCmd.Flags().StringVarP(&execUnitName, "unit", "u", "", "Unit name. If empty the first unit in the pod will be used")
	execCmd.Flags().BoolVarP(&execInteractive, "stdin", "i", false, "Pass stdin to the unit")
	execCmd.Flags().BoolVarP(&execTTY, "tty", "t", false, "Stdin is a TTY")

	return execCmd
}
