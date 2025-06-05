package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type RolesController interface {
}

type rolesController struct {
	useCase usecase.RolesUsecase
}

func NewRolesController(rolesUsecase usecase.RolesUsecase) *rolesController {
	return &rolesController{rolesUsecase}
}

func (u *rolesController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")

	roles, total, err := u.useCase.Browse(page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  roles,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *rolesController) GetAllPermissions(c echo.Context) error {
	permissions, err := u.useCase.GetAllPermissions()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": permissions,
	})
}

func (u *rolesController) Create(c echo.Context) error {
	role := dto.CreateRolesRequest{}

	if err := c.Bind(&role); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *rolesController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	roles, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": roles,
	})
}

func (u *rolesController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	role := dto.CreateRolesRequest{}

	if err := c.Bind(&role); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	roleUpdated, err := u.useCase.Update(id, role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    roleUpdated,
		"message": "Data updated successfully",
	})
}

func (u *rolesController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := u.useCase.Delete(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data deleted successfully",
	})
}
