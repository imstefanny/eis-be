package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type ClassroomsRepository interface {
	GetAll() ([]models.Classrooms, error)
	Browse(page, limit int, search string) ([]models.Classrooms, int64, error)
	Create(classrooms models.Classrooms) error
	Find(id int) (models.Classrooms, error)
	Update(id int, classroom models.Classrooms) error
	Delete(id int) error
	UnDelete(id int) error
}

type classroomsRepository struct {
	db *gorm.DB
}

func NewClassroomsRepository(db *gorm.DB) *classroomsRepository {
	return &classroomsRepository{db}
}

func (r *classroomsRepository) GetAll() ([]models.Classrooms, error) {
	var classrooms []models.Classrooms
	err := r.db.Find(&classrooms)
	if err.Error != nil {
		return nil, err.Error
	}
	return classrooms, nil
}

func (r *classroomsRepository) Browse(page, limit int, search string) ([]models.Classrooms, int64, error) {
	var classrooms []models.Classrooms
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(display_name) LIKE ?", search).Limit(limit).Offset(offset).Unscoped().Preload("Level").Find(&classrooms).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.Classrooms{}).Where("LOWER(display_name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return classrooms, total, nil
}

func (r *classroomsRepository) Create(classrooms models.Classrooms) error {
	err := r.db.Create(&classrooms)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *classroomsRepository) CreateBatch(classrooms []models.Classrooms) error {
	if len(classrooms) == 0 {
		return nil
	}
	err := r.db.Create(&classrooms)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *classroomsRepository) Find(id int) (models.Classrooms, error) {
	classroom := models.Classrooms{}
	if err := r.db.Preload("Level").Unscoped().First(&classroom, id).Error; err != nil {
		return classroom, err
	}
	return classroom, nil
}

func (r *classroomsRepository) Update(id int, classroom models.Classrooms) error {
	query := r.db.Model(&classroom).Updates(classroom)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *classroomsRepository) Delete(id int) error {
	classroom := models.Classrooms{}
	if err := r.db.Delete(&classroom, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *classroomsRepository) UnDelete(id int) error {
	classroom := models.Classrooms{}
	if err := r.db.Unscoped().Model(&classroom).Where("id = ?", id).Update("deleted_at", nil).Error; err != nil {
		return err
	}
	return nil
}
