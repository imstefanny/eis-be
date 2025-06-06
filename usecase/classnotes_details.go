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
	studentAttsRepository       repository.StudentAttsRepository
}

func NewClassNotesDetailsUsecase(classNotesDetailsRepo repository.ClassNotesDetailsRepository, teachersRepo repository.TeachersRepository, studentAttsRepo repository.StudentAttsRepository) *classNotesDetailsUsecase {
	return &classNotesDetailsUsecase{
		classNotesDetailsRepository: classNotesDetailsRepo,
		teachersRepository:          teachersRepo,
		studentAttsRepository:       studentAttsRepo,
	}
}

func (s *classNotesDetailsUsecase) GetAllByTeacher(id int, date string) ([]dto.GetTeacherSchedsResponse, error) {
	teacherID, err := s.teachersRepository.GetByToken(id)
	if err != nil {
		return nil, err
	}

	details, err := s.classNotesDetailsRepository.GetAllByTeacher(int(teacherID.ID), date)
	if err != nil {
		return nil, err
	}

	var response []dto.GetTeacherSchedsResponse
	for _, detail := range details {
		absences, _ := s.studentAttsRepository.FindByAcademicDate(int(detail.AcademicID), detail.Date.Format("2006-01-02"))
		absenceCount := []dto.GetClassNoteAbsenceResponse{}
		absenceDetails := []dto.GetClassNoteAbsenceDetails{}
		absenceCountMap := make(map[string]int)
		for _, absence := range absences {
			absenceCountMap[absence.Status]++
			if absence.Status == "Permission" || absence.Status == "Alpha" || absence.Status == "Sick" {
				absenceCountMap["Leaves"]++
				absenceDetails = append(absenceDetails, dto.GetClassNoteAbsenceDetails{
					ID:        absence.ID,
					StudentID: absence.StudentID,
					FullName:  absence.Student.FullName,
					Status:    absence.Status,
					Remarks:   absence.Remarks,
				})
			}
		}
		for status, count := range absenceCountMap {
			absenceCount = append(absenceCount, dto.GetClassNoteAbsenceResponse{
				Status: status,
				Total:  count,
			})
		}
		response = append(response, dto.GetTeacherSchedsResponse{
			ID:             detail.ID,
			AcademicID: 	 	detail.AcademicID,
			NoteID:         detail.NoteID,
			Date:           detail.Date,
			Day:            detail.Day,
			Class:          detail.Class,
			Subject:        detail.Subject,
			SubjSchedID:    detail.SubjSchedID,
			TeacherID:      detail.TeacherID,
			Teacher:        detail.Teacher,
			TeacherActID:   detail.TeacherActID,
			StartHour:      detail.StartHour,
			EndHour:        detail.EndHour,
			Materials:      detail.Materials,
			Notes:          detail.Notes,
			AbsenceCount:   absenceCount,
			AbsenceDetails: absenceDetails,
		})
	}

	return response, nil
}
