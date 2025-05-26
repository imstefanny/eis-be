package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type UsersRepository interface {
	Create(tx *gorm.DB, user models.Users) (uint, error)
	// Create(user models.Users) (uint, error)
	Find(id int) (models.Users, error)
	Login(data models.Users) (models.Users, error)
	Browse(page, limit int, search string) ([]models.Users, int64, error)
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

func (r *usersRepository) Find(id int) (models.Users, error) {
	user := models.Users{}
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *usersRepository) Browse(page, limit int, search string) ([]models.Users, int64, error) {
	var users []models.Users
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(name) LIKE ?", search).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.Users{}).Where("LOWER(name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}
