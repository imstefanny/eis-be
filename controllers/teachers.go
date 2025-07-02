package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type TeachersController interface {
}

type teachersController struct {
	useCase usecase.TeachersUsecase
}

func NewTeachersController(teachersUsecase usecase.TeachersUsecase) *teachersController {
	return &teachersController{teachersUsecase}
}

func (u *teachersController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")

	teachers, total, err := u.useCase.Browse(page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  teachers,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *teachersController) GetAvailableHomeroomTeachers(c echo.Context) error {
	start_year := c.QueryParam("start_year")
	end_year := c.QueryParam("end_year")
	academic_id, err := strconv.Atoi(c.QueryParam("academic_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid academic_id",
		})
	}

	teachers, err := u.useCase.GetAvailableHomeroomTeachers(start_year, end_year, academic_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": teachers,
	})
}

func (u *teachersController) GetByToken(c echo.Context) error {
	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errToken.Error(),
		})
	}

	var id int = int(claims["userId"].(float64))
	teacher, err := u.useCase.GetByToken(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": teacher,
	})
}

func (u *teachersController) Create(c echo.Context) error {
	teacher := dto.CreateTeachersRequest{}

	if err := c.Bind(&teacher); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errToken.Error(),
		})
	}
	err := u.useCase.Create(teacher, claims)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *teachersController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	teachers, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": teachers,
	})
}

func (u *teachersController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	teacher := dto.CreateTeachersRequest{}

	if err := c.Bind(&teacher); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	teacherUpdated, err := u.useCase.Update(id, teacher)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    teacherUpdated,
		"message": "Data updated successfully",
	})
}

func (u *teachersController) Delete(c echo.Context) error {
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
