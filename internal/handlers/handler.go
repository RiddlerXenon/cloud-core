package handlers

import (
	"github.com/RiddlerXenon/cloud-core/internal/config"
	"github.com/RiddlerXenon/cloud-core/internal/repository"
)

type Handler struct {
	DB  *repository.Database
	Cfg *config.Config
}

func New(d *repository.Database, cfg *config.Config) *Handler {
	return &Handler{DB: d, Cfg: cfg}
}
