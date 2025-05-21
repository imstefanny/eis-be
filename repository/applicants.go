package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type ApplicantsRepository interface {
	Browse(page, limit int, search string) ([]models.Applicants, int64, error)
	Create(applicants models.Applicants) error
	Find(id int) (models.Applicants, error)
	GetByToken(id int) (models.Applicants, error)
	Update(id int, applicant models.Applicants) error
	Delete(id int) error
}

type applicantsRepository struct {
	db *gorm.DB
}

func NewApplicantsRepository(db *gorm.DB) *applicantsRepository {
	return &applicantsRepository{db}
}

func (r *applicantsRepository) Create(applicants models.Applicants) error {
	err := r.db.Create(&applicants)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *applicantsRepository) Browse(page, limit int, search string) ([]models.Applicants, int64, error) {
	var applicants []models.Applicants
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(full_name) LIKE ?", search).Limit(limit).Offset(offset).Preload("Level").Preload("CreatedByName").Preload("UpdatedByName").Find(&applicants).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.Applicants{}).Where("LOWER(full_name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return applicants, total, nil
}

func (r *applicantsRepository) Find(id int) (models.Applicants, error) {
	applicant := models.Applicants{}
	if err := r.db.Preload("Level").Preload("CreatedByName").Preload("UpdatedByName").First(&applicant, id).Error; err != nil {
		return applicant, err
	}
	return applicant, nil
}

func (r *applicantsRepository) GetByToken(id int) (models.Applicants, error) {
	applicant := models.Applicants{}
	if err := r.db.Where("created_by = ?", id).Preload("Level").Preload("CreatedByName").Preload("UpdatedByName").First(&applicant).Error; err != nil {
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
