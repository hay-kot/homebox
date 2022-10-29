package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/rs/zerolog/log"
)

func (ctrl *V1Controller) routeID(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	ID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		log.Err(err).Msg("failed to parse id")
		return uuid.Nil, validate.ErrInvalidID
	}

	return ID, nil
}
