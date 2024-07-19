package config

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"time"
)

type PSQL struct {
	DSN                string
	MaxConnections     string
	MaxIdleConnections string
	ConnMaxLifetime    string
}

func (d PSQL) Connect() (*sql.DB, error) {
	if d.DSN == "" {
		return nil, errors.New("dsn is empty")
	}

	maxConn, err := strconv.Atoi(d.MaxConnections)
	if err != nil {
		return nil, fmt.Errorf("invalid MaxConnections: %s", err.Error())
	}

	idleConn, err := strconv.Atoi(d.MaxIdleConnections)
	if err != nil {
		return nil, fmt.Errorf("invalid MaxIdleConnections: %s", err.Error())
	}

	connMaxLifetime, err := time.ParseDuration(d.ConnMaxLifetime)
	if err != nil {
		return nil, fmt.Errorf("invalid ConnMaxLifetime format: %s", err.Error())
	}

	db, err := sql.Open("postgres", d.DSN)
	if err != nil {
		return nil, err
	}

	//Set configuration pool settings
	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(idleConn)
	db.SetConnMaxLifetime(connMaxLifetime)

	//Verify connections
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
