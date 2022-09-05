package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/ent/user"
	"github.com/hay-kot/content/backend/internal/types"
)

type EntUserRepository struct {
	db *ent.Client
}

func (e *EntUserRepository) GetOneId(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return e.db.User.Query().
		Where(user.ID(id)).
		WithGroup().
		Only(ctx)
}

func (e *EntUserRepository) GetOneEmail(ctx context.Context, email string) (*ent.User, error) {
	return e.db.User.Query().
		Where(user.Email(email)).
		WithGroup().
		Only(ctx)
}

func (e *EntUserRepository) GetAll(ctx context.Context) ([]*ent.User, error) {
	return e.db.User.Query().WithGroup().All(ctx)
}

func (e *EntUserRepository) Create(ctx context.Context, usr types.UserCreate) (*ent.User, error) {
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

func (e *EntUserRepository) Update(ctx context.Context, ID uuid.UUID, data types.UserUpdate) error {
	bldr := e.db.User.Update().Where(user.ID(ID))

	if data.Name != nil {
		bldr = bldr.SetName(*data.Name)
	}

	if data.Email != nil {
		bldr = bldr.SetEmail(*data.Email)
	}

	// TODO: FUTURE
	// if data.Password != nil {
	// 	bldr = bldr.SetPassword(*data.Password)
	// }

	// if data.IsSuperuser != nil {
	// 	bldr = bldr.SetIsSuperuser(*data.IsSuperuser)
	// }

	_, err := bldr.Save(ctx)
	return err
}

func (e *EntUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := e.db.User.Delete().Where(user.ID(id)).Exec(ctx)
	return err
}

func (e *EntUserRepository) DeleteAll(ctx context.Context) error {
	_, err := e.db.User.Delete().Exec(ctx)
	return err
}

func (e *EntUserRepository) GetSuperusers(ctx context.Context) ([]*ent.User, error) {
	users, err := e.db.User.Query().Where(user.IsSuperuser(true)).All(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
