package service

import (
	"github.com/Higor-Edgar/booktime.git/internal/contract"
	"github.com/Higor-Edgar/booktime.git/internal/domain"
	"github.com/Higor-Edgar/booktime.git/internal/repository"
	"github.com/Higor-Edgar/booktime.git/util"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(newUserDTO *contract.NewUserDTO) error
}

type userService struct {
	userRepo repository.UserRepository
}

func (u *userService) CreateUser(newUserDTO *contract.NewUserDTO) error {
	if err := util.ValidateStruct(newUserDTO); err != nil {
		return err
	}

	user := &domain.User{
		ID:       uuid.New(),
		Name:     newUserDTO.Name,
		Email:    newUserDTO.Email,
		Password: newUserDTO.Password,
	}

	return u.userRepo.Create(user)
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}
