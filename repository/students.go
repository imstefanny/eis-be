package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type StudentsRepository interface {
	Browse(page, limit int, search string) ([]models.Students, int64, error)
	Create(students models.Students) error
	GetByIds(ids []int) ([]models.Students, error)
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

func (r *studentsRepository) Browse(page, limit int, search string) ([]models.Students, int64, error) {
	var students []models.Students
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(full_name) LIKE ?", search).Limit(limit).Offset(offset).Preload("Applicant").Preload("User").Preload("Guardians").Preload("Documents").Find(&students).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.Students{}).Where("LOWER(full_name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return students, total, nil
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
	if err := r.db.Preload("Applicant").Preload("User").Preload("Guardians").Preload("Documents").First(&student, id).Error; err != nil {
		return student, err
	}
	return student, nil
}

func (r *studentsRepository) GetByIds(ids []int) ([]models.Students, error) {
	students := []models.Students{}
	if err := r.db.Where("id IN ?", ids).Find(&students).Error; err != nil {
		return students, err
	}
	return students, nil
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
