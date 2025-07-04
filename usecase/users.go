package usecase

import (
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/middlewares"
	"eis-be/models"
	"eis-be/repository"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsersUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	Register(data dto.RegisterUsersRequest) error
	Login(data dto.LoginUsersRequest) (interface{}, error)
	Update(id uint, data dto.UpdateUsersRequest) error
	ChangePassword(id uint, data dto.ChangePasswordRequest) error
	Undelete(id int) error
	Delete(id int) error
}

type usersUsecase struct {
	usersRepository repository.UsersRepository
	rolesRepository repository.RolesRepository
	db              *gorm.DB
}

func NewUsersUsecase(usersRepo repository.UsersRepository, rolesRepo repository.RolesRepository, db *gorm.DB) *usersUsecase {
	return &usersUsecase{
		usersRepository: usersRepo,
		rolesRepository: rolesRepo,
		db:              db,
	}
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	role, _ := s.rolesRepository.FindByName("Applicant")
	userData := models.Users{
		ProfilePic: data.ProfilePic,
		Name:       data.Name,
		Email:      data.Email,
		Password:   string(hashedPassword),
		RoleID:     role.ID,
	}

	_, err = s.usersRepository.Create(s.db, userData)

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
		Email: data.Email,
	}

	user, err := s.usersRepository.Login(userData)

	if err != nil {
		return nil, fmt.Errorf("Invalid email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return nil, fmt.Errorf("Invalid password")
	}

	token, err := middlewares.CreateToken(user.Email, user.ID)

	if err != nil {
		return nil, err
	}

	if user.RoleID == 0 {
		return nil, errors.New("user role not found")
	}
	if user.Role.Permissions == nil {
		return nil, errors.New("user permissions not found")
	}
	permissions := make([]string, len(user.Role.Permissions))
	for i, perm := range user.Role.Permissions {
		permissions[i] = perm.Name
	}

	userResponse := dto.LoginUsersResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Token:       token,
		RoleID:      user.RoleID,
		RoleName:    user.Role.Name,
		Permissions: permissions,
		ProfilePic:  user.ProfilePic,
	}

	return userResponse, nil
}

func (s *usersUsecase) Update(id uint, data dto.UpdateUsersRequest) error {
	userData, _ := s.usersRepository.Find(int(id))

	if data.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		userData.Password = string(hashedPassword)
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

func (s *usersUsecase) ChangePassword(id uint, data dto.ChangePasswordRequest) error {
	userData, err := s.usersRepository.Find(int(id))
	if (data.NewPassword == "") || (err != nil) {
		return errors.New("new password cannot be empty or user not found")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userData.Password = string(hashedPassword)
	err = s.usersRepository.Update(userData)
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

func (s *usersUsecase) Undelete(id int) error {
	err := s.usersRepository.Undelete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *usersUsecase) Delete(id int) error {
	err := s.usersRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
