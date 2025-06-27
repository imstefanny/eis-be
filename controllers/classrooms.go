package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type ClassroomsController interface {
}

type classroomsController struct {
	useCase usecase.ClassroomsUsecase
}

func NewClassroomsController(classroomsUsecase usecase.ClassroomsUsecase) *classroomsController {
	return &classroomsController{classroomsUsecase}
}

func (u *classroomsController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")

	classrooms, total, err := u.useCase.Browse(page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  classrooms,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *classroomsController) Create(c echo.Context) error {
	classroom := dto.CreateClassroomsRequest{}

	if err := c.Bind(&classroom); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(classroom)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *classroomsController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	classrooms, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": classrooms,
	})
}

func (u *classroomsController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	classroom := dto.CreateClassroomsRequest{}

	if err := c.Bind(&classroom); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	classroomUpdated, err := u.useCase.Update(id, classroom)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    classroomUpdated,
		"message": "Data updated successfully",
	})
}

func (u *classroomsController) Delete(c echo.Context) error {
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

func (u *classroomsController) UnDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := u.useCase.UnDelete(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data undeleted successfully",
	})
}
