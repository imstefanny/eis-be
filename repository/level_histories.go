package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type LevelHistoriesRepository interface {
	Create(levelID uint, levelHistories models.LevelHistories) error
}

type levelHistoriesRepository struct {
	db *gorm.DB
}

func NewLevelHistoriesRepository(db *gorm.DB) *levelHistoriesRepository {
	return &levelHistoriesRepository{db}
}

func (r *levelHistoriesRepository) Create(levelID uint, levelHistories models.LevelHistories) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.LevelHistories{}, "level_id = ?", levelID).Error; err != nil {
			return err
		}
		if err := tx.Create(&levelHistories).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
