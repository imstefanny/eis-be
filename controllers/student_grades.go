package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type StudentGradesController interface {
}

type studentGradesController struct {
	useCase usecase.StudentGradesUsecase
}

func NewStudentGradesController(studentGradesUsecase usecase.StudentGradesUsecase) *studentGradesController {
	return &studentGradesController{studentGradesUsecase}
}

func (u *studentGradesController) GetAll(c echo.Context) error {
	termID, err := strconv.Atoi(c.Param("term_id"))
	if err != nil || termID < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid term ID",
		})
	}

	studentGrades, err := u.useCase.GetAll(termID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": studentGrades,
	})
}

func (u *studentGradesController) Create(c echo.Context) error {
	studentGrade := dto.CreateStudentGradesRequest{}

	if err := c.Bind(&studentGrade); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(studentGrade)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *studentGradesController) UpdateByTermID(c echo.Context) error {
	termID, err := strconv.Atoi(c.Param("term_id"))
	if err != nil || termID < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid term ID",
		})
	}

	studentGrade := dto.UpdateStudentGradesRequest{}

	if err := c.Bind(&studentGrade); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if studentGrade.TermID != uint(termID) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Term ID mismatch",
		})
	}

	studentGradeUpdated, err := u.useCase.UpdateByTermID(termID, studentGrade)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    studentGradeUpdated,
		"message": "Data updated successfully",
	})
}

func (u *studentGradesController) GetAllByStudent(c echo.Context) error {
	studentIDsString := strings.Split(c.Param("student_ids"), ",")
	studentIDs := make([]int, 0, len(studentIDsString))
	for _, idStr := range studentIDsString {
		id, err := strconv.Atoi(idStr)
		if err != nil || id < 1 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": fmt.Sprintf("Invalid student ID: %s", idStr),
			})
		}
		studentIDs = append(studentIDs, id)
	}
	termID, err := strconv.Atoi(c.Param("term_id"))
	if err != nil || termID < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid term ID",
		})
	}
	studentGrades, err := u.useCase.GetAllByStudent(termID, studentIDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": studentGrades,
	})
}

func (u *studentGradesController) GetReport(c echo.Context) error {
	academicYear := c.QueryParam("academic_year")
	if academicYear == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Academic year is required",
		})
	}
	levelID, err := strconv.Atoi(c.QueryParam("level_id"))
	if err != nil || levelID < 1 {
		levelID = 0
	}
	academicID, err := strconv.Atoi(c.QueryParam("academic_id"))
	if err != nil || academicID < 1 {
		academicID = 0
	}
	termID, err := strconv.Atoi(c.QueryParam("term_id"))
	if err != nil || termID < 1 {
		termID = 0
	}
	studentGrades, err := u.useCase.GetReport(academicYear, levelID, academicID, termID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": studentGrades,
	})
}

// Students specific methods
func (u *studentGradesController) GetStudentScoreByStudent(c echo.Context) error {
	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errToken.Error(),
		})
	}

	userID := int(claims["userId"].(float64))
	termID, err := strconv.Atoi(c.Param("term_id"))
	if err != nil || termID < 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid term ID",
		})
	}
	scores, err := u.useCase.GetStudentScoreByStudent(userID, termID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": scores,
	})
}
