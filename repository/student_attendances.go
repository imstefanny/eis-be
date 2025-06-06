package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type StudentAttsRepository interface {
	BrowseByAcademicID(academicID, page, limit int, search string, date string) ([]models.StudentAttendances, int64, error)
	CreateBatch(studentAtts []models.StudentAttendances) error
	FindByAcademicDate(academicID int, date string) ([]models.StudentAttendances, error)
	UpdateByAcademicID(academicID int, studentAtt []models.StudentAttendances) error

	// Students specific methods
	GetAttendanceByStudent(studentID int, start, end string) ([]models.StudentAttendances, error)
}

type studentAttsRepository struct {
	db *gorm.DB
}

func NewStudentAttsRepository(db *gorm.DB) *studentAttsRepository {
	return &studentAttsRepository{db}
}

func (r *studentAttsRepository) BrowseByAcademicID(academicID, page, limit int, search string, date string) ([]models.StudentAttendances, int64, error) {
	var studentAtts []models.StudentAttendances
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	query := r.db.
		Where("academic_id = ?", academicID).
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
	if err := r.db.Model(&models.StudentAttendances{}).Where("academic_id = ? AND LOWER(display_name) LIKE ?", academicID, search).Count(&total).Error; err != nil {
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

func (r *studentAttsRepository) UpdateByAcademicID(academicID int, studentAtts []models.StudentAttendances) error {
	query := r.db.Save(studentAtts)
	if err := query.Error; err != nil {
		return err
	}
	return nil
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
