package repository

import (
	"eis-be/dto"
	"eis-be/models"
	"strings"

	"gorm.io/gorm"
)

type StudentsRepository interface {
	Browse(page, limit int, search string) ([]models.Students, int64, error)
	Create(students models.Students) (uint, error)
	GetByIds(ids []int) ([]models.Students, error)
	GetByToken(id int) (models.Students, error)
	GetStudentScoreByUserId(userId int) (interface{}, error)
	Find(id int) (models.Students, error)
	Update(id int, student models.Students) error
	UpdateStudentAcademicId(academic_id int, student []uint) error
	Undelete(id int) error
	Delete(id int) error
}

type studentsRepository struct {
	db *gorm.DB
}

func NewStudentsRepository(db *gorm.DB) *studentsRepository {
	return &studentsRepository{db}
}

func (r *studentsRepository) Browse(page, limit int, search string) ([]models.Students, int64, error) {
	var students []models.Students
	var total int64
	offset := (page - 1) * limit
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(full_name) LIKE ?", search).Limit(limit).Offset(offset).Preload("Applicant").Preload("User").Preload("Guardians").Preload("Documents").Unscoped().Find(&students).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Model(&models.Students{}).Where("LOWER(full_name) LIKE ?", search).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return students, total, nil
}

func (r *studentsRepository) Create(students models.Students) (uint, error) {
	result := r.db.Create(&students)
	if result.Error != nil {
		return 0, result.Error
	}
	return students.ID, nil
}

func (r *studentsRepository) Find(id int) (models.Students, error) {
	student := models.Students{}
	if err := r.db.Preload("Applicant").Preload("User").Preload("Guardians").Preload("Documents").First(&student, id).Error; err != nil {
		return student, err
	}
	return student, nil
}

func (r *studentsRepository) GetByToken(id int) (models.Students, error) {
	student := models.Students{}
	if err := r.db.Where("user_id = ?", id).
		Preload("Guardians").
		Preload("Academics").
		Preload("Academics.Classroom").
		Preload("Academics.HomeroomTeacher").
		Preload("User").
		Unscoped().
		First(&student).Error; err != nil {
		return student, err
	}
	return student, nil
}

func (r *studentsRepository) GetStudentScoreByUserId(userId int) (interface{}, error) {
	result := []dto.StudentScoreResponse{}
	rawSql := `
		SELECT
			subjects.name as subject_name,
			student_grades.first_month,
			student_grades.second_month,
			student_grades.first_quiz,
			student_grades.second_quiz,
			student_grades.finals
		FROM student_grades 
		JOIN students ON students.id = student_grades.student_id
		JOIN academics ON academics.id = students.current_academic_id AND student_grades.academic_id = academics.id
		JOIN subject_schedules ON subject_schedules.academic_id = students.current_academic_id AND subject_schedules.subject_id = student_grades.subject_id
		JOIN subjects ON subjects.id = subject_schedules.subject_id
		WHERE students.user_id = ?
		GROUP BY subject_schedules.subject_id, student_grades.first_month, student_grades.second_month, student_grades.first_quiz, student_grades.second_quiz, student_grades.finals
	`
	if err := r.db.Raw(rawSql, userId).Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *studentsRepository) GetByIds(ids []int) ([]models.Students, error) {
	students := []models.Students{}
	if err := r.db.Where("id IN ?", ids).Find(&students).Error; err != nil {
		return students, err
	}
	return students, nil
}

func (r *studentsRepository) Update(id int, student models.Students) error {
	query := r.db.Model(&student).Updates(student)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (r *studentsRepository) UpdateStudentAcademicId(academic_id int, studentIDs []uint) error {
	var academic models.Academics
	if err := r.db.First(&academic, academic_id).Error; err != nil {
		return err
	}
	var students []models.Students
	if err := r.db.Where("id IN ?", studentIDs).Find(&students).Error; err != nil {
		return err
	}
	if err := r.db.Model(&models.Students{}).Where("id IN ?", studentIDs).Update("current_academic_id", academic_id).Error; err != nil {
		return err
	}
	if err := r.db.Model(&academic).Association("Students").Append(students); err != nil {
		return err
	}

	return nil
}

func (r *studentsRepository) Undelete(id int) error {
	result := r.db.Model(&models.Students{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *studentsRepository) Delete(id int) error {
	student := models.Students{}
	if err := r.db.Delete(&student, id).Error; err != nil {
		return err
	}
	return nil
}
