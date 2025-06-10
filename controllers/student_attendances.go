package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type StudentAttsController interface {
}

type studentAttsController struct {
	useCase usecase.StudentAttsUsecase
}

func NewStudentAttsController(studentAttsUsecase usecase.StudentAttsUsecase) *studentAttsController {
	return &studentAttsController{studentAttsUsecase}
}

func (u *studentAttsController) BrowseByAcademicID(c echo.Context) error {
	academicID, err := strconv.Atoi(c.Param("academic_id"))
	if err != nil || academicID < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid academic ID",
		})
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}
	search := c.QueryParam("search")
	date := c.QueryParam("date")

	studentAtts, total, err := u.useCase.BrowseByAcademicID(academicID, page, limit, search, date)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  studentAtts,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *studentAttsController) CreateBatch(c echo.Context) error {
	studentAtt := dto.CreateBatchStudentAttsRequest{}

	if err := c.Bind(&studentAtt); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.CreateBatch(studentAtt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *studentAttsController) UpdateByAcademicID(c echo.Context) error {
	academicID, _ := strconv.Atoi(c.Param("academic_id"))
	studentAtt := dto.UpdateStudentAttsRequest{}

	if err := c.Bind(&studentAtt); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	studentAttUpdated, err := u.useCase.UpdateByAcademicID(academicID, studentAtt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    studentAttUpdated,
		"message": "Data updated successfully",
	})
}

func (u *studentAttsController) GetReport(c echo.Context) error {
	academicID, err := strconv.Atoi(c.QueryParam("academic_id"))
	if err != nil || academicID < 1 {
		academicID = 0
	}
	levelID, err := strconv.Atoi(c.QueryParam("level_id"))
	if err != nil || levelID < 1 {
		levelID = 0
	}
	classID, err := strconv.Atoi(c.QueryParam("class_id"))
	if err != nil || classID < 1 {
		classID = 0
	}
	search := c.QueryParam("search")
	start_date := c.QueryParam("start_date")
	end_date := c.QueryParam("end_date")
	studentAtts, err := u.useCase.GetReport(academicID, levelID, classID, search, start_date, end_date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  studentAtts,
	})
}

// Students specific methods
func (u *studentAttsController) GetAttendanceByStudent(c echo.Context) error {
	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errToken.Error(),
		})
	}

	id := int(claims["userId"].(float64))
	month, err := strconv.Atoi(c.QueryParam("month"))
	if month < 1 || month > 12 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Invalid month parameter, must be between 1 and 12",
		})
	}
	studentAtts, err := u.useCase.GetAttendanceByStudent(id, month)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": studentAtts,
	})
}
