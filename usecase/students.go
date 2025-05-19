package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type StudentsUsecase interface {
	GetAll() (interface{}, error)
	Create(student dto.CreateStudentsRequest, c echo.Context) error
	Find(id int) (interface{}, error)
	Update(id int, student dto.CreateStudentsRequest) (models.Students, error)
	Delete(id int) error
}

type studentsUsecase struct {
	studentsRepository repository.StudentsRepository
}

func NewStudentsUsecase(studentsRepo repository.StudentsRepository) *studentsUsecase {
	return &studentsUsecase{
		studentsRepository: studentsRepo,
	}
}

func validateCreateStudentsRequest(req dto.CreateStudentsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *studentsUsecase) GetAll() (interface{}, error) {
	students, err := s.studentsRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s *studentsUsecase) Create(student dto.CreateStudentsRequest, c echo.Context) error {
	e := validateCreateStudentsRequest(student)

	if e != nil {
		return e
	}

	parsedDate, edate := time.Parse("2006-01-02", student.DateOfBirth)
	if edate != nil {
		return edate
	}

	studentData := models.Students{
		ApplicantID:       student.ApplicantID,
		CurrentAcademicID: student.CurrentAcademicID,
		UserID:            student.UserID,
		ProfilePicture:    student.ProfilePicture,
		FullName:          student.FullName,
		IdentityNo:        student.IdentityNo,
		NIS:               student.NIS,
		NISN:              student.NISN,
		PlaceOfBirth:      student.PlaceOfBirth,
		DateOfBirth:       parsedDate,
		Address:           student.Address,
		Phone:             student.Phone,
		Religion:          student.Religion,
		ChildSequence:     student.ChildSequence,
		NumberOfSiblings:  student.NumberOfSiblings,
		LivingWith:        student.LivingWith,
		ChildStatus:       student.ChildStatus,
	}

	err := s.studentsRepository.Create(studentData)

	if err != nil {
		return err
	}

	return nil
}

func (s *studentsUsecase) Find(id int) (interface{}, error) {
	student, err := s.studentsRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return student, nil
}

func (s *studentsUsecase) Update(id int, student dto.CreateStudentsRequest) (models.Students, error) {
	studentData, err := s.studentsRepository.Find(id)

	if err != nil {
		return models.Students{}, err
	}

	studentData.ApplicantID = student.ApplicantID
	studentData.CurrentAcademicID = student.CurrentAcademicID
	studentData.UserID = student.UserID
	studentData.ProfilePicture = student.ProfilePicture
	studentData.FullName = student.FullName
	studentData.IdentityNo = student.IdentityNo
	studentData.NIS = student.NIS
	studentData.NISN = student.NISN
	studentData.PlaceOfBirth = student.PlaceOfBirth
	if student.DateOfBirth != "" {
		parsedDate, edate := time.Parse("2006-01-02", student.DateOfBirth)
		if edate != nil {
			return models.Students{}, edate
		}
		studentData.DateOfBirth = parsedDate
	}
	studentData.Address = student.Address
	studentData.Phone = student.Phone
	studentData.Religion = student.Religion
	studentData.ChildSequence = student.ChildSequence
	studentData.NumberOfSiblings = student.NumberOfSiblings
	studentData.LivingWith = student.LivingWith
	studentData.ChildStatus = student.ChildStatus

	e := s.studentsRepository.Update(id, studentData)

	if e != nil {
		return models.Students{}, e
	}

	studentUpdated, err := s.studentsRepository.Find(id)

	if err != nil {
		return models.Students{}, err
	}

	return studentUpdated, nil
}

func (s *studentsUsecase) Delete(id int) error {
	err := s.studentsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
