package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type StudentsController interface {
}

type studentsController struct {
	useCase usecase.StudentsUsecase
}

func NewStudentsController(studentsUsecase usecase.StudentsUsecase) *studentsController {
	return &studentsController{studentsUsecase}
}

func (u *studentsController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")

	students, total, err := u.useCase.Browse(page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  students,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *studentsController) Create(c echo.Context) error {
	student := dto.CreateStudentsRequest{}

	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	studentId, err := u.useCase.Create(student, c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Data created successfully",
		"created_id": studentId,
	})
}

func (u *studentsController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	students, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": students,
	})
}

func (u *studentsController) GetByToken(c echo.Context) error {
	claims, errToken := helpers.GetTokenClaims(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errToken.Error(),
		})
	}

	var id int = int(claims["userId"].(float64))
	teacher, err := u.useCase.GetByToken(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": teacher,
	})
}

func (u *studentsController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	student := dto.CreateStudentsRequest{}

	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	studentUpdated, err := u.useCase.Update(id, student)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    studentUpdated,
		"message": "Data updated successfully",
	})
}

func (u *studentsController) UpdateStudentAcademicId(c echo.Context) error {
	academic_id, _ := strconv.Atoi(c.Param("academic_id"))
	student := dto.UpdateStudentAcademicIdRequest{}

	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.UpdateStudentAcademicId(academic_id, student.StudentIDs)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data updated successfully",
	})
}

func (u *studentsController) Delete(c echo.Context) error {
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
