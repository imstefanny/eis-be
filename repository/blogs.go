package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type BlogsRepository interface {
	Browse(page, limit int, search string) ([]models.Blogs, int64, error)
	Create(blogs models.Blogs) error
	Find(id int) (models.Blogs, error)
	Update(id int, blog models.Blogs) error
	Delete(id int) error
}

type blogsRepository struct {
	db *gorm.DB
}

func NewBlogsRepository(db *gorm.DB) *blogsRepository {
	return &blogsRepository{db}
}

func (r *blogsRepository) Browse(page, limit int, search string) ([]models.Blogs, int64, error) {
	var blogs []models.Blogs
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(title) LIKE ?", search).Limit(limit).Offset(offset).Find(&blogs).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.Blogs{}).Where("LOWER(title) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return blogs, total, nil
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

func (r *blogsRepository) Update(id int, blog models.Blogs) error {
	query := r.db.Model(&blog).Updates(blog)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *blogsRepository) Delete(id int) error {
	blog := models.Blogs{}
	if err := r.db.Delete(&blog, id).Error; err != nil {
		return err
	}
	return nil
}
