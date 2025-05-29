package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type AcademicsRepository interface {
	Browse(page, limit int, search string) ([]models.Academics, int64, error)
	GetAll() ([]models.Academics, error)
	Create(academics models.Academics) error
	CreateBatch(academics []models.Academics) error
	Find(id int) (models.Academics, error)
	Update(id int, academic models.Academics) error
	Delete(id int) error
}

type academicsRepository struct {
	db *gorm.DB
}

func NewAcademicsRepository(db *gorm.DB) *academicsRepository {
	return &academicsRepository{db}
}

func (r *academicsRepository) Browse(page, limit int, search string) ([]models.Academics, int64, error) {
	var academics []models.Academics
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(display_name) LIKE ?", search).Limit(limit).Offset(offset).Preload("Classroom").Preload("HomeroomTeacher").Preload("Students").Find(&academics).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.Academics{}).Where("LOWER(display_name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return academics, total, nil
}

func (r *academicsRepository) GetAll() ([]models.Academics, error) {
	var academics []models.Academics
	if err := r.db.Preload("SubjScheds").Find(&academics).Error; err != nil {
		return nil, err
	}
	return academics, nil
}

func (r *academicsRepository) Create(academics models.Academics) error {
	err := r.db.Create(&academics)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *academicsRepository) CreateBatch(academics []models.Academics) error {
	if len(academics) == 0 {
		return nil
	}
	err := r.db.Create(&academics)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *academicsRepository) Find(id int) (models.Academics, error) {
	academic := models.Academics{}
	if err := r.db.Preload("Classroom").Preload("HomeroomTeacher").Preload("Students").Preload("SubjScheds").Preload("SubjScheds.Teacher").Preload("SubjScheds.Subject").First(&academic, id).Error; err != nil {
		return academic, err
	}
	return academic, nil
}

func (r *academicsRepository) Update(id int, academic models.Academics) error {
	oldAcademic := models.Academics{}
	if e := r.db.Find(&oldAcademic, id).Error; e != nil {
		return e
	}
	r.db.Model(&oldAcademic).Association("Students").Clear()
	query := r.db.Model(&academic).Updates(academic)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *academicsRepository) Delete(id int) error {
	academic := models.Academics{}
	if err := r.db.Delete(&academic, id).Error; err != nil {
		return err
	}
	return nil
}
