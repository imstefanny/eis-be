package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type GuardiansRepository interface {
	GetAll() ([]models.Guardians, error)
	Create(guardians models.Guardians) error
	Find(id int) (models.Guardians, error)
	Update(id int, guardian models.Guardians) error
	Delete(id int) error
}

type guardiansRepository struct {
	db *gorm.DB
}

func NewGuardiansRepository(db *gorm.DB) *guardiansRepository {
	return &guardiansRepository{db}
}

func (r *guardiansRepository) GetAll() ([]models.Guardians, error) {
	guardians := []models.Guardians{}
	if err := r.db.Find(&guardians).Error; err != nil {
		return nil, err
	}
	return guardians, nil
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
	if err := r.db.First(&guardian, id).Error; err != nil {
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
