package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/ardanlabs/conf/v2"
	"github.com/ardanlabs/conf/v2/yaml"

	"os"
)

const (
	ModeDevelopment = "development"
	ModeProduction  = "production"
)

type Config struct {
	Mode     string      `yaml:"mode" conf:"default:development"` // development or production
	Web      WebConfig   `yaml:"web"`
	Database Database    `yaml:"database"`
	Log      LoggerConf  `yaml:"logger"`
	Mailer   MailerConf  `yaml:"mailer"`
	Seed     Seed        `yaml:"seed"`
	Swagger  SwaggerConf `yaml:"swagger"`
}

type SwaggerConf struct {
	Host   string `yaml:"host" conf:"default:localhost:7745"`
	Scheme string `yaml:"scheme" conf:"default:http"`
}

type WebConfig struct {
	Port string `yaml:"port" conf:"default:7745"`
	Host string `yaml:"host" conf:"default:127.0.0.1"`
}

// NewConfig parses the CLI/Config file and returns a Config struct. If the file argument is an empty string, the
// file is not read. If the file is not empty, the file is read and the Config struct is returned.
func NewConfig(file string) (*Config, error) {
	var cfg Config

	const prefix = "API"

	help, err := func() (string, error) {
		if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
			return conf.Parse(prefix, &cfg)
		} else {
			yamlData, err := ioutil.ReadFile(file)
			if err != nil {
				return "", err
			}
			return conf.Parse(prefix, &cfg, yaml.WithData(yamlData))
		}
	}()

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
