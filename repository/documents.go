package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type DocumentsRepository interface {
	Browse(page, limit int, search string) ([]models.Documents, int64, error)
	Create(documents models.Documents) error
	Find(id int) (models.Documents, error)
	FindByApplicantId(id int) ([]models.Documents, error)
	Update(id int, document models.Documents) error
	Delete(id int) error
}

type documentsRepository struct {
	db *gorm.DB
}

func NewDocumentsRepository(db *gorm.DB) *documentsRepository {
	return &documentsRepository{db}
}

func (r *documentsRepository) Browse(page, limit int, search string) ([]models.Documents, int64, error) {
	var documents []models.Documents
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(type_id) LIKE ?", search).Limit(limit).Offset(offset).Find(&documents).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.Documents{}).Where("LOWER(type_id) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return documents, total, nil
}

func (r *documentsRepository) Create(documents models.Documents) error {
	err := r.db.Create(&documents)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *documentsRepository) Find(id int) (models.Documents, error) {
	document := models.Documents{}
	if err := r.db.First(&document, id).Error; err != nil {
		return document, err
	}
	return document, nil
}

func (r *documentsRepository) FindByApplicantId(id int) ([]models.Documents, error) {
	documents := []models.Documents{}
	if err := r.db.Where("applicant_id = ?", id).Find(&documents).Error; err != nil {
		return nil, err
	}
	return documents, nil
}

func (r *documentsRepository) Update(id int, document models.Documents) error {
	query := r.db.Model(&document).Updates(document)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *documentsRepository) Delete(id int) error {
	document := models.Documents{}
	if err := r.db.Delete(&document, id).Error; err != nil {
		return err
	}
	return nil
}
