package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"

	"github.com/go-playground/validator/v10"
)

type DocumentsUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	Create(document dto.CreateDocumentsRequest) error
	Find(id int) (interface{}, error)
	FindByApplicantId(id int) (interface{}, error)
	Update(id int, document dto.CreateDocumentsRequest) (models.Documents, error)
	Delete(id int) error
}

type documentsUsecase struct {
	documentsRepository repository.DocumentsRepository
}

func NewDocumentsUsecase(documentsRepo repository.DocumentsRepository) *documentsUsecase {
	return &documentsUsecase{
		documentsRepository: documentsRepo,
	}
}

func validateCreateDocumentsRequest(req dto.CreateDocumentsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *documentsUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	documents, total, err := s.documentsRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return documents, total, nil
}

func (s *documentsUsecase) Create(document dto.CreateDocumentsRequest) error {
	e := validateCreateDocumentsRequest(document)

	if e != nil {
		return e
	}

	documentData := models.Documents{
		TypeID:       document.TypeID,
		ApplicantID:  document.ApplicantID,
		StudentID:    document.StudentID,
		UploadedFile: document.UploadedFile,
		Description:  document.Description,
	}

	err := s.documentsRepository.Create(documentData)

	if err != nil {
		return err
	}

	return nil
}

func (s *documentsUsecase) Find(id int) (interface{}, error) {
	document, err := s.documentsRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return document, nil
}

func (s *documentsUsecase) FindByApplicantId(id int) (interface{}, error) {
	documents, err := s.documentsRepository.FindByApplicantId(id)

	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (s *documentsUsecase) Update(id int, document dto.CreateDocumentsRequest) (models.Documents, error) {
	documentData, err := s.documentsRepository.Find(id)

	if err != nil {
		return models.Documents{}, err
	}

	documentData.TypeID = document.TypeID
	documentData.ApplicantID = document.ApplicantID
	documentData.StudentID = document.StudentID
	documentData.UploadedFile = document.UploadedFile
	documentData.Description = document.Description

	e := s.documentsRepository.Update(id, documentData)

	if e != nil {
		return models.Documents{}, e
	}

	documentUpdated, err := s.documentsRepository.Find(id)

	if err != nil {
		return models.Documents{}, err
	}

	return documentUpdated, nil
}

func (s *documentsUsecase) Delete(id int) error {
	err := s.documentsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
