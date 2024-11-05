package db

import (
	"database/sql"
	"fmt"
	cnf "training-pj/src/config"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}


func LoadDBConfig() *Config {
	cfg := &Config{
		Host:     cnf.LoadEnv("DB_HOST", "db"),
		Port:     cnf.LoadEnv("DB_PORT", "5432"),
		User:     cnf.LoadEnv("DB_USER", "user"),
		Password: cnf.LoadEnv("DB_PASSWORD", "password"),
		DBName:   cnf.LoadEnv("DB_NAME", "mydb"),
	}	
	return cfg
}

func NewSQLDB(driver string) (*sql.DB, error) {
	cfg := LoadDBConfig()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	db, err := sql.Open(driver, connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}