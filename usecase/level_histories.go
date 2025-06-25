package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
)

type LevelHistoriesUsecase interface {
	Create(levelHistory dto.CreateLevelHistoriesRequest) error
}

type levelHistoriesUsecase struct {
	levelHistoriesRepository repository.LevelHistoriesRepository
}

func NewLevelHistoriesUsecase(levelHistoriesRepo repository.LevelHistoriesRepository) *levelHistoriesUsecase {
	return &levelHistoriesUsecase{
		levelHistoriesRepository: levelHistoriesRepo,
	}
}

func (s *levelHistoriesUsecase) Create(levelHistory dto.CreateLevelHistoriesRequest) error {
	levelHistoryData := models.LevelHistories{
		LevelID:       levelHistory.LevelID,
		OpCertNum:     levelHistory.OpCertNum,
		NPSN:          levelHistory.NPSN,
		Accreditation: levelHistory.Accreditation,
		Curriculum:    levelHistory.Curriculum,
		Email:         levelHistory.Email,
		Phone:         levelHistory.Phone,
		PrincipleID:   levelHistory.PrincipleID,
		OperatorID:    levelHistory.OperatorID,
	}

	err := s.levelHistoriesRepository.Create(levelHistory.LevelID, levelHistoryData)

	if err != nil {
		return err
	}

	return nil
}
