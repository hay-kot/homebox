// Package cgofreesqlite package provides a CGO free implementation of the sqlite3 driver. This wraps the
// modernc.org/sqlite driver and adds the PRAGMA foreign_keys = ON; statement to the connection
// initialization as well as registering the driver with the sql package as "sqlite3" for compatibility
// with entgo.io
//
// NOTE: This does come with around a 30% performance hit compared to the CGO version of the driver.
// however it greatly simplifies the build process and allows for cross compilation.
package cgofreesqlite

import (
	"database/sql"
	"database/sql/driver"

	"modernc.org/sqlite"
)

type CGOFreeSqliteDriver struct {
	*sqlite.Driver
}

type sqlite3DriverConn interface {
	Exec(string, []driver.Value) (driver.Result, error)
}

func (d CGOFreeSqliteDriver) Open(name string) (conn driver.Conn, err error) {
	conn, err = d.Driver.Open(name)
	if err != nil {
		return nil, err
	}
	_, err = conn.(sqlite3DriverConn).Exec("PRAGMA foreign_keys = ON;", nil)
	if err != nil {
		_ = conn.Close()
		return nil, err
	}
	return conn, err
}

func init() { //nolint:gochecknoinits
	sql.Register("sqlite3", CGOFreeSqliteDriver{Driver: &sqlite.Driver{}})
}
