package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	ENV         string      `env:"ENV" envDefault:"dev"`
	PORT        string      `env:"PORT" envDefault:"8080"`
	MySQLConfig MySQLConfig `envPrefix:"MYSQL_"`
}

type MySQLConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"3306"`
	User     string `env:"USER" envDefault:"root"`
	Password string `env:"PASSWORD"`
	Database string `env:"DATABASE"`
}

func NewConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	err = env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil

}
