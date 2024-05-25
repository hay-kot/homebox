package adapters

import (
	"net/http"

	"github.com/google/uuid"
)

type (
	AdapterFunc[T any, Y any] func(*http.Request, T) (Y, error)
	IDFunc[T any, Y any]      func(*http.Request, uuid.UUID, T) (Y, error)
)
