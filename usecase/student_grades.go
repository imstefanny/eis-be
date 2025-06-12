package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"fmt"
	"math"

	"github.com/go-playground/validator/v10"
)

type StudentGradesUsecase interface {
	GetAll(academicID int) (dto.GetStudentGradesResponse, error)
	Create(studentGrade dto.CreateStudentGradesRequest) error
	UpdateByAcademicID(academicID int, studentGrade dto.UpdateStudentGradesRequest) (dto.GetStudentGradesResponse, error)
	// GetReport(academicYear string, levelID, academicID int) (dto.StudentGradesReport, error)
	GetReport(academicYear string, levelID, academicID int) ([]dto.StudentGradesReport, error)
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

func validateUpdateStudentGradesRequest(req dto.UpdateStudentGradesRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *studentGradesUsecase) GetAll(academicID int) (dto.GetStudentGradesResponse, error) {
	academic, err := s.academicsRepository.Find(academicID)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("academic with ID %d not found: %w", academicID, err)
	}

	studentGrades, err := s.studentGradesRepository.GetAll(academicID)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("error browsing student grades: %w", err)
	}

	details := make(map[uint]*dto.GetStudentGradesDetailResponse)
	for _, grade := range studentGrades {
		if _, exists := details[grade.SubjectID]; !exists {
			details[grade.SubjectID] = &dto.GetStudentGradesDetailResponse{
				SubjectID: grade.SubjectID,
				Subject:   grade.Subject.Name,
				Students:  []dto.GetStudentGradesEntryResponse{},
			}
		}
		details[grade.SubjectID].Students = append(details[grade.SubjectID].Students, dto.GetStudentGradesEntryResponse{
			ID:          grade.ID,
			StudentID:   grade.StudentID,
			StudentName: grade.Student.FullName,
			DisplayName: grade.DisplayName,
			Quiz:        grade.Quiz,
			FirstMonth:  grade.FirstMonth,
			SecondMonth: grade.SecondMonth,
			Finals:      grade.Finals,
			FinalGrade:  grade.FinalGrade,
			Remarks:     grade.Remarks,
		})
	}
	var detailsList []dto.GetStudentGradesDetailResponse
	for _, detail := range details {
		detailsList = append(detailsList, *detail)
	}
	response := dto.GetStudentGradesResponse{
		AcademicID: uint(academicID),
		Academic:   academic.DisplayName,
		Details:    detailsList,
	}

	return response, nil
}

func (s *studentGradesUsecase) Create(studentGrade dto.CreateStudentGradesRequest) error {
	e := validateCreateStudentGradesRequest(studentGrade)
	if e != nil {
		return e
	}

	academic, err := s.academicsRepository.Find(int(studentGrade.AcademicID))
	if err != nil {
		return fmt.Errorf("academic with ID %d not found: %w", studentGrade.AcademicID, err)
	}

	studentGradesData := []models.StudentGrades{}
	for _, detail := range studentGrade.Details {
		subject, err := s.subjectsRepository.Find(int(detail.SubjectID))
		if err != nil {
			return fmt.Errorf("subject with ID %d not found: %w", detail.SubjectID, err)
		}
		for _, entry := range detail.Students {
			student, err := s.studentsRepository.Find(int(entry.StudentID))
			if err != nil {
				return fmt.Errorf("student with ID %d not found: %w", entry.StudentID, err)
			}
			finals := math.Round(((entry.FirstMonth+entry.SecondMonth+entry.Quiz)/2 + entry.Finals) / 2)
			studentGradesData = append(studentGradesData, models.StudentGrades{
				DisplayName: academic.DisplayName + " - " + subject.Name + " - " + student.FullName,
				AcademicID:  studentGrade.AcademicID,
				StudentID:   entry.StudentID,
				SubjectID:   detail.SubjectID,
				Quiz:        entry.Quiz,
				FirstMonth:  entry.FirstMonth,
				SecondMonth: entry.SecondMonth,
				Finals:      entry.Finals,
				FinalGrade:  finals,
				Remarks:     entry.Remarks,
			})
		}
	}

	err = s.studentGradesRepository.Create(studentGradesData)

	if err != nil {
		return err
	}

	return nil
}

