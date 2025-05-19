package repository

import (
	"github.com/tetzng/golang-blog/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	GetUserById(user *model.User, id uint) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserById(user *model.User, id uint) error {
	if err := ur.db.First(user, id).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
