//go:build tools
// +build tools

// Package tools contains tools related code
package tools

import (
	// Adding this to support go generate
	_ "github.com/99designs/gqlgen"
	// Adding this to support go generate
	_ "github.com/99designs/gqlgen/graphql"
	// Adding this to support go generate
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
