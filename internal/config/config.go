package config

import "os"

type Config struct {
    ServerAddress string
    DatabaseURL   string
}

func Load() Config {
    return Config{
        ServerAddress: os.Getenv("SERVER_ADDRESS"),
        DatabaseURL:   os.Getenv("DATABASE_URL"),
    }
}
