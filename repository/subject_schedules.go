package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type SubjSchedsRepository interface {
	Browse(page, limit int, search string) ([]models.SubjectSchedules, int64, error)
	Create(subjScheds models.SubjectSchedules) error
	Find(id int) (models.SubjectSchedules, error)
	Update(id int, subjSched models.SubjectSchedules) error
	Delete(id int) error
}

type subjSchedsRepository struct {
	db *gorm.DB
}

func NewSubjSchedsRepository(db *gorm.DB) *subjSchedsRepository {
	return &subjSchedsRepository{db}
}

func (r *subjSchedsRepository) Browse(page, limit int, search string) ([]models.SubjectSchedules, int64, error) {
	var subjScheds []models.SubjectSchedules
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.
		Where("LOWER(display_name) LIKE ?", search).
		Preload("Academic").
		Preload("Subject").
		Preload("Teacher").
		Limit(limit).
		Offset(offset).
		Find(&subjScheds).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.SubjectSchedules{}).Where("LOWER(display_name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return subjScheds, total, nil
}

func (r *subjSchedsRepository) Create(subjScheds models.SubjectSchedules) error {
	err := r.db.Create(&subjScheds)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *subjSchedsRepository) Find(id int) (models.SubjectSchedules, error) {
	subjSched := models.SubjectSchedules{}
	if err := r.db.Preload("Academic").Preload("Subject").Preload("Teacher").First(&subjSched, id).Error; err != nil {
		return subjSched, err
	}
	return subjSched, nil
}

func (r *subjSchedsRepository) Update(id int, subjSched models.SubjectSchedules) error {
	query := r.db.Model(&subjSched).Updates(subjSched)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *subjSchedsRepository) Delete(id int) error {
	subjSched := models.SubjectSchedules{}
	if err := r.db.Delete(&subjSched, id).Error; err != nil {
		return err
	}
	return nil
}
