package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type TeacherAttsRepository interface {
	Browse(page, limit int, search, date string, userId *int) ([]models.TeacherAttendances, int64, error)
	Create(teacherAtts models.TeacherAttendances) error
	CreateBatch(teacherAtts []models.TeacherAttendances) error
	Find(id int) (models.TeacherAttendances, error)
	Update(id int, teacherAtt models.TeacherAttendances) error
	Delete(id int) error
	BrowseReport(search, startDate, endDate string, userId *int) ([]models.TeacherAttendances, error)
	FindByTeacherIdDate(teacherId int, date string) (models.TeacherAttendances, error)
}

type teacherAttsRepository struct {
	db *gorm.DB
}

func NewTeacherAttsRepository(db *gorm.DB) *teacherAttsRepository {
	return &teacherAttsRepository{db}
}

func (r *teacherAttsRepository) Browse(page, limit int, search, date string, userId *int) ([]models.TeacherAttendances, int64, error) {
	var teacherAtts []models.TeacherAttendances
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	query := r.db.
		Where("LOWER(display_name) LIKE ?", search).
		Limit(limit).
		Offset(offset).
		Preload("Teacher").
		Preload("WorkingSchedule").
		Preload("WorkingSchedule.Details")

	if date != "" {
		query = query.Where("DATE(date) = ?", date)
	}
	if userId != nil {
		query = query.
			Joins("JOIN teachers ON teachers.id = teacher_attendances.teacher_id").
			Where("teachers.user_id = ?", &userId)
	}

	if err := query.Find(&teacherAtts).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Model(&models.TeacherAttendances{}).Where("LOWER(display_name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return teacherAtts, total, nil
}

func (r *teacherAttsRepository) BrowseReport(search, startDate, endDate string, userId *int) ([]models.TeacherAttendances, error) {
	var teacherAtts []models.TeacherAttendances
	search = "%" + strings.ToLower(search) + "%"
	query := r.db.
		Where("LOWER(display_name) LIKE ?", search).
		Preload("Teacher").
		Preload("WorkingSchedule").
		Preload("WorkingSchedule.Details")

	if startDate != "" && endDate != "" {
		query = query.Where("DATE(date) BETWEEN ? AND ?", startDate, endDate)
	}
	if userId != nil {
		query = query.
			Joins("JOIN teachers ON teachers.id = teacher_attendances.teacher_id").
			Where("teachers.user_id = ?", &userId)
	}
	if err := query.Find(&teacherAtts).Error; err != nil {
		return nil, err
	}

	return teacherAtts, nil
}

func (r *teacherAttsRepository) Create(teacherAtts models.TeacherAttendances) error {
	err := r.db.Create(&teacherAtts)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *teacherAttsRepository) CreateBatch(teacherAtts []models.TeacherAttendances) error {
	err := r.db.Create(&teacherAtts)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *teacherAttsRepository) Find(id int) (models.TeacherAttendances, error) {
	teacherAtt := models.TeacherAttendances{}
	if err := r.db.Preload("Teacher").Preload("WorkingSchedule").First(&teacherAtt, id).Error; err != nil {
		return teacherAtt, err
	}
	return teacherAtt, nil
}

func (r *teacherAttsRepository) Update(id int, teacherAtt models.TeacherAttendances) error {
	query := r.db.Model(&teacherAtt).Updates(teacherAtt)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *teacherAttsRepository) Delete(id int) error {
	teacherAtt := models.TeacherAttendances{}
	if err := r.db.Delete(&teacherAtt, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *teacherAttsRepository) FindByTeacherIdDate(teacherId int, date string) (models.TeacherAttendances, error) {
	teacherAtt := models.TeacherAttendances{}
	if err := r.db.Where("teacher_id = ? AND DATE(date) = ?", teacherId, date).First(&teacherAtt).Error; err != nil {
		return teacherAtt, err
	}
	return teacherAtt, nil
}
