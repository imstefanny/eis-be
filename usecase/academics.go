package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type AcademicsUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	Create(academic dto.CreateAcademicsRequest) error
	CreateBatch(batch dto.CreateBatchAcademicsRequest) error
	Find(id int) (interface{}, error)
	Update(id int, academic dto.CreateAcademicsRequest) (models.Academics, error)
	Delete(id int) error
}

type academicsUsecase struct {
	academicsRepository  repository.AcademicsRepository
	studentsRepository   repository.StudentsRepository
	classroomsRepository repository.ClassroomsRepository
}

func NewAcademicsUsecase(academicsRepo repository.AcademicsRepository, studentsRepo repository.StudentsRepository, classroomsRepo repository.ClassroomsRepository) *academicsUsecase {
	return &academicsUsecase{
		academicsRepository:  academicsRepo,
		studentsRepository:   studentsRepo,
		classroomsRepository: classroomsRepo,
	}
}

func validateCreateAcademicsRequest(req dto.CreateAcademicsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func validateBatchCreateAcademicsRequest(req dto.CreateBatchAcademicsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *academicsUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	academics, total, err := s.academicsRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return academics, total, nil
}

func (s *academicsUsecase) Create(academic dto.CreateAcademicsRequest) error {
	e := validateCreateAcademicsRequest(academic)

	if e != nil {
		return e
	}

	students := []models.Students{}
	if len(academic.Students) > 0 {
		studentsData, err := s.studentsRepository.GetByIds(academic.Students)
		if err != nil {
			return err
		}
		if len(studentsData) == 0 {
			return fmt.Errorf("students not found")
		}
		students = studentsData
	}

	academicData := models.Academics{
		DisplayName:       academic.DisplayName,
		StartYear:         academic.StartYear,
		EndYear:           academic.EndYear,
		ClassroomID:       academic.ClassroomID,
		Major:             academic.Major,
		HomeroomTeacherID: academic.HomeroomTeacherID,
		Students:          students,
	}

	err := s.academicsRepository.Create(academicData)

	if err != nil {
		return err
	}

	return nil
}

func (s *academicsUsecase) CreateBatch(batch dto.CreateBatchAcademicsRequest) error {
	e := validateBatchCreateAcademicsRequest(batch)
	if e != nil {
		return e
	}

	classrooms, eClass := s.classroomsRepository.GetAll()
	if eClass != nil {
		return eClass
	}

	academicsData := []models.Academics{}

	for _, classroom := range classrooms {
		academic := models.Academics{
			DisplayName:       classroom.DisplayName + " " + batch.StartYear + " - " + batch.EndYear,
			StartYear:         batch.StartYear,
			EndYear:           batch.EndYear,
			ClassroomID:       classroom.ID,
			Major:             "General",
		}
		academicsData = append(academicsData, academic)
	}

	err := s.academicsRepository.CreateBatch(academicsData)
	if err != nil {
		return err
	}

	return nil
}

func (s *academicsUsecase) Find(id int) (interface{}, error) {
	academic, err := s.academicsRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return academic, nil
}

func (s *academicsUsecase) Update(id int, academic dto.CreateAcademicsRequest) (models.Academics, error) {
	academicData, err := s.academicsRepository.Find(id)

	if err != nil {
		return models.Academics{}, err
	}

	students := []models.Students{}
	if len(academic.Students) > 0 {
		studentsData, e := s.studentsRepository.GetByIds(academic.Students)
		if e != nil {
			return models.Academics{}, e
		}
		if len(studentsData) == 0 {
			return models.Academics{}, fmt.Errorf("Students not found")
		}
		students = studentsData
	}

	academicData.DisplayName = academic.DisplayName
	academicData.StartYear = academic.StartYear
	academicData.EndYear = academic.EndYear
	academicData.ClassroomID = academic.ClassroomID
	academicData.Major = academic.Major
	academicData.HomeroomTeacherID = academic.HomeroomTeacherID
	academicData.Students = students

	e := s.academicsRepository.Update(id, academicData)

	if e != nil {
		return models.Academics{}, e
	}

	academicUpdated, err := s.academicsRepository.Find(id)

	if err != nil {
		return models.Academics{}, err
	}

	return academicUpdated, nil
}

func (s *academicsUsecase) Delete(id int) error {
	err := s.academicsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
