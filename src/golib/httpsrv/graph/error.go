// Package gql contains GraphQL related code
package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type gqlError struct {
	*gqlerror.Error
	cause error // ok to be nil
}

var (
	_ error = &gqlError{}
)

// ConvertError converts the given error into graphql error
func ConvertError(ctx context.Context, code, message string, err error) error {
	return &gqlError{
		Error: &gqlerror.Error{
			Path:    graphql.GetPath(ctx),
			Message: message,
			Extensions: map[string]interface{}{
				"code": code,
			},
		},
		cause: err,
	}
}

func errorPresenter(isIntrospectionEnabled bool) graphql.ErrorPresenterFunc {
	return func(ctx context.Context, err error) *gqlerror.Error {
		if err == nil {
			return nil
		}

		gqlErr := graphql.DefaultErrorPresenter(ctx, err)

		// Don't expose any schema-identifiable info when introspection is disabled
		if !isIntrospectionEnabled {
			gqlErr.Locations = nil
			gqlErr.Path = nil
		}

		// TODO: Add logging & alerting

		return gqlErr
	}
}
