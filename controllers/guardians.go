package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type GuardiansController interface {
}

type guardiansController struct {
	useCase usecase.GuardiansUsecase
}

func NewGuardiansController(guardiansUsecase usecase.GuardiansUsecase) *guardiansController {
	return &guardiansController{guardiansUsecase}
}

func (u *guardiansController) GetAll(c echo.Context) error {
	guardians, err := u.useCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": guardians,
	})
}

func (u *guardiansController) Create(c echo.Context) error {
	guardian := dto.CreateGuardiansRequest{}

	if err := c.Bind(&guardian); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(guardian)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *guardiansController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	guardians, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": guardians,
	})
}

func (u *guardiansController) GetGuardianInformationByApplicantId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	guardians, err := u.useCase.FindByApplicantId(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": guardians,
	})
}

func (u *guardiansController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	guardian := dto.CreateGuardiansRequest{}

	if err := c.Bind(&guardian); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	guardianUpdated, err := u.useCase.Update(id, guardian)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    guardianUpdated,
		"message": "Data updated successfully",
	})
}

func (u *guardiansController) Delete(c echo.Context) error {
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
