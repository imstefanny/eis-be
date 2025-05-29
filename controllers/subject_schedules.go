package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type SubjSchedsController interface {
}

type subjSchedsController struct {
	useCase usecase.SubjSchedsUsecase
}

func NewSubjSchedsController(subjSchedsUsecase usecase.SubjSchedsUsecase) *subjSchedsController {
	return &subjSchedsController{subjSchedsUsecase}
}

func (u *subjSchedsController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")

	subjScheds, total, err := u.useCase.Browse(page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  subjScheds,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *subjSchedsController) Create(c echo.Context) error {
	subjSched := dto.CreateSubjSchedsRequest{}

	if err := c.Bind(&subjSched); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(subjSched)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *subjSchedsController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	subjScheds, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": subjScheds,
	})
}

func (u *subjSchedsController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	subjSched := dto.UpdateSubjSchedsRequest{}

	if err := c.Bind(&subjSched); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	subjSchedUpdated, err := u.useCase.Update(id, subjSched)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    subjSchedUpdated,
		"message": "Data updated successfully",
	})
}

func (u *subjSchedsController) Delete(c echo.Context) error {
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
