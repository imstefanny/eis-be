package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type GuardiansUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	Create(guardian dto.CreateGuardiansRequest) error
	Find(id int) (interface{}, error)
	FindByApplicantId(id int) (interface{}, error)
	Update(id int, guardian dto.CreateGuardiansRequest) (models.Guardians, error)
	Delete(id int) error
}

type guardiansUsecase struct {
	guardiansRepository  repository.GuardiansRepository
	applicantsRepository repository.ApplicantsRepository
}

func NewGuardiansUsecase(guardiansRepo repository.GuardiansRepository, applicantsRepo repository.ApplicantsRepository) *guardiansUsecase {
	return &guardiansUsecase{
		guardiansRepository:  guardiansRepo,
		applicantsRepository: applicantsRepo,
	}
}

func validateCreateGuardiansRequest(req dto.CreateGuardiansRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *guardiansUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	guardians, total, err := s.guardiansRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return guardians, total, nil
}

func (s *guardiansUsecase) Create(guardian dto.CreateGuardiansRequest) error {
	e := validateCreateGuardiansRequest(guardian)

	if e != nil {
		return e
	}

	parsedDate, edate := time.Parse("2006-01-02", guardian.DateOfBirth)
	if edate != nil {
		return edate
	}

	guardianData := models.Guardians{
		ApplicantID:      guardian.ApplicantID,
		StudentID:        guardian.StudentID,
		Relation:         guardian.Relation,
		Name:             guardian.Name,
		Religion:         guardian.Religion,
		Job:              guardian.Job,
		Address:          guardian.Address,
		Phone:            guardian.Phone,
		PlaceOfBirth:     guardian.PlaceOfBirth,
		DateOfBirth:      parsedDate,
		HighestEducation: guardian.HighestEducation,
	}

	err := s.guardiansRepository.Create(guardianData)

	if err != nil {
		return err
	}

	return nil
}

func (s *guardiansUsecase) Find(id int) (interface{}, error) {
	guardian, err := s.guardiansRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return guardian, nil
}

func (s *guardiansUsecase) FindByApplicantId(id int) (interface{}, error) {
	guardian, err := s.guardiansRepository.FindByApplicantId(id)

	if err != nil {
		return nil, err
	}

	return guardian, nil
}

func (s *guardiansUsecase) Update(id int, guardian dto.CreateGuardiansRequest) (models.Guardians, error) {
	guardianData, err := s.guardiansRepository.Find(id)

	if err != nil {
		return models.Guardians{}, err
	}

	applicantData, _ := s.applicantsRepository.Find(int(guardianData.ApplicantID))

	if applicantData.State != "approved" && applicantData.State != "draft" {
		applicantData.State = "draft"
		err = s.applicantsRepository.Update(int(guardianData.ApplicantID), applicantData)
		if err != nil {
			return models.Guardians{}, err
		}
	}

	guardianData.ApplicantID = guardian.ApplicantID
	guardianData.StudentID = guardian.StudentID
	guardianData.Relation = guardian.Relation
	guardianData.Name = guardian.Name
	guardianData.Religion = guardian.Religion
	guardianData.Job = guardian.Job
	guardianData.Address = guardian.Address
	guardianData.Phone = guardian.Phone
	guardianData.PlaceOfBirth = guardian.PlaceOfBirth
	if guardian.DateOfBirth != "" {
		parsedDate, edate := time.Parse("2006-01-02", guardian.DateOfBirth)
		if edate != nil {
			return models.Guardians{}, edate
		}
		guardianData.DateOfBirth = parsedDate
	}
	guardianData.HighestEducation = guardian.HighestEducation

	e := s.guardiansRepository.Update(id, guardianData)

	if e != nil {
		return models.Guardians{}, e
	}

	guardianUpdated, err := s.guardiansRepository.Find(id)

	if err != nil {
		return models.Guardians{}, err
	}

	return guardianUpdated, nil
}

func (s *guardiansUsecase) Delete(id int) error {
	err := s.guardiansRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
