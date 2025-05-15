package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type SubjectsRepository interface {
	GetAll() ([]models.Subjects, error)
	Create(subjects models.Subjects) error
	Find(id int) (models.Subjects, error)
	Update(id int, subject models.Subjects) error
	Delete(id int) error
}

type subjectsRepository struct {
	db *gorm.DB
}

func NewSubjectsRepository(db *gorm.DB) *subjectsRepository {
	return &subjectsRepository{db}
}

func (r *subjectsRepository) GetAll() ([]models.Subjects, error) {
	subjects := []models.Subjects{}
	if err := r.db.Find(&subjects).Error; err != nil {
		return nil, err
	}
	return subjects, nil
}

func (r *subjectsRepository) Create(subjects models.Subjects) error {
	err := r.db.Create(&subjects)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *subjectsRepository) Find(id int) (models.Subjects, error) {
	subject := models.Subjects{}
	if err := r.db.First(&subject, id).Error; err != nil {
		return subject, err
	}
	return subject, nil
}

func (r *subjectsRepository) Update(id int, subject models.Subjects) error {
	query := r.db.Model(&subject).Updates(subject)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *subjectsRepository) Delete(id int) error {
	subject := models.Subjects{}
	if err := r.db.Delete(&subject, id).Error; err != nil {
		return err
	}
	return nil
}
