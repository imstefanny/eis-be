package repository

import (
	"eis-be/models"
	"fmt"

	"gorm.io/gorm"
)

type SubjectsRepository interface {
	Browse(page, limit int, search, sortColumn, sortOrder string, isExtracurricular *bool) ([]models.Subjects, int64, error)
	Create(subjects models.Subjects) error
	Find(id int) (models.Subjects, error)
	FindByCode(code string) models.Subjects
	Update(id int, subject models.Subjects) error
	Delete(id int) error
}

type subjectsRepository struct {
	db *gorm.DB
}

func NewSubjectsRepository(db *gorm.DB) *subjectsRepository {
	return &subjectsRepository{db}
}

func (r *subjectsRepository) Browse(page, limit int, search, sortColumn, sortOrder string, isExtracurricular *bool) ([]models.Subjects, int64, error) {
	var subjects []models.Subjects
	var total int64
	offset := (page - 1) * limit

	allowedColumns := map[string]bool{
		"id": true, "display_name": true, "code": true, "created_at": true,
	}
	if !allowedColumns[sortColumn] {
		sortColumn = "created_at"
	}

	orderClause := fmt.Sprintf("%s %s", sortColumn, sortOrder)

	query := r.db.Model(&models.Subjects{})
	if search != "" {
		query = query.Where("display_name LIKE ?", "%"+search+"%")
	}
	if isExtracurricular != nil {
		query = query.Where("is_extracurricular = ?", *isExtracurricular)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Order(orderClause).Limit(limit).Offset(offset).Find(&subjects).Error; err != nil {
		return nil, 0, err
	}

	return subjects, total, nil
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

func (r *subjectsRepository) FindByCode(code string) models.Subjects {
	subject := models.Subjects{}
	if err := r.db.Where("code = ?", code).First(&subject).Error; err != nil {
		return subject
	}
	return subject
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
