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

func (u *applicantsController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")

	applicants, total, err := u.useCase.Browse(page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  applicants,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *applicantsController) GetByToken(c echo.Context) error {
	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errToken.Error(),
		})
	}

	var id int = int(claims["userId"].(float64))
	applicant, err := u.useCase.GetByToken(id)

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

	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errToken.Error(),
		})
	}
	err := u.useCase.Create(applicant, claims)

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

	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errToken.Error(),
		})
	}
	applicantUpdated, err := u.useCase.Update(id, claims, applicant)

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

func (u *applicantsController) ApproveRegistration(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errToken.Error(),
		})
	}

	err := u.useCase.ApproveRegistration(id, claims)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Registration approved successfully",
	})
}

func (u *applicantsController) ApproveDocument(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errToken.Error(),
		})
	}

	err := u.useCase.ApproveDocument(id, claims)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Registration Document approved successfully",
	})
}

func (u *applicantsController) RejectRegistration(c echo.Context) error {
	applicant := dto.RejectApplicantsRequest{}
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(&applicant); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errToken.Error(),
		})
	}

	err := u.useCase.RejectRegistration(id, claims, applicant)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Registration rejected successfully",
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
