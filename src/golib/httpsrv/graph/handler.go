// Package graph contains GraphQL related code
package graph

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"golang.org/x/exp/slog"
)

// Handler returns the gqlgen Handler
func Handler(schema graphql.ExecutableSchema, isIntrospectionEnabled bool) http.Handler {
	srv := handler.New(schema)
	srv.AddTransport(transport.POST{})
	srv.SetErrorPresenter(errorPresenter(isIntrospectionEnabled))
	if isIntrospectionEnabled {
		srv.Use(extension.Introspection{})
	}
	srv.SetRecoverFunc(recoverFunc)
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		logger := slog.FromContext(ctx)

		opCtx := graphql.GetOperationContext(ctx)
		opName := opCtx.OperationName
		if opName == "" {
			opName = "NO_NAME"
		}

		logger = logger.With(
			slog.String("gql.req.type", string(opCtx.Operation.Operation)),
			slog.String("gql.req.operation", opName),
		)
		ctx = slog.NewContext(ctx, logger)

		logger.LogAttrs(
			slog.LevelInfo,
			fmt.Sprintf("START %s/%s", opCtx.Operation.Operation, opName),
			slog.Time("gql.req.start", opCtx.Stats.OperationStart),
		)

		logger.Info(fmt.Sprintf("Raw Query: [%s]. Variables: [%s]", opCtx.RawQuery, opCtx.Variables)) // TODO: Add redaction to variables

		return next(ctx)
	})
	srv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		res := next(ctx)

		logger := slog.FromContext(ctx)

		opCtx := graphql.GetOperationContext(ctx)
		opName := opCtx.OperationName
		if opName == "" {
			opName = "NO_NAME"
		}

		logger.Info(fmt.Sprintf("Data: [%s]. Errors: [%s]", res.Data, res.Errors.Error())) // TODO: Add redaction

		logger.LogAttrs(
			slog.LevelInfo,
			fmt.Sprintf("END %s/%s", opCtx.Operation.Operation, opName),
			slog.String(
				"gql.resp.read_duration",
				fmt.Sprintf(
					"%dms",
					opCtx.Stats.Read.End.Sub(opCtx.Stats.Read.End).Milliseconds(),
				),
			),
			slog.String(
				"gql.resp.parsing_duration",
				fmt.Sprintf(
					"%dms",
					opCtx.Stats.Parsing.End.Sub(opCtx.Stats.Parsing.End).Milliseconds(),
				),
			),
			slog.String(
				"gql.resp.validation_duration",
				fmt.Sprintf(
					"%dms",
					opCtx.Stats.Validation.End.Sub(opCtx.Stats.Validation.End).Milliseconds(),
				),
			),
		)

		return res
	})
	return srv
}
