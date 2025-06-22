package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type TermsController interface {
}

type termsController struct {
	useCase usecase.TermsUsecase
}

func NewTermsController(termsUsecase usecase.TermsUsecase) *termsController {
	return &termsController{termsUsecase}
}

func (u *termsController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	term := dto.UpdateTermRequest{}

	if err := c.Bind(&term); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	termUpdated, err := u.useCase.Update(id, term)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    termUpdated,
		"message": "Data updated successfully",
	})
}
