package main

import (
	"log/slog"
	"os"
	"time"
	"user-service/cmd"
)

func main() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		slog.Error("Error loading timezone", "error", err)
		os.Exit(1)
	}
	time.Local = loc

	if err := cmd.Execute(); err != nil {
		slog.Error("Execute failed", "error", err)
		os.Exit(1)
	}
}
