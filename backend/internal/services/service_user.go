package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/pkgs/hasher"
	"github.com/rs/zerolog/log"
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

type (
	UserRegistration struct {
		GroupToken string `json:"token"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		Password   string `json:"password"`
	}
	UserAuthTokenDetail struct {
		Raw       string    `json:"raw"`
		ExpiresAt time.Time `json:"expiresAt"`
	}
	LoginForm struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

// RegisterUser creates a new user and group in the data with the provided data. It also bootstraps the user's group
// with default Labels and Locations.
func (svc *UserService) RegisterUser(ctx context.Context, data UserRegistration) (repo.UserOut, error) {
	log.Debug().
		Str("name", data.Name).
		Str("email", data.Email).
		Str("groupToken", data.GroupToken).
		Msg("Registering new user")

	var (
		err   error
		group repo.Group
		token repo.GroupInvitation
	)

	if data.GroupToken == "" {
		group, err = svc.repos.Groups.GroupCreate(ctx, "Home")
		if err != nil {
			log.Err(err).Msg("Failed to create group")
			return repo.UserOut{}, err
		}
	} else {
		token, err = svc.repos.Groups.InvitationGet(ctx, hasher.HashToken(data.GroupToken))
		if err != nil {
			log.Err(err).Msg("Failed to get invitation token")
			return repo.UserOut{}, err
		}
		group = token.Group
	}

	hashed, _ := hasher.HashPassword(data.Password)
	usrCreate := repo.UserCreate{
		Name:        data.Name,
		Email:       data.Email,
		Password:    hashed,
		IsSuperuser: false,
		GroupID:     group.ID,
	}

	usr, err := svc.repos.Users.Create(ctx, usrCreate)
	if err != nil {
		return repo.UserOut{}, err
	}

	for _, label := range defaultLabels() {
		_, err := svc.repos.Labels.Create(ctx, group.ID, label)
		if err != nil {
			return repo.UserOut{}, err
		}
	}

	for _, location := range defaultLocations() {
		_, err := svc.repos.Locations.Create(ctx, group.ID, location)
		if err != nil {
			return repo.UserOut{}, err
		}
	}

	// Decrement the invitation token if it was used
	if token.ID != uuid.Nil {
		err = svc.repos.Groups.InvitationUpdate(ctx, token.ID, token.Uses-1)
		if err != nil {
			log.Err(err).Msg("Failed to update invitation token")
			return repo.UserOut{}, err
		}
	}

	return usr, nil
}

// GetSelf returns the user that is currently logged in based of the token provided within
func (svc *UserService) GetSelf(ctx context.Context, requestToken string) (repo.UserOut, error) {
	hash := hasher.HashToken(requestToken)
	return svc.repos.AuthTokens.GetUserFromToken(ctx, hash)
}

func (svc *UserService) UpdateSelf(ctx context.Context, ID uuid.UUID, data repo.UserUpdate) (repo.UserOut, error) {
	err := svc.repos.Users.Update(ctx, ID, data)
	if err != nil {
		return repo.UserOut{}, err
	}

	return svc.repos.Users.GetOneId(ctx, ID)
}

// ============================================================================
// User Authentication

func (svc *UserService) createToken(ctx context.Context, userId uuid.UUID) (UserAuthTokenDetail, error) {
	newToken := hasher.GenerateToken()

	created, err := svc.repos.AuthTokens.CreateToken(ctx, repo.UserAuthTokenCreate{
		UserID:    userId,
		TokenHash: newToken.Hash,
		ExpiresAt: time.Now().Add(oneWeek),
	})

	return UserAuthTokenDetail{Raw: newToken.Raw, ExpiresAt: created.ExpiresAt}, err
}

func (svc *UserService) Login(ctx context.Context, username, password string) (UserAuthTokenDetail, error) {
	usr, err := svc.repos.Users.GetOneEmail(ctx, username)

	if err != nil || !hasher.CheckPasswordHash(password, usr.PasswordHash) {
		return UserAuthTokenDetail{}, ErrorInvalidLogin
	}

	return svc.createToken(ctx, usr.ID)
}

func (svc *UserService) Logout(ctx context.Context, token string) error {
	hash := hasher.HashToken(token)
	err := svc.repos.AuthTokens.DeleteToken(ctx, hash)
	return err
}

func (svc *UserService) RenewToken(ctx context.Context, token string) (UserAuthTokenDetail, error) {
	hash := hasher.HashToken(token)

	dbToken, err := svc.repos.AuthTokens.GetUserFromToken(ctx, hash)

	if err != nil {
		return UserAuthTokenDetail{}, ErrorInvalidToken
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

func (svc *UserService) NewInvitation(ctx Context, uses int, expiresAt time.Time) (string, error) {
	token := hasher.GenerateToken()

	_, err := svc.repos.Groups.InvitationCreate(ctx, ctx.GID, repo.GroupInvitationCreate{
		Token:     token.Hash,
		Uses:      uses,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return "", err
	}

	return token.Raw, nil
}
