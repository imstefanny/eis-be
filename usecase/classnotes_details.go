package usecase

import (
	"eis-be/dto"
	"eis-be/repository"
)

type ClassNotesDetailsUsecase interface {
	GetAllByTeacher(id int, date string) ([]dto.GetTeacherSchedsResponse, error)
}

type classNotesDetailsUsecase struct {
	classNotesDetailsRepository repository.ClassNotesDetailsRepository
	teachersRepository          repository.TeachersRepository
}

func NewClassNotesDetailsUsecase(classNotesDetailsRepo repository.ClassNotesDetailsRepository, teachersRepo repository.TeachersRepository) *classNotesDetailsUsecase {
	return &classNotesDetailsUsecase{
		classNotesDetailsRepository: classNotesDetailsRepo,
		teachersRepository:          teachersRepo,
	}
}

func (u *classNotesDetailsUsecase) GetAllByTeacher(id int, date string) ([]dto.GetTeacherSchedsResponse, error) {
	teacherID, err := u.teachersRepository.GetByToken(id)
	if err != nil {
		return nil, err
	}

	details, err := u.classNotesDetailsRepository.GetAllByTeacher(int(teacherID.ID), date)
	if err != nil {
		return nil, err
	}

	var response []dto.GetTeacherSchedsResponse
	for _, detail := range details {
		response = append(response, dto.GetTeacherSchedsResponse{
			ID:          detail.ID,
			NoteID:      detail.NoteID,
			Date:        detail.Note.Date,
			Day:         detail.SubjSched.Day,
			Class:       detail.SubjSched.Academic.Classroom.Name,
			Subject:     detail.SubjSched.Subject.Name,
			SubjSchedID: detail.SubjSchedID,
			Teacher:     detail.Teacher.Name,
			StartHour:   detail.SubjSched.StartHour,
			EndHour:     detail.SubjSched.EndHour,
			Materials:   detail.Materials,
			Notes:       detail.Notes,
		})
	}

	return response, nil
}
