package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type DocumentsController interface {
}

type documentsController struct {
	useCase usecase.DocumentsUsecase
}

func NewDocumentsController(documentsUsecase usecase.DocumentsUsecase) *documentsController {
	return &documentsController{documentsUsecase}
}

func (u *documentsController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")

	documents, total, err := u.useCase.Browse(page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  documents,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *documentsController) Create(c echo.Context) error {
	document := dto.CreateDocumentsRequest{}

	if err := c.Bind(&document); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(document)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *documentsController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	documents, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": documents,
	})
}

func (u *documentsController) GetDocumentsByApplicantId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	documents, err := u.useCase.FindByApplicantId(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": documents,
	})
}

func (u *documentsController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	document := dto.CreateDocumentsRequest{}

	if err := c.Bind(&document); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	documentUpdated, err := u.useCase.Update(id, document)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    documentUpdated,
		"message": "Data updated successfully",
	})
}

func (u *documentsController) Delete(c echo.Context) error {
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
