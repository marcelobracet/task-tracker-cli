package database

import "task-tracker/database"

func GetDialect(dbType string) (Dialect, error) {
	switch(dbType){
	case "mysql":
			return database.MySQLDialect{}, nil 
	case "postgres":
			return database.PostgresDialect{}, nil
	}
}