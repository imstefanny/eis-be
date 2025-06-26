package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type CurriculumsController interface {
}

type curriculumsController struct {
	useCase usecase.CurriculumsUsecase
}

func NewCurriculumsController(curriculumsUsecase usecase.CurriculumsUsecase) *curriculumsController {
	return &curriculumsController{curriculumsUsecase}
}

func (u *curriculumsController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")

	curriculums, total, err := u.useCase.Browse(page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  curriculums,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *curriculumsController) Create(c echo.Context) error {
	curriculum := dto.CreateCurriculumsRequest{}

	if err := c.Bind(&curriculum); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(curriculum)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *curriculumsController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	curriculums, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": curriculums,
	})
}

func (u *curriculumsController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	curriculum := dto.CreateCurriculumsRequest{}

	if err := c.Bind(&curriculum); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	curriculumUpdated, err := u.useCase.Update(id, curriculum)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    curriculumUpdated,
		"message": "Data updated successfully",
	})
}

func (u *curriculumsController) Delete(c echo.Context) error {
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

func (u *curriculumsController) UnDelete(c echo.Context) error {
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
