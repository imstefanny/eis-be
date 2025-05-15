package usecase

import (
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/models"
	"eis-be/repository"
	"errors"
	"reflect"
)

type SubjectsUsecase interface {
	GetAll() (interface{}, error)
	Create(subject dto.CreateSubjectsRequest) error
	Find(id int) (interface{}, error)
	Update(id int, subject dto.CreateSubjectsRequest) (models.Subjects, error)
	Delete(id int) error
}

type subjectsUsecase struct {
	subjectsRepository repository.SubjectsRepository
}

func NewSubjectsUsecase(subjectsRepo repository.SubjectsRepository) *subjectsUsecase {
	return &subjectsUsecase{
		subjectsRepository: subjectsRepo,
	}
}

func validateCreateSubjectsRequest(req dto.CreateSubjectsRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
		if helpers.IsEmptyField(val.Field(i)) {
			return errors.New("Field can't be empty")
		}
	}
	return nil
}

func (s *subjectsUsecase) GetAll() (interface{}, error) {
	subjects, err := s.subjectsRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (s *subjectsUsecase) Create(subject dto.CreateSubjectsRequest) error {
	e := validateCreateSubjectsRequest(subject)

	if e != nil {
		return e
	}

	subjectData := models.Subjects{
		Name: subject.Name,
		Code: subject.Code,
	}

	err := s.subjectsRepository.Create(subjectData)

	if err != nil {
		return err
	}

	return nil
}

func (s *subjectsUsecase) Find(id int) (interface{}, error) {
	subject, err := s.subjectsRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return subject, nil
}

func (s *subjectsUsecase) Update(id int, subject dto.CreateSubjectsRequest) (models.Subjects, error) {
	subjectData, err := s.subjectsRepository.Find(id)

	if err != nil {
		return models.Subjects{}, err
	}

	subjectData.Name = subject.Name
	subjectData.Code = subject.Code

	e := s.subjectsRepository.Update(id, subjectData)

	if e != nil {
		return models.Subjects{}, e
	}

	subjectUpdated, err := s.subjectsRepository.Find(id)

	if err != nil {
		return models.Subjects{}, err
	}

	return subjectUpdated, nil
}

func (s *subjectsUsecase) Delete(id int) error {
	err := s.subjectsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
