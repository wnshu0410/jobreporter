package main

import (
	"fmt"
	"os"

	"github.com/wnshu0410/jobreporter/internal/log"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	logger, err := log.NewAtLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return err
	}

	defer func() {
		err = logger.Sync()
	}()

	logger.Info("Hello world!", zap.String("location", "world"))

	return err
}
