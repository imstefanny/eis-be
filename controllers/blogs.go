package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type BlogsController interface {
}

type blogsController struct {
	useCase usecase.BlogsUsecase
}

func NewBlogsController(blogsUsecase usecase.BlogsUsecase) *blogsController {
	return &blogsController{blogsUsecase}
}

func (u *blogsController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 6
	}

	search := c.QueryParam("search")

	blogs, total, err := u.useCase.Browse(page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  blogs,
		"page":  page,
		"limit": limit,
		"total": total,
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
		"message": "Data created successfully",
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
		"data":    blogUpdated,
		"message": "Data updated successfully",
	})
}

func (u *blogsController) Delete(c echo.Context) error {
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
