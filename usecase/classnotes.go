package usecase

import (
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/models"
	"eis-be/repository"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type ClassNotesUsecase interface {
	Browse(page, limit int, search string) ([]dto.BrowseClassNotesResponse, int64, error)
	BrowseByAcademicID(academicID, page, limit int, search string) ([]dto.GetClassNotesResponse, int64, error)
	Create(classNote dto.CreateClassNotesRequest) error
	CreateBatch(classNote dto.CreateBatchClassNotesRequest) error
	Find(id int) (dto.GetClassNotesResponse, error)
	Update(id int, classNote dto.CreateClassNotesRequest) (dto.GetClassNotesResponse, error)
	UpdateDetail(id int, classNote dto.CreateClassNotesDetailsRequest) (dto.GetClassNoteEntryResponse, error)
	Delete(id int) error

	// Teacher methods
	FindByTeacher(teacherUserID, schedID int, date string) ([]dto.GetClassNotesResponse, error)
}

type classNotesUsecase struct {
	classNotesRepository  repository.ClassNotesRepository
	academicsRepository   repository.AcademicsRepository
	studentAttsRepository repository.StudentAttsRepository
	teachersRepository    repository.TeachersRepository
}

func NewClassNotesUsecase(classNotesRepo repository.ClassNotesRepository, academicsRepo repository.AcademicsRepository, studentAttsRepo repository.StudentAttsRepository, teachersRepo repository.TeachersRepository) *classNotesUsecase {
	return &classNotesUsecase{
		classNotesRepository:  classNotesRepo,
		academicsRepository:   academicsRepo,
		studentAttsRepository: studentAttsRepo,
		teachersRepository:    teachersRepo,
	}
}

func validateCreateClassNotesRequest(req dto.CreateClassNotesRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func validateCreateBatchClassNotesRequest(req dto.CreateBatchClassNotesRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *classNotesUsecase) Browse(page, limit int, search string) ([]dto.BrowseClassNotesResponse, int64, error) {
	classNotes, total, err := s.classNotesRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	responses := []dto.BrowseClassNotesResponse{}
	for _, classNote := range classNotes {
		academic, err := s.academicsRepository.Find(int(classNote.AcademicID))
		if err != nil {
			return nil, total, err
		}
		response := dto.BrowseClassNotesResponse{
			ID:          classNote.ID,
			DisplayName: classNote.DisplayName,
			AcademicID:  classNote.AcademicID,
			Academic:    academic.DisplayName,
			Date:        classNote.Date,
			CreatedAt:   classNote.CreatedAt,
			UpdatedAt:   classNote.UpdatedAt,
			DeletedAt:   classNote.DeletedAt,
		}
		responses = append(responses, response)
	}

	return responses, total, nil
}

func (s *classNotesUsecase) BrowseByAcademicID(academicID, page, limit int, search string) ([]dto.GetClassNotesResponse, int64, error) {
	classNotes, total, err := s.classNotesRepository.BrowseByAcademicID(academicID, page, limit, search)

	if err != nil {
		return nil, total, err
	}

	responses := []dto.GetClassNotesResponse{}
	for _, classNote := range classNotes {
		details := []dto.GetClassNoteEntryResponse{}
		for _, detail := range classNote.Details {
			detailData := dto.GetClassNoteEntryResponse{
				ID:                detail.ID,
				Subject:           detail.SubjSched.Subject.Name,
				SubjectScheduleId: detail.SubjSchedID,
				Teacher:           detail.SubjSched.Teacher.Name,
				TeacherID:         detail.SubjSched.TeacherID,
				TeacherAct:        detail.Teacher.Name,
				TeacherActID:      detail.TeacherID,
				Materials:         detail.Materials,
				Notes:             detail.Notes,
			}
			details = append(details, detailData)
		}
		absences, _ := s.studentAttsRepository.FindByAcademicDate(academicID, classNote.Date.Format("2006-01-02"))
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
		responses = append(responses, dto.GetClassNotesResponse{
			ID:             classNote.ID,
			AcademicID:     classNote.AcademicID,
			Date:           classNote.Date,
			Details:        details,
			AbsenceCount:   absenceCount,
			AbsenceDetails: absenceDetails,
			CreatedAt:      classNote.CreatedAt,
			UpdatedAt:      classNote.UpdatedAt,
			DeletedAt:      classNote.DeletedAt,
		})
	}

	return responses, total, nil
}

func (s *classNotesUsecase) Create(classNote dto.CreateClassNotesRequest) error {
	e := validateCreateClassNotesRequest(classNote)

	if e != nil {
		return e
	}

	details := []models.ClassNotesDetails{}
	if len(classNote.Details) > 0 {
		for _, detail := range classNote.Details {
			detailData := models.ClassNotesDetails{
				SubjSchedID: detail.SubjSchedID,
				TeacherID:   detail.TeacherID,
				Materials:   detail.Materials,
				Notes:       detail.Notes,
			}
			details = append(details, detailData)
		}
	}

	parsedDate, edate := time.Parse("2006-01-02", classNote.Date)
	if edate != nil {
		return edate
	}

	academic, eAcademic := s.academicsRepository.Find(int(classNote.AcademicID))
	if eAcademic != nil {
		return eAcademic
	}

	classNoteData := models.ClassNotes{
		DisplayName: academic.DisplayName + " - " + classNote.Date,
		AcademicID:  classNote.AcademicID,
		Date:        parsedDate,
		Details:     details,
	}

	err := s.classNotesRepository.Create(classNoteData)

	if err != nil {
		return err
	}

	return nil
}

func (s *classNotesUsecase) CreateBatch(classNote dto.CreateBatchClassNotesRequest) error {
	e := validateCreateBatchClassNotesRequest(classNote)

	if e != nil {
		return e
	}

	parsedDate, edate := time.Parse("2006-01-02", classNote.Date)
	if edate != nil {
		return edate
	}

	academics, err := s.academicsRepository.GetAll()
	if err != nil {
		return err
	}

	classNoteData := []models.ClassNotes{}
	for _, academic := range academics {
		details := []models.ClassNotesDetails{}
		for _, subjSched := range academic.SubjScheds {
			if subjSched.Day != parsedDate.Weekday().String() {
				continue
			}
			detailData := models.ClassNotesDetails{
				SubjSchedID: subjSched.ID,
				TeacherID:   subjSched.TeacherID,
			}
			details = append(details, detailData)
		}
		classNoteData = append(classNoteData, models.ClassNotes{
			DisplayName: academic.DisplayName + " - " + classNote.Date,
			AcademicID:  academic.ID,
			Date:        parsedDate,
			Details:     details,
		})
	}

	err = s.classNotesRepository.CreateBatch(classNoteData)

	if err != nil {
		return err
	}

	return nil
}

func (s *classNotesUsecase) Find(id int) (dto.GetClassNotesResponse, error) {
	classNote, err := s.classNotesRepository.Find(id)

	if err != nil {
		return dto.GetClassNotesResponse{}, err
	}

	details := []dto.GetClassNoteEntryResponse{}
	for _, detail := range classNote.Details {
		detailData := dto.GetClassNoteEntryResponse{
			ID:                detail.ID,
			Subject:           detail.SubjSched.Subject.Name,
			SubjectScheduleId: detail.SubjSchedID,
			Teacher:           detail.SubjSched.Teacher.Name,
			TeacherID:         detail.SubjSched.TeacherID,
			TeacherAct:        detail.Teacher.Name,
			TeacherActID:      detail.TeacherID,
			Materials:         detail.Materials,
			Notes:             detail.Notes,
		}
		details = append(details, detailData)
	}
	response := dto.GetClassNotesResponse{
		ID:         classNote.ID,
		AcademicID: classNote.AcademicID,
		Date:       classNote.Date,
		Details:    details,
		CreatedAt:  classNote.CreatedAt,
		UpdatedAt:  classNote.UpdatedAt,
		DeletedAt:  classNote.DeletedAt,
	}

	return response, nil
}

func (s *classNotesUsecase) Update(id int, classNote dto.CreateClassNotesRequest) (dto.GetClassNotesResponse, error) {
	classNoteData, err := s.classNotesRepository.Find(id)

	if err != nil {
		return dto.GetClassNotesResponse{}, err
	}

	existing := classNoteData.Details
	existingIDs := []int{}
	for _, eDetail := range existing {
		existingIDs = append(existingIDs, int(eDetail.ID))
	}
	incomingDetails := classNote.Details
	incomingIDs := []int{}
	addIDs := []models.ClassNotesDetails{}
	for _, iDetail := range incomingDetails {
		if iDetail.ID != 0 {
			incomingIDs = append(incomingIDs, int(iDetail.ID))
		} else {
			addData := models.ClassNotesDetails{
				NoteID:      classNoteData.ID,
				SubjSchedID: iDetail.SubjSchedID,
				TeacherID:   iDetail.TeacherID,
				Materials:   iDetail.Materials,
				Notes:       iDetail.Notes,
			}
			addIDs = append(addIDs, addData)
		}
	}
	removeIDs := helpers.Difference(existingIDs, incomingIDs)
	updateIDs := helpers.Intersection(incomingIDs, existingIDs)
	incomingUpdate := []models.ClassNotesDetails{}
	for _, iDetail := range incomingDetails {
		for _, id := range updateIDs {
			if int(iDetail.ID) == id {
				incomingUpdate = append(incomingUpdate, models.ClassNotesDetails{
					ID:          iDetail.ID,
					NoteID:      classNoteData.ID,
					SubjSchedID: iDetail.SubjSchedID,
					TeacherID:   iDetail.TeacherID,
					Materials:   iDetail.Materials,
					Notes:       iDetail.Notes,
				})
			}
		}
	}
	if len(addIDs) == 0 && len(updateIDs) == 0 && len(removeIDs) == 0 {
		return dto.GetClassNotesResponse{}, fmt.Errorf("no changes detected")
	}

	details := map[string]interface{}{
		"addIDs":         addIDs,
		"updateIDs":      updateIDs,
		"removeIDs":      removeIDs,
		"incomingUpdate": incomingUpdate,
	}
	eTrx := s.classNotesRepository.Update(id, details)
	if eTrx != nil {
		return dto.GetClassNotesResponse{}, eTrx
	}

	classNoteUpdated, err := s.classNotesRepository.Find(id)

	if err != nil {
		return dto.GetClassNotesResponse{}, err
	}

	updatedDetails := []dto.GetClassNoteEntryResponse{}
	for _, detail := range classNoteUpdated.Details {
		detailData := dto.GetClassNoteEntryResponse{
			ID:                detail.ID,
			Subject:           detail.SubjSched.Subject.Name,
			SubjectScheduleId: detail.SubjSchedID,
			Teacher:           detail.SubjSched.Teacher.Name,
			TeacherID:         detail.SubjSched.TeacherID,
			TeacherAct:        detail.Teacher.Name,
			TeacherActID:      detail.TeacherID,
			Materials:         detail.Materials,
			Notes:             detail.Notes,
		}
		updatedDetails = append(updatedDetails, detailData)
	}
	response := dto.GetClassNotesResponse{
		ID:         classNoteUpdated.ID,
		AcademicID: classNoteUpdated.AcademicID,
		Date:       classNoteUpdated.Date,
		Details:    updatedDetails,
		CreatedAt:  classNoteUpdated.CreatedAt,
		UpdatedAt:  classNoteUpdated.UpdatedAt,
		DeletedAt:  classNoteUpdated.DeletedAt,
	}

	return response, nil
}

func (s *classNotesUsecase) UpdateDetail(id int, classNote dto.CreateClassNotesDetailsRequest) (dto.GetClassNoteEntryResponse, error) {
	classNoteData, err := s.classNotesRepository.FindClassNoteDetail(id)

	if err != nil {
		return dto.GetClassNoteEntryResponse{}, err
	}

	var teacherId uint
	var noteId uint
	if classNote.TeacherID == 0 {
		teacherId = classNoteData.TeacherID
	} else {
		teacherId = classNote.TeacherID
	}

	if classNoteData.NoteID == 0 {
		noteId = classNote.NoteID
	} else {
		noteId = classNoteData.NoteID
	}

	detail := models.ClassNotesDetails{
		ID:          classNote.ID,
		NoteID:      noteId,
		SubjSchedID: classNote.SubjSchedID,
		TeacherID:   teacherId,
		Materials:   classNote.Materials,
		Notes:       classNote.Notes,
	}

	var errTrx error
	if classNote.ID == 0 {
		errTrx = s.classNotesRepository.CreateDetail(detail)
	} else {
		errTrx = s.classNotesRepository.UpdateDetail(detail)
	}

	if errTrx != nil {
		return dto.GetClassNoteEntryResponse{}, errTrx
	}

	classNoteUpdated, err := s.classNotesRepository.FindClassNoteDetail(id)

	if err != nil {
		return dto.GetClassNoteEntryResponse{}, err
	}

	response := dto.GetClassNoteEntryResponse{
		ID:        classNoteUpdated.ID,
		Materials: classNoteUpdated.Materials,
		Notes:     classNoteUpdated.Notes,
	}

	return response, nil
}

func (s *classNotesUsecase) Delete(id int) error {
	err := s.classNotesRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

// Teacher methods
func (s *classNotesUsecase) FindByTeacher(teacherUserID, schedID int, date string) ([]dto.GetClassNotesResponse, error) {
	teacherID, err := s.teachersRepository.GetByToken(teacherUserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get teacher by user ID: %w", err)
	}
	classNotes, err := s.classNotesRepository.FindByTeacher(int(teacherID.ID), schedID, date)

	responses := []dto.GetClassNotesResponse{}
	for _, classNote := range classNotes {
		details := []dto.GetClassNoteEntryResponse{}
		for _, detail := range classNote.Details {
			if detail.TeacherID != teacherID.ID || detail.SubjSchedID != uint(schedID) {
				continue
			}
			detailData := dto.GetClassNoteEntryResponse{
				ID:                detail.ID,
				Subject:           detail.SubjSched.Subject.Name,
				SubjectScheduleId: detail.SubjSchedID,
				Teacher:           detail.SubjSched.Teacher.Name,
				TeacherID:         detail.SubjSched.TeacherID,
				TeacherAct:        detail.Teacher.Name,
				TeacherActID:      detail.TeacherID,
				Materials:         detail.Materials,
				Notes:             detail.Notes,
			}
			details = append(details, detailData)
		}
		response := dto.GetClassNotesResponse{
			ID:         classNote.ID,
			AcademicID: classNote.AcademicID,
			Date:       classNote.Date,
			Details:    details,
			CreatedAt:  classNote.CreatedAt,
			UpdatedAt:  classNote.UpdatedAt,
			DeletedAt:  classNote.DeletedAt,
		}
		responses = append(responses, response)
	}

	return responses, err
}
