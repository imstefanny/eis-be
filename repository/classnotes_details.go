package repository

import (
	"eis-be/dto"
	"time"

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
	// if err := r.db.Table("class_notes_details").
	// 	Select(`
	// 		class_notes_details.id,
	// 		class_notes_details.note_id,
	// 		subject_schedules.day,
	// 		class_notes.date,
	// 		classrooms.display_name AS class,
	// 		subjects.name AS subject,
	// 		subject_schedules.id AS subj_sched_id,
	// 		subject_schedules.academic_id,
	// 		teachers.name AS teacher,
	// 		teachers.id AS teacher_id,
	// 		subject_schedules.start_hour,
	// 		subject_schedules.end_hour,
	// 		class_notes_details.materials,
	// 		class_notes_details.notes
	// 	`).
	// 	Joins("JOIN class_notes ON class_notes_details.note_id = class_notes.id").
	// 	Joins("RIGHT JOIN subject_schedules ON class_notes_details.subj_sched_id = subject_schedules.id").
	// 	Joins("JOIN subjects ON subject_schedules.subject_id = subjects.id").
	// 	Joins("JOIN academics ON subject_schedules.academic_id = academics.id").
	// 	Joins("JOIN classrooms ON academics.classroom_id = classrooms.id").
	// 	Joins("JOIN teachers ON class_notes_details.teacher_id = teachers.id OR subject_schedules.teacher_id = teachers.id").
	// 	Where("teachers.id = ?", teacherID).
	// 	Where("DATE(class_notes.date) = ? OR DATE(class_notes.date) IS NULL", date).
	// 	Scan(&details).Error; err != nil {
	// 	return nil, err
	// }
	parsedDate, _ := time.Parse("2006-01-02", date)
	weekday := parsedDate.Weekday().String()

	// subQuery := r.db.Table("class_notes").Select("id").Where("DATE(class_notes.date) = ?", date)

	// if err := r.db.Debug().Table("subject_schedules").
	// 	Select(`
	// 		class_notes_details.id,
	// 		CASE WHEN DATE(class_notes.date) = ? THEN class_notes_details.note_id ELSE NULL END AS note_id,
	// 		subject_schedules.day,
	// 		class_notes.date,
	// 		classrooms.display_name AS class,
	// 		subjects.name AS subject,
	// 		subject_schedules.id AS subj_sched_id,
	// 		subject_schedules.academic_id,
	// 		teachers.name AS teacher,
	// 		teachers.id AS teacher_id,
	// 		subject_schedules.start_hour,
	// 		subject_schedules.end_hour,
	// 		CASE WHEN DATE(class_notes.date) = ? THEN class_notes_details.materials ELSE NULL END AS materials,
	// 		class_notes_details.notes
	// 	`, date, date).
	// 	Joins(`LEFT JOIN class_notes_details 
	// 		ON subject_schedules.id = class_notes_details.subj_sched_id 
	// 		AND class_notes_details.note_id IN ?`, gorm.Expr("(?)", subQuery)).
	// 	Joins(`LEFT JOIN class_notes 
	// 		ON class_notes_details.note_id = class_notes.id 
	// 		AND DATE(class_notes.date) = ?`, date).
	// 	Joins("JOIN subjects ON subject_schedules.subject_id = subjects.id").
	// 	Joins("JOIN academics ON subject_schedules.academic_id = academics.id").
	// 	Joins("JOIN classrooms ON academics.classroom_id = classrooms.id").
	// 	Joins(`LEFT JOIN teachers 
	// 		ON teachers.id = class_notes_details.teacher_id 
	// 		OR teachers.id = subject_schedules.teacher_id`).
	// 	Where("subject_schedules.day = ?", weekday).
	// 	Where("(teachers.id = ? OR class_notes_details.teacher_id = ?)", teacherID, teacherID).
	// 	Scan(&details).Error; err != nil {
	// 	return nil, err
	// }
	
	rawSQL := `
	SELECT
		class_notes_details.id,
		CASE WHEN DATE(class_notes.date) = ? THEN class_notes_details.note_id ELSE NULL END AS note_id,
		subject_schedules.day,
		class_notes.date,
		classrooms.display_name AS class,
		subjects.name AS subject,
		subject_schedules.id AS subj_sched_id,
		subject_schedules.academic_id,
		teachers.name AS teacher,
		teachers.id AS teacher_id,
		subject_schedules.start_hour,
		subject_schedules.end_hour,
		CASE WHEN DATE(class_notes.date) = ? THEN class_notes_details.materials ELSE NULL END AS materials,
		class_notes_details.notes
	FROM subject_schedules
	LEFT JOIN class_notes_details 
		ON subject_schedules.id = class_notes_details.subj_sched_id
		AND class_notes_details.note_id IN (
			SELECT id FROM class_notes WHERE DATE(class_notes.date) = ?
		)
	LEFT JOIN class_notes 
		ON class_notes_details.note_id = class_notes.id
		AND DATE(class_notes.date) = ?
	JOIN subjects ON subject_schedules.subject_id = subjects.id
	JOIN academics ON subject_schedules.academic_id = academics.id
	JOIN classrooms ON academics.classroom_id = classrooms.id
	LEFT JOIN teachers 
		ON teachers.id = COALESCE(class_notes_details.teacher_id, subject_schedules.teacher_id)
	WHERE subject_schedules.day = ?
		AND (teachers.id = ? OR class_notes_details.teacher_id = ?)
	`

	if err := r.db.Raw(rawSQL, date, date, date, date, weekday, teacherID, teacherID).Scan(&details).Error; err != nil {
		return nil, err
	}
	return details, nil
}
