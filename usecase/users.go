package usecase

import (
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/middlewares"
	"eis-be/models"
	"eis-be/repository"
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UsersUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	Register(data dto.RegisterUsersRequest) error
	Login(data dto.LoginUsersRequest) (interface{}, error)
	Update(id uint, data dto.UpdateUsersRequest) error
}

type usersUsecase struct {
	usersRepository repository.UsersRepository
	db              *gorm.DB
}

func NewUsersUsecase(usersRepo repository.UsersRepository, db *gorm.DB) *usersUsecase {
	return &usersUsecase{usersRepository: usersRepo, db: db}
}

func validateRegisterUsersRequest(req dto.RegisterUsersRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func validateLoginUsersRequest(req dto.LoginUsersRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
		if helpers.IsEmptyField(val.Field(i)) {
			return errors.New("field can't be empty")
		}
	}
	return nil
}

func (s *usersUsecase) Register(data dto.RegisterUsersRequest) error {
	e := validateRegisterUsersRequest(data)

	if e != nil {
		return e
	}

	userData := models.Users{
		ProfilePic: data.ProfilePic,
		Name:       data.Name,
		Email:      data.Email,
		Password:   data.Password,
		RoleID:     data.RoleID,
	}

	_, err := s.usersRepository.Create(s.db, userData)

	if err != nil {
		return err
	}

	return nil
}

func (s *usersUsecase) Login(data dto.LoginUsersRequest) (interface{}, error) {
	e := validateLoginUsersRequest(data)

	if e != nil {
		return nil, e
	}

	userData := models.Users{
		Email:    data.Email,
		Password: data.Password,
	}

	user, err := s.usersRepository.Login(userData)

	if err != nil {
		return nil, err
	}

	token, err := middlewares.CreateToken(user.Email, user.ID)

	if err != nil {
		return nil, err
	}

	userResponse := dto.LoginUsersResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Token:  token,
		RoleID: user.RoleID,
	}

	return userResponse, nil
}

func (s *usersUsecase) Update(id uint, data dto.UpdateUsersRequest) error {
	userData, _ := s.usersRepository.Find(int(id))
	if data.Password != "" {
		userData.Password = data.Password
	}
	if data.ProfilePic != "" {
		userData.ProfilePic = data.ProfilePic
	}
	userData.Name = data.Name
	userData.RoleID = data.RoleID

	err := s.usersRepository.Update(userData)

	if err != nil {
		return err
	}

	return nil
}

func (s *usersUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	teachers, total, err := s.usersRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return teachers, total, nil
}
