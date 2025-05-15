package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type LevelsRepository interface {
	GetAll() ([]models.Levels, error)
	Create(levels models.Levels) error
	Find(id int) (models.Levels, error)
	Update(id int, level models.Levels) error
	Delete(id int) error
}

type levelsRepository struct {
	db *gorm.DB
}

func NewLevelsRepository(db *gorm.DB) *levelsRepository {
	return &levelsRepository{db}
}

func (r *levelsRepository) GetAll() ([]models.Levels, error) {
	levels := []models.Levels{}
	if err := r.db.Find(&levels).Error; err != nil {
		return nil, err
	}
	return levels, nil
}

func (r *levelsRepository) Create(levels models.Levels) error {
	err := r.db.Create(&levels)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *levelsRepository) Find(id int) (models.Levels, error) {
	level := models.Levels{}
	if err := r.db.First(&level, id).Error; err != nil {
		return level, err
	}
	return level, nil
}

func (r *levelsRepository) Update(id int, level models.Levels) error {
	query := r.db.Model(&level).Updates(level)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *levelsRepository) Delete(id int) error {
	level := models.Levels{}
	if err := r.db.Delete(&level, id).Error; err != nil {
		return err
	}
	return nil
}
