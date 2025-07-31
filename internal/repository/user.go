package repository

import (
	"time"
)

type StatusType string

const (
	StatusPending  StatusType = "pending"
	StatusApproved StatusType = "approved"
	StatusRejected StatusType = "rejected"
)

type User struct {
	Id             int
	Name           string
	Email          string
	HashedPassword string
	Status         StatusType
	CrearedAt      time.Time
}

func (d *Database) GetHashedPass(username string) (string, error) {
	row, err := d.DB.QueryRow(
		`
			SELECT hansed_password
			FROM users
			WHERE name = $1
		`, u,
	)
	if err != nil {
		return nil, fmt.Errorf("Could not get user data: %e", err)
	}
	if row == nil {
		return nil, fmt.Error("User not exists!")
	}

	var hash string
	err = row.Scan(&hash)
	if err != nil {
		return nil, fmt.Errorf("Could not get hashed password: %e", err)
	}

	return hash, nil
}

func (d *Database) UserExists(username string) (bool, error) {
	exists := false
	err := d.DB.QueryRow(
		`
			SELECT EXISTS (
				SELECT 1 
				FROM users 
				WHERE name = $1
			)
		`, username,
	).Scan(&exists)

	return exists, err
}

func (d *Database) AddUser(user *User) error {
	_, err := d.DB.Exec(
		`
			INSERT INTO users (name, email, hased_password, status, created_at)
			VALUES ($1, $2, $3, $4, $5);
		`, user.Name, user.Email, user.HashedPassword, user.Status, user.CreatedAt,
	)
	return err
}

func (d *Database) GetUser(id int) (*User, error) {
	row, err := d.DB.QueryRow(
		`
			SELECT id, name, email, hashed_password, status, created_at 
			FROM users 
			WHERE id = $1
		`, id,
	)

	var user User
	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.HashedPassword,
		&user.Status,
		&user.CreatedAt,
	)

	return user, err
}
