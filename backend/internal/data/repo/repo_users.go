package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/hay-kot/homebox/backend/internal/data/ent/user"
)

type UserRepository struct {
	db *ent.Client
}

type (
	// UserCreate is the Data object contain the requirements of creating a user
	// in the database. It should to create users from an API unless the user has
	// rights to create SuperUsers. For regular user in data use the UserIn struct.
	UserCreate struct {
		Name        string    `json:"name"`
		Email       string    `json:"email"`
		Password    string    `json:"password"`
		IsSuperuser bool      `json:"isSuperuser"`
		GroupID     uuid.UUID `json:"groupID"`
		IsOwner     bool      `json:"isOwner"`
	}

	UserUpdate struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	UserOut struct {
		ID           uuid.UUID `json:"id"`
		Name         string    `json:"name"`
		Email        string    `json:"email"`
		IsSuperuser  bool      `json:"isSuperuser"`
		GroupID      uuid.UUID `json:"groupId"`
		GroupName    string    `json:"groupName"`
		PasswordHash string    `json:"-"`
		IsOwner      bool      `json:"isOwner"`
	}
)

var (
	mapUserOutErr  = mapTErrFunc(mapUserOut)
	mapUsersOutErr = mapTEachErrFunc(mapUserOut)
)

func mapUserOut(user *ent.User) UserOut {
	return UserOut{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		IsSuperuser:  user.IsSuperuser,
		GroupID:      user.Edges.Group.ID,
		GroupName:    user.Edges.Group.Name,
		PasswordHash: user.Password,
		IsOwner:      user.Role == "owner",
	}
}

func (r *UserRepository) GetOneID(ctx context.Context, ID uuid.UUID) (UserOut, error) {
	return mapUserOutErr(r.db.User.Query().
		Where(user.ID(ID)).
		WithGroup().
		Only(ctx))
}

func (r *UserRepository) GetOneEmail(ctx context.Context, email string) (UserOut, error) {
	return mapUserOutErr(r.db.User.Query().
		Where(user.EmailEqualFold(email)).
		WithGroup().
		Only(ctx),
	)
}

func (r *UserRepository) GetAll(ctx context.Context) ([]UserOut, error) {
	return mapUsersOutErr(r.db.User.Query().WithGroup().All(ctx))
}

func (r *UserRepository) Create(ctx context.Context, usr UserCreate) (UserOut, error) {
	role := user.RoleUser
	if usr.IsOwner {
		role = user.RoleOwner
	}

	entUser, err := r.db.User.
		Create().
		SetName(usr.Name).
		SetEmail(usr.Email).
		SetPassword(usr.Password).
		SetIsSuperuser(usr.IsSuperuser).
		SetGroupID(usr.GroupID).
		SetRole(role).
		Save(ctx)
	if err != nil {
		return UserOut{}, err
	}

	return r.GetOneID(ctx, entUser.ID)
}

func (r *UserRepository) Update(ctx context.Context, ID uuid.UUID, data UserUpdate) error {
	q := r.db.User.Update().
		Where(user.ID(ID)).
		SetName(data.Name).
		SetEmail(data.Email)

	_, err := q.Save(ctx)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.User.Delete().Where(user.ID(id)).Exec(ctx)
	return err
}

func (r *UserRepository) DeleteAll(ctx context.Context) error {
	_, err := r.db.User.Delete().Exec(ctx)
	return err
}

func (r *UserRepository) GetSuperusers(ctx context.Context) ([]*ent.User, error) {
	users, err := r.db.User.Query().Where(user.IsSuperuser(true)).All(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) ChangePassword(ctx context.Context, UID uuid.UUID, pw string) error {
	return r.db.User.UpdateOneID(UID).SetPassword(pw).Exec(ctx)
}
