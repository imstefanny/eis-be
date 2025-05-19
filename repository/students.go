package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type StudentsRepository interface {
	GetAll() ([]models.Students, error)
	Create(students models.Students) error
	Find(id int) (models.Students, error)
	Update(id int, student models.Students) error
	Delete(id int) error
}

type studentsRepository struct {
	db *gorm.DB
}

func NewStudentsRepository(db *gorm.DB) *studentsRepository {
	return &studentsRepository{db}
}

func (r *studentsRepository) GetAll() ([]models.Students, error) {
	students := []models.Students{}
	if err := r.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *studentsRepository) Create(students models.Students) error {
	err := r.db.Create(&students)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *studentsRepository) Find(id int) (models.Students, error) {
	student := models.Students{}
	if err := r.db.First(&student, id).Error; err != nil {
		return student, err
	}
	return student, nil
}

func (r *studentsRepository) Update(id int, student models.Students) error {
	query := r.db.Model(&student).Updates(student)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *studentsRepository) Delete(id int) error {
	student := models.Students{}
	if err := r.db.Delete(&student, id).Error; err != nil {
		return err
	}
	return nil
}
