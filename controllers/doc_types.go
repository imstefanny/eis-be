package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type DocTypesController interface {
}

type docTypesController struct {
	useCase usecase.DocTypesUsecase
}

func NewDocTypesController(docTypesUsecase usecase.DocTypesUsecase) *docTypesController {
	return &docTypesController{docTypesUsecase}
}

func (u *docTypesController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")

	docTypes, total, err := u.useCase.Browse(page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  docTypes,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *docTypesController) Create(c echo.Context) error {
	docType := dto.CreateDocTypesRequest{}

	if err := c.Bind(&docType); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(docType)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *docTypesController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	docTypes, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": docTypes,
	})
}

func (u *docTypesController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	docType := dto.CreateDocTypesRequest{}

	if err := c.Bind(&docType); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	docTypeUpdated, err := u.useCase.Update(id, docType)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    docTypeUpdated,
		"message": "Data updated successfully",
	})
}

func (u *docTypesController) Delete(c echo.Context) error {
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
