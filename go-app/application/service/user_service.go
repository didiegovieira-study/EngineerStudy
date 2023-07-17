package service

import (
	"github.com/didiegovieira/EngineerStudy/go-app/application/repositories"
	"github.com/didiegovieira/EngineerStudy/go-app/domain/entities"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUserByID(id string) (*entities.User, error) {
	return s.userRepository.GetUserByID(id)
}

func (s *UserService) CreateUser(user *entities.User) error {
	return s.userRepository.CreateUser(&entities.User{})
}

func (s *UserService) DeleteUserByID(id string) *entities.User {
	return s.userRepository.DeleteUserByID(id)
}
