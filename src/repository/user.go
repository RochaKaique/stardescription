package repository

import (
	"bff/src/models/in"
	"database/sql"
)

type Repo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Repo {
	return &Repo{db}
}

func (repo Repo) Create(user in.User) (string, error) {
	statement, err := repo.db.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return "", err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return "", err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return "", err
	}

	return "Usuario Inserido", nil
}

func (repo Repo) FindByEmailForLogin(email string) (in.User, error) {
	line, error := repo.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if error != nil {
		return in.User{}, error
	}
	defer line.Close()

	var user in.User
	if line.Next() {
		if error = line.Scan(
			&user.ID,
			&user.Password,
		); error != nil {
			return in.User{}, error
		}
	}

	return user, nil
}
