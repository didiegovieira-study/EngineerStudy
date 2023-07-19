package mysql

import (
	"database/sql"
	"fmt"

	"github.com/didiegovieira/EngineerStudy/go-app/domain/entities"
	"github.com/didiegovieira/EngineerStudy/go-app/framework/database"
)

func NewMySQLRepository() *UserMySQLRepository {
	db := database.NewMysqlDb()
	return (*UserMySQLRepository)(db)
}

type UserMySQLRepository struct {
	Db *sql.DB
}

func (repository UserMySQLRepository) CreateUser(user *entities.User) error {
	u := user
	query := "INSERT INTO users (id, name, gender, age) VALUES (?, ?, ?, ?)"
	_, err := repository.Db.Exec(query, u.Id, u.Name, u.Gender, u.Age)
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

func (repository UserMySQLRepository) GetUserByID(id string) (*entities.User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := repository.Db.QueryRow(query, id)

	user := &entities.User{}

	err := row.Scan(&user.Id, &user.Name, &user.Gender, &user.Age)
	if err != nil {
		return nil, err
	}

	return user, nil
}
