// Package graph contains GraphQL related code
package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ConvertKnownError converts the known error into *gqlerror.Error
func ConvertKnownError(ctx context.Context, code, message string) *gqlerror.Error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: message,
		Extensions: map[string]interface{}{
			"code": code,
		},
	}
}

// ConvertUnexpectError converts the given unexpected error into *gqlerror.Error
func ConvertUnexpectError(ctx context.Context, err error) *gqlerror.Error {
	gerr := gqlerror.WrapPath(graphql.GetPath(ctx), err)
	gerr.Message = "An unknown error occurred"
	gerr.Extensions = map[string]interface{}{
		"code": "internal_error",
	}
	return gerr
}

func errorPresenter(isIntrospectionEnabled bool) graphql.ErrorPresenterFunc {
	return func(ctx context.Context, err error) *gqlerror.Error {
		if err == nil {
			return nil
		}

		var gerr *gqlerror.Error
		if !errors.As(err, &gerr) {
			gerr = ConvertUnexpectError(ctx, err)
		}

		// Don't expose any schema-identifiable info when introspection is disabled
		if !isIntrospectionEnabled {
			gerr.Locations = nil
			gerr.Path = nil
		}

		if cause := gerr.Unwrap(); cause != nil {
			// TODO: Add logging & alerting
		}

		return gerr
	}
}
