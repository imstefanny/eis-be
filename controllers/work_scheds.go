package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type WorkSchedsController interface {
}

type workSchedsController struct {
	useCase usecase.WorkSchedsUsecase
}

func NewWorkSchedsController(workSchedsUsecase usecase.WorkSchedsUsecase) *workSchedsController {
	return &workSchedsController{workSchedsUsecase}
}

func (u *workSchedsController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")

	workScheds, total, err := u.useCase.Browse(page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  workScheds,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *workSchedsController) Create(c echo.Context) error {
	workSched := dto.CreateWorkSchedsRequest{}

	if err := c.Bind(&workSched); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(workSched)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *workSchedsController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	workScheds, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": workScheds,
	})
}

func (u *workSchedsController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	workSched := dto.CreateWorkSchedsRequest{}

	if err := c.Bind(&workSched); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	workSchedUpdated, err := u.useCase.Update(id, workSched)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    workSchedUpdated,
		"message": "Data updated successfully",
	})
}

func (u *workSchedsController) Delete(c echo.Context) error {
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
