package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type WorkSchedsRepository interface {
	Browse(page, limit int, search string) ([]models.WorkScheds, int64, error)
	Create(workScheds models.WorkScheds) error
	Find(id int) (models.WorkScheds, error)
	Update(id int, workSched models.WorkScheds) error
	Delete(id int) error
}

type workSchedsRepository struct {
	db *gorm.DB
}

func NewWorkSchedsRepository(db *gorm.DB) *workSchedsRepository {
	return &workSchedsRepository{db}
}

func (r *workSchedsRepository) Browse(page, limit int, search string) ([]models.WorkScheds, int64, error) {
	var workScheds []models.WorkScheds
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(name) LIKE ?", search).Limit(limit).Offset(offset).Find(&workScheds).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.WorkScheds{}).Where("LOWER(name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return workScheds, total, nil
}

func (r *workSchedsRepository) Create(workScheds models.WorkScheds) error {
	err := r.db.Create(&workScheds)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *workSchedsRepository) Find(id int) (models.WorkScheds, error) {
	workSched := models.WorkScheds{}
	if err := r.db.Preload("Details").First(&workSched, id).Error; err != nil {
		return workSched, err
	}
	return workSched, nil
}

func (r *workSchedsRepository) Update(id int, workSched models.WorkScheds) error {
	query := r.db.Model(&workSched).Updates(workSched)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *workSchedsRepository) Delete(id int) error {
	workSched := models.WorkScheds{}
	workSchedDetail := models.WorkSchedDetails{}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("work_sched_id = ?", id).Delete(&workSchedDetail).Error; err != nil {
			return err
		}
		if err := tx.Where("id = ?", id).Delete(&workSched).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
