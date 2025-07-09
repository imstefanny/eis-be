package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"errors"
)

type SubjectsUsecase interface {
	Browse(page, limit int, search, sortColumn, sortOrder string, isExtracurricular *bool) ([]dto.GetSubjectsResponse, int64, error)
	Create(subject dto.CreateSubjectsRequest) error
	Find(id int) (dto.GetSubjectsResponse, error)
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

func (s *subjectsUsecase) Browse(page, limit int, search, sortColumn, sortOrder string, isExtracurricular *bool) ([]dto.GetSubjectsResponse, int64, error) {
	subjects, total, err := s.subjectsRepository.Browse(page, limit, search, sortColumn, sortOrder, isExtracurricular)

	if err != nil {
		return nil, total, err
	}

	responses := []dto.GetSubjectsResponse{}
	for _, subject := range subjects {
		responses = append(responses, dto.GetSubjectsResponse{
			ID:                subject.ID,
			Name:              subject.DisplayName,
			Code:              subject.Code,
			IsExtracurricular: subject.IsExtracurricular,
		})
	}

	return responses, total, nil
}

func (s *subjectsUsecase) Create(subject dto.CreateSubjectsRequest) error {
	subjectData := models.Subjects{
		DisplayName:       subject.Code + " - " + subject.Name,
		Code:              subject.Code,
		Name:              subject.Name,
		IsExtracurricular: subject.IsExtracurricular,
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

func (s *subjectsUsecase) Find(id int) (dto.GetSubjectsResponse, error) {
	subject, err := s.subjectsRepository.Find(id)

	if err != nil {
		return dto.GetSubjectsResponse{}, err
	}

	response := dto.GetSubjectsResponse{
		ID:   subject.ID,
		Name: subject.DisplayName,
		Code: subject.Code,
	}

	return response, nil
}

func (s *subjectsUsecase) Update(id int, subject dto.CreateSubjectsRequest) (models.Subjects, error) {
	subjectData, err := s.subjectsRepository.Find(id)

	if err != nil {
		return models.Subjects{}, err
	}

	subjectData.DisplayName = subject.Code + " - " + subject.Name
	subjectData.Name = subject.Name
	subjectData.Code = subject.Code
	subjectData.IsExtracurricular = subject.IsExtracurricular

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
