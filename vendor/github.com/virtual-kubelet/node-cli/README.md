Virtual-Kubelet CLI
==================

This project provides a library for rapid prototyping of a virtual-kubelet node.
It is not intended for production use and may have breaking changes,
but takes as much as made sense from the old command line code from
[github.com/virtual-kubelet/virtual-kubelet][vk].

[vk]: https://github.com/virtual-kubelet/virtual-kubelet


## Usage

```go
package main

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	cli "github.com/virtual-kubelet/node-cli"
	logruscli "github.com/virtual-kubelet/node-cli/logrus"
	"github.com/virtual-kubelet/node-cli/provider"
	"github.com/virtual-kubelet/virtual-kubelet/log"
	logruslogger "github.com/virtual-kubelet/virtual-kubelet/log/logrus"
)

func main() {
	ctx := cli.ContextWithCancelOnSignal(context.Background())
	logger := logrus.StandardLogger()

	log.L = logruslogger.FromLogrus(logrus.NewEntry(logger))
	logConfig := &logruscli.Config{LogLevel: "info"}

	node, err := cli.New(ctx,
		cli.WithProvider("demo", func(cfg provider.InitConfig) (provider.Provider, error) {
			return nil, errors.New("your implementation goes here")
		}),
		// Adds flags and parsing for using logrus as the configured logger
		cli.WithPersistentFlags(logConfig.FlagSet()),
		cli.WithPersistentPreRunCallback(func() error {
			return logruscli.Configure(logConfig, logger)
		}),
	)

	if err != nil {
		panic(err)
	}

	// Notice that the context is not passed in here, this is due to limitations
	// of the underlying CLI library (cobra).
	// contexts get passed through from `cli.New()`
	//
	// Args can be specified here, or os.Args[1:] will be used.
	if err := node.Run(); err != nil {
		panic(err)
	}
}
```
