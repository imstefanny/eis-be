package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"

	"github.com/go-playground/validator/v10"
)

type SubjSchedsUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	Create(subjScheds dto.CreateSubjSchedsRequest) error
	Find(id int) (interface{}, error)
	Update(id int, subjSched dto.UpdateSubjSchedsRequest) (models.SubjectSchedules, error)
	Delete(id int) error
}

type subjSchedsUsecase struct {
	subjSchedsRepository repository.SubjSchedsRepository
}

func NewSubjSchedsUsecase(subjSchedsRepo repository.SubjSchedsRepository) *subjSchedsUsecase {
	return &subjSchedsUsecase{
		subjSchedsRepository: subjSchedsRepo,
	}
}

func validateCreateSubjSchedsRequest(req dto.CreateSubjSchedsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *subjSchedsUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	subjScheds, total, err := s.subjSchedsRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return subjScheds, total, nil
}

func (s *subjSchedsUsecase) Create(subjScheds dto.CreateSubjSchedsRequest) error {
	e := validateCreateSubjSchedsRequest(subjScheds)

	if e != nil {
		return e
	}

	subjSchedsData := []models.SubjectSchedules{}
	for _, sched := range subjScheds.Schedules {
		for _, entry := range sched.Entries {
			subjSchedData := models.SubjectSchedules{
				DisplayName: sched.Day + " - " + entry.StartHour + " to " + entry.EndHour,
				AcademicID:  subjScheds.AcademicID,
				SubjectID:   entry.SubjectID,
				TeacherID:   entry.TeacherID,
				Day:         sched.Day,
				StartHour:   entry.StartHour,
				EndHour:     entry.EndHour,
			}
			subjSchedsData = append(subjSchedsData, subjSchedData)
		}
	}

	err := s.subjSchedsRepository.Create(subjSchedsData)

	if err != nil {
		return err
	}

	return nil
}

func (s *subjSchedsUsecase) Find(id int) (interface{}, error) {
	subjSched, err := s.subjSchedsRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return subjSched, nil
}

func (s *subjSchedsUsecase) Update(id int, subjSched dto.UpdateSubjSchedsRequest) (models.SubjectSchedules, error) {
	subjSchedData, err := s.subjSchedsRepository.Find(id)

	if err != nil {
		return models.SubjectSchedules{}, err
	}

	subjSchedData.AcademicID = subjSched.AcademicID
	subjSchedData.SubjectID = subjSched.SubjectID
	subjSchedData.TeacherID = subjSched.TeacherID
	subjSchedData.Day = subjSched.Day
	subjSchedData.StartHour = subjSched.StartHour
	subjSchedData.EndHour = subjSched.EndHour
	e := s.subjSchedsRepository.Update(id, subjSchedData)

	if e != nil {
		return models.SubjectSchedules{}, e
	}

	subjSchedUpdated, err := s.subjSchedsRepository.Find(id)

	if err != nil {
		return models.SubjectSchedules{}, err
	}

	return subjSchedUpdated, nil
}

func (s *subjSchedsUsecase) Delete(id int) error {
	err := s.subjSchedsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
