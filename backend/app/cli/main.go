package main

import (
	"context"
	"log"
	"os"

	"github.com/hay-kot/git-web-template/backend/ent"
	"github.com/hay-kot/git-web-template/backend/internal/config"
	"github.com/hay-kot/git-web-template/backend/internal/repo"
	_ "github.com/mattn/go-sqlite3"

	"github.com/urfave/cli/v2"
)

func main() {
	cfg, err := config.NewConfig("config.yml")

	if err != nil {
		panic(err)
	}

	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg *config.Config) error {
	// =========================================================================
	// Initialize Database
	c, err := ent.Open(cfg.Database.GetDriver(), cfg.Database.GetUrl())
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer func(c *ent.Client) {
		_ = c.Close()
	}(c)
	if err := c.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Create App
	a := &app{
		repos: repo.EntAllRepos(c),
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "users",
				Aliases: []string{"u"},
				Usage:   "options to manage users",
				Subcommands: []*cli.Command{
					{
						Name:   "list",
						Usage:  "list users in database",
						Action: a.UserList,
					},
					{
						Name:   "add",
						Usage:  "add a new user",
						Action: a.UserCreate,
					},
					{
						Name:   "delete",
						Usage:  "delete user in database",
						Action: a.UserDelete,
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:     "id",
								Usage:    "name of the user to add",
								Required: true,
							},
						},
					},
				},
			},
		},
	}

	return app.Run(os.Args)
}
