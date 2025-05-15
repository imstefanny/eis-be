package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type WorkSchedsRepository interface {
	GetAll() ([]models.WorkScheds, error)
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

func (r *workSchedsRepository) GetAll() ([]models.WorkScheds, error) {
	workScheds := []models.WorkScheds{}
	if err := r.db.Preload("Details").Find(&workScheds).Error; err != nil {
		return nil, err
	}
	return workScheds, nil
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
