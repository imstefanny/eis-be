package repository

import (
	"eis-be/dto"
	"eis-be/models"

	"gorm.io/gorm"
)

type StudentGradesRepository interface {
	GetAll(academicID int) ([]models.StudentGrades, error)
	Create(studentGrades []models.StudentGrades) error
	UpdateByAcademicID(studentGrades []models.StudentGrades) error
	GetReport(startYear, endYear string, levelID, academicID int) ([]dto.StudentGradesReportQuery, error)
}

type studentGradesRepository struct {
	db *gorm.DB
}

func NewStudentGradesRepository(db *gorm.DB) *studentGradesRepository {
	return &studentGradesRepository{db}
}

func (r *studentGradesRepository) GetAll(academicID int) ([]models.StudentGrades, error) {
	var studentGrades []models.StudentGrades
	if err := r.db.Where("academic_id = ?", academicID).
		Preload("Academic").
		Preload("Student").
		Preload("Subject").
		Find(&studentGrades).Error; err != nil {
		return nil, err
	}
	return studentGrades, nil
}

func (r *studentGradesRepository) Create(studentGrades []models.StudentGrades) error {
	err := r.db.Create(&studentGrades)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *studentGradesRepository) UpdateByAcademicID(studentGrades []models.StudentGrades) error {
	query := r.db.Save(studentGrades)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *studentGradesRepository) GetReport(startYear, endYear string, levelID, academicID int) ([]dto.StudentGradesReportQuery, error) {
	var studentGrades []dto.StudentGradesReportQuery

	query := r.db.Table("student_grades").
		Select(`
			students.id AS student_id,
			students.full_name AS student,
			students.nis,
			classrooms.id AS class_id,
			classrooms.display_name AS class,
			ROUND(AVG(student_grades.final_grade), 2) AS finals
		`).
		Joins("JOIN academics ON academics.id = student_grades.academic_id").
		Joins("JOIN students ON students.id = student_grades.student_id").
		Joins("JOIN classrooms ON classrooms.id = academics.classroom_id").
		Where("academics.start_year = ? AND academics.end_year = ?", startYear, endYear).
		Group("students.id, classrooms.id").
		Order("classrooms.id, finals DESC")

	if levelID > 0 {
		query = query.Where("classrooms.level_id = ?", levelID)
	}
	if academicID > 0 {
		query = query.Where("student_grades.academic_id = ?", academicID)
	}
	if err := query.Find(&studentGrades).Error; err != nil {
		return []dto.StudentGradesReportQuery{}, err
	}

	return studentGrades, nil
}
