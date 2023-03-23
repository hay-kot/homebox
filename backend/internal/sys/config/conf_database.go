package config

const (
	DriverSqlite3 = "sqlite3"
)

type Storage struct {
	// Data is the path to the root directory
	Data      string `yaml:"data" conf:"default:./.data"`
	SqliteUrl string `yaml:"sqlite-url" conf:"default:./.data/homebox.db?_pragma=busy_timeout=1000&_pragma=journal_mode=WAL&_fk=1"`
}
