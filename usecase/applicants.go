package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type ApplicantsUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	Create(applicant dto.CreateApplicantsRequest, claims jwt.MapClaims) error
	Find(id int) (interface{}, error)
	GetByToken(id int) (interface{}, error)
	Update(id int, claims jwt.MapClaims, applicant dto.CreateApplicantsRequest) (models.Applicants, error)
	ApproveRegistration(id int, claims jwt.MapClaims) error
	Delete(id int) error
}

type applicantsUsecase struct {
	applicantsRepository repository.ApplicantsRepository
	studentsRepository   repository.StudentsRepository
	guardiansRepository  repository.GuardiansRepository
	usersRepository      repository.UsersRepository
	rolesRepository      repository.RolesRepository
}

func NewApplicantsUsecase(applicantsRepo repository.ApplicantsRepository, studentsRepo repository.StudentsRepository, guardiansRepo repository.GuardiansRepository, usersRepo repository.UsersRepository, rolesRepo repository.RolesRepository) *applicantsUsecase {
	return &applicantsUsecase{
		applicantsRepository: applicantsRepo,
		studentsRepository:   studentsRepo,
		guardiansRepository:  guardiansRepo,
		usersRepository:      usersRepo,
		rolesRepository:      rolesRepo,
	}
}

func validateCreateApplicantsRequest(req dto.CreateApplicantsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *applicantsUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	applicants, total, err := s.applicantsRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return applicants, total, nil
}

func (s *applicantsUsecase) Create(applicant dto.CreateApplicantsRequest, claims jwt.MapClaims) error {
	e := validateCreateApplicantsRequest(applicant)

	if e != nil {
		return e
	}

	parsedDate, edate := time.Parse("2006-01-02", applicant.DateOfBirth)
	if edate != nil {
		return edate
	}

	applicantData := models.Applicants{
		ProfilePic:        applicant.ProfilePic,
		FullName:          applicant.FullName,
		IdentityNo:        applicant.IdentityNo,
		PlaceOfBirth:      applicant.PlaceOfBirth,
		DateOfBirth:       parsedDate,
		Address:           applicant.Address,
		Phone:             applicant.Phone,
		Religion:          applicant.Religion,
		ChildSequence:     applicant.ChildSequence,
		NumberOfSiblings:  applicant.NumberOfSiblings,
		LivingWith:        applicant.LivingWith,
		ChildStatus:       applicant.ChildStatus,
		SchoolOrigin:      applicant.SchoolOrigin,
		LevelID:           applicant.LevelID,
		RegistrationGrade: applicant.RegistrationGrade,
		RegistrationMajor: applicant.RegistrationMajor,
		State:             applicant.State,
		CreatedBy:         uint(claims["userId"].(float64)),
	}

	err := s.applicantsRepository.Create(applicantData)

	if err != nil {
		return err
	}

	return nil
}

func (s *applicantsUsecase) Find(id int) (interface{}, error) {
	applicant, err := s.applicantsRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return applicant, nil
}

func (s *applicantsUsecase) GetByToken(id int) (interface{}, error) {
	applicant, err := s.applicantsRepository.GetByToken(id)
	if err != nil {
		return nil, err
	}

	return applicant, nil
}

func (s *applicantsUsecase) Update(id int, claims jwt.MapClaims, applicant dto.CreateApplicantsRequest) (models.Applicants, error) {
	applicantData, err := s.applicantsRepository.Find(id)

	if err != nil {
		return models.Applicants{}, err
	}

	applicantData.FullName = applicant.FullName
	applicantData.IdentityNo = applicant.IdentityNo
	applicantData.PlaceOfBirth = applicant.PlaceOfBirth
	if applicant.DateOfBirth != "" {
		parsedDate, edate := time.Parse("2006-01-02", applicant.DateOfBirth)
		if edate != nil {
			return models.Applicants{}, edate
		}
		applicantData.DateOfBirth = parsedDate
	}
	applicantData.Address = applicant.Address
	applicantData.Phone = applicant.Phone
	applicantData.Religion = applicant.Religion
	applicantData.ChildSequence = applicant.ChildSequence
	applicantData.NumberOfSiblings = applicant.NumberOfSiblings
	applicantData.LivingWith = applicant.LivingWith
	applicantData.ChildStatus = applicant.ChildStatus
	applicantData.SchoolOrigin = applicant.SchoolOrigin
	applicantData.LevelID = applicant.LevelID
	applicantData.RegistrationGrade = applicant.RegistrationGrade
	applicantData.RegistrationMajor = applicant.RegistrationMajor
	applicantData.State = applicant.State
	applicantData.UpdatedBy = uint(claims["userId"].(float64))

	e := s.applicantsRepository.Update(id, applicantData)

	if e != nil {
		return models.Applicants{}, e
	}

	applicantUpdated, err := s.applicantsRepository.Find(id)

	if err != nil {
		return models.Applicants{}, err
	}

	return applicantUpdated, nil
}

func (s *applicantsUsecase) ApproveRegistration(id int, claims jwt.MapClaims) error {
	applicant, err := s.applicantsRepository.Find(id)

	if err != nil {
		return err
	}

	if applicant.State == "approved" {
		return fmt.Errorf("applicant with ID %d is already approved", id)
	}

	applicant.State = "approved"
	err = s.applicantsRepository.Update(id, applicant)

	if err != nil {
		return err
	}

	userData, _ := s.usersRepository.Find(int(applicant.CreatedBy))
	role, _ := s.rolesRepository.FindByName("Student")
	userData.RoleID = role.ID
	_ = s.usersRepository.Update(userData)
	userUpdated, _ := s.usersRepository.Find(int(applicant.CreatedBy))

	year := applicant.DateOfBirth.Year()
	lastThree := year % 1000
	studentData := models.Students{
		ApplicantID:      uint(id),
		UserID:           userUpdated.ID,
		ProfilePic:       applicant.ProfilePic,
		FullName:         applicant.FullName,
		IdentityNo:       applicant.IdentityNo,
		PlaceOfBirth:     applicant.PlaceOfBirth,
		DateOfBirth:      applicant.DateOfBirth,
		Address:          applicant.Address,
		Email:            applicant.CreatedByName.Email,
		Phone:            applicant.Phone,
		Religion:         applicant.Religion,
		ChildSequence:    applicant.ChildSequence,
		NumberOfSiblings: applicant.NumberOfSiblings,
		LivingWith:       applicant.LivingWith,
		ChildStatus:      applicant.ChildStatus,
	}

	studentID, errStudent := s.studentsRepository.Create(studentData)
	if errStudent != nil {
		return errStudent
	}

	uniqueData := models.Students{
		ID:   studentID,
		NIS:  fmt.Sprintf("%05d", studentID),
		NISN: fmt.Sprintf("%03d%07d", lastThree, studentID),
	}
	eUpdt := s.studentsRepository.Update(int(studentID), uniqueData)
	if eUpdt != nil {
		return eUpdt
	}

	guardians, errGuardians := s.guardiansRepository.FindByApplicantId(id)
	if errGuardians != nil {
		return errGuardians
	}
	for _, guardian := range guardians {
		guardian.StudentID = studentID
		s.guardiansRepository.Update(int(guardian.ID), guardian)
	}

	return nil
}

func (s *applicantsUsecase) Delete(id int) error {
	err := s.applicantsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
