package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"eis-be/helpers"

	"github.com/go-playground/validator/v10"
)

type WorkSchedsUsecase interface {
	GetAll() (interface{}, error)
	Create(workSched dto.CreateWorkSchedsRequest) error
	Find(id int) (interface{}, error)
	Update(id int, workSched dto.CreateWorkSchedsRequest) (models.WorkScheds, error)
	Delete(id int) error
}

type workSchedsUsecase struct {
	workSchedsRepository repository.WorkSchedsRepository
	workSchedDetailsRepository repository.WorkSchedDetailsRepository
}

func NewWorkSchedsUsecase(workSchedsRepo repository.WorkSchedsRepository, workSchedDetailsRepo repository.WorkSchedDetailsRepository) *workSchedsUsecase {
	return &workSchedsUsecase{
		workSchedsRepository: workSchedsRepo,
		workSchedDetailsRepository: workSchedDetailsRepo,
	}
}

func validateCreateWorkSchedsRequest(req dto.CreateWorkSchedsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *workSchedsUsecase) GetAll() (interface{}, error) {
	workScheds, err := s.workSchedsRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return workScheds, nil
}

func (s *workSchedsUsecase) Create(workSched dto.CreateWorkSchedsRequest) error {
	e := validateCreateWorkSchedsRequest(workSched)

	if e != nil {
		return e
	}

	details := []models.WorkSchedDetails{}
	for _, detail := range workSched.Details {
		details = append(details, models.WorkSchedDetails{
			Day:       detail.Day,
			WorkStart: detail.WorkStart,
			WorkEnd:   detail.WorkEnd,
		})
	}

	workSchedData := models.WorkScheds{
		Name: workSched.Name,
		Details: details,
	}

	err := s.workSchedsRepository.Create(workSchedData)

	if err != nil {
		return err
	}

	return nil
}

func (s *workSchedsUsecase) Find(id int) (interface{}, error) {
	workSched, err := s.workSchedsRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return workSched, nil
}

func (s *workSchedsUsecase) Update(id int, workSched dto.CreateWorkSchedsRequest) (models.WorkScheds, error) {
	workSchedData, err := s.workSchedsRepository.Find(id)
	if err != nil {
		return models.WorkScheds{}, err
	}

	existing := workSchedData.Details
	existingIDs := []int{}
	for _, eDetail := range existing {
		existingIDs = append(existingIDs, int(eDetail.ID))
	}
	incomingDetails := workSched.Details
	incomingIDs := []int{}
	addIDs := []models.WorkSchedDetails{}
	for _, iDetail := range incomingDetails {
		if iDetail.ID != 0 {
			incomingIDs = append(incomingIDs, int(iDetail.ID))
		} else {
			addData := models.WorkSchedDetails{
				WorkSchedID: workSchedData.ID,
				Day:       iDetail.Day,
				WorkStart: iDetail.WorkStart,
				WorkEnd:   iDetail.WorkEnd,
			}
			addIDs = append(addIDs, addData)
		}
	}
	removeIDs := helpers.Difference(existingIDs, incomingIDs)
	updateIDs := helpers.Intersection(incomingIDs, existingIDs)
	incomingUpdate := []models.WorkSchedDetails{}
	for _, iDetail := range incomingDetails {
		for _, id := range updateIDs {
			if int(iDetail.ID) == id {
				incomingUpdate = append(incomingUpdate, models.WorkSchedDetails{
					ID:        iDetail.ID,
					WorkSchedID: workSchedData.ID,
					Day:       iDetail.Day,
					WorkStart: iDetail.WorkStart,
					WorkEnd:   iDetail.WorkEnd,
				})
			}
		}
	}
	workSchedData.Name = workSched.Name
	e := s.workSchedsRepository.Update(id, workSchedData)
	e = s.workSchedDetailsRepository.Create(addIDs)
	e = s.workSchedDetailsRepository.Update(updateIDs, incomingUpdate)
	e = s.workSchedDetailsRepository.Delete(removeIDs)
	if e != nil {
		return models.WorkScheds{}, e
	}

	workSchedUpdated, err := s.workSchedsRepository.Find(id)

	if err != nil {
		return models.WorkScheds{}, err
	}

	return workSchedUpdated, nil
}

func (s *workSchedsUsecase) Delete(id int) error {
	err := s.workSchedsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
