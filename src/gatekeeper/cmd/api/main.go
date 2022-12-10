// Binary API
package main

import (
	"context"
	"os"

	"github.com/kneadCODE/bridge/src/golib/app"
	"github.com/kneadCODE/bridge/src/golib/httpsrv"
	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout))
	logger.LogAttrs(slog.InfoLevel, "Initializing app")

	slog.SetDefault(logger)

	ctx := slog.NewContext(context.Background(), logger)

	srv, err := initServer(ctx)
	if err != nil {
		logger.Error("Init server failed", err)
		os.Exit(1)
	}

	logger.LogAttrs(slog.InfoLevel, "App initialized")

	app.Run(
		ctx,
		srv.Start,
	)
}

func initServer(context.Context) (*httpsrv.Server, error) {
	srv, err := httpsrv.New(router())
	if err != nil {
		return nil, err
	}

	return srv, nil
}
