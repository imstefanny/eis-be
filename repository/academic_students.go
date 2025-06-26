package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type AcademicStudentsRepository interface {
	GetByIDs(ids []int) ([]models.AcademicStudents, error)
	Update(academicStudents []models.AcademicStudents) error
	FindByAcademicIDAndStudentID(academicID, studentID uint) (models.AcademicStudents, error)
}

type academicStudentsRepository struct {
	db *gorm.DB
}

func NewAcademicStudentsRepository(db *gorm.DB) *academicStudentsRepository {
	return &academicStudentsRepository{db}
}

func (r *academicStudentsRepository) GetByIDs(ids []int) ([]models.AcademicStudents, error) {
	var academicStudents []models.AcademicStudents
	if err := r.db.Where("id IN ?", ids).Find(&academicStudents).Error; err != nil {
		return nil, err
	}
	return academicStudents, nil
}

func (r *academicStudentsRepository) Update(academicStudents []models.AcademicStudents) error {
	query := r.db.Save(academicStudents)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *academicStudentsRepository) FindByAcademicIDAndStudentID(academicID, studentID uint) (models.AcademicStudents, error) {
	academicStudent := models.AcademicStudents{}
	if err := r.db.Where("academics_id = ? AND students_id = ?", academicID, studentID).First(&academicStudent).Error; err != nil {
		return academicStudent, err
	}
	return academicStudent, nil
}
