package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/ent/user"
	"github.com/hay-kot/homebox/backend/internal/types"
)

type UserRepository struct {
	db *ent.Client
}

func (e *UserRepository) GetOneId(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return e.db.User.Query().
		Where(user.ID(id)).
		WithGroup().
		Only(ctx)
}

func (e *UserRepository) GetOneEmail(ctx context.Context, email string) (*ent.User, error) {
	return e.db.User.Query().
		Where(user.Email(email)).
		WithGroup().
		Only(ctx)
}

func (e *UserRepository) GetAll(ctx context.Context) ([]*ent.User, error) {
	return e.db.User.Query().WithGroup().All(ctx)
}

func (e *UserRepository) Create(ctx context.Context, usr types.UserCreate) (*ent.User, error) {
	err := usr.Validate()
	if err != nil {
		return &ent.User{}, err
	}

	entUser, err := e.db.User.
		Create().
		SetName(usr.Name).
		SetEmail(usr.Email).
		SetPassword(usr.Password).
		SetIsSuperuser(usr.IsSuperuser).
		SetGroupID(usr.GroupID).
		Save(ctx)
	if err != nil {
		return entUser, err
	}

	return e.GetOneId(ctx, entUser.ID)
}

func (e *UserRepository) Update(ctx context.Context, ID uuid.UUID, data types.UserUpdate) error {
	q := e.db.User.Update().
		Where(user.ID(ID)).
		SetName(data.Name).
		SetEmail(data.Email)

	_, err := q.Save(ctx)
	return err
}

func (e *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := e.db.User.Delete().Where(user.ID(id)).Exec(ctx)
	return err
}

func (e *UserRepository) DeleteAll(ctx context.Context) error {
	_, err := e.db.User.Delete().Exec(ctx)
	return err
}

func (e *UserRepository) GetSuperusers(ctx context.Context) ([]*ent.User, error) {
	users, err := e.db.User.Query().Where(user.IsSuperuser(true)).All(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
