package v1

import (
	"net/http"
	"time"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/internal/web/adapters"
	"github.com/hay-kot/homebox/backend/pkgs/server"
)

type (
	GroupInvitationCreate struct {
		Uses      int       `json:"uses" validate:"required,min=1,max=100"`
		ExpiresAt time.Time `json:"expiresAt"`
	}

	GroupInvitation struct {
		Token     string    `json:"token"`
		ExpiresAt time.Time `json:"expiresAt"`
		Uses      int       `json:"uses"`
	}
)

// HandleGroupGet godoc
//
//	@Summary  Get Group
//	@Tags     Group
//	@Produce  json
//	@Success  200 {object} repo.Group
//	@Router   /v1/groups [Get]
//	@Security Bearer
func (ctrl *V1Controller) HandleGroupGet() server.HandlerFunc {
	fn := func(r *http.Request) (repo.Group, error) {
		auth := services.NewContext(r.Context())
		return ctrl.repo.Groups.GroupByID(auth, auth.GID)
	}

	return adapters.Command(fn, http.StatusOK)
}

// HandleGroupUpdate godoc
//
//	@Summary  Update Group
//	@Tags     Group
//	@Produce  json
//	@Param    payload body     repo.GroupUpdate true "User Data"
//	@Success  200     {object} repo.Group
//	@Router   /v1/groups [Put]
//	@Security Bearer
func (ctrl *V1Controller) HandleGroupUpdate() server.HandlerFunc {
	fn := func(r *http.Request, body repo.GroupUpdate) (repo.Group, error) {
		auth := services.NewContext(r.Context())
		return ctrl.svc.Group.UpdateGroup(auth, body)
	}

	return adapters.Action(fn, http.StatusOK)
}

// HandleGroupInvitationsCreate godoc
//
//	@Summary  Create Group Invitation
//	@Tags     Group
//	@Produce  json
//	@Param    payload body     GroupInvitationCreate true "User Data"
//	@Success  200     {object} GroupInvitation
//	@Router   /v1/groups/invitations [Post]
//	@Security Bearer
func (ctrl *V1Controller) HandleGroupInvitationsCreate() server.HandlerFunc {
	fn := func(r *http.Request, body GroupInvitationCreate) (GroupInvitation, error) {
		if body.ExpiresAt.IsZero() {
			body.ExpiresAt = time.Now().Add(time.Hour * 24)
		}

		auth := services.NewContext(r.Context())

		token, err := ctrl.svc.Group.NewInvitation(auth, body.Uses, body.ExpiresAt)

		return GroupInvitation{
			Token:     token,
			ExpiresAt: body.ExpiresAt,
			Uses:      body.Uses,
		}, err
	}

	return adapters.Action(fn, http.StatusCreated)
}
