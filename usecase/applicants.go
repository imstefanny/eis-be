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
}

func NewApplicantsUsecase(applicantsRepo repository.ApplicantsRepository, studentsRepo repository.StudentsRepository) *applicantsUsecase {
	return &applicantsUsecase{
		applicantsRepository: applicantsRepo,
		studentsRepository:   studentsRepo,
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
		return fmt.Errorf("Applicant with ID %d is already approved", id)
	}

	applicant.State = "approved"
	err = s.applicantsRepository.Update(id, applicant)

	if err != nil {
		return err
	}

	year := applicant.DateOfBirth.Year()
	lastThree := year % 1000
	studentData := models.Students{
		ApplicantID:      uint(id),
		UserID:           applicant.CreatedBy,
		ProfilePic:       applicant.ProfilePic,
		FullName:         applicant.FullName,
		IdentityNo:       applicant.IdentityNo,
		NIS:              fmt.Sprintf("%05d", id),
		NISN:             fmt.Sprintf("%03d%07d", lastThree, id),
		PlaceOfBirth:     applicant.PlaceOfBirth,
		DateOfBirth:      applicant.DateOfBirth,
		Address:          applicant.Address,
		Phone:            applicant.Phone,
		Religion:         applicant.Religion,
		ChildSequence:    applicant.ChildSequence,
		NumberOfSiblings: applicant.NumberOfSiblings,
		LivingWith:       applicant.LivingWith,
		ChildStatus:      applicant.ChildStatus,
	}

	eStudent := s.studentsRepository.Create(studentData)

	if eStudent != nil {
		return eStudent
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
