package controllers

import (
	"net/http"
	"strconv"

	"eis-be/usecase"
	"eis-be/dto"

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
	blog := dto.CreateBlogsRequest{}

	if err := c.Bind(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(blog)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": blog,
	})
}

func (u *blogsController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	blogs, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": blogs,
	})
}

func (u *blogsController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	blog := dto.CreateBlogsRequest{}

	if err := c.Bind(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	blogUpdated, err := u.useCase.Update(id, blog)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": blogUpdated,
		"message": "Data updated successfully",
	})
}
