package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/hay-kot/git-web-template/backend/app/api/docs"
	"github.com/hay-kot/git-web-template/backend/ent"
	"github.com/hay-kot/git-web-template/backend/internal/config"
	"github.com/hay-kot/git-web-template/backend/internal/repo"
	"github.com/hay-kot/git-web-template/backend/internal/services"
	"github.com/hay-kot/git-web-template/backend/pkgs/logger"
	"github.com/hay-kot/git-web-template/backend/pkgs/server"
	_ "github.com/mattn/go-sqlite3"
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
	// Setup Logger

	var wrt io.Writer
	wrt = os.Stdout
	if app.conf.Log.File != "" {
		f, err := os.OpenFile(app.conf.Log.File, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer func(f *os.File) {
			_ = f.Close()
		}(f)
		wrt = io.MultiWriter(wrt, f)
	}

	app.logger = logger.New(wrt, logger.LevelDebug)

	// =========================================================================
	// Initialize Database & Repos

	c, err := ent.Open(cfg.Database.GetDriver(), cfg.Database.GetUrl())
	if err != nil {
		app.logger.Fatal(err, logger.Props{
			"details":  "failed to connect to database",
			"database": cfg.Database.GetDriver(),
			"url":      cfg.Database.GetUrl(),
		})
	}
	defer func(c *ent.Client) {
		_ = c.Close()
	}(c)
	if err := c.Schema.Create(context.Background()); err != nil {
		app.logger.Fatal(err, logger.Props{
			"details": "failed to create schema",
		})
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

	app.EnsureAdministrator()
	app.SeedDatabase(app.repos)

	app.logger.Info("Starting HTTP Server", logger.Props{
		"host": app.server.Host,
		"port": app.server.Port,
	})

	// =========================================================================
	// Start Reoccurring Tasks

	go app.StartReoccurringTasks(time.Duration(24)*time.Hour, func() {
		app.repos.AuthTokens.PurgeExpiredTokens(context.Background())
	})

	return app.server.Start(routes)
}
