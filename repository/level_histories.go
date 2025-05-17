package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type LevelHistoriesRepository interface {
	GetAll() ([]models.LevelHistories, error)
	Create(levelHistories models.LevelHistories) error
	Find(id int) (models.LevelHistories, error)
	Update(id int, levelHistory models.LevelHistories) error
	Delete(id int) error
}

type levelHistoriesRepository struct {
	db *gorm.DB
}

func NewLevelHistoriesRepository(db *gorm.DB) *levelHistoriesRepository {
	return &levelHistoriesRepository{db}
}

func (r *levelHistoriesRepository) GetAll() ([]models.LevelHistories, error) {
	levelHistories := []models.LevelHistories{}
	if err := r.db.Find(&levelHistories).Error; err != nil {
		return nil, err
	}
	return levelHistories, nil
}

func (r *levelHistoriesRepository) Create(levelHistories models.LevelHistories) error {
	err := r.db.Create(&levelHistories)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *levelHistoriesRepository) Find(id int) (models.LevelHistories, error) {
	levelHistory := models.LevelHistories{}
	if err := r.db.First(&levelHistory, id).Error; err != nil {
		return levelHistory, err
	}
	return levelHistory, nil
}

func (r *levelHistoriesRepository) Update(id int, levelHistory models.LevelHistories) error {
	query := r.db.Model(&levelHistory).Updates(levelHistory)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *levelHistoriesRepository) Delete(id int) error {
	levelHistory := models.LevelHistories{}
	if err := r.db.Delete(&levelHistory, id).Error; err != nil {
		return err
	}
	return nil
}
