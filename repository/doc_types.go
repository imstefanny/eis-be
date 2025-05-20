package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type DocTypesRepository interface {
	Browse(page, limit int, search string) ([]models.DocTypes, int64, error)
	Create(docTypes models.DocTypes) error
	Find(id int) (models.DocTypes, error)
	Update(id int, docType models.DocTypes) error
	Delete(id int) error
}

type docTypesRepository struct {
	db *gorm.DB
}

func NewDocTypesRepository(db *gorm.DB) *docTypesRepository {
	return &docTypesRepository{db}
}

func (r *docTypesRepository) Browse(page, limit int, search string) ([]models.DocTypes, int64, error) {
	var docTypes []models.DocTypes
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(name) LIKE ?", search).Limit(limit).Offset(offset).Find(&docTypes).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.DocTypes{}).Where("LOWER(name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return docTypes, total, nil
}

func (r *docTypesRepository) Create(docTypes models.DocTypes) error {
	err := r.db.Create(&docTypes)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *docTypesRepository) Find(id int) (models.DocTypes, error) {
	docType := models.DocTypes{}
	if err := r.db.First(&docType, id).Error; err != nil {
		return docType, err
	}
	return docType, nil
}

func (r *docTypesRepository) Update(id int, docType models.DocTypes) error {
	query := r.db.Model(&docType).Updates(docType)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *docTypesRepository) Delete(id int) error {
	docType := models.DocTypes{}
	if err := r.db.Delete(&docType, id).Error; err != nil {
		return err
	}
	return nil
}
