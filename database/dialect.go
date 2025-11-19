package database

import "gorm.io/gorm"

type Dialect interface {
	Open(dsn string) (gorm.Dialector, error)
	BuildDSN(config *Config) string 
}

type Config struct {
	Host string 
	User string 
	Password string 
	DBName string 
	Port string 
	SSLMode string 
	TimeZone string 
	DBTtyoe string 
}