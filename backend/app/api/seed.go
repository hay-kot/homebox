package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/internal/repo"
	"github.com/hay-kot/content/backend/internal/types"
	"github.com/hay-kot/content/backend/pkgs/hasher"
	"github.com/rs/zerolog/log"
)

const (
	DefaultGroup    = "Default"
	DefaultName     = "Admin"
	DefaultEmail    = "admin@admin.com"
	DefaultPassword = "admin"
)

// EnsureAdministrator ensures that there is at least one superuser in the database
// if one isn't found a default is generate using the default credentials
func (a *app) EnsureAdministrator() {
	superusers, err := a.repos.Users.GetSuperusers(context.Background())

	if err != nil {
		log.Fatal().Err(err).Msg("failed to get superusers")
	}
	if len(superusers) > 0 {
		return
	}

	pw, _ := hasher.HashPassword(DefaultPassword)
	newSuperUser := types.UserCreate{
		Name:        DefaultName,
		Email:       DefaultEmail,
		IsSuperuser: true,
		Password:    pw,
	}

	log.Info().
		Str("name", newSuperUser.Name).
		Str("email", newSuperUser.Email).
		Msg("no superusers found, creating default superuser")

	_, err = a.repos.Users.Create(context.Background(), newSuperUser)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to create default superuser")
	}

}

func (a *app) SeedDatabase(repos *repo.AllRepos) {
	if !a.conf.Seed.Enabled {
		return
	}

	group, err := repos.Groups.Create(context.Background(), DefaultGroup)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create default group")
	}

	for _, user := range a.conf.Seed.Users {

		// Check if User Exists
		usr, _ := repos.Users.GetOneEmail(context.Background(), user.Email)

		if usr.ID != uuid.Nil {
			log.Info().Str("email", user.Email).Msg("user already exists, skipping")
			continue
		}

		hashedPw, err := hasher.HashPassword(user.Password)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to hash password")
		}

		_, err = repos.Users.Create(context.Background(), types.UserCreate{
			Name:        user.Name,
			Email:       user.Email,
			IsSuperuser: user.IsSuperuser,
			Password:    hashedPw,
			GroupID:     group.ID,
		})

		if err != nil {
			log.Fatal().Err(err).Msg("failed to create user")
		}

		log.Info().Str("email", user.Email).Msg("created user")
	}
}
