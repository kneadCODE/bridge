//go:build tools
// +build tools

// Package tools contains tools related code
package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
