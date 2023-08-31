package database

import (
	"avitotest/config"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Connect(cfg *config.Config) (*sqlx.DB, error) {
	fmt.Println(cfg.PG.URL)
	base, err := sqlx.Connect("postgres", cfg.PG.URL)
	if err != nil {
		return nil, err
	}
	return base, nil
}
