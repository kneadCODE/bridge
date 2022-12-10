// Package app contains the App configuration
package app

import (
	"errors"
)

var (
	// ErrInvalidConfig means the app config is Invalid
	ErrInvalidConfig = errors.New("invalid app config")
)
