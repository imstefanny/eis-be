package controllers

import (
	"net/http"
	"strconv"

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

func (u *studentGradesController) GetAll(c echo.Context) error {
	academicID, err := strconv.Atoi(c.Param("academic_id"))
	if err != nil || academicID < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid academic ID",
		})
	}

	studentGrades, err := u.useCase.GetAll(academicID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  studentGrades,
	})
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
