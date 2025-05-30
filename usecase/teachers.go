package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type TeachersUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	GetByToken(id int) (interface{}, error)
	Create(teacher dto.CreateTeachersRequest, claims jwt.MapClaims) error
	Find(id int) (interface{}, error)
	Update(id int, teacher dto.CreateTeachersRequest) (models.Teachers, error)
	Delete(id int) error
}

type teachersUsecase struct {
	teachersRepository repository.TeachersRepository
	usersRepository    repository.UsersRepository
	db                 *gorm.DB
}

func NewTeachersUsecase(teachersRepo repository.TeachersRepository, usersRepo repository.UsersRepository, db *gorm.DB) *teachersUsecase {
	return &teachersUsecase{
		teachersRepository: teachersRepo,
		usersRepository:    usersRepo,
		db:                 db,
	}
}

func validateCreateTeachersRequest(req dto.CreateTeachersRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *teachersUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	teachers, total, err := s.teachersRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return teachers, total, nil
}

func (s *teachersUsecase) GetByToken(id int) (interface{}, error) {
	teacher, err := s.teachersRepository.GetByToken(id)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func (s *teachersUsecase) Find(id int) (interface{}, error) {
	teacher, err := s.teachersRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func (s *teachersUsecase) Create(teacher dto.CreateTeachersRequest, claims jwt.MapClaims) error {
	e := validateCreateTeachersRequest(teacher)

	if e != nil {
		return e
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	userData := models.Users{
		ProfilePic: teacher.ProfilePic,
		Name:       teacher.Name,
		Email:      teacher.Email,
		Password:   "123456",
	}
	userID, eUser := s.usersRepository.Create(tx, userData)
	if eUser != nil {
		tx.Rollback()
		return eUser
	}

	teacherData := models.Teachers{
		IdentityNo:  teacher.IdentityNo,
		Name:        teacher.Name,
		NUPTK:       teacher.NUPTK,
		Phone:       teacher.Phone,
		Email:       teacher.Email,
		Address:     teacher.Address,
		JobTitle:    teacher.JobTitle,
		LevelID:     teacher.LevelID,
		UserID:      userID,
		WorkSchedID: teacher.WorkSchedID,
		ProfilePic:  teacher.ProfilePic,
		MachineID:   teacher.MachineID,
	}

	err := s.teachersRepository.Create(tx, teacherData)

	if err != nil {
		tx.Rollback()
		return err
	}

	eCommit := tx.Commit()
	if eCommit.Error != nil {
		tx.Rollback()
		return eCommit.Error
	}

	return nil
}

func (s *teachersUsecase) Update(id int, teacher dto.CreateTeachersRequest) (models.Teachers, error) {
	teacherData, err := s.teachersRepository.Find(id)

	if err != nil {
		return models.Teachers{}, err
	}

	teacherData.Name = teacher.Name
	teacherData.NUPTK = teacher.NUPTK
	teacherData.Phone = teacher.Phone
	teacherData.Email = teacher.Email
	teacherData.Address = teacher.Address
	teacherData.JobTitle = teacher.JobTitle
	teacherData.LevelID = teacher.LevelID
	teacherData.WorkSchedID = teacher.WorkSchedID
	teacherData.ProfilePic = teacher.ProfilePic
	teacherData.DeletedAt = teacher.DeletedAt
	teacherData.MachineID = teacher.MachineID

	errUnscope := s.teachersRepository.UndeleteTeacher(id)
	if errUnscope != nil {
		return models.Teachers{}, errUnscope
	}
	e := s.teachersRepository.Update(id, teacherData)

	if e != nil {
		return models.Teachers{}, e
	}

	teacherUpdated, err := s.teachersRepository.Find(id)

	if err != nil {
		return models.Teachers{}, err
	}

	return teacherUpdated, nil
}

func (s *teachersUsecase) Delete(id int) error {
	err := s.teachersRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
