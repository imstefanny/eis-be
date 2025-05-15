package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type ApplicantsController interface {
}

type applicantsController struct {
	useCase usecase.ApplicantsUsecase
}

func NewApplicantsController(applicantsUsecase usecase.ApplicantsUsecase) *applicantsController {
	return &applicantsController{applicantsUsecase}
}

func (u *applicantsController) GetAll(c echo.Context) error {
	applicants, err := u.useCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": applicants,
	})
}

func (u *applicantsController) GetApplicantInformationByToken(c echo.Context) error {
	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errToken.Error(),
		})
	}

	var id int = int(claims["userId"].(float64))
	applicant, err := u.useCase.FindByCreatedBy(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": applicant,
	})
}

func (u *applicantsController) Create(c echo.Context) error {
	applicant := dto.CreateApplicantsRequest{}

	if err := c.Bind(&applicant); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(applicant, c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *applicantsController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	applicants, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": applicants,
	})
}

func (u *applicantsController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	applicant := dto.CreateApplicantsRequest{}

	if err := c.Bind(&applicant); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	applicantUpdated, err := u.useCase.Update(id, applicant)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    applicantUpdated,
		"message": "Data updated successfully",
	})
}

func (u *applicantsController) Delete(c echo.Context) error {
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
