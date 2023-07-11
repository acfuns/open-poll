package utils

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type config struct {
	HOST     string `env:"DATABASE_HOST"`
	PORT     int    `env:"DATABASE_PORT"`
	USER     string `env:"DATABASE_USER"`
	PASSWORD int    `env:"DATABASE_PASSWORD"`
	DB_NAME  string `env:"DB_NAME"`
}

func Cfg() (config, error) {
	err := godotenv.Load()
	if err != nil {
		return config{}, err
	}

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return config{}, err
	}

	return cfg, nil
}
