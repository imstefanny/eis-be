package usecase

import (
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/models"
	"eis-be/repository"
	"errors"
	"reflect"
)

type DocTypesUsecase interface {
	GetAll() (interface{}, error)
	Create(docType dto.CreateDocTypesRequest) error
	Find(id int) (interface{}, error)
	Update(id int, docType dto.CreateDocTypesRequest) (models.DocTypes, error)
	Delete(id int) error
}

type docTypesUsecase struct {
	docTypesRepository repository.DocTypesRepository
}

func NewDocTypesUsecase(docTypesRepo repository.DocTypesRepository) *docTypesUsecase {
	return &docTypesUsecase{
		docTypesRepository: docTypesRepo,
	}
}

func validateCreateDocTypesRequest(req dto.CreateDocTypesRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
		if helpers.IsEmptyField(val.Field(i)) {
			return errors.New("Field can't be empty")
		}
	}
	return nil
}

func (s *docTypesUsecase) GetAll() (interface{}, error) {
	docTypes, err := s.docTypesRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return docTypes, nil
}

func (s *docTypesUsecase) Create(docType dto.CreateDocTypesRequest) error {
	e := validateCreateDocTypesRequest(docType)

	if e != nil {
		return e
	}

	docTypeData := models.DocTypes{
		Name:        docType.Name,
		Description: docType.Description,
	}

	err := s.docTypesRepository.Create(docTypeData)

	if err != nil {
		return err
	}

	return nil
}

func (s *docTypesUsecase) Find(id int) (interface{}, error) {
	docType, err := s.docTypesRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return docType, nil
}

func (s *docTypesUsecase) Update(id int, docType dto.CreateDocTypesRequest) (models.DocTypes, error) {
	docTypeData, err := s.docTypesRepository.Find(id)

	if err != nil {
		return models.DocTypes{}, err
	}

	docTypeData.Name = docType.Name
	docTypeData.Description = docType.Description

	e := s.docTypesRepository.Update(id, docTypeData)

	if e != nil {
		return models.DocTypes{}, e
	}

	docTypeUpdated, err := s.docTypesRepository.Find(id)

	if err != nil {
		return models.DocTypes{}, err
	}

	return docTypeUpdated, nil
}

func (s *docTypesUsecase) Delete(id int) error {
	err := s.docTypesRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
