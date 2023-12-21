// Package config provides the configuration for the application.
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/ardanlabs/conf/v3"
)

const (
	ModeDevelopment = "development"
	ModeProduction  = "production"
)

type Config struct {
	conf.Version
	Mode    string     `yaml:"mode"    conf:"default:development"` // development or production
	Web     WebConfig  `yaml:"web"`
	Storage Storage    `yaml:"storage"`
	Log     LoggerConf `yaml:"logger"`
	Mailer  MailerConf `yaml:"mailer"`
	Demo    bool       `yaml:"demo"`
	Debug   DebugConf  `yaml:"debug"`
	Options Options    `yaml:"options"`
}

type Options struct {
	AllowRegistration    bool `yaml:"disable_registration"    conf:"default:true"`
	AutoIncrementAssetID bool `yaml:"auto_increment_asset_id" conf:"default:true"`
}

type DebugConf struct {
	Enabled bool   `yaml:"enabled" conf:"default:false"`
	Port    string `yaml:"port"    conf:"default:4000"`
}

type WebConfig struct {
	Port          string `yaml:"port"            conf:"default:7745"`
	Host          string `yaml:"host"`
	MaxUploadSize int64  `yaml:"max_file_upload" conf:"default:10"`
	ReadTimeout   int    `yaml:"read_timeout"    conf:"default:10"`
	WriteTimeout  int    `yaml:"write_timeout"   conf:"default:10"`
	IdleTimeout   int    `yaml:"idle_timeout"    conf:"default:30"`
}

// New parses the CLI/Config file and returns a Config struct. If the file argument is an empty string, the
// file is not read. If the file is not empty, the file is read and the Config struct is returned.
func New(buildstr string, description string) (*Config, error) {
	var cfg Config
	const prefix = "HBOX"

	cfg.Version = conf.Version{
		Build: buildstr,
		Desc:  description,
	}

	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			os.Exit(0)
		}
		return &cfg, fmt.Errorf("parsing config: %w", err)
	}

	return &cfg, nil
}

// Print prints the configuration to stdout as a json indented string
// This is useful for debugging. If the marshaller errors out, it will panic.
func (c *Config) Print() {
	res, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}
