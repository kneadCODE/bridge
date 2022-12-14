// Binary API
package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/kneadCODE/bridge/src/golib/app"
	"github.com/kneadCODE/bridge/src/golib/httpsrv"
	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout))
	logger.LogAttrs(slog.LevelInfo, "Initializing app")

	appCfg := app.Config{
		Name:        os.Getenv("APP_NAME"),
		Environment: app.Env(os.Getenv("APP_ENV")),
		Version:     os.Getenv("APP_VERSION"),
		Server:      os.Getenv("APP_SERVER"),
	}
	if err := appCfg.IsValid(); err != nil {
		logger.Error("Init app config failed", err)
		os.Exit(1)
	}

	logger = appCfg.EnrichLogger(logger)
	slog.SetDefault(logger)
	ctx := slog.NewContext(context.Background(), logger)

	srv, err := initServer(ctx)
	if err != nil {
		logger.Error("Init server failed", err)
		os.Exit(1)
	}

	logger.LogAttrs(slog.LevelInfo, "App initialized")

	appCfg.Run(
		ctx,
		srv.Start,
	)
}

func initServer(ctx context.Context) (*httpsrv.Server, error) {
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid port: [%s], %w", portStr, err)
	}
	srv, err := httpsrv.New(ctx, router(), httpsrv.ServerPort(port))
	if err != nil {
		return nil, err
	}

	return srv, nil
}
