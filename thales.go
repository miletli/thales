package thales

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	db *sql.DB
	driverName, dataSourceName string
}

func (confing *Config) Open(driverName string, dataSourceName string) error {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	confing.db = db
	confing.driverName = driverName
	confing.dataSourceName = dataSourceName
	return nil
}

func (confing *Config) Set(sql string, args ...interface{}) (sql.Result, error) {
	if confing.db == nil {
		return nil, errors.New("database connection is not open")
	}

	result, err := confing.db.Exec(sql, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (confing *Config) Close() error {
	if confing.db == nil {
		return nil
	}

	err := confing.db.Close()
	if err != nil {
		return err
	}

	confing.db = nil
	return nil
}
