// Package graph contains GraphQL related code
package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

// Handler returns the gqlgen Handler
func Handler(schema graphql.ExecutableSchema, isIntrospectionEnabled bool) http.Handler {
	srv := handler.New(schema)
	srv.AddTransport(transport.POST{})
	srv.SetErrorPresenter(errorPresenter(isIntrospectionEnabled))
	if isIntrospectionEnabled {
		srv.Use(extension.Introspection{})
	}
	return srv
}
