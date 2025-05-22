package controllers

import (
	"net/http"
	"strconv"

	// "strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type UsersController interface {
}

type usersController struct {
	useCase usecase.UsersUsecase
}

func NewUsersController(usersUsecase usecase.UsersUsecase) *usersController {
	return &usersController{usersUsecase}
}

func (u *usersController) Register(c echo.Context) error {
	user := dto.RegisterUsersRequest{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Register(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": user,
	})
}

func (u *usersController) Login(c echo.Context) error {
	user := dto.LoginUsersRequest{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userResponse, err := u.useCase.Login(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": userResponse,
	})
}

func (u *usersController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")

	users, total, err := u.useCase.Browse(page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  users,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}
