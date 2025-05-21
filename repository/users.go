package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type UsersRepository interface {
	Create(tx *gorm.DB, user models.Users) (uint, error)
	// Create(user models.Users) (uint, error)
	Login(data models.Users) (models.Users, error)
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *usersRepository {
	return &usersRepository{db}
}

func (r *usersRepository) Create(tx *gorm.DB, user models.Users) (uint, error) {
	err := tx.Create(&user)
	if err.Error != nil {
		return 0, err.Error
	}
	return user.ID, nil
}

func (r *usersRepository) Login(data models.Users) (models.Users, error) {
	user := models.Users{}
	err := r.db.Where("email = ? AND password = ?", data.Email, data.Password).First(&user).Error
	if err != nil {
		return models.Users{}, err
	}
	return user, nil
}
