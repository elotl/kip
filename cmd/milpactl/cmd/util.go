package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
)

func fatal(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}
	fmt.Fprint(os.Stderr, msg)
	os.Exit(1)
}

func dieIfReplyError(cmd string, reply *clientapi.APIReply) {
	if reply.Status < 200 || reply.Status >= 400 {
		fatal("%s returned %d - %s", cmd, reply.Status, reply.Body)
	}
}

func dieIfError(err error, format string, args ...interface{}) {
	if err != nil {
		s := fmt.Sprintf(format, args...)
		msg := s + ": " + err.Error()
		fatal(msg)
	}
}
