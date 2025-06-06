package usecase

import (
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/models"
	"eis-be/repository"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type SubjSchedsUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	Create(subjScheds dto.CreateSubjSchedsRequest) error
	Find(id int) (interface{}, error)
	Update(id int, subjSched dto.UpdateSubjSchedsRequest) (models.SubjectSchedules, error)
	UpdateByAcademicID(subjSched dto.UpdateBatchSubjSchedsRequest) (dto.UpdateSubjSchedsResponse, error)
	Delete(id int) error

	// Teacher specific methods
	GetAllByTeacher(teacherUserID int) ([]dto.TeacherStudentGetSubjScheds, error)

	// Student specific methods
	GetScheduleByStudent(studentUserID int) ([]dto.TeacherStudentGetSubjScheds, error)
}

type subjSchedsUsecase struct {
	subjSchedsRepository repository.SubjSchedsRepository
	academicsRepository  repository.AcademicsRepository
	teachersRepository   repository.TeachersRepository
	studentsRepository   repository.StudentsRepository
}

func NewSubjSchedsUsecase(subjSchedsRepo repository.SubjSchedsRepository, academicsRepo repository.AcademicsRepository, teachersRepo repository.TeachersRepository, studentsRepo repository.StudentsRepository) *subjSchedsUsecase {
	return &subjSchedsUsecase{
		subjSchedsRepository: subjSchedsRepo,
		academicsRepository:  academicsRepo,
		teachersRepository:   teachersRepo,
		studentsRepository:   studentsRepo,
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

func (s *subjSchedsUsecase) UpdateByAcademicID(subjSched dto.UpdateBatchSubjSchedsRequest) (dto.UpdateSubjSchedsResponse, error) {
	academicData, err := s.academicsRepository.Find(int(subjSched.AcademicID))

	if err != nil {
		return dto.UpdateSubjSchedsResponse{}, err
	}

	existing := academicData.SubjScheds
	existingIDs := []int{}
	for _, eDetail := range existing {
		existingIDs = append(existingIDs, int(eDetail.ID))
	}
	incomingDetails := subjSched.Entries
	incomingIDs := []int{}
	addIDs := []models.SubjectSchedules{}
	for _, iDetail := range incomingDetails {
		if iDetail.ID != 0 {
			incomingIDs = append(incomingIDs, int(iDetail.ID))
		} else {
			addData := models.SubjectSchedules{
				DisplayName: iDetail.Day + " - " + iDetail.StartHour + " to " + iDetail.EndHour,
				AcademicID:  subjSched.AcademicID,
				SubjectID:   iDetail.SubjectID,
				TeacherID:   iDetail.TeacherID,
				Day:         iDetail.Day,
				StartHour:   iDetail.StartHour,
				EndHour:     iDetail.EndHour,
			}
			addIDs = append(addIDs, addData)
		}
	}
	removeIDs := helpers.Difference(existingIDs, incomingIDs)
	updateIDs := helpers.Intersection(incomingIDs, existingIDs)
	incomingUpdate := []models.SubjectSchedules{}
	for _, iDetail := range incomingDetails {
		for _, id := range updateIDs {
			if int(iDetail.ID) == id {
				incomingUpdate = append(incomingUpdate, models.SubjectSchedules{
					ID:          iDetail.ID,
					DisplayName: iDetail.Day + " - " + iDetail.StartHour + " to " + iDetail.EndHour,
					AcademicID:  subjSched.AcademicID,
					SubjectID:   iDetail.SubjectID,
					TeacherID:   iDetail.TeacherID,
					Day:         iDetail.Day,
					StartHour:   iDetail.StartHour,
					EndHour:     iDetail.EndHour,
				})
			}
		}
	}
	if len(addIDs) == 0 && len(updateIDs) == 0 && len(removeIDs) == 0 {
		return dto.UpdateSubjSchedsResponse{}, fmt.Errorf("no changes detected")
	}

	details := map[string]interface{}{
		"addIDs":         addIDs,
		"updateIDs":      updateIDs,
		"removeIDs":      removeIDs,
		"incomingUpdate": incomingUpdate,
	}
	eTrx := s.subjSchedsRepository.UpdateBatch(details)
	if eTrx != nil {
		return dto.UpdateSubjSchedsResponse{}, eTrx
	}

	academicUpdated, err := s.academicsRepository.Find(int(subjSched.AcademicID))

	if err != nil {
		return dto.UpdateSubjSchedsResponse{}, err
	}

	scheduleDays := map[string][]dto.GetSubjectScheduleEntryResponse{}
	for _, schedule := range academicUpdated.SubjScheds {
		day := schedule.Day
		if _, exists := scheduleDays[day]; !exists {
			scheduleDays[day] = []dto.GetSubjectScheduleEntryResponse{}
		}
		scheduleEntry := dto.GetSubjectScheduleEntryResponse{
			ID:        schedule.ID,
			Subject:   schedule.Subject.Name,
			Teacher:   schedule.Teacher.Name,
			StartHour: schedule.StartHour,
			EndHour:   schedule.EndHour,
		}
		scheduleDays[day] = append(scheduleDays[day], scheduleEntry)
	}
	response := dto.UpdateSubjSchedsResponse{}
	responses := []dto.GetSubjectScheduleResponse{}
	for day, entries := range scheduleDays {
		schedule := dto.GetSubjectScheduleResponse{
			Day:     day,
			Entries: entries,
		}
		responses = append(responses, schedule)
	}
	response.AcademicID = uint(academicUpdated.ID)
	response.Entries = responses

	return response, err
}

func (s *subjSchedsUsecase) Delete(id int) error {
	err := s.subjSchedsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *subjSchedsUsecase) GetAllByTeacher(teacherUserID int) ([]dto.TeacherStudentGetSubjScheds, error) {
	teacherID, err := s.teachersRepository.GetByToken(teacherUserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get teacher by user ID: %w", err)
	}

	subjScheds, err := s.subjSchedsRepository.GetAllByTeacher(int(teacherID.ID))
	if err != nil {
		return nil, err
	}

	response := []dto.TeacherStudentGetSubjScheds{}
	scheduleDays := map[string][]models.SubjectSchedules{}
	for _, subjSched := range subjScheds {
		day := subjSched.Day
		if _, exists := scheduleDays[day]; !exists {
			scheduleDays[day] = []models.SubjectSchedules{}
		}
		scheduleDays[day] = append(scheduleDays[day], subjSched)
	}
	for day, entries := range scheduleDays {
		schedule := dto.TeacherStudentGetSubjScheds{
			Day:     day,
			Details: []dto.TeacherStudentGetSubjSchedsDetails{},
		}
		for _, entry := range entries {
			schedule.Details = append(schedule.Details, dto.TeacherStudentGetSubjSchedsDetails{
				SubjSchedID: entry.ID,
				ClassID:     entry.Academic.ClassroomID,
				Class:       entry.Academic.Classroom.Name,
				SubjectID:   entry.SubjectID,
				Subject:     entry.Subject.Name,
				TeacherID:   entry.TeacherID,
				StartHour:   entry.StartHour,
				EndHour:     entry.EndHour,
			})
		}
		response = append(response, schedule)
	}

	return response, nil
}

func (s *subjSchedsUsecase) GetScheduleByStudent(studentUserID int) ([]dto.TeacherStudentGetSubjScheds, error) {
	student, err := s.studentsRepository.GetByToken(studentUserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get student by user ID: %w", err)
	}

	subjScheds, err := s.subjSchedsRepository.GetScheduleByStudent(int(student.CurrentAcademicID))
	if err != nil {
		return nil, err
	}

	responses := []dto.TeacherStudentGetSubjScheds{}
	scheduleDays := map[string][]models.SubjectSchedules{}
	for _, subjSched := range subjScheds {
		day := subjSched.Day
		if _, exists := scheduleDays[day]; !exists {
			scheduleDays[day] = []models.SubjectSchedules{}
		}
		scheduleDays[day] = append(scheduleDays[day], subjSched)
	}
	for day, entries := range scheduleDays {
		schedule := dto.TeacherStudentGetSubjScheds{
			Day:     day,
			Details: []dto.TeacherStudentGetSubjSchedsDetails{},
		}
		for _, entry := range entries {
			schedule.Details = append(schedule.Details, dto.TeacherStudentGetSubjSchedsDetails{
				SubjSchedID: entry.ID,
				ClassID:     entry.Academic.ClassroomID,
				Class:       entry.Academic.Classroom.Name,
				SubjectID:   entry.SubjectID,
				Subject:     entry.Subject.Name,
				Teacher: 	 	 entry.Teacher.Name,
				TeacherID:   entry.TeacherID,
				StartHour:   entry.StartHour,
				EndHour:     entry.EndHour,
			})
		}
		responses = append(responses, schedule)
	}

	return responses, nil
}
