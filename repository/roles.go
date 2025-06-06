package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type RolesRepository interface {
	Browse(page, limit int, search string) ([]models.Roles, int64, error)
	Create(roles models.Roles) error
	Find(id int) (models.Roles, error)
	Update(id int, role models.Roles) error
	Delete(id int) error
	FindByName(name string) (models.Roles, error)
}

type rolesRepository struct {
	db *gorm.DB
}

func NewRolesRepository(db *gorm.DB) *rolesRepository {
	return &rolesRepository{db}
}

func (r *rolesRepository) Browse(page, limit int, search string) ([]models.Roles, int64, error) {
	var roles []models.Roles
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.
		Where("LOWER(name) LIKE ?", search).
		Limit(limit).
		Offset(offset).
		Find(&roles).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.Roles{}).Where("LOWER(name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return roles, total, nil
}

func (r *rolesRepository) Create(roles models.Roles) error {
	err := r.db.Create(&roles)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *rolesRepository) Find(id int) (models.Roles, error) {
	role := models.Roles{}
	if err := r.db.Preload("Permissions").First(&role, id).Error; err != nil {
		return role, err
	}
	return role, nil
}

func (r *rolesRepository) Update(id int, role models.Roles) error {
	oldRole := models.Roles{}
	if err := r.db.Find(&oldRole, id).Error; err != nil {
		return err
	}
	r.db.Model(&oldRole).Association("Permissions").Clear()
	if err := r.db.Model(&role).Updates(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *rolesRepository) Delete(id int) error {
	role := models.Roles{}
	if err := r.db.Unscoped().Delete(&role, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *rolesRepository) FindByName(name string) (models.Roles, error) {
	role := models.Roles{}
	if err := r.db.Where("name = ?", name).First(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}
