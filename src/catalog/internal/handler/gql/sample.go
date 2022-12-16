// Package gql contains GQL resolvers
package gql

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/kneadCODE/bridge/src/golib/httpsrv/graph"
)

// NoOp is a sample resolver
func (resolver) NoOp(ctx context.Context) (*Result, error) {
	strPtr := "abc"
	return &Result{Person: &Person{Name: &strPtr}}, nil
}

// KnownError is a sample resolver
func (resolver) KnownError(ctx context.Context) (*bool, error) {
	graphql.AddError(
		ctx,
		graph.ConvertBadRequestError(ctx, "invalid_qty", "Invalid qty given"),
	)
	return nil, graph.ConvertBadRequestError(ctx, "invalid_param", "Invalid param given")
}

// UnknownError is a sample resolver
func (resolver) UnknownError(ctx context.Context) (*bool, error) {
	graphql.AddError(
		ctx,
		graph.ConvertUnexpectError(ctx, errors.New("unknown error")),
	)
	return nil, graph.ConvertUnexpectError(ctx, errors.New("some error"))
}
