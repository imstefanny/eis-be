package controllers

import (
	"net/http"

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
