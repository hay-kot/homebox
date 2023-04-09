package main

import (
	"time"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/internal/sys/config"
	"github.com/hay-kot/homebox/backend/pkgs/mailer"
	"github.com/hay-kot/httpkit/server"
)

type app struct {
	conf     *config.Config
	mailer   mailer.Mailer
	db       *ent.Client
	server   *server.Server
	repos    *repo.AllRepos
	services *services.AllServices
}

func new(conf *config.Config) *app {
	s := &app{
		conf: conf,
	}

	s.mailer = mailer.Mailer{
		Host:     s.conf.Mailer.Host,
		Port:     s.conf.Mailer.Port,
		Username: s.conf.Mailer.Username,
		Password: s.conf.Mailer.Password,
		From:     s.conf.Mailer.From,
	}

	return s
}

func (a *app) startBgTask(t time.Duration, fn func()) {
	timer := time.NewTimer(t)

	for {
		timer.Reset(t)
		a.server.Background(fn)
		<-timer.C
	}
}
