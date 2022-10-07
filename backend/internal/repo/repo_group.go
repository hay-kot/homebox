package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/ent/groupinvitationtoken"
)

type GroupRepository struct {
	db *ent.Client
}

type (
	Group struct {
		ID        uuid.UUID
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
		Currency  string
	}

	GroupInvitationCreate struct {
		Token     []byte    `json:"-"`
		ExpiresAt time.Time `json:"expiresAt"`
		Uses      int       `json:"uses"`
	}

	GroupInvitation struct {
		ID        uuid.UUID `json:"id"`
		ExpiresAt time.Time `json:"expiresAt"`
		Uses      int       `json:"uses"`
		Group     Group     `json:"group"`
	}
)

var (
	mapToGroupErr = mapTErrFunc(mapToGroup)
)

func mapToGroup(g *ent.Group) Group {
	return Group{
		ID:        g.ID,
		Name:      g.Name,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
		Currency:  g.Currency.String(),
	}
}

var (
	mapToGroupInvitationErr = mapTErrFunc(mapToGroupInvitation)
)

func mapToGroupInvitation(g *ent.GroupInvitationToken) GroupInvitation {
	return GroupInvitation{
		ID:        g.ID,
		ExpiresAt: g.ExpiresAt,
		Uses:      g.Uses,
		Group:     mapToGroup(g.Edges.Group),
	}
}

func (r *GroupRepository) GroupCreate(ctx context.Context, name string) (Group, error) {
	return mapToGroupErr(r.db.Group.Create().
		SetName(name).
		Save(ctx))
}

func (r *GroupRepository) GroupByID(ctx context.Context, id uuid.UUID) (Group, error) {
	return mapToGroupErr(r.db.Group.Get(ctx, id))
}

func (r *GroupRepository) InvitationGet(ctx context.Context, token []byte) (GroupInvitation, error) {
	return mapToGroupInvitationErr(r.db.GroupInvitationToken.Query().
		Where(groupinvitationtoken.Token(token)).
		WithGroup().
		Only(ctx))
}

func (r *GroupRepository) InvitationCreate(ctx context.Context, groupID uuid.UUID, invite GroupInvitationCreate) (GroupInvitation, error) {
	entity, err := r.db.GroupInvitationToken.Create().
		SetGroupID(groupID).
		SetToken(invite.Token).
		SetExpiresAt(invite.ExpiresAt).
		SetUses(invite.Uses).
		Save(ctx)

	if err != nil {
		return GroupInvitation{}, err
	}

	return r.InvitationGet(ctx, entity.Token)
}

func (r *GroupRepository) InvitationUpdate(ctx context.Context, id uuid.UUID, uses int) error {
	_, err := r.db.GroupInvitationToken.UpdateOneID(id).SetUses(uses).Save(ctx)
	return err
}

// InvitationPurge removes all expired invitations or those that have been used up.
// It returns the number of deleted invitations.
func (r *GroupRepository) InvitationPurge(ctx context.Context) (amount int, err error) {
	q := r.db.GroupInvitationToken.Delete()
	q.Where(groupinvitationtoken.Or(
		groupinvitationtoken.ExpiresAtLT(time.Now()),
		groupinvitationtoken.UsesLTE(0),
	))

	return q.Exec(ctx)
}
