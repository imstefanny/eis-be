package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type UsersRepository interface {
	Create(user models.Users) error
	Login(data models.Users) (models.Users, error)
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *usersRepository {
	return &usersRepository{db}
}

func (r *usersRepository) Create(user models.Users) error {
	err := r.db.Create(&user)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *usersRepository) Login(data models.Users) (models.Users, error) {
	user := models.Users{}
	err := r.db.Where("email = ? AND password = ?", data.Email, data.Password).First(&user).Error
	if err != nil {
		return models.Users{}, err
	}
	return user, nil
}
