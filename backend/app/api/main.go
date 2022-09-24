package main

import (
	"context"
	"os"
	"time"

	"github.com/hay-kot/homebox/backend/app/api/docs"
	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/internal/config"
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/internal/services"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

var (
	Version   = "0.1.0"
	Commit    = "HEAD"
	BuildTime = "now"
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
	path := ""
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	cfg, err := config.NewConfig(path)
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.Host = cfg.Swagger.Host

	if err := run(cfg); err != nil {
		panic(err)
	}
}

func run(cfg *config.Config) error {
	app := new(cfg)
	app.setupLogger()

	// =========================================================================
	// Initialize Database & Repos

	err := os.MkdirAll(cfg.Storage.Data, 0755)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create data directory")
	}

	c, err := ent.Open("sqlite3", cfg.Storage.SqliteUrl)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("driver", "sqlite").
			Str("url", cfg.Storage.SqliteUrl).
			Msg("failed opening connection to sqlite")
	}
	defer func(c *ent.Client) {
		_ = c.Close()
	}(c)
	if err := c.Schema.Create(context.Background()); err != nil {
		log.Fatal().
			Err(err).
			Str("driver", "sqlite").
			Str("url", cfg.Storage.SqliteUrl).
			Msg("failed creating schema resources")
	}

	app.db = c
	app.repos = repo.EntAllRepos(c)
	app.services = services.NewServices(app.repos, cfg.Storage.Data)

	// =========================================================================
	// Start Server
	app.server = server.NewServer(app.conf.Web.Host, app.conf.Web.Port)
	routes := app.newRouter(app.repos)

	if app.conf.Mode != config.ModeDevelopment {
		app.logRoutes(routes)
	}

	log.Info().Msgf("Starting HTTP Server on %s:%s", app.server.Host, app.server.Port)

	// =========================================================================
	// Start Reoccurring Tasks

	go app.startBgTask(time.Duration(24)*time.Hour, func() {
		_, err := app.repos.AuthTokens.PurgeExpiredTokens(context.Background())
		if err != nil {
			log.Error().
				Err(err).
				Msg("failed to purge expired tokens")
		}
	})

	return app.server.Start(routes)
}
