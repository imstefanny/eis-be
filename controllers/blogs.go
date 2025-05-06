package controllers

import (
	"net/http"

	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type BlogsController interface{
}

type blogsController struct {
	useCase usecase.BlogsUsecase
}

func NewBlogsController(blogsUsecase usecase.BlogsUsecase) *blogsController {
	return &blogsController{blogsUsecase}
}

func (u *blogsController) GetAll(c echo.Context) error {
	blogs, err := u.useCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": blogs,
	})
}

func (u *blogsController) Create(c echo.Context) error {
	studio := dto.CreateStudioRequest{}

	if err := c.Bind(&studio); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	err := u.useCase.Create(studio)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": studio,
	})
}
