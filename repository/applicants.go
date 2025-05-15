package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type ApplicantsRepository interface {
	GetAll() ([]models.Applicants, error)
	Create(applicants models.Applicants) error
	Find(id int) (models.Applicants, error)
	FindByCreatedBy(id int) (models.Applicants, error)
	Update(id int, applicant models.Applicants) error
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
	applicant := models.Applicants{}
	if err := r.db.First(&applicant, id).Error; err != nil {
		return applicant, err
	}
	return applicant, nil
}

func (r *applicantsRepository) FindByCreatedBy(id int) (models.Applicants, error) {
	applicant := models.Applicants{}
	if err := r.db.Where("created_by = ?", id).First(&applicant).Error; err != nil {
		return applicant, err
	}
	return applicant, nil
}

func (r *applicantsRepository) Update(id int, applicant models.Applicants) error {
	query := r.db.Model(&applicant).Updates(applicant)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *applicantsRepository) Delete(id int) error {
	applicant := models.Applicants{}
	if err := r.db.Delete(&applicant, id).Error; err != nil {
		return err
	}
	return nil
}
