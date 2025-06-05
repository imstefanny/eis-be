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
	usersService := usecase.NewUsersUsecase(usersRepository, db)
	usersController := controllers.NewUsersController(usersService)

	rolesRepository := repository.NewRolesRepository(db)
	permissionsRepository := repository.NewPermissionsRepository(db)
	rolesService := usecase.NewRolesUsecase(rolesRepository, permissionsRepository)
	rolesController := controllers.NewRolesController(rolesService)

	blogsRepository := repository.NewBlogsRepository(db)
	blogsService := usecase.NewBlogsUsecase(blogsRepository, usersRepository)
	blogsController := controllers.NewBlogsController(blogsService)

	studentsRepository := repository.NewStudentsRepository(db)
	studentsService := usecase.NewStudentsUsecase(studentsRepository, usersRepository, db)
	studentsController := controllers.NewStudentsController(studentsService)

	teachersRepository := repository.NewTeachersRepository(db)
	teachersService := usecase.NewTeachersUsecase(teachersRepository, usersRepository, db)
	teachersController := controllers.NewTeachersController(teachersService)

	teacherAttsRepository := repository.NewTeacherAttsRepository(db)
	teacherAttsService := usecase.NewTeacherAttsUsecase(teacherAttsRepository, teachersRepository)
	teacherAttsController := controllers.NewTeacherAttsController(teacherAttsService)

	guardiansRepository := repository.NewGuardiansRepository(db)
	guardiansService := usecase.NewGuardiansUsecase(guardiansRepository)
	guardiansController := controllers.NewGuardiansController(guardiansService)

	applicantsRepository := repository.NewApplicantsRepository(db)
	applicantsService := usecase.NewApplicantsUsecase(applicantsRepository, studentsRepository, guardiansRepository)
	applicantsController := controllers.NewApplicantsController(applicantsService)

	docTypesRepository := repository.NewDocTypesRepository(db)
	docTypesService := usecase.NewDocTypesUsecase(docTypesRepository)
	docTypesController := controllers.NewDocTypesController(docTypesService)

	documentsRepository := repository.NewDocumentsRepository(db)
	documentsService := usecase.NewDocumentsUsecase(documentsRepository)
	documentsController := controllers.NewDocumentsController(documentsService)

	workSchedsRepository := repository.NewWorkSchedsRepository(db)
	workSchedDetailsRepository := repository.NewWorkSchedDetailsRepository(db)
	workSchedsService := usecase.NewWorkSchedsUsecase(workSchedsRepository, workSchedDetailsRepository)
	workSchedsController := controllers.NewWorkSchedsController(workSchedsService)

	subjectsRepository := repository.NewSubjectsRepository(db)
	subjectsService := usecase.NewSubjectsUsecase(subjectsRepository)
	subjectsController := controllers.NewSubjectsController(subjectsService)

	levelsRepository := repository.NewLevelsRepository(db)
	levelsService := usecase.NewLevelsUsecase(levelsRepository)
	levelsController := controllers.NewLevelsController(levelsService)

	levelHistoriesRepository := repository.NewLevelHistoriesRepository(db)
	levelHistoriesService := usecase.NewLevelHistoriesUsecase(levelHistoriesRepository)
	levelHistoriesController := controllers.NewLevelHistoriesController(levelHistoriesService)

	classroomsRepository := repository.NewClassroomsRepository(db)
	classroomsService := usecase.NewClassroomsUsecase(classroomsRepository, levelsRepository)
	classroomsController := controllers.NewClassroomsController(classroomsService)

	academicsRepository := repository.NewAcademicsRepository(db)
	academicsService := usecase.NewAcademicsUsecase(academicsRepository, studentsRepository, classroomsRepository)
	academicsController := controllers.NewAcademicsController(academicsService)

	subjSchedsRepository := repository.NewSubjSchedsRepository(db)
	subjSchedsService := usecase.NewSubjSchedsUsecase(subjSchedsRepository, academicsRepository, teachersRepository)
	subjSchedsController := controllers.NewSubjSchedsController(subjSchedsService)

	studentAttsRepository := repository.NewStudentAttsRepository(db)
	studentAttsService := usecase.NewStudentAttsUsecase(studentAttsRepository, studentsRepository, academicsRepository)
	studentAttsController := controllers.NewStudentAttsController(studentAttsService)

	classNotesRepository := repository.NewClassNotesRepository(db)
	classNotesService := usecase.NewClassNotesUsecase(classNotesRepository, academicsRepository, studentAttsRepository, teachersRepository)
	classNotesController := controllers.NewClassNotesController(classNotesService)

	studentGradesRepository := repository.NewStudentGradesRepository(db)
	studentGradesService := usecase.NewStudentGradesUsecase(studentGradesRepository, academicsRepository, studentsRepository, subjectsRepository)
	studentGradesController := controllers.NewStudentGradesController(studentGradesService)

	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/register", usersController.Register)
	e.POST("/login", usersController.Login)

	eUser := e.Group("/users")
	eUser.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eUser.GET("", usersController.Browse)
	eUser.PUT("/:id", usersController.Update)

	eRoles := e.Group("/roles")
	eRoles.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eRoles.GET("", rolesController.Browse)
	eRoles.GET("/permissions", rolesController.GetAllPermissions)
	eRoles.GET("/:id", rolesController.Find)
	eRoles.POST("", rolesController.Create)
	eRoles.PUT("/:id", rolesController.Update)
	eRoles.DELETE("/:id", rolesController.Delete)

	eBlogs := e.Group("/blogs")
	eBlogs.GET("", blogsController.Browse)
	eBlogs.GET("/:id", blogsController.Find)
	eBlogs.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eBlogs.POST("", blogsController.Create)
	eBlogs.PUT("/:id", blogsController.Update)
	eBlogs.DELETE("/:id", blogsController.Delete)

	eApplicants := e.Group("/applicants")
	eApplicants.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eApplicants.GET("", applicantsController.Browse)
	eApplicants.GET("/my", applicantsController.GetByToken)
	eApplicants.GET("/:id", applicantsController.Find)
	eApplicants.POST("", applicantsController.Create)
	eApplicants.PUT("/:id", applicantsController.Update)
	eApplicants.DELETE("/:id", applicantsController.Delete)
	eApplicants.POST("/approve/:id", applicantsController.ApproveRegistration)

	eGuardians := e.Group("/guardians")
	eGuardians.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eGuardians.GET("", guardiansController.Browse)
	eGuardians.GET("/my-information/:id", guardiansController.GetGuardianInformationByApplicantId)
	eGuardians.GET("/:id", guardiansController.Find)
	eGuardians.POST("", guardiansController.Create)
	eGuardians.PUT("/:id", guardiansController.Update)
	eGuardians.DELETE("/:id", guardiansController.Delete)

	eDocTypes := e.Group("/doctypes")
	eDocTypes.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eDocTypes.GET("", docTypesController.Browse)
	eDocTypes.GET("/:id", docTypesController.Find)
	eDocTypes.POST("", docTypesController.Create)
	eDocTypes.PUT("/:id", docTypesController.Update)
	eDocTypes.DELETE("/:id", docTypesController.Delete)

	eDocs := e.Group("/documents")
	eDocs.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eDocs.GET("", documentsController.Browse)
	eDocs.GET("/my-information/:id", documentsController.GetDocumentsByApplicantId)
	eDocs.GET("/:id", documentsController.Find)
	eDocs.POST("", documentsController.Create)
	eDocs.PUT("/:id", documentsController.Update)
	eDocs.DELETE("/:id", documentsController.Delete)

	eWorkScheds := e.Group("/workscheds")
	eWorkScheds.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eWorkScheds.GET("", workSchedsController.Browse)
	eWorkScheds.GET("/:id", workSchedsController.Find)
	eWorkScheds.POST("", workSchedsController.Create)
	eWorkScheds.PUT("/:id", workSchedsController.Update)
	eWorkScheds.PUT("/undelete/:id", workSchedsController.Undelete)
	eWorkScheds.DELETE("/:id", workSchedsController.Delete)

	eSubjects := e.Group("/subjects")
	eSubjects.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eSubjects.GET("", subjectsController.Browse)
	eSubjects.GET("/:id", subjectsController.Find)
	eSubjects.POST("", subjectsController.Create)
	eSubjects.PUT("/:id", subjectsController.Update)
	eSubjects.DELETE("/:id", subjectsController.Delete)

	eLevels := e.Group("/levels")
	eLevels.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eLevels.GET("", levelsController.Browse)
	eLevels.GET("/:id", levelsController.Find)
	eLevels.POST("", levelsController.Create)
	eLevels.PUT("/:id", levelsController.Update)
	eLevels.DELETE("/:id", levelsController.Delete)

	eLevelHistories := e.Group("/levelhistories")
	eLevelHistories.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eLevelHistories.GET("", levelHistoriesController.Browse)
	eLevelHistories.GET("/:id", levelHistoriesController.Find)
	eLevelHistories.POST("", levelHistoriesController.Create)
	eLevelHistories.PUT("/:id", levelHistoriesController.Update)
	eLevelHistories.DELETE("/:id", levelHistoriesController.Delete)

	eClassrooms := e.Group("/classrooms")
	eClassrooms.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eClassrooms.GET("", classroomsController.Browse)
	eClassrooms.GET("/:id", classroomsController.Find)
	eClassrooms.POST("", classroomsController.Create)
	eClassrooms.PUT("/:id", classroomsController.Update)
	eClassrooms.DELETE("/:id", classroomsController.Delete)

	eStudents := e.Group("/students")
	eStudents.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eStudents.GET("", studentsController.Browse)
	// eStudents.GET("/my", studentsController.GetByToken)
	eStudents.GET("/:id", studentsController.Find)
	eStudents.POST("", studentsController.Create)
	eStudents.PUT("/:id", studentsController.Update)
	eStudents.PUT("/update-current-academic/:academic_id", studentsController.UpdateStudentAcademicId)
	eStudents.DELETE("/:id", studentsController.Delete)

	eTeachers := e.Group("/teachers")
	eTeachers.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eTeachers.GET("", teachersController.Browse)
	eTeachers.GET("/my", teachersController.GetByToken)
	eTeachers.GET("/:id", teachersController.Find)
	eTeachers.POST("", teachersController.Create)
	eTeachers.PUT("/:id", teachersController.Update)
	eTeachers.DELETE("/:id", teachersController.Delete)

	eTeacherAtts := eTeachers.Group("/attendances")
	eTeacherAtts.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eTeacherAtts.GET("", teacherAttsController.Browse)
	eTeacherAtts.GET("/:id", teacherAttsController.Find)
	eTeacherAtts.POST("", teacherAttsController.Create)
	eTeacherAtts.POST("/batch", teacherAttsController.CreateBatch)
	eTeacherAtts.PUT("/:id", teacherAttsController.Update)
	eTeacherAtts.DELETE("/:id", teacherAttsController.Delete)

	eAcademics := e.Group("/academics")
	eAcademics.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eAcademics.GET("", academicsController.Browse)
	eAcademics.GET("/:id", academicsController.Find)
	eAcademics.POST("/batch", academicsController.CreateBatch)
	eAcademics.POST("", academicsController.Create)
	eAcademics.PUT("/:id", academicsController.Update)
	eAcademics.DELETE("/:id", academicsController.Delete)

	eAcademicAtts := eAcademics.Group("/:academic_id/attendances")
	eAcademicAtts.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eAcademicAtts.GET("", studentAttsController.BrowseByAcademicID)
	eAcademicAtts.PUT("", studentAttsController.UpdateByAcademicID)

	eStudentAtts := eStudents.Group("/attendances")
	eStudentAtts.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eStudentAtts.POST("/batch", studentAttsController.CreateBatch)

	eSubjScheds := e.Group("/subjectschedules")
	eSubjScheds.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eSubjScheds.GET("", subjSchedsController.Browse)
	eSubjScheds.GET("/:id", subjSchedsController.Find)
	eSubjScheds.POST("", subjSchedsController.Create)
	eSubjScheds.PUT("", subjSchedsController.UpdateByAcademicID)
	eSubjScheds.PUT("/:id", subjSchedsController.Update)
	eSubjScheds.DELETE("/:id", subjSchedsController.Delete)

	eClassNotes := eAcademics.Group("/classnotes")
	eClassNotes.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eClassNotes.GET("", classNotesController.Browse)
	eClassNotes.GET("/:id", classNotesController.Find)
	eClassNotes.POST("", classNotesController.Create)
	eClassNotes.POST("/batch", classNotesController.CreateBatch)
	eClassNotes.PUT("/:id", classNotesController.Update)
	eClassNotes.PUT("/detail/:id", classNotesController.UpdateDetail)
	eClassNotes.DELETE("/:id", classNotesController.Delete)

	eAcademicNotes := eAcademics.Group("/:academic_id/classnotes")
	eAcademicNotes.GET("", classNotesController.BrowseByAcademicID)

	eAcademicGrades := eAcademics.Group("/:academic_id/grades")
	eAcademicGrades.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eAcademicGrades.GET("", studentGradesController.GetAll)
	eAcademicGrades.POST("", studentGradesController.Create)
	eAcademicGrades.PUT("", studentGradesController.UpdateByAcademicID)

	tSchedules := eTeachers.Group("/schedules")
	tSchedules.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	tSchedules.GET("", subjSchedsController.GetAllByTeacher)
	tSchedules.GET("/:sched_id/classnotes", classNotesController.FindByTeacher)
}
