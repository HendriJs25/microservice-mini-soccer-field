package config

import (
	"fmt"
	"strings"
)

const (
	defaultDatabaseHost             = "localhost"
	defaultDatabasePort             = "5432"
	defaultDatabaseSSLMode          = "disable"
	defaultDatabaseMaxOpenConns     = 10
	defaultDatabaseMaxLifetimeconns = 10
	defaultDatabaseIdleConns        = 10
	defaultDatabaseIdleTime         = 10
)

type Database struct {
	Host             string
	Port             string
	Name             string
	Username         string
	Password         string
	SSLMode          string
	MaxOpenConns     int
	MaxLifetimeConns int
	IdleConns        int
	IdleTime         int
}

func loadDatabase() (Database, error) {
	maxOpenConns, err := getEnvInt("DATABASE_MAX_OPEN_CONNECTION", defaultDatabaseMaxOpenConns)
	if err != nil {
		return Database{}, err
	}

	maxLifetimeConns, err := getEnvInt("DATABASE_MAX_LIFETIME_CONNECTION", defaultDatabaseMaxLifetimeconns)
	if err != nil {
		return Database{}, err
	}

	idleConns, err := getEnvInt("DATABASE_IDLE_CONNECTION", defaultDatabaseIdleConns)
	if err != nil {
		return Database{}, err
	}

	idleTime, err := getEnvInt("DATABASE_IDLE_TIME", defaultDatabaseIdleTime)
	if err != nil {
		return Database{}, err
	}

	return Database{
		Host:             getEnv("DATABASE_HOST", defaultDatabaseHost),
		Port:             getEnv("DATABASE_PORT", defaultDatabasePort),
		Name:             getEnv("DATABASE_NAME", ""),
		Username:         getEnv("DATABASE_USERNAME", ""),
		Password:         getEnv("DATABASE_PASSWORD", ""),
		SSLMode:          getEnv("DATABASE_SSL_MODE", defaultDatabaseSSLMode),
		MaxOpenConns:     maxOpenConns,
		MaxLifetimeConns: maxLifetimeConns,
		IdleConns:        idleConns,
		IdleTime:         idleTime,
	}, nil
}

func (d *Database) Validate() error {
	if strings.TrimSpace(d.Host) == "" {
		return fmt.Errorf("database host is required")
	}

	if err := validatePort("DATABASE_PORT", d.Port); err != nil {
		return err
	}

	if strings.TrimSpace(d.Name) == "" {
		return fmt.Errorf("database name is required")
	}

	if strings.TrimSpace(d.Username) == "" {
		return fmt.Errorf("database username is required")
	}

	if strings.TrimSpace(d.Password) == "" {
		return fmt.Errorf("database password is required")
	}

	if strings.TrimSpace(d.SSLMode) == "" {
		return fmt.Errorf("database ssl_mode is required")
	}

	if d.MaxOpenConns <= 0 {
		return fmt.Errorf("DATABASE_MAX_OPEN_CONNECTIONS must be greater than zero")
	}

	if d.MaxLifetimeConns <= 0 {
		return fmt.Errorf("DATABASE_MAX_LIFETIME_CONNECTIONS must be greater than zero")
	}

	if d.IdleConns <= 0 {
		return fmt.Errorf("DATABASE_IDLE_CONNECTIONS must be greater than zero")
	}

	if d.IdleTime <= 0 {
		return fmt.Errorf("DATABASE_IDLE_TIME must be greater than zero")
	}

	return nil
}
