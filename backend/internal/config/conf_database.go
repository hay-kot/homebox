package config

const (
	DriverSqlite3 = "sqlite3"
)

type Storage struct {
	// Data is the path to the root directory
	Data      string `yaml:"data" conf:"default:./homebox-data"`
	SqliteUrl string `yaml:"sqlite-url" conf:"default:./homebox-data/homebox.db?_fk=1"`
}
