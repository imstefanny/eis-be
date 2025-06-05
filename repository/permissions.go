package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type PermissionsRepository interface {
	GetByIds(ids []int) ([]models.Permissions, error)
	GetAll() ([]models.Permissions, error)
}

type permissionsRepository struct {
	db *gorm.DB
}

func NewPermissionsRepository(db *gorm.DB) *permissionsRepository {
	return &permissionsRepository{db}
}

func (r *permissionsRepository) GetAll() ([]models.Permissions, error) {
	var permissions []models.Permissions
	if err := r.db.Find(&permissions).Error; err != nil {
		return nil, err
	}
	if len(permissions) == 0 {
		return nil, models.ErrPermissionsNotFound{}
	}
	return permissions, nil
}

func (r *permissionsRepository) GetByIds(ids []int) ([]models.Permissions, error) {
	var permissions []models.Permissions
	if err := r.db.Where("id IN ?", ids).Find(&permissions).Error; err != nil {
		return nil, err
	}
	if len(permissions) == 0 {
		return nil, models.ErrPermissionsNotFound{}
	}
	return permissions, nil
}
