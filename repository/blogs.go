package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type BlogsRepository interface {
	GetAll() ([]models.Blogs, error)
	Create(blogs models.Blogs) error
	Find(id int) (models.Blogs, error)
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

func (r *blogsRepository) Create(blogs models.Blogs) error {
	err := r.db.Create(&blogs)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *blogsRepository) Find(id int) (models.Blogs, error) {
	blog := models.Blogs{}
	if err := r.db.First(&blog, id).Error; err != nil {
		return blog, err
	}
	return blog, nil
}
