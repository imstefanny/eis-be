package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
)

type AcademicStudentsUsecase interface {
	Update(academicStudents []dto.UpdateAcademicStudentsRequest) error
}

type academicStudentsUsecase struct {
	academicStudentsRepository repository.AcademicStudentsRepository
}

func NewAcademicStudentsUsecase(academicStudentsRepo repository.AcademicStudentsRepository) *academicStudentsUsecase {
	return &academicStudentsUsecase{
		academicStudentsRepository: academicStudentsRepo,
	}
}

func (s *academicStudentsUsecase) Update(academicStudents []dto.UpdateAcademicStudentsRequest) error {
	ids := []int{}
	for _, student := range academicStudents {
		ids = append(ids, student.ID)
	}
	academicStudentDatas, err := s.academicStudentsRepository.GetByIDs(ids)
	if err != nil {
		return err
	}
	academicStudentsData := []models.AcademicStudents{}
	for _, academicStudentData := range academicStudentDatas {
		for _, student := range academicStudents {
			if academicStudentData.ID == uint(student.ID) {
				academicStudentData.FirstTermNotes = student.FirstTermNotes
				academicStudentData.SecondTermNotes = student.SecondTermNotes
			}
		}
		academicStudentsData = append(academicStudentsData, academicStudentData)
	}
	err = s.academicStudentsRepository.Update(academicStudentsData)
	if err != nil {
		return err
	}

	return nil
}
