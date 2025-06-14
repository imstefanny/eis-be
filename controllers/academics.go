package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type AcademicsController interface {
}

type academicsController struct {
	useCase usecase.AcademicsUsecase
}

func NewAcademicsController(academicsUsecase usecase.AcademicsUsecase) *academicsController {
	return &academicsController{academicsUsecase}
}

func (u *academicsController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")
	academicYear := c.QueryParam("academic_year")

	academics, total, err := u.useCase.Browse(page, limit, search, academicYear)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  academics,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *academicsController) Create(c echo.Context) error {
	academic := dto.CreateAcademicsRequest{}

	if err := c.Bind(&academic); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(academic)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *academicsController) CreateBatch(c echo.Context) error {
	academic := dto.CreateBatchAcademicsRequest{}

	if err := c.Bind(&academic); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.CreateBatch(academic)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *academicsController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	academics, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": academics,
	})
}

func (u *academicsController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	academic := dto.CreateAcademicsRequest{}

	if err := c.Bind(&academic); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	academicUpdated, err := u.useCase.Update(id, academic)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    academicUpdated,
		"message": "Data updated successfully",
	})
}

func (u *academicsController) Delete(c echo.Context) error {
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

// Students specific methods
func (u *academicsController) GetAcademicsByStudent(c echo.Context) error {
	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errToken.Error(),
		})
	}

	userID := int(claims["userId"].(float64))
	academics, err := u.useCase.GetAcademicsByStudent(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": academics,
	})
}
