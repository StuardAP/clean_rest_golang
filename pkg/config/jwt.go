package config


type JWTConfig struct {
    JWTExpirationTime int64
    JWTSecret         string
}

func LoadJWTConfig() JWTConfig {
    return JWTConfig{
        JWTExpirationTime: getEnvAsInt("JWT_EXP", 604800), //! ONE WEEK
        JWTSecret:         getEnv("JWT_SECRET", ""),
    }
}
