package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type StudentBehaviourActivitiesRepository interface {
	GetByAcademicIdAndTermId(academicID, termID int) ([]models.StudentBehaviourActivities, error)
	Create(studentBehaviour []models.StudentBehaviourActivities) error
	// UpdateByTermID(studentBehaviour, newStudents []models.StudentBehaviourActivities) error
}

type studentBehaviourActivitiesRepository struct {
	db *gorm.DB
}

func NewStudentBehaviourActivitiesRepository(db *gorm.DB) *studentBehaviourActivitiesRepository {
	return &studentBehaviourActivitiesRepository{db}
}

func (r *studentBehaviourActivitiesRepository) GetByAcademicIdAndTermId(academicID, termID int) ([]models.StudentBehaviourActivities, error) {
	var studentBehaviour []models.StudentBehaviourActivities
	if err := r.db.Where("term_id = ?", termID).
		Where("academic_id = ?", academicID).
		Preload("Academic").
		Preload("Term").
		Preload("Student").
		Find(&studentBehaviour).Error; err != nil {
		return nil, err
	}
	return studentBehaviour, nil
}

func (r *studentBehaviourActivitiesRepository) Create(studentBehaviour []models.StudentBehaviourActivities) error {
	err := r.db.Create(&studentBehaviour)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
