package main

import (
	"log/slog"
	"modcore/cli/cmd"
	"modcore/cli/ipc"
	"os"
)

var logger *slog.Logger

func main() {
	logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
	ipc.Log = logger
	cmd.Log = logger
	if err := cmd.Execute(); err != nil {
		logger.Error("%v", err)
	}
}
