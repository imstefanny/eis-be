package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type StudentBehaviourActivitiesRepository interface {
	GetByAcademicIdAndTermId(academicID, termID int) ([]models.StudentBehaviourActivities, error)
	Create(studentBehaviour []models.StudentBehaviourActivities) error
	Update(studentBehaviour models.StudentBehaviourActivities) error
	FindByStudentIDAndAcademicIDAndTermID(studentID, academicID, termID uint) (*models.StudentBehaviourActivities, error)
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

func (r *studentBehaviourActivitiesRepository) Update(studentBehaviour models.StudentBehaviourActivities) error {
	err := r.db.Save(&studentBehaviour)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *studentBehaviourActivitiesRepository) FindByStudentIDAndAcademicIDAndTermID(studentID, academicID, termID uint) (*models.StudentBehaviourActivities, error) {
	var studentBehaviour models.StudentBehaviourActivities
	if err := r.db.Where("student_id = ?", studentID).
		Where("academic_id = ?", academicID).
		Where("term_id = ?", termID).
		Preload("Academic").
		Preload("Term").
		Preload("Student").
		First(&studentBehaviour).Error; err != nil {
		return nil, err
	}
	return &studentBehaviour, nil
}
