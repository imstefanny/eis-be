package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
)

type ClassroomsUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	Create(classroom dto.CreateClassroomsRequest) error
	Find(id int) (interface{}, error)
	Update(id int, classroom dto.CreateClassroomsRequest) (models.Classrooms, error)
	Delete(id int) error
	UnDelete(id int) error
}

type classroomsUsecase struct {
	classroomsRepository repository.ClassroomsRepository
	levelsRepository     repository.LevelsRepository
}

func NewClassroomsUsecase(classroomsRepo repository.ClassroomsRepository, levelsRepo repository.LevelsRepository) *classroomsUsecase {
	return &classroomsUsecase{
		classroomsRepository: classroomsRepo,
		levelsRepository:     levelsRepo,
	}
}

func (s *classroomsUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	classrooms, total, err := s.classroomsRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return classrooms, total, nil
}

func (s *classroomsUsecase) Create(classroom dto.CreateClassroomsRequest) error {
	level, err := s.levelsRepository.Find(int(classroom.LevelID))
	if err != nil {
		return err
	}

	classroomData := models.Classrooms{
		DisplayName: level.Name + " / " + classroom.Grade + " - " + classroom.Name,
		LevelID:     classroom.LevelID,
		Grade:       classroom.Grade,
		Name:        classroom.Name,
	}

	err = s.classroomsRepository.Create(classroomData)

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

	level, err := s.levelsRepository.Find(int(classroom.LevelID))
	if err != nil {
		return models.Classrooms{}, err
	}

	classroomData.DisplayName = level.Name + " / " + classroom.Grade + " - " + classroom.Name
	classroomData.LevelID = classroom.LevelID
	classroomData.Grade = classroom.Grade
	classroomData.Name = classroom.Name
	classroomData.DeletedAt = classroom.DeletedAt

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

func (s *classroomsUsecase) UnDelete(id int) error {
	err := s.classroomsRepository.UnDelete(id)

	if err != nil {
		return err
	}

	return nil
}
