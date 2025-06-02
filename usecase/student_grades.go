package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"

	"github.com/go-playground/validator/v10"
)

type StudentGradesUsecase interface {
	Create(studentGrade dto.CreateStudentGradesRequest) error
}

type studentGradesUsecase struct {
	studentGradesRepository repository.StudentGradesRepository
	academicsRepository     repository.AcademicsRepository
	studentsRepository      repository.StudentsRepository
	subjectsRepository      repository.SubjectsRepository
}

func NewStudentGradesUsecase(studentGradesRepo repository.StudentGradesRepository, academicsRepo repository.AcademicsRepository, studentsRepo repository.StudentsRepository, subjectsRepo repository.SubjectsRepository) *studentGradesUsecase {
	return &studentGradesUsecase{
		studentGradesRepository: studentGradesRepo,
		academicsRepository:     academicsRepo,
		studentsRepository:      studentsRepo,
		subjectsRepository:      subjectsRepo,
	}
}

func validateCreateStudentGradesRequest(req dto.CreateStudentGradesRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *studentGradesUsecase) Create(studentGrade dto.CreateStudentGradesRequest) error {
	e := validateCreateStudentGradesRequest(studentGrade)
	if e != nil {
		return e
	}

	academic, err := s.academicsRepository.Find(int(studentGrade.AcademicID))
	if err != nil {
		return err
	}

	student, err := s.studentsRepository.Find(int(studentGrade.StudentID))
	if err != nil {
		return err
	}

	subject, err := s.subjectsRepository.Find(int(studentGrade.SubjectID))
	if err != nil {
		return err
	}

	studentGradesData := models.StudentGrades{
		DisplayName: academic.DisplayName + " - " + subject.Name + " - " + student.FullName,
		AcademicID:  studentGrade.AcademicID,
		StudentID:   studentGrade.StudentID,
		SubjectID:   studentGrade.SubjectID,
		Quiz:        studentGrade.Quiz,
		FirstMonth:  studentGrade.FirstMonth,
		SecondMonth: studentGrade.SecondMonth,
		Midterm:     studentGrade.Midterm,
		Finals:      studentGrade.Finals,
		Remarks:     studentGrade.Remarks,
	}

	err = s.studentGradesRepository.Create(studentGradesData)

	if err != nil {
		return err
	}

	return nil
}
