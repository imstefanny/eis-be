package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type StudentAttsUsecase interface {
	BrowseByAcademicID(academicID, page, limit int, search string, date string) (dto.GetAllStudentAttsRequest, int64, error)
	CreateBatch(studentAtts dto.CreateBatchStudentAttsRequest) error
	UpdateByAcademicID(academicID int, studentAtt dto.UpdateStudentAttsRequest) (dto.GetAllStudentAttsRequest, error)

	// Students specific methods
	GetAttendanceByStudent(id, month int) (dto.StudentGetAttendancesResponse, error)
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

// Students specific methods
func (s *studentAttsUsecase) GetAttendanceByStudent(id, month int) (dto.StudentGetAttendancesResponse, error) {
	student, err := s.studentsRepository.GetByToken(id)
	if err != nil {
		return dto.StudentGetAttendancesResponse{}, fmt.Errorf("student with ID %d not found", id)
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	yearStr := student.Academics.StartYear
	if month < 7 {
		yearStr = student.Academics.EndYear
	}
	year, _ := strconv.Atoi(yearStr)
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)
	end := start.AddDate(0, 1, 0)
	studentAtts, err := s.studentAttsRepository.GetAttendanceByStudent(int(student.ID), start.Format("2006-01-02"), end.Format("2006-01-02"))
	if err != nil {
		return dto.StudentGetAttendancesResponse{}, err
	}

	attendances := []dto.StudentAttendanceDetail{}
	attMap := map[string]int{
		"Present":    0,
		"Permission": 0,
		"Sick":       0,
		"Alpha":      0,
	}
	for _, studentAtt := range studentAtts {
		entry := dto.StudentAttendanceDetail{
			ID:      studentAtt.ID,
			Date:    studentAtt.Date.Format("2006-01-02"),
			Status:  studentAtt.Status,
			Remarks: studentAtt.Remarks,
		}
		attendances = append(attendances, entry)
		attMap[studentAtt.Status]++
	}
	response := dto.StudentGetAttendancesResponse{
		Month:           month,
		Student:         student.FullName,
		Academic:        student.Academics.DisplayName,
		PresenceCount:   attMap["Present"],
		PermissionCount: attMap["Permission"],
		SickCount:       attMap["Sick"],
		AlphaCount:      attMap["Alpha"],
		Details:         attendances,
	}

	return response, nil
}
