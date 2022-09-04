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

func Test_DatabaseConfig_Unknown(t *testing.T) {
	dbConf := &Database{
		Driver: "null",
	}

	assert.Panics(t, func() { dbConf.GetUrl() })

}
