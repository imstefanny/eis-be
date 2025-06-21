package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type TermsRepository interface {
	Find(id int) (models.Terms, error)
}

type termsRepository struct {
	db *gorm.DB
}

func NewTermsRepository(db *gorm.DB) *termsRepository {
	return &termsRepository{db}
}

func (r *termsRepository) Find(id int) (models.Terms, error) {
	term := models.Terms{}
	if err := r.db.
		Preload("Academic").
		Preload("Academic").
		Preload("Academic.Classroom").
		Preload("Academic.Classroom.Level").
		Preload("Academic.HomeroomTeacher").
		Preload("Academic.Classroom.Level.Histories.Principle").
		First(&term, id).Error; err != nil {
		return term, err
	}
	return term, nil
}
