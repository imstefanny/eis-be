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
	usersRepository := repository.NewUsersRepository(db)
	usersService := usecase.NewUsersUsecase(usersRepository)
	usersController := controllers.NewUsersController(usersService)

	blogsRepository := repository.NewBlogsRepository(db)
	blogsService := usecase.NewBlogsUsecase(blogsRepository)
	blogsController := controllers.NewBlogsController(blogsService)

	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/register", usersController.Register)

	eBlogs := e.Group("/blogs")
	eBlogs.GET("", blogsController.GetAll)
	eBlogs.GET("/:id", blogsController.Find)
	eBlogs.POST("", blogsController.Create)
	eBlogs.PUT("/:id", blogsController.Update)
	eBlogs.DELETE("/:id", blogsController.Delete)
}
