package repository

import (
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type AcademicsRepository interface {
	Browse(page, limit int, search, startYear, endYear string) ([]models.Academics, int64, error)
	BrowseByTeacherId(page, limit int, search, startYear, endYear string, teacherId int) ([]models.Academics, int64, error)
	GetAll() ([]models.Academics, error)
	Create(academic models.Academics) error
	CreateBatch(academics []models.Academics) error
	Find(id int) (models.Academics, error)
	Update(id int, academic models.Academics) error
	Delete(id int) error
	UpdateNewCurriculum(levelID int, grade string, curriculumID uint) error

	// Students specific methods
	GetAcademicsByStudent(studentID int) ([]models.Academics, error)
}

type academicsRepository struct {
	db *gorm.DB
}

func NewAcademicsRepository(db *gorm.DB) *academicsRepository {
	return &academicsRepository{db}
}

func (r *academicsRepository) Browse(page, limit int, search, startYear, endYear string) ([]models.Academics, int64, error) {
	var academics []models.Academics
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"

	query := r.db.Model(&models.Academics{})
	query = query.Where("LOWER(display_name) LIKE ?", search)

	if startYear != "" && endYear != "" {
		query = query.Where("start_year = ? AND end_year = ?", startYear, endYear)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).
		Offset(offset).
		Preload("Classroom").
		Preload("Classroom.Level").
		Preload("HomeroomTeacher").
		Preload("Students").
		Preload("Terms").
		Preload("Curriculum").
		Unscoped().
		Find(&academics).Error; err != nil {
		return nil, 0, err
	}
	return academics, total, nil
}

func (r *academicsRepository) BrowseByTeacherId(page, limit int, search, startYear, endYear string, teacherId int) ([]models.Academics, int64, error) {
	var academics []models.Academics
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"

	query := r.db.Debug().Model(&models.Academics{}).
		Joins("LEFT JOIN subject_schedules ON subject_schedules.academic_id = academics.id").
		Joins("LEFT JOIN teachers ON teachers.id = subject_schedules.teacher_id").
		Where("LOWER(academics.display_name) LIKE ?", search).
		Where("teachers.id = ? OR academics.homeroom_teacher_id = ?", teacherId, teacherId)

	if startYear != "" && endYear != "" {
		query = query.Where("academics.start_year = ? AND academics.end_year = ?", startYear, endYear)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).
		Offset(offset).
		Preload("Classroom").
		Preload("Classroom.Level").
		Preload("HomeroomTeacher").
		Preload("Students").
		Preload("Terms").
		Preload("Curriculum").
		Unscoped().
		Distinct().
		Find(&academics).Error; err != nil {
		return nil, 0, err
	}
	return academics, total, nil
}

func (r *academicsRepository) GetAll() ([]models.Academics, error) {
	var academics []models.Academics
	if err := r.db.Preload("SubjScheds").Preload("Students").Find(&academics).Error; err != nil {
		return nil, err
	}
	return academics, nil
}

func (r *academicsRepository) Create(academic models.Academics) error {
	err := r.db.Create(&academic)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *academicsRepository) CreateBatch(academics []models.Academics) error {
	if len(academics) == 0 {
		return nil
	}
	err := r.db.Create(&academics)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *academicsRepository) Find(id int) (models.Academics, error) {
	academic := models.Academics{}
	if err := r.db.Preload("Classroom").
		Preload("Classroom.Level").
		Preload("HomeroomTeacher").
		Preload("Terms").
		Preload("Students").
		Preload("SubjScheds").
		Preload("SubjScheds.Teacher").
		Unscoped().
		Preload("SubjScheds.Subject").
		Preload("ClassNotes").
		Preload("ClassNotes.Details").
		Preload("ClassNotes.Details.Teacher").
		Preload("ClassNotes.Details.SubjSched.Teacher").
		Preload("ClassNotes.Details.SubjSched.Subject").
		Preload("Curriculum").
		Preload("Curriculum.CurriculumSubjects").
		Preload("Curriculum.CurriculumSubjects.Subject").
		First(&academic, id).Error; err != nil {
		return academic, err
	}
	return academic, nil
}

func (r *academicsRepository) Update(id int, academic models.Academics) error {
	oldAcademic := models.Academics{}
	if e := r.db.Find(&oldAcademic, id).Error; e != nil {
		return e
	}
	r.db.Model(&oldAcademic).Association("Students").Clear()
	query := r.db.Model(&academic).Updates(academic)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *academicsRepository) Delete(id int) error {
	academic := models.Academics{}
	if err := r.db.Delete(&academic, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *academicsRepository) GetAcademicsByStudent(studentID int) ([]models.Academics, error) {
	academics := []models.Academics{}
	if err := r.db.Table("academics").
		Select("academics.*").
		Joins("JOIN academic_students ON academic_students.academics_id = academics.id").
		Where("students_id = ?", studentID).
		Preload("Terms").
		Find(&academics).Error; err != nil {
		return nil, err
	}
	return academics, nil
}

func (r *academicsRepository) UpdateNewCurriculum(levelID int, grade string, curriculumID uint) error {
	ids := []int{}
	r.db.Table("academics").
		Joins("JOIN classrooms ON classrooms.id = academics.classroom_id").
		Where("classrooms.level_id = ? AND classrooms.grade = ?", levelID, grade).
		Where("curriculum_id IS NULL").
		Pluck("academics.id", &ids)
	if len(ids) > 0 {
		if err := r.db.Model(&models.Academics{}).Where("id IN ?", ids).Update("curriculum_id", curriculumID).Error; err != nil {
			return err
		}
	}
	return nil
}
