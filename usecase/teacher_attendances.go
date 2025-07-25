package usecase

import (
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/models"
	"eis-be/repository"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type TeacherAttsUsecase interface {
	Browse(page, limit int, search string, date string, userId *int) (interface{}, int64, error)
	Create(teacherAtt dto.CreateTeacherAttsRequest) error
	CreateBatch(teacherAtts dto.CreateBatchTeacherAttsRequest) error
	Find(id int) (interface{}, error)
	Update(id int, teacherAtt dto.CreateTeacherAttsRequest) (models.TeacherAttendances, error)
	Delete(id int) error

	GetReport(search, start_date, end_date string, userId *int) ([]dto.GetTeacherAttsReport, error)
}

type teacherAttsUsecase struct {
	teacherAttsRepository repository.TeacherAttsRepository
	teachersRepository    repository.TeachersRepository
	workSchedsRepository  repository.WorkSchedsRepository
}

func NewTeacherAttsUsecase(teacherAttsRepo repository.TeacherAttsRepository, teachersRepo repository.TeachersRepository, workSchedsRepo repository.WorkSchedsRepository) *teacherAttsUsecase {
	return &teacherAttsUsecase{
		teacherAttsRepository: teacherAttsRepo,
		teachersRepository:    teachersRepo,
		workSchedsRepository:  workSchedsRepo,
	}
}

func validateCreateTeacherAttsRequest(req dto.CreateTeacherAttsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func validateCreateBatchTeacherAttsRequest(req dto.CreateBatchTeacherAttsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *teacherAttsUsecase) Browse(page, limit int, search string, date string, userId *int) (interface{}, int64, error) {
	teacherAtts, total, err := s.teacherAttsRepository.Browse(page, limit, search, date, userId)

	if err != nil {
		return nil, total, err
	}

	return teacherAtts, total, nil
}

func (s *teacherAttsUsecase) Create(teacherAtt dto.CreateTeacherAttsRequest) error {
	e := validateCreateTeacherAttsRequest(teacherAtt)

	if e != nil {
		return e
	}

	exist, _ := s.teacherAttsRepository.FindByTeacherIdDate(int(teacherAtt.TeacherID), teacherAtt.Date)
	if exist.ID > 0 {
		return fmt.Errorf("attendance for teacher with ID %d on date %s already exists", teacherAtt.TeacherID, teacherAtt.Date)
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	parseDate, err := time.Parse("2006-01-02", teacherAtt.Date)
	if err != nil {
		return err
	}
	parseInTime, err := time.Parse("15:04:05", teacherAtt.LogInTime)
	if err != nil {
		return err
	}
	if teacherAtt.LogOutTime == "" {
		teacherAtt.LogOutTime = "00:00:00"
	}
	parseOutTime, err := time.Parse("15:04:05", teacherAtt.LogOutTime)
	if err != nil {
		return err
	}

	teacher, err := s.teachersRepository.Find(int(teacherAtt.TeacherID))
	if err != nil {
		return err
	}
	teacherAttData := models.TeacherAttendances{
		DisplayName:       strconv.Itoa(int(teacherAtt.TeacherID)) + " - " + teacher.Name + " / " + parseDate.Format("2006-01-02"),
		TeacherID:         teacherAtt.TeacherID,
		WorkingScheduleID: teacher.WorkSchedID,
		Date:              time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseDate.Hour(), parseDate.Minute(), parseDate.Second(), 0, loc),
		LogInTime:         time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseInTime.Hour(), parseInTime.Minute(), parseInTime.Second(), 0, loc),
		LogOutTime:        time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseOutTime.Hour(), parseOutTime.Minute(), parseOutTime.Second(), 0, loc),
		Remark:            teacherAtt.Remark,
		Note:              teacherAtt.Note,
	}

	err = s.teacherAttsRepository.Create(teacherAttData)

	if err != nil {
		return err
	}

	return nil
}

