package graph

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/exp/slog"
)

var (
	_ graphql.RecoverFunc = recoverFunc
)

func recoverFunc(ctx context.Context, rcv interface{}) error {
	if rcv == nil {
		return nil
	}

	if err, ok := rcv.(error); ok {
		slog.FromContext(ctx).Error(fmt.Sprintf("PANIC RECOVERED. Stack: [%s]", string(debug.Stack())), err)
	} else {
		slog.FromContext(ctx).LogAttrs(
			slog.ErrorLevel,
			fmt.Sprintf("PANIC RECOVERED: [%+v]. Stack: [%s]", rcv, string(debug.Stack())),
		)
	}

	return &gqlerror.Error{
		Message: "An unknown error occurred",
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]interface{}{
			"code": errCodeInternal.String(),
		},
	}
}
