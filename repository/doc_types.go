package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type DocTypesRepository interface {
	GetAll() ([]models.DocTypes, error)
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

func (r *docTypesRepository) GetAll() ([]models.DocTypes, error) {
	docTypes := []models.DocTypes{}
	if err := r.db.Find(&docTypes).Error; err != nil {
		return nil, err
	}
	return docTypes, nil
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
