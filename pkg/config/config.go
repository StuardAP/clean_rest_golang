package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
    ServerConfig  ServerConfig
    DatabaseConfig DatabaseConfig
    JWTConfig     JWTConfig
}

var Envs *Config

func init() {
    var err error
    Envs, err = LoadConfig()
    if err != nil {
        log.Fatalf("ERROR INITIALIZING CONFIGURATION: %v", err)
    }
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        log.Printf("WARNING: ERROR LOADING .ENV FILE: %v", err)
    }

    return &Config{
        ServerConfig:  LoadServerConfig(),
        DatabaseConfig: LoadDatabaseConfig(),
        JWTConfig:     LoadJWTConfig(),
    }, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
    if value, ok := os.LookupEnv(key); ok {
        i, err := strconv.ParseInt(value, 10, 64)
        if err != nil {
            return fallback
        }
        return i
    }
    return fallback
}
