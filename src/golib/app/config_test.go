package app

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// TODO: Look into adding fuzzy tests.

func TestConfig_IsValid(t *testing.T) {
	type testCase struct {
		givenCfg Config
		expErr   error
	}
	tcs := map[string]testCase{
		"valid, dev": {
			givenCfg: Config{
				Name:        "name",
				Environment: EnvDev,
				Version:     "v123",
				Server:      "server",
			},
		},
		"valid, staging": {
			givenCfg: Config{
				Name:        "name",
				Environment: EnvStaging,
				Version:     "v123",
				Server:      "server",
			},
		},
		"valid, prod": {
			givenCfg: Config{
				Name:        "name",
				Environment: EnvProd,
				Version:     "v123",
				Server:      "server",
			},
		},
		"no name": {
			givenCfg: Config{
				Environment: EnvProd,
				Version:     "v123",
				Server:      "server",
			},
			expErr: fmt.Errorf("invalid name: [], %w", ErrInvalidConfig),
		},
		"name too long": {
			givenCfg: Config{
				Name:        "abcdefghijkl",
				Environment: EnvProd,
				Version:     "v123",
				Server:      "server",
			},
			expErr: fmt.Errorf("invalid name: [abcdefghijkl], %w", ErrInvalidConfig),
		},
		"no version": {
			givenCfg: Config{
				Name:        "name",
				Environment: EnvProd,
				Server:      "server",
			},
			expErr: fmt.Errorf("invalid version: [], %w", ErrInvalidConfig),
		},
		"version too long": {
			givenCfg: Config{
				Name:        "name",
				Environment: EnvProd,
				Version:     "vXXXX.XXX.XXX",
				Server:      "server",
			},
			expErr: fmt.Errorf("invalid version: [vXXXX.XXX.XXX], %w", ErrInvalidConfig),
		},
		"no server": {
			givenCfg: Config{
				Name:        "name",
				Environment: EnvProd,
				Version:     "v123",
			},
			expErr: fmt.Errorf("invalid server: [], %w", ErrInvalidConfig),
		},
		"server too long": {
			givenCfg: Config{
				Name:        "name",
				Environment: EnvProd,
				Version:     "v123",
				Server:      "serverserverserverserverserverserverserverserverserverserverserverserverserverserver",
			},
			expErr: fmt.Errorf("invalid server: [serverserverserverserverserverserverserverserverserverserverserverserverserverserver], %w", ErrInvalidConfig),
		},
		"no env": {
			givenCfg: Config{
				Name:    "name",
				Version: "v123",
				Server:  "server",
			},
			expErr: fmt.Errorf("invalid env: [], %w", ErrInvalidConfig),
		},
		"invalid env": {
			givenCfg: Config{
				Name:        "name",
				Environment: "env",
				Version:     "v123",
				Server:      "server",
			},
			expErr: fmt.Errorf("invalid env: [env], %w", ErrInvalidConfig),
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given, When & Then:
			require.Equal(t, tc.expErr, tc.givenCfg.IsValid())
		})
	}
}
