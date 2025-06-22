package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"time"
)

type TermsUsecase interface {
	Update(id int, term dto.UpdateTermRequest) (models.Terms, error)
}

type termsUsecase struct {
	termsRepository repository.TermsRepository
}

func NewTermsUsecase(termsRepo repository.TermsRepository) *termsUsecase {
	return &termsUsecase{
		termsRepository: termsRepo,
	}
}

func (u *termsUsecase) Update(id int, term dto.UpdateTermRequest) (models.Terms, error) {
	termData, err := u.termsRepository.Find(id)
	if err != nil {
		return models.Terms{}, err
	}

	if term.FirstEndDate != "" && term.FirstStartDate != "" {
		parsedFirstStartDate, _ := time.Parse("2006-01-02", term.FirstStartDate)
		termData.FirstStartDate = parsedFirstStartDate
		parsedFirstEndDate, _ := time.Parse("2006-01-02", term.FirstEndDate)
		termData.FirstEndDate = parsedFirstEndDate
	}
	if term.SecondStartDate != "" && term.SecondEndDate != "" {
		parsedSecondStartDate, _ := time.Parse("2006-01-02", term.SecondStartDate)
		termData.SecondStartDate = parsedSecondStartDate
		parsedSecondEndDate, _ := time.Parse("2006-01-02", term.SecondEndDate)
		termData.SecondEndDate = parsedSecondEndDate
	}

	if err := u.termsRepository.Update(termData); err != nil {
		return models.Terms{}, err
	}

	return termData, nil
}
