package config

import (
	"errors"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	App    AppConfig
	Server ServerConfig
	Pow    PowConfig
}

type AppConfig struct {
	Name        string `env:"APP_NAME"`
	Environment string `env:"APP_ENVIRONMENT"`
}

type ServerConfig struct {
	Port int `env:"SERVER_PORT"`
}

type PowConfig struct {
	ChallengeDifficulty int `env:"POW_CHALLENGE_DIFFICULTY"`
}

func Init(path string) (Config, error) {
	err := godotenv.Load(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return Config{}, err
	}

	cfg := Config{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
