package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"

	"github.com/go-playground/validator/v10"
)

type DocTypesUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
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
	validate := validator.New()
	return validate.Struct(req)
}

func (s *docTypesUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	docTypes, total, err := s.docTypesRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return docTypes, total, nil
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
