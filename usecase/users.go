package usecase

import (
	"eis-be/models"
	"eis-be/repository"
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/middlewares"
	"reflect"
	"errors"
)

type UsersUsecase interface {
	Register(data dto.RegisterUsersRequest) error
	Login(data dto.LoginUsersRequest) (interface{}, error)
}

type usersUsecase struct {
	usersRepository repository.UsersRepository
}

func NewUsersUsecase(usersRepo repository.UsersRepository) *usersUsecase {
	return &usersUsecase{usersRepository: usersRepo}
}

func validateRegisterUsersRequest(req dto.RegisterUsersRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
		if helpers.IsEmptyField(val.Field(i)) {
			return errors.New("Field can't be empty")
		}
	}
	return nil
}

func validateLoginUsersRequest(req dto.LoginUsersRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
		if helpers.IsEmptyField(val.Field(i)) {
			return errors.New("Field can't be empty")
		}
	}
	return nil
}

func (s *usersUsecase) Register(data dto.RegisterUsersRequest) error {
	e := validateRegisterUsersRequest(data)
	
	if e!= nil {
		return e
	}
	
	userData := models.Users{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
		RoleID: data.RoleID,
	}

	err := s.usersRepository.Create(userData)

	if err != nil {
		return err
	}

	return nil
}

func (s *usersUsecase) Login(data dto.LoginUsersRequest) (interface{}, error) {
	e := validateLoginUsersRequest(data)
	
	if e!= nil {
		return nil, e
	}

	userData := models.Users{
		Email: data.Email,
		Password: data.Password,
	}

	user, err := s.usersRepository.Login(userData)

	if err != nil {
		return nil, err
	}

	token, err := middlewares.CreateToken(user.Email, user.Password)

	if err != nil {
		return nil, err
	}

	userResponse := dto.LoginUsersResponse{
		ID		: user.ID,
		Name 	: user.Name,
		Email	: user.Email,
		Token 	: token,
		RoleID	: user.RoleID,
	}

	return userResponse, nil
}
