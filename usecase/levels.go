package usecase

import (
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/models"
	"eis-be/repository"
	"errors"
	"reflect"
)

type LevelsUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	Create(level dto.CreateLevelsRequest) error
	Find(id int) (interface{}, error)
	Update(id int, level dto.CreateLevelsRequest) (models.Levels, error)
	Delete(id int) error
}

type levelsUsecase struct {
	levelsRepository repository.LevelsRepository
}

func NewLevelsUsecase(levelsRepo repository.LevelsRepository) *levelsUsecase {
	return &levelsUsecase{
		levelsRepository: levelsRepo,
	}
}

func validateCreateLevelsRequest(req dto.CreateLevelsRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
		if helpers.IsEmptyField(val.Field(i)) {
			return errors.New("Field can't be empty")
		}
	}
	return nil
}

func (s *levelsUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	levels, total, err := s.levelsRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return levels, total, nil
}

func (s *levelsUsecase) Create(level dto.CreateLevelsRequest) error {
	e := validateCreateLevelsRequest(level)

	if e != nil {
		return e
	}

	levelData := models.Levels{
		Name: level.Name,
	}

	err := s.levelsRepository.Create(levelData)

	if err != nil {
		return err
	}

	return nil
}

func (s *levelsUsecase) Find(id int) (interface{}, error) {
	level, err := s.levelsRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return level, nil
}

func (s *levelsUsecase) Update(id int, level dto.CreateLevelsRequest) (models.Levels, error) {
	levelData, err := s.levelsRepository.Find(id)

	if err != nil {
		return models.Levels{}, err
	}

	levelData.Name = level.Name
	e := s.levelsRepository.Update(id, levelData)

	if e != nil {
		return models.Levels{}, e
	}

	levelUpdated, err := s.levelsRepository.Find(id)

	if err != nil {
		return models.Levels{}, err
	}

	return levelUpdated, nil
}

func (s *levelsUsecase) Delete(id int) error {
	err := s.levelsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
