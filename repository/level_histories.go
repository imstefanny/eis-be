package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type LevelHistoriesRepository interface {
	Browse(page, limit int, search string) ([]models.LevelHistories, int64, error)
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

func (r *levelHistoriesRepository) Browse(page, limit int, search string) ([]models.LevelHistories, int64, error) {
	var levelHistories []models.LevelHistories
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(op_cert_num) LIKE ?", search).Limit(limit).Offset(offset).Find(&levelHistories).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.LevelHistories{}).Where("LOWER(op_cert_num) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return levelHistories, total, nil
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
