package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/hay-kot/homebox/backend/internal/core/currencies"
	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/core/services/reporting/eventbus"
	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/hay-kot/homebox/backend/internal/data/migrations"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/internal/sys/config"
	"github.com/hay-kot/homebox/backend/internal/web/mid"
	"github.com/hay-kot/httpkit/errchain"
	"github.com/hay-kot/httpkit/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	_ "github.com/hay-kot/homebox/backend/pkgs/cgofreesqlite"
)

var (
	version   = "nightly"
	commit    = "HEAD"
	buildTime = "now"
)

func build() string {
	short := commit
	if len(short) > 7 {
		short = short[:7]
	}

	return fmt.Sprintf("%s, commit %s, built at %s", version, short, buildTime)
}

// @title                      Homebox API
// @version                    1.0
// @description                Track, Manage, and Organize your Things.
// @contact.name               Don't
// @BasePath                   /api
// @securityDefinitions.apikey Bearer
// @in                         header
// @name                       Authorization
// @description                "Type 'Bearer TOKEN' to correctly set the API Key"
func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	cfg, err := config.New(build(), "Homebox inventory management system")
	if err != nil {
		panic(err)
	}

	if err := run(cfg); err != nil {
		panic(err)
	}
}

func run(cfg *config.Config) error {
	app := new(cfg)
	app.setupLogger()

	// =========================================================================
	// Initialize Database & Repos

	c, err := ent.Open("sqlite3", cfg.Storage.SqliteURL)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("driver", "sqlite").
			Str("url", cfg.Storage.SqliteURL).
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
			Str("url", cfg.Storage.SqliteURL).
			Msg("failed creating schema resources")
	}

	err = os.RemoveAll(temp)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to remove temporary directory for database migrations")
		return err
	}

	collectFuncs := []currencies.CollectorFunc{
		currencies.CollectDefaults(),
	}

	if cfg.Options.CurrencyConfig != "" {
		log.Info().
			Str("path", cfg.Options.CurrencyConfig).
			Msg("loading currency config file")

		content, err := os.ReadFile(cfg.Options.CurrencyConfig)
		if err != nil {
			log.Fatal().
				Err(err).
				Str("path", cfg.Options.CurrencyConfig).
				Msg("failed to read currency config file")
		}

		collectFuncs = append(collectFuncs, currencies.CollectJSON(bytes.NewReader(content)))
	}

	currencies, err := currencies.CollectionCurrencies(collectFuncs...)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to collect currencies")
	}

	app.bus = eventbus.New()
	app.db = c
	app.repos = repo.New(c, app.bus, cfg.Storage.Data)
	app.services = services.New(
		app.repos,
		services.WithAutoIncrementAssetID(cfg.Options.AutoIncrementAssetID),
		services.WithCurrencies(currencies),
	)

	// =========================================================================
	// Start Server

	logger := log.With().Caller().Logger()

	router := chi.NewMux()
	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		mid.Logger(logger),
		middleware.Recoverer,
		middleware.StripSlashes,
	)

	chain := errchain.New(mid.Errors(app.server, logger))

	app.mountRoutes(router, chain, app.repos)

	app.server = server.NewServer(
		server.WithHost(app.conf.Web.Host),
		server.WithPort(app.conf.Web.Port),
		server.WithReadTimeout(app.conf.Web.ReadTimeout),
		server.WithWriteTimeout(app.conf.Web.WriteTimeout),
		server.WithIdleTimeout(app.conf.Web.IdleTimeout),
	)
	log.Info().Msgf("Starting HTTP Server on %s:%s", app.server.Host, app.server.Port)

	// =========================================================================
	// Start Reoccurring Tasks

	go app.bus.Run()

	go app.startBgTask(time.Duration(24)*time.Hour, func() {
		_, err := app.repos.AuthTokens.PurgeExpiredTokens(context.Background())
		if err != nil {
			log.Error().
				Err(err).
				Msg("failed to purge expired tokens")
		}
	})
	go app.startBgTask(time.Duration(24)*time.Hour, func() {
		_, err := app.repos.Groups.InvitationPurge(context.Background())
		if err != nil {
			log.Error().
				Err(err).
				Msg("failed to purge expired invitations")
		}
	})
	go app.startBgTask(time.Duration(1)*time.Hour, func() {
		now := time.Now()

		if now.Hour() == 8 {
			fmt.Println("run notifiers")
			err := app.services.BackgroundService.SendNotifiersToday(context.Background())
			if err != nil {
				log.Error().
					Err(err).
					Msg("failed to send notifiers")
			}
		}
	})

	// TODO: Remove through external API that does setup
	if cfg.Demo {
		log.Info().Msg("Running in demo mode, creating demo data")
		app.SetupDemo()
	}

	if cfg.Debug.Enabled {
		debugrouter := app.debugRouter()
		go func() {
			if err := http.ListenAndServe(":"+cfg.Debug.Port, debugrouter); err != nil {
				log.Fatal().Err(err).Msg("failed to start debug server")
			}
		}()
	}

	return app.server.Start(router)
}
