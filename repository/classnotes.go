package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type ClassNotesRepository interface {
	Browse(page, limit int, search string) ([]models.ClassNotes, int64, error)
	BrowseByTermID(termID, page, limit int, search string) ([]models.ClassNotes, int64, error)
	Create(classNotes models.ClassNotes) error
	CreateBatch(classNotes []models.ClassNotes) error
	Find(id int) (models.ClassNotes, error)
	FindClassNoteDetail(id int) (models.ClassNotesDetails, error)
	Update(id int, params map[string]interface{}) error
	CreateDetail(models.ClassNotesDetails) error
	UpdateDetail(models.ClassNotesDetails) error
	Delete(id int) error

	// Teacher specific methods
	FindByTeacher(teacherID, schedID int, date string) ([]models.ClassNotes, error)
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
	if err := r.db.Where("LOWER(display_name) LIKE ?", search).Limit(limit).Offset(offset).Preload("Term").Preload("Academic").Find(&classNotes).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.ClassNotes{}).Where("LOWER(display_name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return classNotes, total, nil
}

func (r *classNotesRepository) BrowseByTermID(termID, page, limit int, search string) ([]models.ClassNotes, int64, error) {
	var classNotes []models.ClassNotes
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("term_id = ? AND LOWER(display_name) LIKE ?", termID, search).
		Limit(limit).
		Offset(offset).
		Preload("Academic").
		Preload("Details").
		Preload("Details.SubjSched.Teacher").
		Preload("Details.SubjSched.Subject").
		Preload("Details.Teacher").
		Find(&classNotes).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.ClassNotes{}).Where("term_id = ? AND LOWER(display_name) LIKE ?", termID, search).Count(&total).Error; err != nil {
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
	if err := r.db.Preload("Academic").
		Preload("Details").
		Preload("Details.SubjSched.Teacher").
		Preload("Details.SubjSched.Subject").
		Preload("Details.Teacher").
		First(&classNote, id).Error; err != nil {
		return classNote, err
	}
	return classNote, nil
}

func (r *classNotesRepository) FindClassNoteDetail(id int) (models.ClassNotesDetails, error) {
	classNoteDetail := models.ClassNotesDetails{}
	if err := r.db.
		Where("id = ?", id).
		Order("id").
		Limit(1).
		Find(&classNoteDetail).Error; err != nil {
		return classNoteDetail, err
	}
	return classNoteDetail, nil
}

func (r *classNotesRepository) Update(id int, params map[string]interface{}) error {
	classNote := models.ClassNotesDetails{}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		addIDs := params["addIDs"].([]models.ClassNotesDetails)
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

func (r *classNotesRepository) CreateDetail(classNoteDetail models.ClassNotesDetails) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&classNoteDetail).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *classNotesRepository) UpdateDetail(classNoteDetail models.ClassNotesDetails) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&classNoteDetail).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *classNotesRepository) Delete(id int) error {
	classNote := models.ClassNotes{}
	classNoteDetails := models.ClassNotesDetails{}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("note_id = ?", id).Delete(&classNoteDetails).Error; err != nil {
			return err
		}
		if err := tx.Where("id = ?", id).Delete(&classNote).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// Teacher specific methods
func (r *classNotesRepository) FindByTeacher(teacherID, schedID int, date string) ([]models.ClassNotes, error) {
	var classNotes []models.ClassNotes

	if err := r.db.Preload("Academic").
		Preload("Details").
		Preload("Details.SubjSched.Teacher").
		Preload("Details.SubjSched.Subject").
		Preload("Details.Teacher").
		Joins("JOIN class_notes_details ON class_notes.id = class_notes_details.note_id").
		Joins("JOIN subject_schedules ON class_notes_details.subj_sched_id = subject_schedules.id").
		Where("class_notes_details.teacher_id = ?", teacherID).
		Where("subject_schedules.id = ?", schedID).
		Where("DATE(class_notes.date) = ?", date).Find(&classNotes).Error; err != nil {
		return nil, err
	}

	return classNotes, nil
}
