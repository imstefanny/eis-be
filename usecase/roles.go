package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type RolesUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	GetAllPermissions() ([]dto.GetPermissionsResponse, error)
	Create(role dto.CreateRolesRequest) error
	Find(id int) (dto.GetRolesResponse, error)
	Update(id int, role dto.CreateRolesRequest) (models.Roles, error)
	Delete(id int) error
}

type rolesUsecase struct {
	rolesRepository       repository.RolesRepository
	permissionsRepository repository.PermissionsRepository
}

func NewRolesUsecase(rolesRepo repository.RolesRepository, permissionsRepo repository.PermissionsRepository) *rolesUsecase {
	return &rolesUsecase{
		rolesRepository:       rolesRepo,
		permissionsRepository: permissionsRepo,
	}
}

func validateCreateRolesRequest(req dto.CreateRolesRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *rolesUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	roles, total, err := s.rolesRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return roles, total, nil
}

func (s *rolesUsecase) GetAllPermissions() ([]dto.GetPermissionsResponse, error) {
	permissions, err := s.permissionsRepository.GetAll()

	if err != nil {
		return nil, err
	}

	response := []dto.GetPermissionsResponse{}
	for _, permission := range permissions {
		response = append(response, dto.GetPermissionsResponse{
			ID:   int(permission.ID),
			Name: permission.Name,
		})
	}

	return response, nil
}

func (s *rolesUsecase) Create(role dto.CreateRolesRequest) error {
	e := validateCreateRolesRequest(role)

	if e != nil {
		return e
	}

	permissions := []models.Permissions{}
	if len(role.Permissions) > 0 {
		permissionsData, e := s.permissionsRepository.GetByIds(role.Permissions)
		if e != nil {
			return e
		}
		if len(permissionsData) == 0 {
			return models.ErrPermissionsNotFound{}
		}
		permissions = permissionsData
	}
	roleData := models.Roles{
		Name:        role.Name,
		Permissions: permissions,
	}

	err := s.rolesRepository.Create(roleData)

	if err != nil {
		return err
	}

	return nil
}

func (s *rolesUsecase) Find(id int) (dto.GetRolesResponse, error) {
	role, err := s.rolesRepository.Find(id)

	permissions := []dto.GetPermissionsResponse{}
	for _, permission := range role.Permissions {
		permissions = append(permissions, dto.GetPermissionsResponse{
			ID:   int(permission.ID),
			Name: permission.Name,
		})
	}

	response := dto.GetRolesResponse{
		ID:          int(role.ID),
		Name:        role.Name,
		Permissions: permissions,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
		DeletedAt:   role.DeletedAt,
	}

	if err != nil {
		return dto.GetRolesResponse{}, err
	}

	return response, nil
}

func (s *rolesUsecase) Update(id int, role dto.CreateRolesRequest) (models.Roles, error) {
	roleData, err := s.rolesRepository.Find(id)

	if err != nil {
		return models.Roles{}, err
	}

	if roleData.Name == "Applicant" || roleData.Name == "Student" || roleData.Name == "Admin" {
		return models.Roles{}, models.ErrCannotUpdateRole{}
	}

	permissions := []models.Permissions{}
	if len(role.Permissions) > 0 {
		fmt.Println("more than 1 permissions")
		permissionsData, e := s.permissionsRepository.GetByIds(role.Permissions)
		if e != nil {
			return models.Roles{}, e
		}
		if len(permissionsData) == 0 {
			return models.Roles{}, models.ErrPermissionsNotFound{}
		}
		permissions = permissionsData
	}

	roleData.Name = role.Name
	roleData.Permissions = permissions
	e := s.rolesRepository.Update(id, roleData)

	if e != nil {
		return models.Roles{}, e
	}

	roleUpdated, err := s.rolesRepository.Find(id)

	if err != nil {
		return models.Roles{}, err
	}

	return roleUpdated, nil
}

func (s *rolesUsecase) Delete(id int) error {
	roleData, _ := s.rolesRepository.Find(id)
	if roleData.Name == "Applicant" || roleData.Name == "Student" || roleData.Name == "Admin" || roleData.Name == "Principal" || roleData.Name == "Teacher" {
		return fmt.Errorf("cannot delete role %s", roleData.Name)
	}

	err := s.rolesRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
