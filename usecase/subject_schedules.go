package usecase

import (
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/models"
	"eis-be/repository"
	"errors"
	"reflect"
)

type SubjSchedsUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	Create(subjSched dto.CreateSubjSchedsRequest) error
	Find(id int) (interface{}, error)
	Update(id int, subjSched dto.CreateSubjSchedsRequest) (models.SubjectSchedules, error)
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
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
		if helpers.IsEmptyField(val.Field(i)) {
			return errors.New("field can't be empty")
		}
	}
	return nil
}

func (s *subjSchedsUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	subjScheds, total, err := s.subjSchedsRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return subjScheds, total, nil
}

func (s *subjSchedsUsecase) Create(subjSched dto.CreateSubjSchedsRequest) error {
	e := validateCreateSubjSchedsRequest(subjSched)

	if e != nil {
		return e
	}

	subjSchedData := models.SubjectSchedules{
		DisplayName: subjSched.DisplayName,
		AcademicID:  subjSched.AcademicID,
		SubjectID:   subjSched.SubjectID,
		TeacherID:   subjSched.TeacherID,
		Day:         subjSched.Day,
		StartHour:   subjSched.StartHour,
		EndHour:     subjSched.EndHour,
	}

	err := s.subjSchedsRepository.Create(subjSchedData)

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

func (s *subjSchedsUsecase) Update(id int, subjSched dto.CreateSubjSchedsRequest) (models.SubjectSchedules, error) {
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
