package v1

import (
	"net/http"
	"time"

	"github.com/hay-kot/homebox/backend/internal/services"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog/log"
)

type (
	GroupInvitationCreate struct {
		Uses      int       `json:"uses"`
		ExpiresAt time.Time `json:"expiresAt"`
	}

	GroupInvitation struct {
		Token     string    `json:"token"`
		ExpiresAt time.Time `json:"expiresAt"`
		Uses      int       `json:"uses"`
	}
)

// HandleUserSelf godoc
// @Summary   Get the current user
// @Tags      User
// @Produce   json
// @Param     payload  body      GroupInvitationCreate  true  "User Data"
// @Success   200      {object}  GroupInvitation
// @Router    /v1/groups/invitations [Post]
// @Security  Bearer
func (ctrl *V1Controller) HandleGroupInvitationsCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := GroupInvitationCreate{}

		if err := server.Decode(r, &data); err != nil {
			log.Err(err).Msg("failed to decode user registration data")
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		if data.ExpiresAt.IsZero() {
			data.ExpiresAt = time.Now().Add(time.Hour * 24)
		}

		ctx := services.NewContext(r.Context())

		token, err := ctrl.svc.User.NewInvitation(ctx, data.Uses, data.ExpiresAt)
		if err != nil {
			log.Err(err).Msg("failed to create new token")
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		server.Respond(w, http.StatusCreated, GroupInvitation{
			Token:     token,
			ExpiresAt: data.ExpiresAt,
			Uses:      data.Uses,
		})
	}
}
