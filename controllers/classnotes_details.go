package controllers

import (
	"net/http"

	"eis-be/helpers"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type ClassNotesDetailsController interface {
}

type classNotesDetailsController struct {
	useCase usecase.ClassNotesDetailsUsecase
}

func NewClassNotesDetailsController(classNotesDetailsUsecase usecase.ClassNotesDetailsUsecase) *classNotesDetailsController {
	return &classNotesDetailsController{classNotesDetailsUsecase}
}

func (u *classNotesDetailsController) GetAllByTeacher(c echo.Context) error {
	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errToken.Error(),
		})
	}

	id := int(claims["userId"].(float64))
	date := c.QueryParam("date")

	subjScheds, err := u.useCase.GetAllByTeacher(id, date)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": subjScheds,
	})
}
