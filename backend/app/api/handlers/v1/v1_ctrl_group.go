package v1

import (
	"net/http"
	"time"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
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
// @Success  200 {object} repo.GroupStatistics
// @Router   /v1/groups/statistics [Get]
// @Security Bearer
func (ctrl *V1Controller) HandleGroupStatistics() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())

		stats, err := ctrl.repo.Groups.GroupStatistics(ctx, ctx.GID)
		if err != nil {
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusOK, stats)
	}
}

// HandleGroupGet godoc
// @Summary  Get the current user's group
// @Tags     Group
// @Produce  json
// @Success  200 {object} repo.Group
// @Router   /v1/groups [Get]
// @Security Bearer
func (ctrl *V1Controller) HandleGroupGet() server.HandlerFunc {
	return ctrl.handleGroupGeneral()
}

// HandleGroupUpdate godoc
// @Summary  Updates some fields of the current users group
// @Tags     Group
// @Produce  json
// @Param    payload body     repo.GroupUpdate true "User Data"
// @Success  200     {object} repo.Group
// @Router   /v1/groups [Put]
// @Security Bearer
func (ctrl *V1Controller) HandleGroupUpdate() server.HandlerFunc {
	return ctrl.handleGroupGeneral()
}

func (ctrl *V1Controller) handleGroupGeneral() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())

		switch r.Method {
		case http.MethodGet:
			group, err := ctrl.repo.Groups.GroupByID(ctx, ctx.GID)
			if err != nil {
				log.Err(err).Msg("failed to get group")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}

			return server.Respond(w, http.StatusOK, group)

		case http.MethodPut:
			data := repo.GroupUpdate{}
			if err := server.Decode(r, &data); err != nil {
				return validate.NewRequestError(err, http.StatusBadRequest)
			}

			group, err := ctrl.svc.Group.UpdateGroup(ctx, data)
			if err != nil {
				log.Err(err).Msg("failed to update group")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}

			return server.Respond(w, http.StatusOK, group)
		}

		return nil
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
func (ctrl *V1Controller) HandleGroupInvitationsCreate() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		data := GroupInvitationCreate{}
		if err := server.Decode(r, &data); err != nil {
			log.Err(err).Msg("failed to decode user registration data")
			return validate.NewRequestError(err, http.StatusBadRequest)
		}

		if data.ExpiresAt.IsZero() {
			data.ExpiresAt = time.Now().Add(time.Hour * 24)
		}

		ctx := services.NewContext(r.Context())

		token, err := ctrl.svc.Group.NewInvitation(ctx, data.Uses, data.ExpiresAt)
		if err != nil {
			log.Err(err).Msg("failed to create new token")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusCreated, GroupInvitation{
			Token:     token,
			ExpiresAt: data.ExpiresAt,
			Uses:      data.Uses,
		})
	}
}
