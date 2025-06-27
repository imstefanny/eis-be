package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type CurriculumsRepository interface {
	Browse(page, limit int, search string) ([]models.Curriculums, int64, error)
	Create(curriculums models.Curriculums) (uint, error)
	Find(id int) (models.Curriculums, error)
	Update(id int, params map[string]interface{}) error
	Delete(id int) error
	UnDelete(id int) error

	// Helper methods
	GetCurriculumnsByLevelIDandGrade(levelID int, grade string) ([]models.Curriculums, error)
}

type curriculumsRepository struct {
	db *gorm.DB
}

func NewCurriculumsRepository(db *gorm.DB) *curriculumsRepository {
	return &curriculumsRepository{db}
}

func (r *curriculumsRepository) Browse(page, limit int, search string) ([]models.Curriculums, int64, error) {
	var curriculums []models.Curriculums
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(display_name) LIKE ?", search).Preload("Level").Unscoped().Limit(limit).Offset(offset).Find(&curriculums).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.Curriculums{}).Unscoped().Where("LOWER(display_name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return curriculums, total, nil
}

func (r *curriculumsRepository) Create(curriculums models.Curriculums) (uint, error) {
	err := r.db.Create(&curriculums)
	if err.Error != nil {
		return 0, err.Error
	}
	return curriculums.ID, nil
}

func (r *curriculumsRepository) Find(id int) (models.Curriculums, error) {
	curriculum := models.Curriculums{}
	if err := r.db.Preload("Level").Preload("CurriculumSubjects").Preload("CurriculumSubjects.Subject").Unscoped().First(&curriculum, id).Error; err != nil {
		return curriculum, err
	}
	return curriculum, nil
}

func (r *curriculumsRepository) Update(id int, params map[string]interface{}) error {
	curriculumSubject := models.CurriculumSubjects{}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		parents := params["parents"].(models.Curriculums)
		if err := tx.Save(&parents).Error; err != nil {
			return err
		}
		if len(params["removeIDs"].([]int)) > 0 {
			if err := tx.Unscoped().Delete(&curriculumSubject, params["removeIDs"]).Error; err != nil {
				return err
			}
		}
		addIDs := params["addIDs"].([]models.CurriculumSubjects)
		if len(addIDs) > 0 {
			if err := tx.Create(&addIDs).Error; err != nil {
				return err
			}
		}
		if len(params["updateIDs"].([]int)) > 0 {
			if err := tx.Save(params["incomingUpdates"]).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *curriculumsRepository) Delete(id int) error {
	curriculum := models.Curriculums{}
	if err := r.db.Delete(&curriculum, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *curriculumsRepository) UnDelete(id int) error {
	curriculum := models.Curriculums{}
	if err := r.db.Unscoped().Model(&curriculum).Where("id = ?", id).Update("deleted_at", nil).Error; err != nil {
		return err
	}
	return nil
}

func (r *curriculumsRepository) GetCurriculumnsByLevelIDandGrade(levelID int, grade string) ([]models.Curriculums, error) {
	var curriculums []models.Curriculums
	if err := r.db.Where("level_id = ? AND grade = ?", levelID, grade).Order("created_at DESC").Find(&curriculums).Error; err != nil {
		return nil, err
	}
	return curriculums, nil
}
