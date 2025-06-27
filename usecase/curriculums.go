package usecase

import (
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/models"
	"eis-be/repository"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CurriculumsUsecase interface {
	Browse(page, limit int, search string) ([]dto.GetCurriculumsResponse, int64, error)
	Create(curriculum dto.CreateCurriculumsRequest) error
	Find(id int) (dto.GetCurriculumsResponse, error)
	Update(id int, curriculum dto.CreateCurriculumsRequest) (dto.GetCurriculumsResponse, error)
	Delete(id int) error
	UnDelete(id int) error
}

type curriculumsUsecase struct {
	curriculumsRepository repository.CurriculumsRepository
	levelsRepository      repository.LevelsRepository
	academicsRepository   repository.AcademicsRepository
}

func NewCurriculumsUsecase(curriculumsRepo repository.CurriculumsRepository, levelsRepo repository.LevelsRepository, academicsRepo repository.AcademicsRepository) *curriculumsUsecase {
	return &curriculumsUsecase{
		curriculumsRepository: curriculumsRepo,
		levelsRepository:      levelsRepo,
		academicsRepository:   academicsRepo,
	}
}

func validateCreateCurriculumsRequest(req dto.CreateCurriculumsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *curriculumsUsecase) Browse(page, limit int, search string) ([]dto.GetCurriculumsResponse, int64, error) {
	curriculums, total, err := s.curriculumsRepository.Browse(page, limit, search)

	if err != nil {
		return []dto.GetCurriculumsResponse{}, total, err
	}

	responses := []dto.GetCurriculumsResponse{}
	for _, curriculum := range curriculums {
		subjects := []dto.GetCurriculumSubjectsResponse{}
		response := dto.GetCurriculumsResponse{
			ID:                 curriculum.ID,
			DisplayName:        curriculum.DisplayName,
			Name:               curriculum.Name,
			LevelID:            curriculum.LevelID,
			Level:              curriculum.Level.Name,
			Grade:              curriculum.Grade,
			CurriculumSubjects: subjects,
			DeletedAt:          curriculum.DeletedAt,
		}
		responses = append(responses, response)
	}

	return responses, total, nil
}

func (s *curriculumsUsecase) Create(curriculum dto.CreateCurriculumsRequest) error {
	e := validateCreateCurriculumsRequest(curriculum)

	if e != nil {
		return e
	}

	subjects := []models.CurriculumSubjects{}
	for _, subject := range curriculum.CurriculumSubjects {
		subjects = append(subjects, models.CurriculumSubjects{
			SubjectID:  subject.SubjectID,
			Competence: subject.Competence,
		})
	}
	level, _ := s.levelsRepository.Find(int(curriculum.LevelID))
	curriculumData := models.Curriculums{
		DisplayName:        level.Name + " - " + curriculum.Grade + " / " + curriculum.Name,
		Name:               curriculum.Name,
		LevelID:            curriculum.LevelID,
		Grade:              curriculum.Grade,
		CurriculumSubjects: subjects,
	}

	curriculumID, err := s.curriculumsRepository.Create(curriculumData)
	if err != nil {
		return err
	}

	_ = s.academicsRepository.UpdateNewCurriculum(int(curriculum.LevelID), curriculum.Grade, curriculumID)

	return nil
}

func (s *curriculumsUsecase) Find(id int) (dto.GetCurriculumsResponse, error) {
	curriculum, err := s.curriculumsRepository.Find(id)

	if err != nil {
		return dto.GetCurriculumsResponse{}, err
	}

	subjects := []dto.GetCurriculumSubjectsResponse{}
	for _, subject := range curriculum.CurriculumSubjects {
		subjects = append(subjects, dto.GetCurriculumSubjectsResponse{
			ID:         subject.ID,
			SubjectID:  subject.SubjectID,
			Subject:    subject.Subject.Name,
			Competence: subject.Competence,
		})
	}
	response := dto.GetCurriculumsResponse{
		ID:                 curriculum.ID,
		DisplayName:        curriculum.DisplayName,
		Name:               curriculum.Name,
		LevelID:            curriculum.LevelID,
		Level:              curriculum.Level.Name,
		Grade:              curriculum.Grade,
		CurriculumSubjects: subjects,
	}

	return response, nil
}

func (s *curriculumsUsecase) Update(id int, curriculum dto.CreateCurriculumsRequest) (dto.GetCurriculumsResponse, error) {
	curriculumData, err := s.curriculumsRepository.Find(id)

	if err != nil {
		return dto.GetCurriculumsResponse{}, err
	}

	existing := curriculumData.CurriculumSubjects
	existingIDs := []int{}
	for _, eDetail := range existing {
		existingIDs = append(existingIDs, int(eDetail.ID))
	}
	incomingDetails := curriculum.CurriculumSubjects
	incomingIDs := []int{}
	addIDs := []models.CurriculumSubjects{}
	for _, iDetail := range incomingDetails {
		if iDetail.ID == 0 {
			addIDs = append(addIDs, models.CurriculumSubjects{
				CurriculumID: curriculumData.ID,
				SubjectID:    iDetail.SubjectID,
				Competence:   iDetail.Competence,
			})
		} else {
			incomingIDs = append(incomingIDs, int(iDetail.ID))
		}
	}
	removeIDs := helpers.Difference(existingIDs, incomingIDs)
	updateIDs := helpers.Intersection(incomingIDs, existingIDs)
	incomingUpdates := []models.CurriculumSubjects{}
	for _, iDetail := range incomingDetails {
		for _, uID := range updateIDs {
			if iDetail.ID == uint(uID) {
				incomingUpdates = append(incomingUpdates, models.CurriculumSubjects{
					ID:           iDetail.ID,
					CurriculumID: curriculumData.ID,
					SubjectID:    iDetail.SubjectID,
					Competence:   iDetail.Competence,
				})
			}
		}
	}
	if len(addIDs) == 0 && len(updateIDs) == 0 && len(removeIDs) == 0 {
		return dto.GetCurriculumsResponse{}, fmt.Errorf("no changes detected")
	}

	level, _ := s.levelsRepository.Find(int(curriculum.LevelID))
	curriculumData.DisplayName = level.Name + " - " + curriculum.Grade + " / " + curriculum.Name
	curriculumData.LevelID = curriculum.LevelID
	curriculumData.Grade = curriculum.Grade
	curriculumData.Name = curriculum.Name
	details := map[string]interface{}{
		"parents":         curriculumData,
		"addIDs":          addIDs,
		"updateIDs":       updateIDs,
		"removeIDs":       removeIDs,
		"incomingUpdates": incomingUpdates,
	}
	eTrx := s.curriculumsRepository.Update(id, details)
	if eTrx != nil {
		return dto.GetCurriculumsResponse{}, eTrx
	}

	curriculumUpdated, err := s.curriculumsRepository.Find(id)

	if err != nil {
		return dto.GetCurriculumsResponse{}, err
	}

	subjects := []dto.GetCurriculumSubjectsResponse{}
	for _, subject := range curriculumUpdated.CurriculumSubjects {
		subjects = append(subjects, dto.GetCurriculumSubjectsResponse{
			ID:         subject.ID,
			SubjectID:  subject.SubjectID,
			Subject:    subject.Subject.Name,
			Competence: subject.Competence,
		})
	}
	response := dto.GetCurriculumsResponse{
		ID:                 curriculumUpdated.ID,
		DisplayName:        curriculumUpdated.DisplayName,
		Name:               curriculumUpdated.Name,
		LevelID:            curriculumUpdated.LevelID,
		Level:              curriculumUpdated.Level.Name,
		Grade:              curriculumUpdated.Grade,
		CurriculumSubjects: subjects,
	}

	return response, nil
}

func (s *curriculumsUsecase) Delete(id int) error {
	err := s.curriculumsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *curriculumsUsecase) UnDelete(id int) error {
	err := s.curriculumsRepository.UnDelete(id)

	if err != nil {
		return err
	}

	return nil
}
