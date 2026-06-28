package option

import (
	"errors"
	"fmt"
	"time"
)

type dbConfig struct {
	host            string
	port            int
	user            string
	password        string
	dbName          string
	sslMode         string
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifetime time.Duration
}

type DBOption func(*dbConfig)

func WithHost(host string) DBOption {
	return func(c *dbConfig) { c.host = host }
}

func WithPort(port int) DBOption {
	return func(c *dbConfig) { c.port = port }
}

func WithUser(user string) DBOption {
	return func(c *dbConfig) { c.user = user }
}

func WithPassword(password string) DBOption {
	return func(c *dbConfig) { c.password = password }
}

func WithDBName(dbName string) DBOption {
	return func(c *dbConfig) { c.dbName = dbName }
}

func WithSSLMode(mode string) DBOption {
	return func(c *dbConfig) { c.sslMode = mode }
}

func WithMaxOpenConns(n int) DBOption {
	return func(c *dbConfig) { c.maxOpenConns = n }
}

func WithMaxIdleConns(n int) DBOption {
	return func(c *dbConfig) { c.maxIdleConns = n }
}

func WithConnMaxLifetime(d time.Duration) DBOption {
	return func(c *dbConfig) { c.connMaxLifetime = d }
}

type DBConnection struct {
	config dbConfig
}

func NewDBConnection(opts ...DBOption) (*DBConnection, error) {
	cfg := dbConfig{
		host:            "localhost",
		port:            5432,
		sslMode:         "disable",
		maxOpenConns:    10,
		maxIdleConns:    5,
		connMaxLifetime: 30 * time.Minute,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	if cfg.user == "" {
		return nil, errors.New("database user is required")
	}
	if cfg.dbName == "" {
		return nil, errors.New("database name is required")
	}

	return &DBConnection{config: cfg}, nil
}

func (c *DBConnection) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.config.host,
		c.config.port,
		c.config.user,
		c.config.password,
		c.config.dbName,
		c.config.sslMode,
	)
}

func (c *DBConnection) Connect() error {
	return nil
}

func (c *DBConnection) PoolConfig() (maxOpen, maxIdle int, maxLifetime time.Duration) {
	return c.config.maxOpenConns, c.config.maxIdleConns, c.config.connMaxLifetime
}
