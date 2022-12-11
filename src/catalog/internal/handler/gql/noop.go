// Package gql contains GQL resolvers
package gql

import (
	"context"
)

// NoOp is a sample resolver
func (resolver) NoOp(context.Context) (bool, error) {
	return false, nil
}
