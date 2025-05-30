package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type TeacherAttsUsecase interface {
	Browse(page, limit int, search string, date string) (interface{}, int64, error)
	Create(teacherAtt dto.CreateTeacherAttsRequest) error
	CreateBatch(teacherAtts dto.CreateBatchTeacherAttsRequest) error
	Find(id int) (interface{}, error)
	Update(id int, teacherAtt dto.CreateTeacherAttsRequest) (models.TeacherAttendances, error)
	Delete(id int) error
}

type teacherAttsUsecase struct {
	teacherAttsRepository repository.TeacherAttsRepository
	teacherRepository     repository.TeachersRepository
}

func NewTeacherAttsUsecase(teacherAttsRepo repository.TeacherAttsRepository, teachersRepo repository.TeachersRepository) *teacherAttsUsecase {
	return &teacherAttsUsecase{
		teacherAttsRepository: teacherAttsRepo,
		teacherRepository:     teachersRepo,
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

func (s *teacherAttsUsecase) Browse(page, limit int, search string, date string) (interface{}, int64, error) {
	blogs, total, err := s.teacherAttsRepository.Browse(page, limit, search, date)

	if err != nil {
		return nil, total, err
	}

	return blogs, total, nil
}

func (s *teacherAttsUsecase) Create(teacherAtt dto.CreateTeacherAttsRequest) error {
	e := validateCreateTeacherAttsRequest(teacherAtt)

	if e != nil {
		return e
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

	teacher, err := s.teacherRepository.Find(int(teacherAtt.TeacherID))
	if err != nil {
		return err
	}
	teacherAttData := models.TeacherAttendances{
		DisplayName:       teacher.Name + " - " + parseDate.Format("2006-01-02"),
		TeacherID:         teacherAtt.TeacherID,
		WorkingScheduleID: teacherAtt.WorkingScheduleID,
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
	for _, teacherAtt := range teacherAtts.Entries {
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
		teacher, err := s.teacherRepository.GetByMachineID(int(teacherAtt.TeacherID))
		if err != nil {
			return err
		}
		teacherAttData := models.TeacherAttendances{
			DisplayName:       teacher.Name + " - " + parseDate.Format("2006-01-02"),
			TeacherID:         teacher.ID,
			WorkingScheduleID: teacherAtt.WorkingScheduleID,
			Date:              parseDate,
			LogInTime:         time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseInTime.Hour(), parseInTime.Minute(), parseInTime.Second(), 0, parseDate.Location()),
			LogOutTime:        time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseOutTime.Hour(), parseOutTime.Minute(), parseOutTime.Second(), 0, parseDate.Location()),
			Remark:            teacherAtt.Remark,
			Note:              teacherAtt.Note,
		}
		teacherAttsData = append(teacherAttsData, teacherAttData)
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

	loc, _ := time.LoadLocation("Asia/Jakarta")
	teacherAttData.TeacherID = teacherAtt.TeacherID
	teacherAttData.WorkingScheduleID = teacherAtt.WorkingScheduleID
	teacherAttData.Date = time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseDate.Hour(), parseDate.Minute(), parseDate.Second(), 0, loc)
	teacherAttData.LogInTime = time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseInTime.Hour(), parseInTime.Minute(), parseInTime.Second(), 0, loc)
	teacherAttData.LogOutTime = time.Date(parseDate.Year(), parseDate.Month(), parseDate.Day(), parseOutTime.Hour(), parseOutTime.Minute(), parseOutTime.Second(), 0, loc)
	teacherAttData.Remark = teacherAtt.Remark
	teacherAttData.Note = teacherAtt.Note

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
