package controllers

import (
	"net/http"
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
