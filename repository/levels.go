package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type LevelsRepository interface {
	Browse(page, limit int, search string) ([]models.Levels, int64, error)
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

func (r *levelsRepository) Browse(page, limit int, search string) ([]models.Levels, int64, error) {
	var levels []models.Levels
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(name) LIKE ?", search).Limit(limit).Offset(offset).Find(&levels).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.Levels{}).Where("LOWER(name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return levels, total, nil
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
