package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type WorkSchedsRepository interface {
	GetAll() ([]models.WorkScheds, error)
	Create(workScheds models.WorkScheds) error
	Find(id int) (models.WorkScheds, error)
	Update(id int, blog models.WorkScheds) error
	Delete(id int) error
}

type workSchedsRepository struct {
	db *gorm.DB
}

func NewWorkSchedsRepository(db *gorm.DB) *workSchedsRepository {
	return &workSchedsRepository{db}
}

func (r *workSchedsRepository) GetAll() ([]models.WorkScheds, error) {
	workScheds := []models.WorkScheds{}
	if err := r.db.Preload("Details").Find(&workScheds).Error; err != nil {
		return nil, err
	}
	return workScheds, nil
}

func (r *workSchedsRepository) Create(workScheds models.WorkScheds) error {
	err := r.db.Create(&workScheds)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *workSchedsRepository) Find(id int) (models.WorkScheds, error) {
	blog := models.WorkScheds{}
	if err := r.db.Preload("Details").First(&blog, id).Error; err != nil {
		return blog, err
	}
	return blog, nil
}

func (r *workSchedsRepository) Update(id int, blog models.WorkScheds) error {
	query := r.db.Model(&blog).Updates(blog)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *workSchedsRepository) Delete(id int) error {
	blog := models.WorkScheds{}
	if err := r.db.Delete(&blog, id).Error; err != nil {
		return err
	}
	return nil
}
