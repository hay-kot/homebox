package main

import (
	"os"
	"strings"

	"github.com/hay-kot/content/backend/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// setupLogger initializes the zerolog config
// for the shared logger.
func (a *app) setupLogger() {
	// Logger Init
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if a.conf.Mode != config.ModeProduction {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	log.Level(getLevel(a.conf.Log.Level))
}

func getLevel(l string) zerolog.Level {
	switch strings.ToLower(l) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}
