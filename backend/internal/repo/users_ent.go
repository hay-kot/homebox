package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/git-web-template/backend/ent"
	"github.com/hay-kot/git-web-template/backend/ent/user"
	"github.com/hay-kot/git-web-template/backend/internal/types"
)

type EntUserRepository struct {
	db *ent.Client
}

func (e *EntUserRepository) toUserOut(usr *types.UserOut, entUsr *ent.User) {
	usr.ID = entUsr.ID
	usr.Password = entUsr.Password
	usr.Name = entUsr.Name
	usr.Email = entUsr.Email
	usr.IsSuperuser = entUsr.IsSuperuser
}

func (e *EntUserRepository) GetOneId(ctx context.Context, id uuid.UUID) (types.UserOut, error) {
	usr, err := e.db.User.Query().Where(user.ID(id)).Only(ctx)

	usrOut := types.UserOut{}

	if err != nil {
		return usrOut, err
	}

	e.toUserOut(&usrOut, usr)

	return usrOut, nil
}

func (e *EntUserRepository) GetOneEmail(ctx context.Context, email string) (types.UserOut, error) {
	usr, err := e.db.User.Query().Where(user.Email(email)).Only(ctx)

	usrOut := types.UserOut{}

	if err != nil {
		return usrOut, err
	}

	e.toUserOut(&usrOut, usr)

	return usrOut, nil
}

func (e *EntUserRepository) GetAll(ctx context.Context) ([]types.UserOut, error) {
	users, err := e.db.User.Query().All(ctx)

	if err != nil {
		return nil, err
	}

	var usrs []types.UserOut

	for _, usr := range users {
		usrOut := types.UserOut{}
		e.toUserOut(&usrOut, usr)
		usrs = append(usrs, usrOut)
	}

	return usrs, nil
}

func (e *EntUserRepository) Create(ctx context.Context, usr types.UserCreate) (types.UserOut, error) {
	err := usr.Validate()
	usrOut := types.UserOut{}

	if err != nil {
		return usrOut, err
	}

	entUser, err := e.db.User.
		Create().
		SetName(usr.Name).
		SetEmail(usr.Email).
		SetPassword(usr.Password).
		SetIsSuperuser(usr.IsSuperuser).
		Save(ctx)

	e.toUserOut(&usrOut, entUser)

	return usrOut, err
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

func (e *EntUserRepository) GetSuperusers(ctx context.Context) ([]types.UserOut, error) {
	users, err := e.db.User.Query().Where(user.IsSuperuser(true)).All(ctx)

	if err != nil {
		return nil, err
	}

	var usrs []types.UserOut

	for _, usr := range users {
		usrOut := types.UserOut{}
		e.toUserOut(&usrOut, usr)
		usrs = append(usrs, usrOut)
	}

	return usrs, nil
}
