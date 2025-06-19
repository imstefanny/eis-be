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
	Browse(page, limit int, search string) (interface{}, int64, error)
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
			return errors.New("field can't be empty")
		}
	}
	return nil
}

func (s *levelHistoriesUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	levelHistories, total, err := s.levelHistoriesRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return levelHistories, total, nil
}

func (s *levelHistoriesUsecase) Create(levelHistory dto.CreateLevelHistoriesRequest) error {
	var principleID *uint
	if levelHistory.PrincipleID != 0 {
		principleID = &levelHistory.PrincipleID
	}
	var operatorID *uint
	if levelHistory.OperatorID != 0 {
		operatorID = &levelHistory.OperatorID
	}

	levelHistorieData := models.LevelHistories{
		LevelID:       levelHistory.LevelID,
		OpCertNum:     levelHistory.OpCertNum,
		NPSN:          levelHistory.NPSN,
		Accreditation: levelHistory.Accreditation,
		Curriculum:    levelHistory.Curriculum,
		Email:         levelHistory.Email,
		Phone:         levelHistory.Phone,
		PrincipleID:   principleID,
		OperatorID:    operatorID,
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
	var principleID *uint
	if levelHistory.PrincipleID != 0 {
		principleID = &levelHistory.PrincipleID
	}
	var operatorID *uint
	if levelHistory.OperatorID != 0 {
		operatorID = &levelHistory.OperatorID
	}

	levelHistorieData.LevelID = levelHistory.LevelID
	levelHistorieData.OpCertNum = levelHistory.OpCertNum
	levelHistorieData.NPSN = levelHistory.NPSN
	levelHistorieData.Accreditation = levelHistory.Accreditation
	levelHistorieData.Curriculum = levelHistory.Curriculum
	levelHistorieData.Email = levelHistory.Email
	levelHistorieData.Phone = levelHistory.Phone
	levelHistorieData.PrincipleID = principleID
	levelHistorieData.OperatorID = operatorID
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
