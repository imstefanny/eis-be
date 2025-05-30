package controllers

import (
	"net/http"
	"strconv"

	"eis-be/dto"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
)

type ClassNotesController interface {
}

type classNotesController struct {
	useCase usecase.ClassNotesUsecase
}

func NewClassNotesController(classNotesUsecase usecase.ClassNotesUsecase) *classNotesController {
	return &classNotesController{classNotesUsecase}
}

func (u *classNotesController) Browse(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.QueryParam("search")

	classNotes, total, err := u.useCase.Browse(page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  classNotes,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *classNotesController) BrowseByAcademicID(c echo.Context) error {
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

	classNotes, total, err := u.useCase.BrowseByAcademicID(academicID, page, limit, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  classNotes,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func (u *classNotesController) Create(c echo.Context) error {
	classNote := dto.CreateClassNotesRequest{}

	if err := c.Bind(&classNote); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.Create(classNote)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *classNotesController) CreateBatch(c echo.Context) error {
	classNote := dto.CreateBatchClassNotesRequest{}

	if err := c.Bind(&classNote); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.useCase.CreateBatch(classNote)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data created successfully",
	})
}

func (u *classNotesController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	classNotes, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": classNotes,
	})
}

func (u *classNotesController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	classNote := dto.CreateClassNotesRequest{}

	if err := c.Bind(&classNote); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	classNoteUpdated, err := u.useCase.Update(id, classNote)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    classNoteUpdated,
		"message": "Data updated successfully",
	})
}

// func (u *classNotesController) Delete(c echo.Context) error {
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
