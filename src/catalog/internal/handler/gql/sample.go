// Package gql contains GQL resolvers
package gql

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/kneadCODE/bridge/src/golib/httpsrv/graph"
)

// NoOp is a sample resolver
func (resolver) NoOp(context.Context) (*bool, error) {
	return nil, nil
}

// KnownError is a sample resolver
func (resolver) KnownError(ctx context.Context) (*bool, error) {
	graphql.AddError(
		ctx,
		graph.ConvertKnownError(ctx, graph.ErrCodeBadRequest, "invalid_qty", "Invalid qty given"),
	)
	return nil, graph.ConvertKnownError(ctx, graph.ErrCodeBadRequest, "invalid_param", "Invalid param given")
}

// UnknownError is a sample resolver
func (resolver) UnknownError(ctx context.Context) (*bool, error) {
	graphql.AddError(
		ctx,
		graph.ConvertUnexpectError(ctx, errors.New("unknown error")),
	)
	return nil, graph.ConvertUnexpectError(ctx, errors.New("some error"))
}
