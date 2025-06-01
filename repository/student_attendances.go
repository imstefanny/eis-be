package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type StudentAttsRepository interface {
	// Browse(page, limit int, search string) ([]models.StudentAttendances, int64, error)
	BrowseByAcademicID(academicID, page, limit int, search string, date string) ([]models.StudentAttendances, int64, error)
	// Create(studentAtts models.StudentAttendances) error
	CreateBatch(studentAtts []models.StudentAttendances) error
	FindByAcademicDate(academicID int, date string) ([]models.StudentAttendances, error)
	// Find(id int) (models.StudentAttendances, error)
	UpdateByAcademicID(academicID int, studentAtt []models.StudentAttendances) error
	// Delete(id int) error
}

type studentAttsRepository struct {
	db *gorm.DB
}

func NewStudentAttsRepository(db *gorm.DB) *studentAttsRepository {
	return &studentAttsRepository{db}
}

// func (r *studentAttsRepository) Browse(page, limit int, search string) ([]models.StudentAttendances, int64, error) {
// 	var studentAtts []models.StudentAttendances
// 	var total int64
// 	offset := (page - 1) * limit
// 	search = "%" + strings.ToLower(search) + "%"
// 	if err := r.db.Where("LOWER(display_name) LIKE ?", search).Limit(limit).Offset(offset).Preload("Academic").Preload("Details").Find(&studentAtts).Error; err != nil {
// 		return nil, 0, err
// 	}
// 	if err := r.db.Model(&models.StudentAttendances{}).Where("LOWER(display_name) LIKE ?", search).Count(&total).Error; err != nil {
// 		return nil, 0, err
// 	}
// 	return studentAtts, total, nil
// }

func (r *studentAttsRepository) BrowseByAcademicID(academicID, page, limit int, search string, date string) ([]models.StudentAttendances, int64, error) {
	var studentAtts []models.StudentAttendances
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	query := r.db.
		Where("academic_id = ?", academicID).
		Limit(limit).
		Offset(offset).
		Preload("Academic").
		Preload("Student")
	if search != "" {
		query = query.Where("LOWER(display_name) LIKE ?", search)
	}
	if date != "" {
		query = query.Where("DATE(date) = ?", date)
	}

	if err := query.Find(&studentAtts).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.StudentAttendances{}).Where("academic_id = ? AND LOWER(display_name) LIKE ?", academicID, search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return studentAtts, total, nil
}

func (r *studentAttsRepository) FindByAcademicDate(academicID int, date string) ([]models.StudentAttendances, error) {
	var studentAtts []models.StudentAttendances
	query := r.db.Where("academic_id = ?", academicID).Where("DATE(date) = ?", date).
		Preload("Academic").
		Preload("Student")
	if err := query.Find(&studentAtts).Error; err != nil {
		return nil, err
	}
	if len(studentAtts) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return studentAtts, nil
}

// func (r *studentAttsRepository) Create(studentAtts models.StudentAttendances) error {
// 	err := r.db.Create(&studentAtts)
// 	if err.Error != nil {
// 		return err.Error
// 	}
// 	return nil
// }

func (r *studentAttsRepository) CreateBatch(studentAtts []models.StudentAttendances) error {
	err := r.db.Create(&studentAtts)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// func (r *studentAttsRepository) Find(id int) (models.StudentAttendances, error) {
// 	studentAtt := models.StudentAttendances{}
// 	if err := r.db.Preload("Academic").
// 		Preload("Details").
// 		Preload("Details.SubjSched.Teacher").
// 		Preload("Details.SubjSched.Subject").
// 		Preload("Details.Teacher").
// 		First(&studentAtt, id).Error; err != nil {
// 		return studentAtt, err
// 	}
// 	return studentAtt, nil
// }

func (r *studentAttsRepository) UpdateByAcademicID(academicID int, studentAtts []models.StudentAttendances) error {
	query := r.db.Save(studentAtts)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

// func (r *studentAttsRepository) Update(id int, params map[string]interface{}) error {
// 	studentAtt := models.StudentAttendancesDetails{}
// 	err := r.db.Transaction(func(tx *gorm.DB) error {
// 		addIDs := params["addIDs"].([]models.StudentAttendancesDetails)
// 		if len(addIDs) > 0 {
// 			if err := tx.Create(&addIDs).Error; err != nil {
// 				return err
// 			}
// 		}
// 		if len(params["updateIDs"].([]int)) > 0 {
// 			if err := tx.Save(params["incomingUpdate"]).Error; err != nil {
// 				return err
// 			}
// 		}
// 		if len(params["removeIDs"].([]int)) > 0 {
// 			if err := tx.Unscoped().Delete(&studentAtt, params["removeIDs"]).Error; err != nil {
// 				return err
// 			}
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *studentAttsRepository) Delete(id int) error {
// 	studentAtt := models.StudentAttendances{}
// 	studentAttDetails := models.StudentAttendancesDetails{}
// 	err := r.db.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Where("note_id = ?", id).Delete(&studentAttDetails).Error; err != nil {
// 			return err
// 		}
// 		if err := tx.Where("id = ?", id).Delete(&studentAtt).Error; err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
