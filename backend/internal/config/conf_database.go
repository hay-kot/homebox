package config

const (
	DriverSqlite3 = "sqlite3"
)

type Database struct {
	Driver    string `yaml:"driver" conf:"default:sqlite3"`
	SqliteUrl string `yaml:"sqlite-url" conf:"default:file:ent?mode=memory&cache=shared&_fk=1"`
}

func (d *Database) GetDriver() string {
	return d.Driver
}

func (d *Database) GetUrl() string {
	switch d.Driver {
	case DriverSqlite3:
		return d.SqliteUrl
	default:
		panic("unknown database driver")
	}
}
