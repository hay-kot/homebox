package config

const (
	DriverSqlite3 = "sqlite3"
)

type Storage struct {
	// Data is the path to the root directory
	Data      string `yaml:"data" conf:"default:./.data"`
	SqliteUrl string `yaml:"sqlite-url" conf:"default:./.data/homebox.db?_fk=1"`
}
