package route

import (
	"eis-be/controllers"
	"eis-be/repository"
	"eis-be/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func Route(e *echo.Echo, db *gorm.DB) {
	blogsRepository := repository.NewBlogsRepository(db)
	blogsService := usecase.NewBlogsUsecase(blogsRepository)
	blogsController := controllers.NewBlogsController(blogsService)

	e.Pre(middleware.RemoveTrailingSlash())

	eBlogs := e.Group("/blogs")
	eBlogs.GET("", blogsController.GetAll)
}
