package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type ClassNotesDetailsRepository interface {
	GetAllByTeacher(teacherID int, date string) ([]models.ClassNotesDetails, error)
}

type classNotesDetailsRepository struct {
	db *gorm.DB
}

func NewClassNotesDetailsRepository(db *gorm.DB) *classNotesDetailsRepository {
	return &classNotesDetailsRepository{db}
}

func (r *classNotesDetailsRepository) GetAllByTeacher(teacherID int, date string) ([]models.ClassNotesDetails, error) {
	var details []models.ClassNotesDetails
	if err := r.db.Preload("Note").
		Preload("SubjSched").
		Preload("SubjSched.Academic.Classroom").
		Preload("SubjSched.Subject").
		Preload("Teacher").
		Joins("JOIN class_notes ON class_notes_details.note_id = class_notes.id").
		Where("teacher_id = ?", teacherID).
		Where("DATE(class_notes.date) = ?", date).
		Find(&details).Error; err != nil {
		return nil, err
	}
	return details, nil
}
