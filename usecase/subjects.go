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
	Browse(page, limit int, search, sortColumn, sortOrder string) (interface{}, int64, error)
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
			return errors.New("field can't be empty")
		}
	}
	return nil
}

func (s *subjectsUsecase) Browse(page, limit int, search, sortColumn, sortOrder string) (interface{}, int64, error) {
	blogs, total, err := s.subjectsRepository.Browse(page, limit, search, sortColumn, sortOrder)

	if err != nil {
		return nil, total, err
	}

	return blogs, total, nil
}

func (s *subjectsUsecase) Create(subject dto.CreateSubjectsRequest) error {
	e := validateCreateSubjectsRequest(subject)

	if e != nil {
		return e
	}

	subjectData := models.Subjects{
		Code: subject.Code,
		Name: subject.Name,
	}
	subjectResult := s.subjectsRepository.FindByCode(subject.Code)
	if subjectResult.ID != 0 {
		return errors.New("mata pelajaran dengan kode ini sudah ada")
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
