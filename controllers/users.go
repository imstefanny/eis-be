package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type UsersController interface {
}

type usersController struct {
	useCase usecase.UsersUsecase
}
