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

func (u *documentsController) GetAll(c echo.Context) error {
	documents, err := u.useCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": documents,
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
		"data": document,
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
