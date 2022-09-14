package config

const (
	LogFormatJSON = "json"
	LogFormatText = "text"
)

type LoggerConf struct {
	Level  string `conf:"default:info"`
	Format string `conf:"default:text"`
}
