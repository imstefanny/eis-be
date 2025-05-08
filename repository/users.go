package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type UsersRepository interface {
	Create(user models.Users) error
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
