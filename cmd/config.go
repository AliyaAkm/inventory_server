package main

import (
	"github.com/caarlos0/env/v11"
)

type DbConfig struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	DBName   string `env:"DBNAME"`
	SSLMode  string `env:"SSLMODE"`
}

type Config struct {
	DbConfig DbConfig
	HTTPPort string `env:"HTTP_PORT"`
}

func ReadEnv() (Config, error) {
	opts := env.Options{
		RequiredIfNoDef: true,
	}

	cfg := new(Config)
	err := env.ParseWithOptions(cfg, opts)
	if err != nil {
		return Config{}, err
	}

	return *cfg, nil
}
