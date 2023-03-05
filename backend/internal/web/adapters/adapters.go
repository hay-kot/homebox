// Package adapters provides functions to adapt functions to the server.Handler interface.
package adapters

import (
	"context"

	"github.com/google/uuid"
)

type AdapterFunc[T any, Y any] func(context.Context, T) (Y, error)
type IDFunc[T any, Y any] func(context.Context, uuid.UUID, T) (Y, error)
