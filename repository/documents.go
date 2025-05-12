package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type DocumentsRepository interface {
	GetAll() ([]models.Documents, error)
	Create(documents models.Documents) error
	Find(id int) (models.Documents, error)
	Update(id int, document models.Documents) error
	Delete(id int) error
}

type documentsRepository struct {
	db *gorm.DB
}

func NewDocumentsRepository(db *gorm.DB) *documentsRepository {
	return &documentsRepository{db}
}

func (r *documentsRepository) GetAll() ([]models.Documents, error) {
	documents := []models.Documents{}
	if err := r.db.Find(&documents).Error; err != nil {
		return nil, err
	}
	return documents, nil
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
