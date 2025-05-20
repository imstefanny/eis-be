package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type LevelHistoriesController interface {
}

type levelHistoriesController struct {
	useCase usecase.LevelHistoriesUsecase
}

func NewLevelHistoriesController(levelHistoriesUsecase usecase.LevelHistoriesUsecase) *levelHistoriesController {
	return &levelHistoriesController{levelHistoriesUsecase}
}

func (u *levelHistoriesController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")

	levelHistories, total, err := u.useCase.Browse(page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  levelHistories,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *levelHistoriesController) Create(c echo.Context) error {
	levelHistory := dto.CreateLevelHistoriesRequest{}

	if err := c.Bind(&levelHistory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(levelHistory)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *levelHistoriesController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	levelHistories, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": levelHistories,
	})
}

func (u *levelHistoriesController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	levelHistory := dto.CreateLevelHistoriesRequest{}

	if err := c.Bind(&levelHistory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	levelHistorieUpdated, err := u.useCase.Update(id, levelHistory)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    levelHistorieUpdated,
		"message": "Data updated successfully",
	})
}

func (u *levelHistoriesController) Delete(c echo.Context) error {
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
