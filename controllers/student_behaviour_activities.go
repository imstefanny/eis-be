package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type StudentBehaviourActivitiesController interface {
}

type studentBehaviourActivitiesController struct {
	useCase usecase.StudentBehaviourActivitiesUsecase
}

func NewStudentBehaviourActivitiesController(studentBehaviourUsecase usecase.StudentBehaviourActivitiesUsecase) *studentBehaviourActivitiesController {
	return &studentBehaviourActivitiesController{studentBehaviourUsecase}
}

func (u *studentBehaviourActivitiesController) GetByAcademicIdAndTermId(c echo.Context) error {
	termID, err := strconv.Atoi(c.Param("term_id"))
	if err != nil || termID < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid term ID",
		})
	}
	academicId, err := strconv.Atoi(c.Param("academic_id"))
	if err != nil || academicId < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid academic ID",
		})
	}

	studentBehaviour, err := u.useCase.GetByAcademicIdAndTermId(academicId, termID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": studentBehaviour,
	})
}

func (u *studentBehaviourActivitiesController) Create(c echo.Context) error {
	studentBehaviour := []dto.StudentBehaviourActivityRequest{}

	if err := c.Bind(&studentBehaviour); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(studentBehaviour)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Student behaviour activities created successfully",
	})
}

func (u *studentBehaviourActivitiesController) Update(c echo.Context) error {
	termID, err := strconv.Atoi(c.Param("term_id"))
	if err != nil || termID < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid term ID",
		})
	}

	studentBehaviour := []dto.StudentBehaviourActivityRequest{}

	if err := c.Bind(&studentBehaviour); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	errUpdate := u.useCase.Update(studentBehaviour)

	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data updated successfully",
	})
}
