package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


type MySQLDialect struct{}

func (m *MySQLDialect) Open(dsn string) (gorm.Dialector, error) {
	return mysql.Open(dsn), nil
}

func (m *MySQLDialect) BuildDSN(config *Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		config.User, config.Password, config.Host, config.Port, config.DBName, config.TimeZone,
	)
}