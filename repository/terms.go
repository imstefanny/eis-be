package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type TermsRepository interface {
	Find(id int) (models.Terms, error)
	Update(term models.Terms) error
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

func (r *termsRepository) Update(term models.Terms) error {
	updateData := map[string]interface{}{}

	if !term.FirstStartDate.IsZero() {
		updateData["first_start_date"] = term.FirstStartDate
	}
	if !term.FirstEndDate.IsZero() {
		updateData["first_end_date"] = term.FirstEndDate
	}
	if !term.SecondStartDate.IsZero() {
		updateData["second_start_date"] = term.SecondStartDate
	}
	if !term.SecondEndDate.IsZero() {
		updateData["second_end_date"] = term.SecondEndDate
	}

	if len(updateData) == 0 {
		return nil
	}
	if err := r.db.Model(&term).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}
