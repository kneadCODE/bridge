// Package app contains the App configuration
package app

import (
	"fmt"
	"strings"
)

// Config represents the configuration for the app
type Config struct {
	// Name of the application
	Name string
	// Environment is the environment in which the app is running
	Environment Env
	// Version is the version of the app
	Version string
	// Server is the instance on which the app is running
	Server string
}

// IsValid checks if the app is valid, else returns an error if it isn't
func (c Config) IsValid() error {
	if c.Name == "" || len(c.Name) > 10 {
		return fmt.Errorf("invalid name: [%s], %w", c.Name, ErrInvalidConfig)
	}

	switch c.Environment {
	case EnvDev, EnvStaging, EnvProd:
	default:
		return fmt.Errorf("invalid env: [%s], %w", c.Environment, ErrInvalidConfig)
	}

	if !strings.HasPrefix(c.Version, "v") || len(c.Version) > 12 { // vXXX.XXX.XXX max
		return fmt.Errorf("invalid version: [%s], %w", c.Version, ErrInvalidConfig)
	}

	if c.Server == "" || len(c.Server) > 50 {
		return fmt.Errorf("invalid server: [%s], %w", c.Server, ErrInvalidConfig)
	}

	return nil
}
