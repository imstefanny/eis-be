package repository

import (
	"eis-be/dto"

	"gorm.io/gorm"
)

type ClassNotesDetailsRepository interface {
	GetAllByTeacher(teacherID int, date string) ([]dto.ClassNotesRepoRes, error)
}

type classNotesDetailsRepository struct {
	db *gorm.DB
}

func NewClassNotesDetailsRepository(db *gorm.DB) *classNotesDetailsRepository {
	return &classNotesDetailsRepository{db}
}

func (r *classNotesDetailsRepository) GetAllByTeacher(teacherID int, date string) ([]dto.ClassNotesRepoRes, error) {
	details := []dto.ClassNotesRepoRes{}
	if err := r.db.Table("class_notes_details").
		Select(`
			class_notes_details.id,
			class_notes_details.note_id,
			subject_schedules.day,
			class_notes.date,
			classrooms.name AS class,
			subjects.name AS subject,
			class_notes_details.subj_sched_id,
			subject_schedules.academic_id,
			teachers.name AS teacher,
			subject_schedules.start_hour,
			subject_schedules.end_hour,
			class_notes_details.materials,
			class_notes_details.notes
		`).
		Joins("JOIN class_notes ON class_notes_details.note_id = class_notes.id").
		Joins("RIGHT JOIN subject_schedules ON class_notes_details.subj_sched_id = subject_schedules.id").
		Joins("JOIN subjects ON subject_schedules.subject_id = subjects.id").
		Joins("JOIN academics ON subject_schedules.academic_id = academics.id").
		Joins("JOIN classrooms ON academics.classroom_id = classrooms.id").
		Joins("JOIN teachers ON class_notes_details.teacher_id = teachers.id OR subject_schedules.teacher_id = teachers.id").
		Where("teachers.id = ?", teacherID).
		Where("DATE(class_notes.date) = ? OR DATE(class_notes.date) IS NULL", date).
		Scan(&details).Error; err != nil {
		return nil, err
	}
	// if err := r.db.Model(&models.ClassNotesDetails{}).
	// 	Preload("Note").
	// 	Preload("SubjSched").
	// 	Preload("SubjSched.Academic.Classroom").
	// 	Preload("SubjSched.Subject").
	// 	Preload("Teacher").
	// 	Joins("JOIN class_notes ON class_notes_details.note_id = class_notes.id").
	// 	Where("teacher_id = ?", teacherID).
	// 	Where("DATE(class_notes.date) = ?", date).
	// 	Find(&details).Error; err != nil {
	// 	return nil, err
	// }
	return details, nil
}
