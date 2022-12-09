package main

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/kneadCODE/bridge/src/golib/app"
	"golang.org/x/exp/slog"
)

func main() {
	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout))
	ctx = slog.NewContext(ctx, logger)

	logger.Info("API Initialized")

	app.Run(ctx, func(ctx context.Context) error {
		time.Sleep(5 * time.Second)
		return errors.New("some err")
	})
}
