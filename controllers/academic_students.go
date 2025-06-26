package controllers

import (
	"net/http"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type AcademicStudentsController interface {
}

type academicStudentsController struct {
	useCase usecase.AcademicStudentsUsecase
}

func NewAcademicStudentsController(academicStudentsUsecase usecase.AcademicStudentsUsecase) *academicStudentsController {
	return &academicStudentsController{academicStudentsUsecase}
}

func (u *academicStudentsController) Update(c echo.Context) error {
	academicStudents := []dto.UpdateAcademicStudentsRequest{}

	if err := c.Bind(&academicStudents); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Update(academicStudents)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data updated successfully",
	})
}
