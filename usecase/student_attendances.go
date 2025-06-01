package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type StudentAttsUsecase interface {
	BrowseByAcademicID(academicID, page, limit int, search string, date string) (dto.GetAllStudentAttsRequest, int64, error)
	CreateBatch(studentAtts dto.CreateBatchStudentAttsRequest) error
	UpdateByAcademicID(academicID int, studentAtt dto.UpdateStudentAttsRequest) (dto.GetAllStudentAttsRequest, error)
}

type studentAttsUsecase struct {
	studentAttsRepository repository.StudentAttsRepository
	studentsRepository    repository.StudentsRepository
	academicsRepository   repository.AcademicsRepository
}

func NewStudentAttsUsecase(studentAttsRepo repository.StudentAttsRepository, studentsRepo repository.StudentsRepository, academicsRepo repository.AcademicsRepository) *studentAttsUsecase {
	return &studentAttsUsecase{
		studentAttsRepository: studentAttsRepo,
		studentsRepository:    studentsRepo,
		academicsRepository:   academicsRepo,
	}
}

func validateCreateBatchStudentAttsRequest(req dto.CreateBatchStudentAttsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *studentAttsUsecase) BrowseByAcademicID(academicID, page, limit int, search string, date string) (dto.GetAllStudentAttsRequest, int64, error) {
	studentAtts, total, err := s.studentAttsRepository.BrowseByAcademicID(academicID, page, limit, search, date)

	if err != nil {
		return dto.GetAllStudentAttsRequest{}, total, err
	}

	academic, err := s.academicsRepository.Find(academicID)
	if err != nil {
		return dto.GetAllStudentAttsRequest{}, total, fmt.Errorf("academic with ID %d not found", academicID)
	}
	students := []dto.GetAllStudentAttsEntryRequest{}
	for _, studentAtt := range studentAtts {
		student, err := s.studentsRepository.Find(int(studentAtt.StudentID))
		if err != nil {
			return dto.GetAllStudentAttsRequest{}, total, fmt.Errorf("failed to find student with ID %d: %w", studentAtt.StudentID, err)
		}
		entry := dto.GetAllStudentAttsEntryRequest{
			ID:          studentAtt.ID,
			StudentID:   student.ID,
			Student:     student.FullName,
			DisplayName: studentAtt.DisplayName,
			Status:      studentAtt.Status,
			Remarks:     studentAtt.Remarks,
			CreatedAt:   studentAtt.CreatedAt,
			UpdatedAt:   studentAtt.UpdatedAt,
			DeletedAt:   studentAtt.DeletedAt,
		}
		students = append(students, entry)
	}

	response := dto.GetAllStudentAttsRequest{
		AcademicID: academic.ID,
		Academic:   academic.DisplayName,
		Date:       date,
		Students:   students,
	}

	return response, total, nil
}

func (s *studentAttsUsecase) CreateBatch(studentAtt dto.CreateBatchStudentAttsRequest) error {
	e := validateCreateBatchStudentAttsRequest(studentAtt)

	if e != nil {
		return e
	}

	parsedDate, edate := time.Parse("2006-01-02", studentAtt.Date)
	if edate != nil {
		return edate
	}

	academics, err := s.academicsRepository.GetAll()
	if err != nil {
		return err
	}

	studentAttsData := []models.StudentAttendances{}
	for _, academic := range academics {
		for _, student := range academic.Students {
			studentAttData := models.StudentAttendances{
				DisplayName: student.FullName + " - " + parsedDate.Format("2006-01-02"),
				AcademicID:  academic.ID,
				StudentID:   student.ID,
				Date:        parsedDate,
				Status:      "Present",
				Remarks:     "",
			}
			studentAttsData = append(studentAttsData, studentAttData)
		}
	}

	err = s.studentAttsRepository.CreateBatch(studentAttsData)

	if err != nil {
		return err
	}

	return nil
}

func (s *studentAttsUsecase) UpdateByAcademicID(academicID int, studentAtt dto.UpdateStudentAttsRequest) (dto.GetAllStudentAttsRequest, error) {
	academic, err := s.academicsRepository.Find(academicID)
	if err != nil {
		return dto.GetAllStudentAttsRequest{}, fmt.Errorf("academic with ID %d not found", academicID)
	}
	parseDate, err := time.Parse("2006-01-02", studentAtt.Date)
	if err != nil {
		return dto.GetAllStudentAttsRequest{}, err
	}

	studentAttsData, err := s.studentAttsRepository.FindByAcademicDate(academicID, parseDate.Format("2006-01-02"))
	if len(studentAttsData) == 0 {
		toBeCreated := []models.StudentAttendances{}
		for _, student := range studentAtt.Students {
			studentDetail, _ := s.studentsRepository.Find(int(student.StudentID))
			toBeCreated = append(toBeCreated, models.StudentAttendances{
				DisplayName: studentDetail.FullName,
				AcademicID:  uint(academicID),
				StudentID:   student.StudentID,
				Date:        parseDate,
				Status:      student.Status,
				Remarks:     student.Remarks,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			})
		}
		s.studentAttsRepository.CreateBatch(toBeCreated)
		return dto.GetAllStudentAttsRequest{
			AcademicID: uint(academicID),
			Academic:   academic.DisplayName,
			Date:       studentAtt.Date,
			Students:   []dto.GetAllStudentAttsEntryRequest{},
		}, nil
	}

	if err != nil {
		return dto.GetAllStudentAttsRequest{}, err
	}

	toBeUpdated := []models.StudentAttendances{}
	for _, studentAttData := range studentAttsData {
		for _, student := range studentAtt.Students {
			if studentAttData.StudentID == student.StudentID {
				studentAttData.Status = student.Status
				studentAttData.Remarks = student.Remarks
			}
		}
		toBeUpdated = append(toBeUpdated, studentAttData)
	}

	e := s.studentAttsRepository.UpdateByAcademicID(academicID, toBeUpdated)
	if e != nil {
		return dto.GetAllStudentAttsRequest{}, e
	}

	studentAttUpdated, err := s.studentAttsRepository.FindByAcademicDate(academicID, parseDate.Format("2006-01-02"))
	if err != nil {
		return dto.GetAllStudentAttsRequest{}, err
	}

	students := []dto.GetAllStudentAttsEntryRequest{}
	for _, studentAtt := range studentAttUpdated {
		student, err := s.studentsRepository.Find(int(studentAtt.StudentID))
		if err != nil {
			return dto.GetAllStudentAttsRequest{}, fmt.Errorf("failed to find student with ID %d: %w", studentAtt.StudentID, err)
		}
		entry := dto.GetAllStudentAttsEntryRequest{
			ID:          studentAtt.ID,
			StudentID:   student.ID,
			Student:     student.FullName,
			DisplayName: studentAtt.DisplayName,
			Status:      studentAtt.Status,
			Remarks:     studentAtt.Remarks,
			CreatedAt:   studentAtt.CreatedAt,
			UpdatedAt:   studentAtt.UpdatedAt,
			DeletedAt:   studentAtt.DeletedAt,
		}
		students = append(students, entry)
	}

	response := dto.GetAllStudentAttsRequest{
		AcademicID: uint(academicID),
		Academic:   academic.DisplayName,
		Date:       studentAtt.Date,
		Students:   students,
	}

	return response, nil
}
