// Package httpsrv contains HTTP Server
package httpsrv

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"golang.org/x/exp/slog"
)

// New returns a new instance of Server.
func New(ctx context.Context, router Router, options ...ServerOption) (*Server, error) {
	s := &Server{
		srv: &http.Server{
			Addr:         ":9000",
			Handler:      router.Handler(),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
			BaseContext: func(net.Listener) context.Context {
				return slog.NewContext(context.Background(), slog.FromContext(ctx))
			},
		},
		gracefulShutdownTimeout: 10 * time.Second,
	}

	for _, opt := range options {
		if err := opt(s); err != nil {
			return nil, err
		}
	}

	return s, nil
}

// Server is the server instance
type Server struct {
	srv                     *http.Server
	gracefulShutdownTimeout time.Duration
}

// Start starts the server and is context aware and shuts down when the context gets cancelled.
func (s *Server) Start(ctx context.Context) error {
	logger := slog.FromContext(ctx)

	startErrChan := make(chan error, 1)

	go func() {
		logger.LogAttrs(slog.LevelInfo, fmt.Sprintf("Starting HTTP server on [%s]", s.srv.Addr))
		startErrChan <- s.srv.ListenAndServe()
	}()

	for {
		select {
		case <-ctx.Done():
			return s.stop(logger)
		case err := <-startErrChan:
			if err != http.ErrServerClosed {
				return fmt.Errorf("http server startup failed: %w", err)
			}
			return nil
		}
	}
}

func (s *Server) stop(logger *slog.Logger) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.gracefulShutdownTimeout) // Cannot rely on root context as that might have been cancelled.
	defer cancel()

	logger.LogAttrs(slog.LevelInfo, "Attempting HTTP server graceful shutdown")
	if err := s.srv.Shutdown(ctx); err != nil {
		logger.Error("HTTP server graceful shutdown failed", err)

		logger.LogAttrs(slog.LevelInfo, "Attempting HTTP server force shutdown")
		if err = s.srv.Close(); err != nil {
			logger.Error("HTTP server graceful shutdown failed", err)
			return err
		}
	}

	logger.LogAttrs(slog.LevelInfo, "HTTP server shutdown complete")

	return nil
}

// ServerOption customizes the Server
type ServerOption = func(*Server) error

// ServerPort sets the server port to the given port
func ServerPort(port int) ServerOption {
	return func(s *Server) error {
		s.srv.Addr = fmt.Sprintf(":%d", port)
		return nil
	}
}
