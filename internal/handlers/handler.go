package handlers

import (
	"database/sql"

	"github.com/RiddlerXenon/cloud-core/internal/config"
)

type Handler struct {
	DB  *sql.DB
	Cfg *config.Config
}

func New(db *sql.DB, cfg *config.Config) *Handler {
	return &Handler{DB: db, Cfg: cfg}
}
