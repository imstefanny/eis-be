package usecase

import (
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/models"
	"eis-be/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ApplicantsUsecase interface {
	GetAll() (interface{}, error)
	Create(applicant dto.CreateApplicantsRequest, c echo.Context) error
	Find(id int) (interface{}, error)
	FindByCreatedBy(id int) (interface{}, error)
	Update(id int, applicant dto.CreateApplicantsRequest) (models.Applicants, error)
	Delete(id int) error
}

type applicantsUsecase struct {
	applicantsRepository repository.ApplicantsRepository
}

func NewApplicantsUsecase(applicantsRepo repository.ApplicantsRepository) *applicantsUsecase {
	return &applicantsUsecase{
		applicantsRepository: applicantsRepo,
	}
}

func validateCreateApplicantsRequest(req dto.CreateApplicantsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *applicantsUsecase) GetAll() (interface{}, error) {
	applicants, err := s.applicantsRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return applicants, nil
}

func (s *applicantsUsecase) Create(applicant dto.CreateApplicantsRequest, c echo.Context) error {
	e := validateCreateApplicantsRequest(applicant)

	if e != nil {
		return e
	}

	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return errToken
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

func (s *applicantsUsecase) FindByCreatedBy(id int) (interface{}, error) {
	applicant, err := s.applicantsRepository.FindByCreatedBy(id)
	if err != nil {
		return nil, err
	}

	return applicant, nil
}

func (s *applicantsUsecase) Update(id int, applicant dto.CreateApplicantsRequest) (models.Applicants, error) {
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

func (s *applicantsUsecase) Delete(id int) error {
	err := s.applicantsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
