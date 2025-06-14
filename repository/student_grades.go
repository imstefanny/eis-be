package repository

import (
	"eis-be/dto"
	"eis-be/models"

	"gorm.io/gorm"
)

type StudentGradesRepository interface {
	GetAll(termID int) ([]models.StudentGrades, error)
	Create(studentGrades []models.StudentGrades) error
	UpdateByTermID(studentGrades, newStudents []models.StudentGrades) error
	GetReport(startYear, endYear string, levelID, academicID int) ([]dto.StudentGradesReportQuery, error)

	// Students specific methods
	GetStudentScoreByStudent(studentID, termID int) ([]models.StudentGrades, error)
}

type studentGradesRepository struct {
	db *gorm.DB
}

func NewStudentGradesRepository(db *gorm.DB) *studentGradesRepository {
	return &studentGradesRepository{db}
}

func (r *studentGradesRepository) GetAll(termID int) ([]models.StudentGrades, error) {
	var studentGrades []models.StudentGrades
	if err := r.db.Where("term_id = ?", termID).
		Preload("Academic").
		Preload("Term").
		Preload("Student").
		Preload("Subject").
		Find(&studentGrades).Error; err != nil {
		return nil, err
	}
	return studentGrades, nil
}

func (r *studentGradesRepository) Create(studentGrades []models.StudentGrades) error {
	err := r.db.Create(&studentGrades)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *studentGradesRepository) UpdateByTermID(studentGrades, newStudents []models.StudentGrades) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if len(newStudents) > 0 {
			if err := tx.Create(&newStudents).Error; err != nil {
				return err
			}
		}
		if len(studentGrades) > 0 {
			if err := tx.Save(&studentGrades).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *studentGradesRepository) GetReport(startYear, endYear string, levelID, academicID int) ([]dto.StudentGradesReportQuery, error) {
	var studentGrades []dto.StudentGradesReportQuery

	query := r.db.Table("student_grades").
		Select(`
			students.id AS student_id,
			students.full_name AS student,
			students.nis,
			classrooms.id AS class_id,
			classrooms.display_name AS class,
			ROUND(AVG(student_grades.final_grade), 2) AS finals
		`).
		Joins("JOIN academics ON academics.id = student_grades.academic_id").
		Joins("JOIN students ON students.id = student_grades.student_id").
		Joins("JOIN classrooms ON classrooms.id = academics.classroom_id").
		Where("academics.start_year = ? AND academics.end_year = ?", startYear, endYear).
		Group("students.id, academics.id").
		Order("academics.id, finals DESC")

	if levelID > 0 {
		query = query.Where("classrooms.level_id = ?", levelID)
	}
	if academicID > 0 {
		query = query.Where("student_grades.academic_id = ?", academicID)
	}
	if err := query.Find(&studentGrades).Error; err != nil {
		return []dto.StudentGradesReportQuery{}, err
	}

	return studentGrades, nil
}

// Students specific methods
func (r *studentGradesRepository) GetStudentScoreByStudent(studentID, termID int) ([]models.StudentGrades, error) {
	// result := []dto.StudentScoreResponse{}
	// rawSql := `
	// 	SELECT
	// 		subjects.name as subject_name,
	// 		student_grades.first_month,
	// 		student_grades.second_month,
	// 		student_grades.first_quiz,
	// 		student_grades.second_quiz,
	// 		student_grades.finals
	// 	FROM student_grades
	// 	JOIN students ON students.id = student_grades.student_id
	// 	JOIN academics ON academics.id = students.current_academic_id AND student_grades.academic_id = academics.id
	// 	JOIN subject_schedules ON subject_schedules.academic_id = students.current_academic_id AND subject_schedules.subject_id = student_grades.subject_id
	// 	JOIN subjects ON subjects.id = subject_schedules.subject_id
	// 	WHERE students.user_id = ?
	// 	GROUP BY subject_schedules.subject_id, student_grades.first_month, student_grades.second_month, student_grades.first_quiz, student_grades.second_quiz, student_grades.finals
	// `
	// if err := r.db.Raw(rawSql, userID).Scan(&result).Error; err != nil {
	// 	return nil, err
	// }
	studentGrades := []models.StudentGrades{}
	if err := r.db.Preload("Subject").Where("student_id = ? AND term_id = ?", studentID, termID).Order("subject_id").Find(&studentGrades).Error; err != nil {
		return nil, err
	}
	return studentGrades, nil
}
