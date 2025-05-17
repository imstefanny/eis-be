package usecase

import (
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/models"
	"eis-be/repository"
	"errors"
	"reflect"
)

type LevelHistoriesUsecase interface {
	GetAll() (interface{}, error)
	Create(levelHistory dto.CreateLevelHistoriesRequest) error
	Find(id int) (interface{}, error)
	Update(id int, levelHistory dto.CreateLevelHistoriesRequest) (models.LevelHistories, error)
	Delete(id int) error
}

type levelHistoriesUsecase struct {
	levelHistoriesRepository repository.LevelHistoriesRepository
}

func NewLevelHistoriesUsecase(levelHistoriesRepo repository.LevelHistoriesRepository) *levelHistoriesUsecase {
	return &levelHistoriesUsecase{
		levelHistoriesRepository: levelHistoriesRepo,
	}
}

func validateCreateLevelHistoriesRequest(req dto.CreateLevelHistoriesRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
		if helpers.IsEmptyField(val.Field(i)) {
			return errors.New("Field can't be empty")
		}
	}
	return nil
}

func (s *levelHistoriesUsecase) GetAll() (interface{}, error) {
	levelHistories, err := s.levelHistoriesRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return levelHistories, nil
}

func (s *levelHistoriesUsecase) Create(levelHistory dto.CreateLevelHistoriesRequest) error {
	e := validateCreateLevelHistoriesRequest(levelHistory)

	if e != nil {
		return e
	}

	levelHistorieData := models.LevelHistories{
		LevelID: levelHistory.LevelID,
		OpCertNum: levelHistory.OpCertNum,
		Accreditation: levelHistory.Accreditation,
		Curriculum: levelHistory.Curriculum,
		Email: levelHistory.Email,
		Phone: levelHistory.Phone,
		PrincipleID: levelHistory.PrincipleID,
		OperatorID: levelHistory.OperatorID,
		State: levelHistory.State,
	}

	err := s.levelHistoriesRepository.Create(levelHistorieData)

	if err != nil {
		return err
	}

	return nil
}

func (s *levelHistoriesUsecase) Find(id int) (interface{}, error) {
	levelHistory, err := s.levelHistoriesRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return levelHistory, nil
}

func (s *levelHistoriesUsecase) Update(id int, levelHistory dto.CreateLevelHistoriesRequest) (models.LevelHistories, error) {
	levelHistorieData, err := s.levelHistoriesRepository.Find(id)

	if err != nil {
		return models.LevelHistories{}, err
	}

	levelHistorieData.LevelID = levelHistory.LevelID
	levelHistorieData.OpCertNum = levelHistory.OpCertNum
	levelHistorieData.Accreditation = levelHistory.Accreditation
	levelHistorieData.Curriculum = levelHistory.Curriculum
	levelHistorieData.Email = levelHistory.Email
	levelHistorieData.Phone = levelHistory.Phone
	levelHistorieData.PrincipleID = levelHistory.PrincipleID
	levelHistorieData.OperatorID = levelHistory.OperatorID
	levelHistorieData.State = levelHistory.State
	e := s.levelHistoriesRepository.Update(id, levelHistorieData)

	if e != nil {
		return models.LevelHistories{}, e
	}

	levelHistorieUpdated, err := s.levelHistoriesRepository.Find(id)

	if err != nil {
		return models.LevelHistories{}, err
	}

	return levelHistorieUpdated, nil
}

func (s *levelHistoriesUsecase) Delete(id int) error {
	err := s.levelHistoriesRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
