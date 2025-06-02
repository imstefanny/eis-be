package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type StudentGradesRepository interface {
	Create(studentGrades models.StudentGrades) error
}

type studentGradesRepository struct {
	db *gorm.DB
}

func NewStudentGradesRepository(db *gorm.DB) *studentGradesRepository {
	return &studentGradesRepository{db}
}

func (r *studentGradesRepository) Create(studentGrades models.StudentGrades) error {
	err := r.db.Create(&studentGrades)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
