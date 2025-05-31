package controllers

import (
	"net/http"

	"eis-be/dto"
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

// func (u *studentAttsController) Browse(c echo.Context) error {
// 	page, err := strconv.Atoi(c.QueryParam("page"))
// 	if err != nil || page < 1 {
// 		page = 1
// 	}
// 	limit, err := strconv.Atoi(c.QueryParam("limit"))
// 	if err != nil || limit < 1 {
// 		limit = 10
// 	}
// 	search := c.QueryParam("search")
// 	sortColumn := c.QueryParam("sortColumn")
// 	if sortColumn == "" {
// 		sortColumn = "created_at"
// 	}
// 	sortOrder := c.QueryParam("sortOrder")
// 	if sortOrder != "asc" && sortOrder != "desc" {
// 		sortOrder = "desc"
// 	}
// 	date := c.QueryParam("date")

// 	blogs, total, err := u.useCase.Browse(page, limit, search, date)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"data":  blogs,
// 		"page":  page,
// 		"limit": limit,
// 		"total": total,
// 	})
// }

// func (u *studentAttsController) Create(c echo.Context) error {
// 	studentAtt := dto.CreateStudentAttsRequest{}

// 	if err := c.Bind(&studentAtt); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 	}

// 	err := u.useCase.Create(studentAtt)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Data created successfully",
// 	})
// }

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

// func (u *studentAttsController) Find(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	studentAtts, err := u.useCase.Find(id)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"data": studentAtts,
// 	})
// }

// func (u *studentAttsController) Update(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	studentAtt := dto.CreateStudentAttsRequest{}

// 	if err := c.Bind(&studentAtt); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 	}

// 	studentAttUpdated, err := u.useCase.Update(id, studentAtt)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"data":    studentAttUpdated,
// 		"message": "Data updated successfully",
// 	})
// }

// func (u *studentAttsController) Delete(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	err := u.useCase.Delete(id)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Data deleted successfully",
// 	})
// }
