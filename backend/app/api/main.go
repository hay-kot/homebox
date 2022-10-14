package main

import (
	"context"
	"os"
	"path/filepath"
	"time"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/hay-kot/homebox/backend/app/api/docs"
	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/internal/config"
	"github.com/hay-kot/homebox/backend/internal/migrations"
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/internal/services"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

var (
	version   = "nightly"
	commit    = "HEAD"
	buildTime = "now"
)

// @title                      Go API Templates
// @version                    1.0
// @description                This is a simple Rest API Server Template that implements some basic User and Authentication patterns to help you get started and bootstrap your next project!.
// @contact.name               Don't
// @license.name               MIT
// @BasePath                   /api
// @securityDefinitions.apikey Bearer
// @in                         header
// @name                       Authorization
// @description                "Type 'Bearer TOKEN' to correctly set the API Key"
func main() {
	cfg, err := config.New()
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
		err := c.Close()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to close database connection")
		}
	}(c)

	temp := filepath.Join(os.TempDir(), "migrations")

	err = migrations.Write(temp)
	if err != nil {
		return err
	}

	dir, err := atlas.NewLocalDir(temp)
	if err != nil {
		return err
	}

	options := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	}

	err = c.Schema.Create(context.Background(), options...)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("driver", "sqlite").
			Str("url", cfg.Storage.SqliteUrl).
			Msg("failed creating schema resources")
	}

	err = os.RemoveAll(temp)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to remove temporary directory for database migrations")
		return err
	}

	app.db = c
	app.repos = repo.New(c, cfg.Storage.Data)
	app.services = services.NewServices(app.repos)

	// =========================================================================
	// Start Server
	app.server = server.NewServer(
		server.WithHost(app.conf.Web.Host),
		server.WithPort(app.conf.Web.Port),
	)

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

	// TODO: Remove through external API that does setup
	if cfg.Demo {
		log.Info().Msg("Running in demo mode, creating demo data")
		app.SetupDemo()
	}

	return app.server.Start(routes)
}
