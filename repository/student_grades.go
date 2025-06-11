package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type StudentGradesRepository interface {
	GetAll(academicID int) ([]models.StudentGrades, error)
	Create(studentGrades []models.StudentGrades) error
	UpdateByAcademicID(studentGrades []models.StudentGrades) error
	GetReport(startYear, endYear string, levelID int) ([]models.StudentGrades, error)
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

func (r *studentGradesRepository) GetReport(startYear, endYear string, levelID int) ([]models.StudentGrades, error) {
	var studentGrades []models.StudentGrades

	query := r.db.
		Joins("JOIN academics ON academics.id = student_grades.academic_id").
		Where("academics.start_year = ? AND academics.end_year = ?", startYear, endYear).
		Preload("Academic").
		Preload("Academic.Classroom").
		Preload("Academic.Classroom.Level").
		Preload("Student").
		Preload("Subject")

	if levelID > 0 {
		query = query.Joins("JOIN classrooms ON classrooms.id = academics.classroom_id").Where("classrooms.level_id = ?", levelID)
	}
	if err := query.Find(&studentGrades).Error; err != nil {
		return nil, err
	}

	return studentGrades, nil
}
