// Package graph contains GraphQL related code
package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ErrorCode represents the code returned to the client
type ErrorCode string

// String returns the string representation of the ErrorCode
func (e ErrorCode) String() string {
	return string(e)
}

// Based on Apollo's spec - https://www.apollographql.com/docs/apollo-server/data/errors
var (
	// ErrCodeInternal means an internal error occurred
	ErrCodeInternal = ErrorCode("INTERNAL_SERVER_ERROR")
	// ErrCodeBadRequest means a bad request was sent
	ErrCodeBadRequest = ErrorCode("BAD_REQUEST")
	// ErrCodeUnauthenticated means the request was not authenticated
	ErrCodeUnauthenticated = ErrorCode("UNAUTHENTICATED")
	// ErrCodeForbidden means the request was not authorized
	ErrCodeForbidden = ErrorCode("FORBIDDEN")
)

// ConvertKnownError converts the known error into *gqlerror.Error
func ConvertKnownError(ctx context.Context, code ErrorCode, cause, message string) *gqlerror.Error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: message,
		Extensions: map[string]interface{}{
			"code":  code.String(),
			"cause": cause,
		},
	}
}

// ConvertUnexpectError converts the given unexpected error into *gqlerror.Error
func ConvertUnexpectError(ctx context.Context, err error) *gqlerror.Error {
	gerr := gqlerror.WrapPath(graphql.GetPath(ctx), err)
	gerr.Message = "An unknown error occurred"
	gerr.Extensions = map[string]interface{}{
		"code": ErrCodeInternal,
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

		if underlyingErr := gerr.Unwrap(); underlyingErr != nil {
			// TODO: Add logging & alerting
		}

		return gerr
	}
}
