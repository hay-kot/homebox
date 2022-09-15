package mappers

import (
	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/internal/types"
)

func ToOutUser(user *ent.User, err error) (*types.UserOut, error) {
	if err != nil {
		return &types.UserOut{}, err
	}
	return &types.UserOut{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		IsSuperuser: user.IsSuperuser,
		GroupName:   user.Edges.Group.Name,
		GroupID:     user.Edges.Group.ID,
	}, nil
}
