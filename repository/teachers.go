package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type TeachersRepository interface {
	Browse(page, limit int, search string) ([]models.Teachers, int64, error)
	Find(id int) (models.Teachers, error)
	GetByMachineID(machineID int) (models.Teachers, error)
	GetByToken(id int) (models.Teachers, error)
	GetAvailableHomeroomTeachers(start_year, end_year string, academic_id int) ([]models.Teachers, error)
	Create(tx *gorm.DB, teachers models.Teachers) error
	Update(id int, teacher models.Teachers) error
	UndeleteTeacher(id int) error
	Delete(id int) error
}

type teachersRepository struct {
	db *gorm.DB
}

func NewTeachersRepository(db *gorm.DB) *teachersRepository {
	return &teachersRepository{db}
}

func (r *teachersRepository) Create(tx *gorm.DB, teachers models.Teachers) error {
	err := tx.Create(&teachers)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *teachersRepository) Browse(page, limit int, search string) ([]models.Teachers, int64, error) {
	var teachers []models.Teachers
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(name) LIKE ?", search).Limit(limit).Offset(offset).Preload("Level").Preload("WorkSched").Preload("User").Unscoped().Find(&teachers).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.Teachers{}).Where("LOWER(name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return teachers, total, nil
}

func (r *teachersRepository) Find(id int) (models.Teachers, error) {
	teacher := models.Teachers{}
	if err := r.db.Preload("Level").Preload("WorkSched").Preload("User").Unscoped().First(&teacher, id).Error; err != nil {
		return teacher, err
	}
	return teacher, nil
}

func (r *teachersRepository) GetByMachineID(machineID int) (models.Teachers, error) {
	teacher := models.Teachers{}
	if err := r.db.Where("machine_id = ?", machineID).Unscoped().First(&teacher).Error; err != nil {
		return teacher, err
	}
	return teacher, nil
}

func (r *teachersRepository) GetByToken(id int) (models.Teachers, error) {
	teacher := models.Teachers{}
	if err := r.db.Where("user_id = ?", id).Preload("Level").Preload("WorkSched").Preload("User").Preload("User.Role").Unscoped().First(&teacher).Error; err != nil {
		return teacher, err
	}
	return teacher, nil
}

func (r *teachersRepository) Update(id int, teacher models.Teachers) error {
	query := r.db.Model(&teacher).Updates(teacher)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *teachersRepository) UndeleteTeacher(id int) error {
	result := r.db.Model(&models.Teachers{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *teachersRepository) Delete(id int) error {
	teacher := models.Teachers{}
	if err := r.db.Delete(&teacher, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *teachersRepository) GetAvailableHomeroomTeachers(start_year, end_year string, academic_id int) ([]models.Teachers, error) {
	var teachers []models.Teachers
	var rawSQL = `
		SELECT 
			teachers.id,
			teachers.name,
			teachers.nuptk,
			users.email,
			users.role_id,
			acd.id as academic_id,
			acd.display_name
		FROM teachers
		INNER JOIN users ON users.id = teachers.user_id
		LEFT JOIN academics acd ON acd.homeroom_teacher_id = teachers.id AND start_year = ? AND end_year = ?
		WHERE role_id != 2 AND role_id != 1 AND role_id != 5
		AND acd.display_name IS NULL OR acd.id = ?
	`
	if err := r.db.Raw(rawSQL, start_year, end_year, academic_id).Scan(&teachers).Error; err != nil {
		return nil, err
	}

	return teachers, nil
}
