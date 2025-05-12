package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type ApplicantsRepository interface {
	GetAll() ([]models.Applicants, error)
	Create(applicants models.Applicants) error
	Find(id int) (models.Applicants, error)
	Update(id int, blog models.Applicants) error
	Delete(id int) error
}

type applicantsRepository struct {
	db *gorm.DB
}

func NewApplicantsRepository(db *gorm.DB) *applicantsRepository {
	return &applicantsRepository{db}
}

func (r *applicantsRepository) GetAll() ([]models.Applicants, error) {
	applicants := []models.Applicants{}
	if err := r.db.Find(&applicants).Error; err != nil {
		return nil, err
	}
	return applicants, nil
}

func (r *applicantsRepository) Create(applicants models.Applicants) error {
	err := r.db.Create(&applicants)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *applicantsRepository) Find(id int) (models.Applicants, error) {
	blog := models.Applicants{}
	if err := r.db.First(&blog, id).Error; err != nil {
		return blog, err
	}
	return blog, nil
}

func (r *applicantsRepository) Update(id int, blog models.Applicants) error {
	query := r.db.Model(&blog).Updates(blog)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *applicantsRepository) Delete(id int) error {
	blog := models.Applicants{}
	if err := r.db.Delete(&blog, id).Error; err != nil {
		return err
	}
	return nil
}
