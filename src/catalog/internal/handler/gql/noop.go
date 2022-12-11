package gql

import (
	"context"
)

func (resolver) NoOp(context.Context) (bool, error) {
	return false, nil
}
