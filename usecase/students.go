package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StudentsUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	Create(student dto.CreateStudentsRequest, c echo.Context) (uint, error)
	Find(id int) (interface{}, error)
	Update(id int, student dto.CreateStudentsRequest) (models.Students, error)
	UpdateStudentAcademicId(id int, academic []uint) error
	Delete(id int) error
}

type studentsUsecase struct {
	studentsRepository repository.StudentsRepository
	usersRepository    repository.UsersRepository
	db                 *gorm.DB
}

func NewStudentsUsecase(studentsRepo repository.StudentsRepository, usersRepo repository.UsersRepository, db *gorm.DB) *studentsUsecase {
	return &studentsUsecase{
		studentsRepository: studentsRepo,
		usersRepository:    usersRepo,
		db:                 db,
	}
}

func validateCreateStudentsRequest(req dto.CreateStudentsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *studentsUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	students, total, err := s.studentsRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return students, total, nil
}

func (s *studentsUsecase) Create(student dto.CreateStudentsRequest, c echo.Context) (uint, error) {
	e := validateCreateStudentsRequest(student)

	if e != nil {
		return 0, e
	}

	parsedDate, edate := time.Parse("2006-01-02", student.DateOfBirth)
	if edate != nil {
		return 0, edate
	}

	userData := models.Users{
		Name:       student.FullName,
		Email:      student.Email,
		Password:   "123456",
		RoleID:     1,
		ProfilePic: student.ProfilePic,
	}

	userId, errUser := s.usersRepository.Create(s.db, userData)
	if errUser != nil {
		return 0, errUser
	}

	studentData := models.Students{
		ApplicantID:       student.ApplicantID,
		CurrentAcademicID: student.CurrentAcademicID,
		UserID:            userId,
		ProfilePic:        student.ProfilePic,
		FullName:          student.FullName,
		IdentityNo:        student.IdentityNo,
		NIS:               student.NIS,
		NISN:              student.NISN,
		PlaceOfBirth:      student.PlaceOfBirth,
		DateOfBirth:       parsedDate,
		Address:           student.Address,
		Email:             student.Email,
		Phone:             student.Phone,
		Religion:          student.Religion,
		ChildSequence:     student.ChildSequence,
		NumberOfSiblings:  student.NumberOfSiblings,
		LivingWith:        student.LivingWith,
		ChildStatus:       student.ChildStatus,
	}

	studentId, err := s.studentsRepository.Create(studentData)

	if err != nil {
		return 0, err
	}

	year := parsedDate.Year()
	lastThree := year % 1000
	uniqueData := models.Students{
		ID:   studentId,
		NIS:  fmt.Sprintf("%05d", studentId),
		NISN: fmt.Sprintf("%03d%07d", lastThree, studentId),
	}

	eUpdt := s.studentsRepository.Update(int(studentId), uniqueData)
	if eUpdt != nil {
		return 0, eUpdt
	}

	return studentId, nil
}

func (s *studentsUsecase) Find(id int) (interface{}, error) {
	student, err := s.studentsRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return student, nil
}

func (s *studentsUsecase) Update(id int, student dto.CreateStudentsRequest) (models.Students, error) {
	errUnscope := s.studentsRepository.Undelete(id)
	if errUnscope != nil {
		return models.Students{}, errUnscope
	}

	studentData, err := s.studentsRepository.Find(id)

	if err != nil {
		return models.Students{}, err
	}

	studentData.ApplicantID = student.ApplicantID
	studentData.CurrentAcademicID = student.CurrentAcademicID
	studentData.UserID = student.UserID
	studentData.ProfilePic = student.ProfilePic
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
	studentData.Email = student.Email
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

func (s *studentsUsecase) UpdateStudentAcademicId(academic_id int, studentIDs []uint) error {
	err := s.studentsRepository.UpdateStudentAcademicId(academic_id, studentIDs)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentsUsecase) Delete(id int) error {
	err := s.studentsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
