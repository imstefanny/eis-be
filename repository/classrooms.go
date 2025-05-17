package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type ClassroomsRepository interface {
	GetAll() ([]models.Classrooms, error)
	Create(classrooms models.Classrooms) error
	Find(id int) (models.Classrooms, error)
	Update(id int, classroom models.Classrooms) error
	Delete(id int) error
}

type classroomsRepository struct {
	db *gorm.DB
}

func NewClassroomsRepository(db *gorm.DB) *classroomsRepository {
	return &classroomsRepository{db}
}

func (r *classroomsRepository) GetAll() ([]models.Classrooms, error) {
	classrooms := []models.Classrooms{}
	if err := r.db.Find(&classrooms).Error; err != nil {
		return nil, err
	}
	return classrooms, nil
}

func (r *classroomsRepository) Create(classrooms models.Classrooms) error {
	err := r.db.Create(&classrooms)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *classroomsRepository) Find(id int) (models.Classrooms, error) {
	classroom := models.Classrooms{}
	if err := r.db.First(&classroom, id).Error; err != nil {
		return classroom, err
	}
	return classroom, nil
}

func (r *classroomsRepository) Update(id int, classroom models.Classrooms) error {
	query := r.db.Model(&classroom).Updates(classroom)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *classroomsRepository) Delete(id int) error {
	classroom := models.Classrooms{}
	if err := r.db.Delete(&classroom, id).Error; err != nil {
		return err
	}
	return nil
}
