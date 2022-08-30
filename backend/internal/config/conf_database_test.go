package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DatabaseConfig_Sqlite(t *testing.T) {
	dbConf := &Database{
		Driver:    DriverSqlite3,
		SqliteUrl: "file:ent?mode=memory&cache=shared&_fk=1",
	}

	assert.Equal(t, "sqlite3", dbConf.GetDriver())
	assert.Equal(t, "file:ent?mode=memory&cache=shared&_fk=1", dbConf.GetUrl())
}

func Test_DatabaseConfig_Postgres(t *testing.T) {
	dbConf := &Database{
		Driver:      DriverPostgres,
		PostgresUrl: "postgres://user:pass@host:port/dbname?sslmode=disable",
	}

	assert.Equal(t, "postgres", dbConf.GetDriver())
	assert.Equal(t, "postgres://user:pass@host:port/dbname?sslmode=disable", dbConf.GetUrl())
}

func Test_DatabaseConfig_Unknown(t *testing.T) {
	dbConf := &Database{
		Driver: "null",
	}

	assert.Panics(t, func() { dbConf.GetUrl() })

}
