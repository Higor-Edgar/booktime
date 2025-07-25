package repository

import (
	"github.com/Higor-Edgar/booktime.git/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
}

type userRepository struct {
	database *gorm.DB
}

func (r *userRepository) Create(user *domain.User) error {
	return r.database.Create(user).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		database: db,
	}
}
