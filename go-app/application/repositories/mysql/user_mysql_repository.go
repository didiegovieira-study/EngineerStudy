package mysql

import (
	"database/sql"
	"fmt"

	"github.com/didiegovieira/EngineerStudy/go-app/domain/entities"
	"github.com/didiegovieira/EngineerStudy/go-app/domain/entities/mysql"
)

type UserMySQLRepository struct {
	Db *sql.DB
}

func (repository UserMySQLRepository) CreateUser(user mysql.User) error {
	u := user
	query := "INSERT INTO users (id, username, email) VALUES (?, ?, ?)"
	_, err := repository.Db.Exec(query, u.ID, u.Username, u.Email)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}

func (repository UserMySQLRepository) DeleteUserByID(id string) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := repository.Db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

func (repository UserMySQLRepository) GetByID(id string) (mysql.User, error) {
	query := "SELECT id, username, email FROM users WHERE id = ?"
	row := repository.Db.QueryRow(query, id)

	user := &entities.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}
