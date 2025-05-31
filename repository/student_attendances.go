package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type StudentAttsRepository interface {
	// Browse(page, limit int, search string) ([]models.StudentAttendances, int64, error)
	// BrowseByAcademicID(academicID, page, limit int, search string) ([]models.StudentAttendances, int64, error)
	// Create(studentAtts models.StudentAttendances) error
	CreateBatch(studentAtts []models.StudentAttendances) error
	// Find(id int) (models.StudentAttendances, error)
	// Update(id int, params map[string]interface{}) error
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

// func (r *studentAttsRepository) BrowseByAcademicID(academicID, page, limit int, search string) ([]models.StudentAttendances, int64, error) {
// 	var studentAtts []models.StudentAttendances
// 	var total int64
// 	offset := (page - 1) * limit
// 	search = "%" + strings.ToLower(search) + "%"
// 	if err := r.db.Where("academic_id = ? AND LOWER(display_name) LIKE ?", academicID, search).
// 		Limit(limit).
// 		Offset(offset).
// 		Preload("Academic").
// 		Preload("Details").
// 		Preload("Details.SubjSched.Teacher").
// 		Preload("Details.SubjSched.Subject").
// 		Preload("Details.Teacher").
// 		Find(&studentAtts).Error; err != nil {
// 		return nil, 0, err
// 	}
// 	if err := r.db.Model(&models.StudentAttendances{}).Where("academic_id = ? AND LOWER(display_name) LIKE ?", academicID, search).Count(&total).Error; err != nil {
// 		return nil, 0, err
// 	}
// 	return studentAtts, total, nil
// }

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
