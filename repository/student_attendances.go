package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type StudentAttsRepository interface {
	BrowseByTermID(termID, page, limit int, search string, date string) ([]models.StudentAttendances, int64, error)
	CreateBatch(studentAtts []models.StudentAttendances) error
	FindByAcademicDate(academicID int, date string) ([]models.StudentAttendances, error)
	UpdateByTermID(termID int, studentAtt []models.StudentAttendances) error
	Browse(academicID, levelID, classID, termID int, search, start_date, end_date string) ([]models.StudentAttendances, error)

	// Students specific methods
	GetAttendanceByStudent(studentID int, start, end string) ([]models.StudentAttendances, error)
}

type studentAttsRepository struct {
	db *gorm.DB
}

func NewStudentAttsRepository(db *gorm.DB) *studentAttsRepository {
	return &studentAttsRepository{db}
}

func (r *studentAttsRepository) BrowseByTermID(termID, page, limit int, search string, date string) ([]models.StudentAttendances, int64, error) {
	var studentAtts []models.StudentAttendances
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	query := r.db.
		Where("term_id = ?", termID).
		Limit(limit).
		Offset(offset).
		Preload("Academic").
		Preload("Student")
	if search != "" {
		query = query.Where("LOWER(display_name) LIKE ?", search)
	}
	if date != "" {
		query = query.Where("DATE(date) = ?", date)
	}

	if err := query.Find(&studentAtts).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.StudentAttendances{}).Where("term_id = ? AND LOWER(display_name) LIKE ?", termID, search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return studentAtts, total, nil
}

func (r *studentAttsRepository) FindByAcademicDate(academicID int, date string) ([]models.StudentAttendances, error) {
	var studentAtts []models.StudentAttendances
	query := r.db.Where("academic_id = ?", academicID).Where("DATE(date) = ?", date).
		Preload("Academic").
		Preload("Student")
	if err := query.Find(&studentAtts).Error; err != nil {
		return nil, err
	}
	if len(studentAtts) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return studentAtts, nil
}

func (r *studentAttsRepository) CreateBatch(studentAtts []models.StudentAttendances) error {
	err := r.db.Create(&studentAtts)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *studentAttsRepository) UpdateByTermID(termID int, studentAtts []models.StudentAttendances) error {
	query := r.db.Save(studentAtts)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *studentAttsRepository) Browse(academicID, levelID, classID, termID int, search, start_date, end_date string) ([]models.StudentAttendances, error) {
	var studentAtts []models.StudentAttendances
	search = "%" + strings.ToLower(search) + "%"

	query := r.db.
		Preload("Academic").
		Preload("Academic.Classroom.Level").
		Preload("Student")

	if academicID > 0 {
		query = query.Where("academic_id = ?", academicID)
	}
	if levelID > 0 {
		query = query.
			Joins("JOIN academics ON academics.id = student_attendances.academic_id").
			Joins("JOIN classrooms ON classrooms.id = academics.classroom_id").
			Where("classrooms.level_id = ?", levelID)
	}
	if classID > 0 {
		query = query.Where("class_id = ?", classID)
	}
	if termID > 0 {
		query = query.Where("term_id = ?", termID)
	}
	if search != "" {
		query = query.Where("LOWER(student_attendances.display_name) LIKE ?", search)
	}
	if start_date != "" && end_date != "" {
		query = query.Where("DATE(date) BETWEEN ? AND ?", start_date, end_date)
	}

	if err := query.Find(&studentAtts).Error; err != nil {
		return nil, err
	}

	return studentAtts, nil
}

// Students specific methods
func (r *studentAttsRepository) GetAttendanceByStudent(studentID int, start, end string) ([]models.StudentAttendances, error) {
	var studentAtts []models.StudentAttendances
	if err := r.db.Where("student_id = ?", studentID).
		Where("academic_id = students.current_academic_id").
		Where("date BETWEEN ? AND ?", start, end).
		Preload("Academic").
		Joins("JOIN students ON students.id = student_attendances.student_id").
		Order("date ASC").
		Find(&studentAtts).Error; err != nil {
		return nil, err
	}
	if len(studentAtts) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return studentAtts, nil
}
