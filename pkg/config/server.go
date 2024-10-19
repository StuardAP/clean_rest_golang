package config

type ServerConfig struct {
    PublicHost string
    Port       string
}

func LoadServerConfig() ServerConfig {
    return ServerConfig{
        PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
        Port:       getEnv("PORT", "8080"),
    }
}
