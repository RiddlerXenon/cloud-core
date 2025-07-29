package repository

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

func InitDB() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=CloudStorage password=12345 dbname=cloud-core sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("База данных не подключена!")
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("База данных не отвечает!")
	}
	zap.S().Info("База данных подключена успешно.")

	return db, nil
}

type User struct {
	Id             int
	Name           string
	Email          string
	HashedPassword string
	CrearedAt      string
}

func getUser(id int) User {
}
