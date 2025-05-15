package controllers

import (
	"net/http"
	"strconv"

	"eis-be/usecase"
	"eis-be/dto"

	"github.com/labstack/echo/v4"
)

type SubjectsController interface{
}

type subjectsController struct {
	useCase usecase.SubjectsUsecase
}

func NewSubjectsController(subjectsUsecase usecase.SubjectsUsecase) *subjectsController {
	return &subjectsController{subjectsUsecase}
}

func (u *subjectsController) GetAll(c echo.Context) error {
	subjects, err := u.useCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": subjects,
	})
}

func (u *subjectsController) Create(c echo.Context) error {
	subject := dto.CreateSubjectsRequest{}

	if err := c.Bind(&subject); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(subject)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *subjectsController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	subjects, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": subjects,
	})
}

func (u *subjectsController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	subject := dto.CreateSubjectsRequest{}

	if err := c.Bind(&subject); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	subjectUpdated, err := u.useCase.Update(id, subject)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": subjectUpdated,
		"message": "Data updated successfully",
	})
}

func (u *subjectsController) Delete(c echo.Context) error {
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
