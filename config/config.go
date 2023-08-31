package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Host       `yaml:"host"`
		Port       `yaml:"port"`
		PG         `yaml:"postgres"`
		Migrations `yaml:"migrationpath"`
	}

	Host struct {
		Host string `env-required:"true"    env:"HTTP_HOST"`
	}

	Port struct {
		Port string `env-required:"true" yaml:"port"    env:"HTTP_PORT"`
	}

	PG struct {
		URL string `env-required:"true"                      env:"PG_URL"`
	}

	Migrations struct {
		MigrationsPath string `env-required:"true"    env:"DB_MIGRATION_PATH"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	//err := cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error updating env: %w", err)
	}

	return cfg, nil
}
