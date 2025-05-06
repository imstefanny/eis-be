package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type BlogsRepository interface {
	GetAll() ([]models.Blogs, error)
}

type blogsRepository struct {
	db *gorm.DB
}

func NewBlogsRepository(db *gorm.DB) *blogsRepository {
	return &blogsRepository{db}
}

func (r *blogsRepository) GetAll() ([]models.Blogs, error) {
	blogs := []models.Blogs{}
	if err := r.db.Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}
