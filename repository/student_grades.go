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
	GetReport(startYear, endYear string, levelID, academicID, termID int) ([]dto.StudentGradesReportQuery, error)
	GetMonthlyReportByStudent(academicID, studentID int) ([]dto.GetPrintMonthlyReportGrade, error)

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

func (r *studentGradesRepository) GetMonthlyReportByStudent(academicID, studentID int) ([]dto.GetPrintMonthlyReportGrade, error) {
	studentGrades := []dto.GetPrintMonthlyReportGrade{}
	rawSQL := `
	WITH first AS (
		SELECT
			terms.id,
			subject_id,
			subjects.name,
			student_grades.first_quiz,
			student_grades.second_quiz,
			student_grades.first_month,
			student_grades.second_month
		FROM student_grades
			JOIN subjects ON student_grades.subject_id = subjects.id
			JOIN academics ON student_grades.academic_id = academics.id
			JOIN terms ON academics.id = terms.academic_id
		WHERE
			student_grades.student_id = ? AND terms.name = 'Semester 1' AND academics.id = ?
	), second AS (
		SELECT
			terms.id,
			subject_id,
			subjects.name,
			student_grades.first_quiz,
			student_grades.second_quiz,
			student_grades.first_month,
			student_grades.second_month
		FROM student_grades
			JOIN subjects ON student_grades.subject_id = subjects.id
			JOIN academics ON student_grades.academic_id = academics.id
			JOIN terms ON academics.id = terms.academic_id
		WHERE
			student_grades.student_id = ? AND terms.name = 'Semester 2' AND academics.id = ?
	) SELECT
		first.name AS subject,
		first.first_quiz AS st_first_quiz,
		first.second_quiz AS st_second_quiz,
		first.first_month AS st_first_month,
		first.second_month AS st_second_month,
		second.first_quiz AS nd_first_quiz,
		second.second_quiz AS nd_second_quiz,
		second.first_month AS nd_first_month,
		second.second_month AS nd_second_month
	FROM first JOIN second ON first.subject_id = second.subject_id
	ORDER BY first.subject_id
	`

	if err := r.db.Raw(rawSQL, studentID, academicID, studentID, academicID).Scan(&studentGrades).Error; err != nil {
		return nil, err
	}
	return studentGrades, nil
}

func (r *studentGradesRepository) GetReport(startYear, endYear string, levelID, academicID, termID int) ([]dto.StudentGradesReportQuery, error) {
	var studentGrades []dto.StudentGradesReportQuery

	query := r.db.Table("student_grades").
		Select(`
			students.id AS student_id,
			students.full_name AS student,
			students.nis,
			classrooms.id AS class_id,
    		CONCAT(classrooms.display_name, " : ", terms.name) AS class,
			ROUND(AVG(student_grades.final_grade), 2) AS finals
		`).
		Joins("JOIN terms ON terms.id = student_grades.term_id").
		Joins("JOIN academics ON terms.academic_id = academics.id").
		Joins("JOIN students ON students.id = student_grades.student_id").
		Joins("JOIN classrooms ON classrooms.id = academics.classroom_id").
		Where("academics.start_year = ? AND academics.end_year = ?", startYear, endYear).
		Group("students.id, terms.id").
		Order("terms.id, finals DESC, students.id")

	if levelID > 0 {
		query = query.Where("classrooms.level_id = ?", levelID)
	}
	if academicID > 0 {
		query = query.Where("student_grades.academic_id = ?", academicID)
	}
	if termID > 0 {
		query = query.Where("student_grades.term_id = ?", termID)
	}
	if err := query.Find(&studentGrades).Error; err != nil {
		return []dto.StudentGradesReportQuery{}, err
	}

	return studentGrades, nil
}

// Students specific methods
func (r *studentGradesRepository) GetStudentScoreByStudent(studentID, termID int) ([]models.StudentGrades, error) {
	studentGrades := []models.StudentGrades{}
	if err := r.db.
		Preload("Subject").
		Where("student_id = ? AND term_id = ?", studentID, termID).
		Order("subject_id").
		Find(&studentGrades).Error; err != nil {
		return nil, err
	}
	return studentGrades, nil
}
