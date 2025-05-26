package usecase

import (
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/models"
	"eis-be/repository"
	"errors"
	"reflect"
)

type ClassroomsUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	Create(classroom dto.CreateClassroomsRequest) error
	Find(id int) (interface{}, error)
	Update(id int, classroom dto.CreateClassroomsRequest) (models.Classrooms, error)
	Delete(id int) error
}

type classroomsUsecase struct {
	classroomsRepository repository.ClassroomsRepository
}

func NewClassroomsUsecase(classroomsRepo repository.ClassroomsRepository) *classroomsUsecase {
	return &classroomsUsecase{
		classroomsRepository: classroomsRepo,
	}
}

func validateCreateClassroomsRequest(req dto.CreateClassroomsRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
		if helpers.IsEmptyField(val.Field(i)) {
			return errors.New("field can't be empty")
		}
	}
	return nil
}

func (s *classroomsUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	classrooms, total, err := s.classroomsRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return classrooms, total, nil
}

func (s *classroomsUsecase) Create(classroom dto.CreateClassroomsRequest) error {
	e := validateCreateClassroomsRequest(classroom)

	if e != nil {
		return e
	}

	classroomData := models.Classrooms{
		DisplayName: classroom.DisplayName,
		LevelID:     classroom.LevelID,
		Grade:       classroom.Grade,
		Name:        classroom.Name,
	}

	err := s.classroomsRepository.Create(classroomData)

	if err != nil {
		return err
	}

	return nil
}

func (s *classroomsUsecase) Find(id int) (interface{}, error) {
	classroom, err := s.classroomsRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return classroom, nil
}

func (s *classroomsUsecase) Update(id int, classroom dto.CreateClassroomsRequest) (models.Classrooms, error) {
	classroomData, err := s.classroomsRepository.Find(id)

	if err != nil {
		return models.Classrooms{}, err
	}

	classroomData.DisplayName = classroom.DisplayName
	classroomData.LevelID = classroom.LevelID
	classroomData.Grade = classroom.Grade
	classroomData.Name = classroom.Name

	e := s.classroomsRepository.Update(id, classroomData)

	if e != nil {
		return models.Classrooms{}, e
	}

	classroomUpdated, err := s.classroomsRepository.Find(id)

	if err != nil {
		return models.Classrooms{}, err
	}

	return classroomUpdated, nil
}

func (s *classroomsUsecase) Delete(id int) error {
	err := s.classroomsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
