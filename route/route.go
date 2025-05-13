package route

import (
	"eis-be/constants"
	"eis-be/controllers"
	"eis-be/repository"
	"eis-be/usecase"

	echojwt "github.com/labstack/echo-jwt/v4"
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

	applicantsRepository := repository.NewApplicantsRepository(db)
	applicantsService := usecase.NewApplicantsUsecase(applicantsRepository)
	applicantsController := controllers.NewApplicantsController(applicantsService)

	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/register", usersController.Register)
	e.POST("/login", usersController.Login)

	eBlogs := e.Group("/blogs")
	eBlogs.GET("", blogsController.Browse)
	eBlogs.GET("/:id", blogsController.Find)
	eBlogs.POST("", blogsController.Create)
	eBlogs.PUT("/:id", blogsController.Update)
	eBlogs.DELETE("/:id", blogsController.Delete)

	eApplicants := e.Group("/applicants")
	eApplicants.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eApplicants.GET("", applicantsController.GetAll)
	eApplicants.GET("/:id", applicantsController.Find)
	eApplicants.POST("", applicantsController.Create)
	eApplicants.PUT("/:id", applicantsController.Update)
	eApplicants.DELETE("/:id", applicantsController.Delete)
}
