package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type TeacherAttsRepository interface {
	Browse(page, limit int, search string) ([]models.TeacherAttendances, int64, error)
	Create(teacherAtts models.TeacherAttendances) error
	CreateBatch(teacherAtts []models.TeacherAttendances) error
	Find(id int) (models.TeacherAttendances, error)
	Update(id int, teacherAtt models.TeacherAttendances) error
	Delete(id int) error
}

type teacherAttsRepository struct {
	db *gorm.DB
}

func NewTeacherAttsRepository(db *gorm.DB) *teacherAttsRepository {
	return &teacherAttsRepository{db}
}

func (r *teacherAttsRepository) Browse(page, limit int, search string) ([]models.TeacherAttendances, int64, error) {
	var teacherAtts []models.TeacherAttendances
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(display_name) LIKE ?", search).Limit(limit).Offset(offset).Preload("Teacher").Preload("WorkingSchedule").Find(&teacherAtts).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.TeacherAttendances{}).Where("LOWER(display_name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return teacherAtts, total, nil
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
