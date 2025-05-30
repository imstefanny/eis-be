package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type ClassNotesRepository interface {
	Browse(page, limit int, search string) ([]models.ClassNotes, int64, error)
	BrowseByAcademicID(academicID, page, limit int, search string) ([]models.ClassNotes, int64, error)
	Create(classNotes models.ClassNotes) error
	CreateBatch(classNotes []models.ClassNotes) error
	Find(id int) (models.ClassNotes, error)
	Update(id int, params map[string]interface{}) error
	// Delete(id int) error
}

type classNotesRepository struct {
	db *gorm.DB
}

func NewClassNotesRepository(db *gorm.DB) *classNotesRepository {
	return &classNotesRepository{db}
}

func (r *classNotesRepository) Browse(page, limit int, search string) ([]models.ClassNotes, int64, error) {
	var classNotes []models.ClassNotes
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(display_name) LIKE ?", search).Limit(limit).Offset(offset).Preload("Academic").Preload("Details").Find(&classNotes).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.ClassNotes{}).Where("LOWER(display_name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return classNotes, total, nil
}

func (r *classNotesRepository) BrowseByAcademicID(academicID, page, limit int, search string) ([]models.ClassNotes, int64, error) {
	var classNotes []models.ClassNotes
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("academic_id = ? AND LOWER(display_name) LIKE ?", academicID, search).Limit(limit).Offset(offset).Preload("Academic").Preload("Details").Find(&classNotes).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.ClassNotes{}).Where("academic_id = ? AND LOWER(display_name) LIKE ?", academicID, search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return classNotes, total, nil
}

func (r *classNotesRepository) Create(classNotes models.ClassNotes) error {
	err := r.db.Create(&classNotes)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *classNotesRepository) CreateBatch(classNotes []models.ClassNotes) error {
	err := r.db.Create(&classNotes)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *classNotesRepository) Find(id int) (models.ClassNotes, error) {
	classNote := models.ClassNotes{}
	if err := r.db.Preload("Academic").Preload("Details").First(&classNote, id).Error; err != nil {
		return classNote, err
	}
	return classNote, nil
}

func (r *classNotesRepository) Update(id int, params map[string]interface{}) error {
	classNote := models.ClassNotes{}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		addIDs := params["addIDs"].([]models.ClassNotes)
		if len(addIDs) > 0 {
			if err := tx.Create(&addIDs).Error; err != nil {
				return err
			}
		}
		if len(params["updateIDs"].([]int)) > 0 {
			if err := tx.Save(params["incomingUpdate"]).Error; err != nil {
				return err
			}
		}
		if len(params["removeIDs"].([]int)) > 0 {
			if err := tx.Unscoped().Delete(&classNote, params["removeIDs"]).Error; err != nil {
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

// func (r *classNotesRepository) Delete(id int) error {
// 	classNote := models.ClassNotes{}
// 	if err := r.db.Delete(&classNote, id).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
