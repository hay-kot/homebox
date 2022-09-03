package main

import (
	"context"
	"os"
	"time"

	"github.com/hay-kot/content/backend/app/api/docs"
	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/internal/config"
	"github.com/hay-kot/content/backend/internal/repo"
	"github.com/hay-kot/content/backend/internal/services"
	"github.com/hay-kot/content/backend/pkgs/server"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// @title                       Go API Templates
// @version                     1.0
// @description                 This is a simple Rest API Server Template that implements some basic User and Authentication patterns to help you get started and bootstrap your next project!.
// @contact.name                Don't
// @license.name                MIT
// @BasePath                    /api
// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description                 "Type 'Bearer TOKEN' to correctly set the API Key"
func main() {
	// Logger Init
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfgFile := "config.yml"

	cfg, err := config.NewConfig(cfgFile)
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.Host = cfg.Swagger.Host

	if err := run(cfg); err != nil {
		panic(err)
	}
}

func run(cfg *config.Config) error {
	app := NewApp(cfg)

	// =========================================================================
	// Initialize Database & Repos

	c, err := ent.Open(cfg.Database.GetDriver(), cfg.Database.GetUrl())
	if err != nil {
		log.Fatal().
			Err(err).
			Str("driver", cfg.Database.GetDriver()).
			Str("url", cfg.Database.GetUrl()).
			Msg("failed opening connection to sqlite")
	}
	defer func(c *ent.Client) {
		_ = c.Close()
	}(c)
	if err := c.Schema.Create(context.Background()); err != nil {
		log.Fatal().
			Err(err).
			Msg("failed creating schema resources")
	}

	app.db = c
	app.repos = repo.EntAllRepos(c)
	app.services = services.NewServices(app.repos)

	// =========================================================================
	// Start Server

	app.conf.Print()

	app.server = server.NewServer(app.conf.Web.Host, app.conf.Web.Port)

	routes := app.newRouter(app.repos)
	app.LogRoutes(routes)

	app.SeedDatabase(app.repos)

	log.Info().Msgf("Starting HTTP Server on %s:%s", app.server.Host, app.server.Port)

	// =========================================================================
	// Start Reoccurring Tasks

	go app.StartReoccurringTasks(time.Duration(24)*time.Hour, func() {
		_, err := app.repos.AuthTokens.PurgeExpiredTokens(context.Background())
		if err != nil {
			log.Error().
				Err(err).
				Msg("failed to purge expired tokens")
		}
	})

	return app.server.Start(routes)
}
