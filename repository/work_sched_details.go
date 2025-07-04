package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type WorkSchedDetailsRepository interface {
	Create(workSchedDetails []models.WorkSchedDetails) error
	Find(id []int) ([]models.WorkSchedDetails, error)
	Update(id []int, workSchedDetail []models.WorkSchedDetails) error
	Delete(id []int) error
	Undelete(workSchedId int) error
}

type workSchedDetailsRepository struct {
	db *gorm.DB
}

func NewWorkSchedDetailsRepository(db *gorm.DB) *workSchedDetailsRepository {
	return &workSchedDetailsRepository{db}
}

func (r *workSchedDetailsRepository) Create(workSchedDetails []models.WorkSchedDetails) error {
	err := r.db.Create(&workSchedDetails)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *workSchedDetailsRepository) Find(id []int) ([]models.WorkSchedDetails, error) {
	workSchedDetails := []models.WorkSchedDetails{}
	if err := r.db.Unscoped().Find(&workSchedDetails, id).Error; err != nil {
		return workSchedDetails, err
	}
	return workSchedDetails, nil
}

func (r *workSchedDetailsRepository) Update(id []int, workSchedDetail []models.WorkSchedDetails) error {
	query := r.db.Save(workSchedDetail)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *workSchedDetailsRepository) Undelete(workSchedId int) error {
	result := r.db.Model(&models.WorkSchedDetails{}).Unscoped().Where("work_sched_id = ?", workSchedId).Update("deleted_at", nil)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *workSchedDetailsRepository) Delete(id []int) error {
	workSchedDetail := models.WorkSchedDetails{}
	if err := r.db.Unscoped().Delete(&workSchedDetail, id).Error; err != nil {
		return err
	}
	return nil
}
