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
	Mode    string      `yaml:"mode" conf:"default:development"` // development or production
	Web     WebConfig   `yaml:"web"`
	Storage Storage     `yaml:"storage"`
	Log     LoggerConf  `yaml:"logger"`
	Mailer  MailerConf  `yaml:"mailer"`
	Swagger SwaggerConf `yaml:"swagger"`
	Demo    bool        `yaml:"demo"`
	Debug   DebugConf   `yaml:"debug"`
	Options Options     `yaml:"options"`
}

type Options struct {
	AllowRegistration    bool `yaml:"disable_registration" conf:"default:true"`
	AutoIncrementAssetID bool `yaml:"auto_increment_asset_id" conf:"default:true"`
}

type DebugConf struct {
	Enabled bool   `yaml:"enabled" conf:"default:false"`
	Port    string `yaml:"port" conf:"default:4000"`
}

type SwaggerConf struct {
	Host   string `yaml:"host" conf:"default:localhost:7745"`
	Scheme string `yaml:"scheme" conf:"default:http"`
}

type WebConfig struct {
	Port          string `yaml:"port" conf:"default:7745"`
	Host          string `yaml:"host"`
	MaxUploadSize int64  `yaml:"max_file_upload" conf:"default:10"`
}

// New parses the CLI/Config file and returns a Config struct. If the file argument is an empty string, the
// file is not read. If the file is not empty, the file is read and the Config struct is returned.
func New() (*Config, error) {
	var cfg Config
	const prefix = "HBOX"

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
