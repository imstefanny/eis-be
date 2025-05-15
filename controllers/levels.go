package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type LevelsController interface {
}

type levelsController struct {
	useCase usecase.LevelsUsecase
}

func NewLevelsController(levelsUsecase usecase.LevelsUsecase) *levelsController {
	return &levelsController{levelsUsecase}
}

func (u *levelsController) GetAll(c echo.Context) error {
	levels, err := u.useCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": levels,
	})
}

func (u *levelsController) Create(c echo.Context) error {
	level := dto.CreateLevelsRequest{}

	if err := c.Bind(&level); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(level)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *levelsController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	levels, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": levels,
	})
}

func (u *levelsController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	level := dto.CreateLevelsRequest{}

	if err := c.Bind(&level); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	levelUpdated, err := u.useCase.Update(id, level)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    levelUpdated,
		"message": "Data updated successfully",
	})
}

func (u *levelsController) Delete(c echo.Context) error {
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
