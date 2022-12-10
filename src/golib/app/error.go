package app

import (
	"errors"
)

var (
	// ErrInvalidConfig means the app config is Invalid
	ErrInvalidConfig = errors.New("invalid app config")
)
