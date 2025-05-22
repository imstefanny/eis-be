package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type GuardiansRepository interface {
	Browse(page, limit int, search string) ([]models.Guardians, int64, error)
	Create(guardians models.Guardians) error
	Find(id int) (models.Guardians, error)
	Update(id int, guardian models.Guardians) error
	FindByApplicantId(id int) ([]models.Guardians, error)
	Delete(id int) error
}

type guardiansRepository struct {
	db *gorm.DB
}

func NewGuardiansRepository(db *gorm.DB) *guardiansRepository {
	return &guardiansRepository{db}
}

func (r *guardiansRepository) Browse(page, limit int, search string) ([]models.Guardians, int64, error) {
	var guardians []models.Guardians
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(name) LIKE ?", search).Limit(limit).Offset(offset).Preload("Applicant").Preload("Student").Find(&guardians).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.Guardians{}).Where("LOWER(name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return guardians, total, nil
}

func (r *guardiansRepository) Create(guardians models.Guardians) error {
	err := r.db.Create(&guardians)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *guardiansRepository) Find(id int) (models.Guardians, error) {
	guardian := models.Guardians{}
	if err := r.db.Preload("Applicant").Preload("Student").First(&guardian, id).Error; err != nil {
		return guardian, err
	}
	return guardian, nil
}

func (r *guardiansRepository) FindByApplicantId(id int) ([]models.Guardians, error) {
	guardian := []models.Guardians{}
	if err := r.db.Where("applicant_id = ?", id).Preload("Applicant").Preload("Student").Find(&guardian).Error; err != nil {
		return guardian, err
	}
	return guardian, nil
}

func (r *guardiansRepository) Update(id int, guardian models.Guardians) error {
	query := r.db.Model(&guardian).Updates(guardian)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *guardiansRepository) Delete(id int) error {
	guardian := models.Guardians{}
	if err := r.db.Delete(&guardian, id).Error; err != nil {
		return err
	}
	return nil
}
