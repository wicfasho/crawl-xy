// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	InsertPage(ctx context.Context, arg InsertPageParams) (uuid.UUID, error)
}

var _ Querier = (*Queries)(nil)
