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
	GetAll(termID int) (dto.GetStudentGradesResponse, error)
	Create(studentGrade dto.CreateStudentGradesRequest) error
	UpdateByTermID(termID int, studentGrade dto.UpdateStudentGradesRequest) (dto.GetStudentGradesResponse, error)
	GetReport(academicYear string, levelID, academicID int) ([]dto.StudentGradesReport, error)
}

type studentGradesUsecase struct {
	studentGradesRepository repository.StudentGradesRepository
	academicsRepository     repository.AcademicsRepository
	termsRepository         repository.TermsRepository
	studentsRepository      repository.StudentsRepository
	subjectsRepository      repository.SubjectsRepository
}

func NewStudentGradesUsecase(studentGradesRepo repository.StudentGradesRepository, academicsRepo repository.AcademicsRepository, termsRepo repository.TermsRepository, studentsRepo repository.StudentsRepository, subjectsRepo repository.SubjectsRepository) *studentGradesUsecase {
	return &studentGradesUsecase{
		studentGradesRepository: studentGradesRepo,
		academicsRepository:     academicsRepo,
		termsRepository:         termsRepo,
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

func (s *studentGradesUsecase) GetAll(termID int) (dto.GetStudentGradesResponse, error) {
	term, err := s.termsRepository.Find(termID)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("term with ID %d not found: %w", termID, err)
	}

	studentGrades, err := s.studentGradesRepository.GetAll(termID)
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
			FirstQuiz:   grade.FirstQuiz,
			SecondQuiz:  grade.SecondQuiz,
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
		AcademicID: term.AcademicID,
		Academic:   term.Academic.DisplayName,
		TermID:     term.ID,
		Term:       term.Name,
		Details:    detailsList,
	}

	return response, nil
}

func (s *studentGradesUsecase) Create(studentGrade dto.CreateStudentGradesRequest) error {
	e := validateCreateStudentGradesRequest(studentGrade)
	if e != nil {
		return e
	}

	term, err := s.termsRepository.Find(int(studentGrade.TermID))
	if err != nil {
		return fmt.Errorf("term with ID %d not found: %w", studentGrade.TermID, err)
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
			finals := math.Round((((entry.FirstMonth+entry.SecondMonth)/2+(entry.FirstQuiz+entry.SecondQuiz)/2)/2 + entry.Finals) / 2)
			studentGradesData = append(studentGradesData, models.StudentGrades{
				DisplayName: term.Academic.DisplayName + " - " + subject.Name + " - " + student.FullName,
				AcademicID:  studentGrade.AcademicID,
				TermID:      term.ID,
				StudentID:   entry.StudentID,
				SubjectID:   detail.SubjectID,
				FirstQuiz:   entry.FirstQuiz,
				SecondQuiz:  entry.SecondQuiz,
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

func (s *studentGradesUsecase) UpdateByTermID(termID int, studentGrade dto.UpdateStudentGradesRequest) (dto.GetStudentGradesResponse, error) {
	e := validateUpdateStudentGradesRequest(studentGrade)
	if e != nil {
		return dto.GetStudentGradesResponse{}, e
	}
	term, err := s.termsRepository.Find(termID)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("term with ID %d not found: %w", termID, err)
	}

	studentGradeData, err := s.studentGradesRepository.GetAll(termID)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("error browsing student grades: %w", err)
	}

	existingGrades := make(map[uint]models.StudentGrades)
	for _, grade := range studentGradeData {
		existingGrades[grade.ID] = grade
	}
	studentGradesData := []models.StudentGrades{}
	newStudents := []models.StudentGrades{}
	for _, detail := range studentGrade.Details {
		for _, grade := range detail.Students {
			finals := math.Round((((grade.FirstMonth+grade.SecondMonth)/2+(grade.FirstQuiz+grade.SecondQuiz)/2)/2 + grade.Finals) / 2)
			if grade.ID != 0 {
				studentGradesData = append(studentGradesData, models.StudentGrades{
					ID:          grade.ID,
					DisplayName: existingGrades[grade.ID].DisplayName,
					AcademicID:  existingGrades[grade.ID].AcademicID,
					TermID:      existingGrades[grade.ID].TermID,
					StudentID:   existingGrades[grade.ID].StudentID,
					SubjectID:   detail.SubjectID,
					FirstQuiz:   grade.FirstQuiz,
					SecondQuiz:  grade.SecondQuiz,
					FirstMonth:  grade.FirstMonth,
					SecondMonth: grade.SecondMonth,
					Finals:      grade.Finals,
					FinalGrade:  finals,
					Remarks:     grade.Remarks,
				})
			} else {
				student, _ := s.studentsRepository.Find(int(grade.StudentID))
				newStudents = append(newStudents, models.StudentGrades{
					DisplayName: term.Academic.DisplayName + " - " + detail.Subject + " - " + student.FullName,
					AcademicID:  studentGrade.AcademicID,
					TermID:      term.ID,
					StudentID:   grade.StudentID,
					SubjectID:   detail.SubjectID,
					FirstQuiz:   grade.FirstQuiz,
					SecondQuiz:  grade.SecondQuiz,
					FirstMonth:  grade.FirstMonth,
					SecondMonth: grade.SecondMonth,
					Finals:      grade.Finals,
					FinalGrade:  finals,
					Remarks:     grade.Remarks,
				})
			}
		}
	}

	err = s.studentGradesRepository.UpdateByTermID(studentGradesData, newStudents)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("error updating student grades for term ID %d: %w", termID, err)
	}

	studentGradesUpdated, err := s.studentGradesRepository.GetAll(termID)
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
			FirstQuiz:   grade.FirstQuiz,
			SecondQuiz:  grade.SecondQuiz,
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
		AcademicID: term.AcademicID,
		Academic:   term.Academic.DisplayName,
		TermID:     term.ID,
		Term:       term.Name,
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
