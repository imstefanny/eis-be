package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type GuardiansRepository interface {
	GetAll() ([]models.Guardians, error)
	Create(guardians models.Guardians) error
	Find(id int) (models.Guardians, error)
	Update(id int, blog models.Guardians) error
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
	blog := models.Guardians{}
	if err := r.db.First(&blog, id).Error; err != nil {
		return blog, err
	}
	return blog, nil
}

func (r *guardiansRepository) Update(id int, blog models.Guardians) error {
	query := r.db.Model(&blog).Updates(blog)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *guardiansRepository) Delete(id int) error {
	blog := models.Guardians{}
	if err := r.db.Delete(&blog, id).Error; err != nil {
		return err
	}
	return nil
}
