package repository

import (
	"eis-be/models"

	"gorm.io/gorm"
)

type CurriculumSubjectsRepository interface {
	GetByCurriculumSubjectID(curriculumID, subjectID uint) (models.CurriculumSubjects, error)
}

type curriculumSubjectsRepository struct {
	db *gorm.DB
}

func NewCurriculumSubjectsRepository(db *gorm.DB) *curriculumSubjectsRepository {
	return &curriculumSubjectsRepository{db}
}

func (r *curriculumSubjectsRepository) GetByCurriculumSubjectID(curriculumID, subjectID uint) (models.CurriculumSubjects, error) {
	curriculumSubject := models.CurriculumSubjects{}
	if err := r.db.Where("curriculum_id = ? AND subject_id = ?", curriculumID, subjectID).First(&curriculumSubject).Error; err != nil {
		return curriculumSubject, err
	}
	return curriculumSubject, nil
}
