package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDialect struct {}

func (p *PostgresDialect) Open(dsn string) (gorm.Dialector, error) {
	return postgres.Open(dsn), nil
}

func (p *PostgresDialect) BuildDSN(config *Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode, config.TimeZone,
	)
}