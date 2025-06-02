package controllers

import (
	"net/http"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type StudentGradesController interface {
}

type studentGradesController struct {
	useCase usecase.StudentGradesUsecase
}

func NewStudentGradesController(studentGradesUsecase usecase.StudentGradesUsecase) *studentGradesController {
	return &studentGradesController{studentGradesUsecase}
}

func (u *studentGradesController) Create(c echo.Context) error {
	studentGrade := dto.CreateStudentGradesRequest{}

	if err := c.Bind(&studentGrade); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(studentGrade)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}