func (s *teacherAttsUsecase) CreateBatch(teacherAtts dto.CreateBatchTeacherAttsRequest) error {
	e := validateCreateBatchTeacherAttsRequest(teacherAtts)

	if e != nil {
		return e
	}

	teacherAttsData := []models.TeacherAttendances{}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	for _, teacherAtt := range teacherAtts.Entries {
		exist, _ := s.teacherAttsRepository.FindByTeacherIdDate(int(teacherAtt.TeacherID), teacherAtt.Date)
		if exist.ID > 0 {
			continue
		}
		parseDate, err := time.Parse("2006-01-02", teacherAtt.Date)
		if err != nil {
			return err
		}
		parseInTime, err := time.Parse("15:04:05", teacherAtt.LogInTime)
		if err != nil {
			return err
		}
		if teacherAtt.LogOutTime == "" {
			teacherAtt.LogOutTime = "00:00:00"
		}
		parseOutTime, err := time.Parse("15:04:05", teacherAtt.LogOutTime)
		if err != nil {
			return err
		}
		teacher, err := s.teachersRepository.GetByMachineID(int(teacherAtt.TeacherID))
		if err != nil {
			return err
		}
		if teacher.WorkSchedID == 0 {
			return fmt.Errorf("teacher with ID %d does not have a working schedule", teacher.ID)
		}
		workSched, _ := s.workSchedsRepository.Find(int(teacher.WorkSchedID))
		teacherAttData := models.TeacherAttendances{
			DisplayName:       strconv.Itoa(int(teacherAtt.TeacherID)) + " - " + teacher.Name + " / " + parseDate.Format("2006-01-02"),
			TeacherID:         teacher.ID,
			WorkingScheduleID: teacher.WorkSchedID,
			Date:              time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseDate.Hour(), parseDate.Minute(), parseDate.Second(), 0, loc),
			LogInTime:         time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseInTime.Hour(), parseInTime.Minute(), parseInTime.Second(), 0, loc),
			LogOutTime:        time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseOutTime.Hour(), parseOutTime.Minute(), parseOutTime.Second(), 0, loc),
			Remark:            teacherAtt.Remark,
			Note:              teacherAtt.Note,
		}
		remark := helpers.TeacherAttsRemark(teacherAttData, workSched)
		teacherAttData.Remark = remark
		teacherAttsData = append(teacherAttsData, teacherAttData)
	}

	if len(teacherAttsData) == 0 {
		return nil
	}
	err := s.teacherAttsRepository.CreateBatch(teacherAttsData)

	if err != nil {
		return err
	}

	return nil
}

func (s *teacherAttsUsecase) Find(id int) (interface{}, error) {
	teacherAtt, err := s.teacherAttsRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return teacherAtt, nil
}

func (s *teacherAttsUsecase) Update(id int, teacherAtt dto.CreateTeacherAttsRequest) (models.TeacherAttendances, error) {
	teacherAttData, err := s.teacherAttsRepository.Find(id)

	if err != nil {
		return models.TeacherAttendances{}, err
	}

	parseDate, err := time.Parse("2006-01-02", teacherAtt.Date)
	if err != nil {
		return models.TeacherAttendances{}, err
	}
	parseInTime, err := time.Parse("15:04:05", teacherAtt.LogInTime)
	if err != nil {
		return models.TeacherAttendances{}, err
	}
	parseOutTime, err := time.Parse("15:04:05", teacherAtt.LogOutTime)
	if err != nil {
		return models.TeacherAttendances{}, err
	}
	workSched, _ := s.workSchedsRepository.Find(int(teacherAttData.WorkingScheduleID))

	loc, _ := time.LoadLocation("Asia/Jakarta")
	teacherAttData.TeacherID = teacherAtt.TeacherID
	teacherAttData.WorkingScheduleID = teacherAtt.WorkingScheduleID
	teacherAttData.Date = time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseDate.Hour(), parseDate.Minute(), parseDate.Second(), 0, loc)
	teacherAttData.LogInTime = time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseInTime.Hour(), parseInTime.Minute(), parseInTime.Second(), 0, loc)
	teacherAttData.LogOutTime = time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseOutTime.Hour(), parseOutTime.Minute(), parseOutTime.Second(), 0, loc)
	teacherAttData.Note = teacherAtt.Note

	remark := helpers.TeacherAttsRemark(teacherAttData, workSched)
	teacherAttData.Remark = remark

	e := s.teacherAttsRepository.Update(id, teacherAttData)

	if e != nil {
		return models.TeacherAttendances{}, e
	}

	teacherAttUpdated, err := s.teacherAttsRepository.Find(id)

	if err != nil {
		return models.TeacherAttendances{}, err
	}

	return teacherAttUpdated, nil
}

func (s *teacherAttsUsecase) Delete(id int) error {
	err := s.teacherAttsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *teacherAttsUsecase) GetReport(search, start_date, end_date string, userId *int) ([]dto.GetTeacherAttsReport, error) {
	teacherAtts, err := s.teacherAttsRepository.BrowseReport(search, start_date, end_date, userId)
	if err != nil {
		return []dto.GetTeacherAttsReport{}, err
	}

	attMap := map[string]*dto.GetTeacherAttsReport{}
	for _, att := range teacherAtts {
		if _, exists := attMap[att.Teacher.Name]; !exists {
			overall := helpers.CountWorkdays(start_date, end_date, att.WorkingSchedule)
			attMap[att.Teacher.Name] = &dto.GetTeacherAttsReport{
				Teacher:         att.Teacher.Name,
				LateCount:       0,
				EarlyLeaveCount: 0,
				AbsenceCount:    0,
				PresentCount:    0,
				TotalCount:      overall,
			}
		}
		if strings.Contains(att.Remark, "Terlambat") {
			attMap[att.Teacher.Name].LateCount++
		}
		if strings.Contains(att.Remark, "Pulang Cepat") {
			attMap[att.Teacher.Name].EarlyLeaveCount++
		}
		attMap[att.Teacher.Name].PresentCount++
	}

	responses := []dto.GetTeacherAttsReport{}
	for _, entries := range attMap {
		if entries.TotalCount != 0 {
			entries.AbsenceCount = entries.TotalCount - entries.PresentCount
		}
		responses = append(responses, *entries)
	}

	return responses, nil
}
