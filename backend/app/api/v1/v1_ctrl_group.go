package v1

import (
	"net/http"
	"strings"
	"time"

	"github.com/hay-kot/homebox/backend/internal/repo"
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

// HandleGroupGet godoc
// @Summary  Get the current user's group
// @Tags     Group
// @Produce  json
// @Success  200 {object} repo.Group
// @Router   /v1/groups [Get]
// @Security Bearer
func (ctrl *V1Controller) HandleGroupGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := services.NewContext(r.Context())

		group, err := ctrl.svc.Group.Get(ctx)
		if err != nil {
			log.Err(err).Msg("failed to get group")
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		server.Respond(w, http.StatusOK, group)

	}
}

// HandleGroupUpdate godoc
// @Summary  Updates some fields of the current users group
// @Tags     Group
// @Produce  json
// @Param    payload body     repo.GroupUpdate true "User Data"
// @Success  200     {object} repo.Group
// @Router   /v1/groups [Put]
// @Security Bearer
func (ctrl *V1Controller) HandleGroupUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := repo.GroupUpdate{}

		if err := server.Decode(r, &data); err != nil {
			server.RespondError(w, http.StatusBadRequest, err)
			return
		}

		ctx := services.NewContext(r.Context())

		group, err := ctrl.svc.Group.UpdateGroup(ctx, data)
		if err != nil {
			log.Err(err).Msg("failed to update group")
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}
		group.Currency = strings.ToUpper(group.Currency) // TODO: Hack to fix the currency enums being lower case
		server.Respond(w, http.StatusOK, group)
	}
}

// HandleGroupInvitationsCreate godoc
// @Summary  Get the current user
// @Tags     Group
// @Produce  json
// @Param    payload body     GroupInvitationCreate true "User Data"
// @Success  200     {object} GroupInvitation
// @Router   /v1/groups/invitations [Post]
// @Security Bearer
func (ctrl *V1Controller) HandleGroupInvitationsCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := GroupInvitationCreate{}

		if err := server.Decode(r, &data); err != nil {
			log.Err(err).Msg("failed to decode user registration data")
			server.RespondError(w, http.StatusBadRequest, err)
			return
		}

		if data.ExpiresAt.IsZero() {
			data.ExpiresAt = time.Now().Add(time.Hour * 24)
		}

		ctx := services.NewContext(r.Context())

		token, err := ctrl.svc.Group.NewInvitation(ctx, data.Uses, data.ExpiresAt)
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
