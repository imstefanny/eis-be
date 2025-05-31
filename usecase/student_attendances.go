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
	// Create(studentAtt dto.CreateStudentAttsRequest) error
	CreateBatch(studentAtts dto.CreateBatchStudentAttsRequest) error
	// Find(id int) (interface{}, error)
	// Update(id int, studentAtt dto.CreateStudentAttsRequest) (models.StudentAttendances, error)
	// Delete(id int) error
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

// func validateCreateStudentAttsRequest(req dto.CreateStudentAttsRequest) error {
// 	validate := validator.New()
// 	return validate.Struct(req)
// }

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
		Academic: academic.DisplayName,
		Date:     date,
		Students: students,
	}

	return response, total, nil
}

// func (s *studentAttsUsecase) Create(studentAtt dto.CreateStudentAttsRequest) error {
// 	e := validateCreateStudentAttsRequest(studentAtt)

// 	if e != nil {
// 		return e
// 	}

// 	loc, _ := time.LoadLocation("Asia/Jakarta")
// 	parseDate, err := time.Parse("2006-01-02", studentAtt.Date)
// 	if err != nil {
// 		return err
// 	}
// 	parseInTime, err := time.Parse("15:04:05", studentAtt.LogInTime)
// 	if err != nil {
// 		return err
// 	}
// 	if studentAtt.LogOutTime == "" {
// 		studentAtt.LogOutTime = "00:00:00"
// 	}
// 	parseOutTime, err := time.Parse("15:04:05", studentAtt.LogOutTime)
// 	if err != nil {
// 		return err
// 	}

// 	student, err := s.studentRepository.Find(int(studentAtt.StudentID))
// 	if err != nil {
// 		return err
// 	}
// 	studentAttData := models.StudentAttendances{
// 		DisplayName:       student.Name + " - " + parseDate.Format("2006-01-02"),
// 		StudentID:         studentAtt.StudentID,
// 		WorkingScheduleID: student.WorkSchedID,
// 		Date:              time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseDate.Hour(), parseDate.Minute(), parseDate.Second(), 0, loc),
// 		LogInTime:         time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseInTime.Hour(), parseInTime.Minute(), parseInTime.Second(), 0, loc),
// 		LogOutTime:        time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseOutTime.Hour(), parseOutTime.Minute(), parseOutTime.Second(), 0, loc),
// 		Remark:            studentAtt.Remark,
// 		Note:              studentAtt.Note,
// 	}

// 	err = s.studentAttsRepository.Create(studentAttData)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

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
				Status:      "Present", // Default status
				Remarks:     "",        // Default remarks
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

// func (s *studentAttsUsecase) Find(id int) (interface{}, error) {
// 	studentAtt, err := s.studentAttsRepository.Find(id)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return studentAtt, nil
// }

// func (s *studentAttsUsecase) Update(id int, studentAtt dto.CreateStudentAttsRequest) (models.StudentAttendances, error) {
// 	studentAttData, err := s.studentAttsRepository.Find(id)

// 	if err != nil {
// 		return models.StudentAttendances{}, err
// 	}

// 	parseDate, err := time.Parse("2006-01-02", studentAtt.Date)
// 	if err != nil {
// 		return models.StudentAttendances{}, err
// 	}
// 	parseInTime, err := time.Parse("15:04:05", studentAtt.LogInTime)
// 	if err != nil {
// 		return models.StudentAttendances{}, err
// 	}
// 	parseOutTime, err := time.Parse("15:04:05", studentAtt.LogOutTime)
// 	if err != nil {
// 		return models.StudentAttendances{}, err
// 	}

// 	loc, _ := time.LoadLocation("Asia/Jakarta")
// 	studentAttData.StudentID = studentAtt.StudentID
// 	studentAttData.WorkingScheduleID = studentAtt.WorkingScheduleID
// 	studentAttData.Date = time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseDate.Hour(), parseDate.Minute(), parseDate.Second(), 0, loc)
// 	studentAttData.LogInTime = time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseInTime.Hour(), parseInTime.Minute(), parseInTime.Second(), 0, loc)
// 	studentAttData.LogOutTime = time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseOutTime.Hour(), parseOutTime.Minute(), parseOutTime.Second(), 0, loc)
// 	studentAttData.Remark = studentAtt.Remark
// 	studentAttData.Note = studentAtt.Note

// 	e := s.studentAttsRepository.Update(id, studentAttData)

// 	if e != nil {
// 		return models.StudentAttendances{}, e
// 	}

// 	studentAttUpdated, err := s.studentAttsRepository.Find(id)

// 	if err != nil {
// 		return models.StudentAttendances{}, err
// 	}

// 	return studentAttUpdated, nil
// }

// func (s *studentAttsUsecase) Delete(id int) error {
// 	err := s.studentAttsRepository.Delete(id)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
