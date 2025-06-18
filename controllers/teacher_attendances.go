package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type TeacherAttsController interface {
}

type teacherAttsController struct {
	useCase usecase.TeacherAttsUsecase
}

func NewTeacherAttsController(teacherAttsUsecase usecase.TeacherAttsUsecase) *teacherAttsController {
	return &teacherAttsController{teacherAttsUsecase}
}

func (u *teacherAttsController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}
	search := c.QueryParam("search")
	date := c.QueryParam("date")

	var userId *int
	idStr := c.QueryParam("userId")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err == nil {
			userId = &id
		}
	}

	teacherAtts, total, err := u.useCase.Browse(page, limit, search, date, userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  teacherAtts,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *teacherAttsController) Create(c echo.Context) error {
	teacherAtt := dto.CreateTeacherAttsRequest{}

	if err := c.Bind(&teacherAtt); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(teacherAtt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *teacherAttsController) CreateBatch(c echo.Context) error {
	teacherAtt := dto.CreateBatchTeacherAttsRequest{}

	if err := c.Bind(&teacherAtt); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.CreateBatch(teacherAtt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *teacherAttsController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	teacherAtts, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": teacherAtts,
	})
}

func (u *teacherAttsController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	teacherAtt := dto.CreateTeacherAttsRequest{}

	if err := c.Bind(&teacherAtt); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	teacherAttUpdated, err := u.useCase.Update(id, teacherAtt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    teacherAttUpdated,
		"message": "Data updated successfully",
	})
}

func (u *teacherAttsController) Delete(c echo.Context) error {
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

func (u *teacherAttsController) GetReport(c echo.Context) error {
	search := c.QueryParam("search")
	start_date := c.QueryParam("start_date")
	end_date := c.QueryParam("end_date")
	var userId *int
	idStr := c.QueryParam("userId")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err == nil {
			userId = &id
		}
	}

	teacherAtts, err := u.useCase.GetReport(search, start_date, end_date, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": teacherAtts,
	})
}
