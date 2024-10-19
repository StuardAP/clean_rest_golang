package config

import (
    "fmt"
)

type DatabaseConfig struct {
    DBUser    string
    DBPassword string
    DBAddress  string
    DBName    string
}

func LoadDatabaseConfig() DatabaseConfig {
    return DatabaseConfig{
        DBUser:    getEnv("DB_USER", ""),
        DBPassword: getEnv("DB_PASSWORD", ""),
        DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
        DBName:    getEnv("DB_NAME", ""),
    }
}

