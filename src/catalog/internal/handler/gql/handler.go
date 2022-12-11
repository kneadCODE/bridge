//go:generate go run github.com/99designs/gqlgen generate

// Package gql contains GQL resolvers
package gql

import (
	"net/http"

	"github.com/kneadCODE/bridge/src/golib/httpsrv/graph"
)

func NewHandler(isIntrospectionEnabled bool) http.Handler {
	cfg := Config{
		Resolvers: resolver{},
		// TODO: Add directives & complexities
	}
	return graph.Handler(NewExecutableSchema(cfg), isIntrospectionEnabled)
}

type resolver struct {
}

func (r resolver) Query() QueryResolver {
	return r
}
