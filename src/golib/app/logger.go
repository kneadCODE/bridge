// Package app contains the App configuration
package app

import (
	"golang.org/x/exp/slog"
)

// EnrichLogger adds fields based on the app config to the logger
func (c Config) EnrichLogger(logger *slog.Logger) *slog.Logger {
	return logger.With(
		slog.String("app.name", c.Name),
		slog.String("app.env", c.Environment.String()),
		slog.String("app.version", c.Version),
		slog.String("app.server", c.Server),
	)
}
