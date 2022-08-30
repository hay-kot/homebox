package config

type LoggerConf struct {
	Level string `conf:"default:debug"`
	File  string `conf:""`
}
