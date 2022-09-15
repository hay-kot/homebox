package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/internal/services/mappers"
	"github.com/hay-kot/homebox/backend/internal/types"
	"github.com/hay-kot/homebox/backend/pkgs/hasher"
)

var (
	oneWeek              = time.Hour * 24 * 7
	ErrorInvalidLogin    = errors.New("invalid username or password")
	ErrorInvalidToken    = errors.New("invalid token")
	ErrorTokenIdMismatch = errors.New("token id mismatch")
)

type UserService struct {
	repos *repo.AllRepos
}

// RegisterUser creates a new user and group in the data with the provided data. It also bootstraps the user's group
// with default Labels and Locations.
func (svc *UserService) RegisterUser(ctx context.Context, data types.UserRegistration) (*types.UserOut, error) {
	group, err := svc.repos.Groups.Create(ctx, data.GroupName)
	if err != nil {
		return &types.UserOut{}, err
	}

	hashed, _ := hasher.HashPassword(data.User.Password)
	usrCreate := types.UserCreate{
		Name:        data.User.Name,
		Email:       data.User.Email,
		Password:    hashed,
		IsSuperuser: false,
		GroupID:     group.ID,
	}

	usr, err := svc.repos.Users.Create(ctx, usrCreate)
	if err != nil {
		return &types.UserOut{}, err
	}

	for _, label := range defaultLabels() {
		_, err := svc.repos.Labels.Create(ctx, group.ID, label)
		if err != nil {
			return &types.UserOut{}, err
		}
	}

	for _, location := range defaultLocations() {
		_, err := svc.repos.Locations.Create(ctx, group.ID, location)
		if err != nil {
			return &types.UserOut{}, err
		}
	}

	return mappers.ToOutUser(usr, nil)
}

// GetSelf returns the user that is currently logged in based of the token provided within
func (svc *UserService) GetSelf(ctx context.Context, requestToken string) (*types.UserOut, error) {
	hash := hasher.HashToken(requestToken)
	return mappers.ToOutUser(svc.repos.AuthTokens.GetUserFromToken(ctx, hash))
}

func (svc *UserService) UpdateSelf(ctx context.Context, ID uuid.UUID, data types.UserUpdate) (*types.UserOut, error) {
	err := svc.repos.Users.Update(ctx, ID, data)
	if err != nil {
		return &types.UserOut{}, err
	}

	return mappers.ToOutUser(svc.repos.Users.GetOneId(ctx, ID))
}

// ============================================================================
// User Authentication

func (svc *UserService) createToken(ctx context.Context, userId uuid.UUID) (types.UserAuthTokenDetail, error) {
	newToken := hasher.GenerateToken()

	created, err := svc.repos.AuthTokens.CreateToken(ctx, types.UserAuthTokenCreate{
		UserID:    userId,
		TokenHash: newToken.Hash,
		ExpiresAt: time.Now().Add(oneWeek),
	})

	return types.UserAuthTokenDetail{Raw: newToken.Raw, ExpiresAt: created.ExpiresAt}, err
}

func (svc *UserService) Login(ctx context.Context, username, password string) (types.UserAuthTokenDetail, error) {
	usr, err := svc.repos.Users.GetOneEmail(ctx, username)

	if err != nil || !hasher.CheckPasswordHash(password, usr.Password) {
		return types.UserAuthTokenDetail{}, ErrorInvalidLogin
	}

	return svc.createToken(ctx, usr.ID)
}

func (svc *UserService) Logout(ctx context.Context, token string) error {
	hash := hasher.HashToken(token)
	err := svc.repos.AuthTokens.DeleteToken(ctx, hash)
	return err
}

func (svc *UserService) RenewToken(ctx context.Context, token string) (types.UserAuthTokenDetail, error) {
	hash := hasher.HashToken(token)

	dbToken, err := svc.repos.AuthTokens.GetUserFromToken(ctx, hash)

	if err != nil {
		return types.UserAuthTokenDetail{}, ErrorInvalidToken
	}

	newToken, _ := svc.createToken(ctx, dbToken.ID)

	return newToken, nil
}

// DeleteSelf deletes the user that is currently logged based of the provided UUID
// There is _NO_ protection against deleting the wrong user, as such this should only
// be used when the identify of the user has been confirmed.
func (svc *UserService) DeleteSelf(ctx context.Context, ID uuid.UUID) error {
	return svc.repos.Users.Delete(ctx, ID)
}
