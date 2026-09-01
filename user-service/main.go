package main

import (
	"log/slog"
	"os"
	"user-service/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		slog.Error("Execute failed", "error", err)
		os.Exit(1)
	}
}
