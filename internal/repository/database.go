package repository

import (
	"database/sql"
	"fmt"

	"github.com/RiddlerXenon/cloud-core/internal/config"
	"go.uber.org/zap"
)

type Database struct {
	DB  *sql.DB
	Cfg *config.Config
}

func InitDB(cfg *config.Config) (*Database, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("База данных не подключена!")
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("База данных не отвечает!")
	}
	zap.S().Info("База данных подключена успешно.")

	return &Database{DB: db, Cfg: cfg}, nil
}
