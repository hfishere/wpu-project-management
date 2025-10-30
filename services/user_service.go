package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hfishere/wpu-project-management/models"
	"github.com/hfishere/wpu-project-management/repositories"
	"github.com/hfishere/wpu-project-management/utils"
)

type UserService interface {
	Register(user *models.User) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(user *models.User) error {
	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser.InternalID != 0 {
		return errors.New("email already registered")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.Role = "user"
	user.PublicID = uuid.New()

	return s.repo.Create(user)
}
