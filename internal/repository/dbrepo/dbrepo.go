package dbrepo

import (
	"booking/internal/config"
	"booking/internal/repository"
	"database/sql"
)

type potgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &potgresDBRepo{
		App: a,
		DB:  conn,
	}
}
