package repositories

import "github.com/didiegovieira/EngineerStudy/go-app/domain/entities"

type UserRepository interface {
	GetUserByID(id string) (*entities.User, error)
	CreateUser(user *entities.User) error
	DeleteUserByID(id string) error
}