func (s *studentGradesUsecase) UpdateByAcademicID(academicID int, studentGrade dto.UpdateStudentGradesRequest) (dto.GetStudentGradesResponse, error) {
	e := validateUpdateStudentGradesRequest(studentGrade)
	if e != nil {
		return dto.GetStudentGradesResponse{}, e
	}
	academic, err := s.academicsRepository.Find(academicID)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("academic with ID %d not found: %w", academicID, err)
	}

	studentGradeData, err := s.studentGradesRepository.GetAll(academicID)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("error browsing student grades: %w", err)
	}

	existingGrades := make(map[uint]models.StudentGrades)
	for _, grade := range studentGradeData {
		existingGrades[grade.ID] = grade
	}
	studentGradesData := []models.StudentGrades{}
	for _, detail := range studentGrade.Details {
		for _, grade := range detail.Students {
			finals := math.Round(((grade.FirstMonth+grade.SecondMonth+grade.Quiz)/2 + grade.Finals) / 2)
			studentGradesData = append(studentGradesData, models.StudentGrades{
				ID:          grade.ID,
				DisplayName: existingGrades[grade.ID].DisplayName,
				AcademicID:  existingGrades[grade.ID].AcademicID,
				StudentID:   existingGrades[grade.ID].StudentID,
				SubjectID:   detail.SubjectID,
				Quiz:        grade.Quiz,
				FirstMonth:  grade.FirstMonth,
				SecondMonth: grade.SecondMonth,
				Finals:      grade.Finals,
				FinalGrade:  finals,
				Remarks:     grade.Remarks,
			})
		}
	}

	err = s.studentGradesRepository.UpdateByAcademicID(studentGradesData)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("error updating student grades for academic ID %d: %w", academicID, err)
	}

	studentGradesUpdated, err := s.studentGradesRepository.GetAll(academicID)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("error browsing student grades: %w", err)
	}

	details := make(map[uint]*dto.GetStudentGradesDetailResponse)
	for _, grade := range studentGradesUpdated {
		if _, exists := details[grade.SubjectID]; !exists {
			details[grade.SubjectID] = &dto.GetStudentGradesDetailResponse{
				SubjectID: grade.SubjectID,
				Subject:   grade.Subject.Name,
				Students:  []dto.GetStudentGradesEntryResponse{},
			}
		}
		details[grade.SubjectID].Students = append(details[grade.SubjectID].Students, dto.GetStudentGradesEntryResponse{
			ID:          grade.ID,
			StudentID:   grade.StudentID,
			StudentName: grade.Student.FullName,
			DisplayName: grade.DisplayName,
			Quiz:        grade.Quiz,
			FirstMonth:  grade.FirstMonth,
			SecondMonth: grade.SecondMonth,
			Finals:      grade.Finals,
			FinalGrade:  grade.FinalGrade,
			Remarks:     grade.Remarks,
		})
	}
	var detailsList []dto.GetStudentGradesDetailResponse
	for _, detail := range details {
		detailsList = append(detailsList, *detail)
	}
	response := dto.GetStudentGradesResponse{
		AcademicID: uint(academicID),
		Academic:   academic.DisplayName,
		Details:    detailsList,
	}

	return response, nil
}

func (s *studentGradesUsecase) GetReport(academicYear string, levelID, academicID int) ([]dto.StudentGradesReport, error) {
	startYear, endYear := "", ""
	if academicYear != "" {
		startYear, endYear = academicYear[:4], academicYear[5:9]
	}

	studentGrades, err := s.studentGradesRepository.GetReport(startYear, endYear, levelID, academicID)
	if err != nil {
		return []dto.StudentGradesReport{}, err
	}
	responses := []dto.StudentGradesReport{}
	classMap := make(map[string][]dto.StudentGradesReportTopStudent)
	for _, grade := range studentGrades {
		if _, exists := classMap[grade.Class]; !exists {
			classMap[grade.Class] = []dto.StudentGradesReportTopStudent{}
		}
		top := dto.StudentGradesReportTopStudent{
			Rank:    0,
			Student: grade.Student,
			NIS:     grade.NIS,
			Class:   grade.Class,
			Finals:  grade.Finals,
		}
		classMap[grade.Class] = append(classMap[grade.Class], top)
	}
	for class, students := range classMap {
		average := 0.0
		for idx, student := range students {
			average += student.Finals
			students[idx].Rank = idx + 1
		}
		if len(students) > 0 {
			average /= float64(len(students))
		}
		responses = append(responses, dto.StudentGradesReport{
			Class:    class,
			Average:  math.Round(average*100) / 100,
			Students: students,
		})
	}
	return responses, nil
}
