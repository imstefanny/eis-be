package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type SubjectsController interface {
}

type subjectsController struct {
	useCase usecase.SubjectsUsecase
}

func NewSubjectsController(subjectsUsecase usecase.SubjectsUsecase) *subjectsController {
	return &subjectsController{subjectsUsecase}
}

func (u *subjectsController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}
	search := c.QueryParam("search")
	sortColumn := c.QueryParam("sortColumn")
	if sortColumn == "" {
		sortColumn = "created_at"
	}
	sortOrder := c.QueryParam("sortOrder")
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	blogs, total, err := u.useCase.Browse(page, limit, search, sortColumn, sortOrder)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  blogs,
		"page":  page,
		"limit": limit,
		"total": total,
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
		"data":    subjectUpdated,
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
