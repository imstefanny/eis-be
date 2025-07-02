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
	rolesRepository := repository.NewRolesRepository(db)
	permissionsRepository := repository.NewPermissionsRepository(db)
	rolesService := usecase.NewRolesUsecase(rolesRepository, permissionsRepository)
	rolesController := controllers.NewRolesController(rolesService)

	usersRepository := repository.NewUsersRepository(db)
	usersService := usecase.NewUsersUsecase(usersRepository, rolesRepository, db)
	usersController := controllers.NewUsersController(usersService)

	blogsRepository := repository.NewBlogsRepository(db)
	blogsService := usecase.NewBlogsUsecase(blogsRepository, usersRepository)
	blogsController := controllers.NewBlogsController(blogsService)

	studentsRepository := repository.NewStudentsRepository(db)
	studentsService := usecase.NewStudentsUsecase(studentsRepository, usersRepository, rolesRepository, db)
	studentsController := controllers.NewStudentsController(studentsService)

	workSchedsRepository := repository.NewWorkSchedsRepository(db)
	workSchedDetailsRepository := repository.NewWorkSchedDetailsRepository(db)
	workSchedsService := usecase.NewWorkSchedsUsecase(workSchedsRepository, workSchedDetailsRepository)
	workSchedsController := controllers.NewWorkSchedsController(workSchedsService)

	teachersRepository := repository.NewTeachersRepository(db)
	teachersService := usecase.NewTeachersUsecase(teachersRepository, usersRepository, db)
	teachersController := controllers.NewTeachersController(teachersService)

	teacherAttsRepository := repository.NewTeacherAttsRepository(db)
	teacherAttsService := usecase.NewTeacherAttsUsecase(teacherAttsRepository, teachersRepository, workSchedsRepository)
	teacherAttsController := controllers.NewTeacherAttsController(teacherAttsService)

	guardiansRepository := repository.NewGuardiansRepository(db)
	applicantsRepository := repository.NewApplicantsRepository(db)

	documentsRepository := repository.NewDocumentsRepository(db)
	documentsService := usecase.NewDocumentsUsecase(documentsRepository, applicantsRepository)
	documentsController := controllers.NewDocumentsController(documentsService)

	applicantsService := usecase.NewApplicantsUsecase(applicantsRepository, studentsRepository, guardiansRepository, usersRepository, rolesRepository, documentsRepository)
	applicantsController := controllers.NewApplicantsController(applicantsService)

	guardiansService := usecase.NewGuardiansUsecase(guardiansRepository, applicantsRepository)
	guardiansController := controllers.NewGuardiansController(guardiansService)

	docTypesRepository := repository.NewDocTypesRepository(db)
	docTypesService := usecase.NewDocTypesUsecase(docTypesRepository)
	docTypesController := controllers.NewDocTypesController(docTypesService)

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

	termsRepository := repository.NewTermsRepository(db)
	termsService := usecase.NewTermsUsecase(termsRepository)
	termsController := controllers.NewTermsController(termsService)

	curriculumSubjectsRepository := repository.NewCurriculumSubjectsRepository(db)
	curriculumsRepository := repository.NewCurriculumsRepository(db)
	academicsRepository := repository.NewAcademicsRepository(db)

	curriculumsService := usecase.NewCurriculumsUsecase(curriculumsRepository, levelsRepository, academicsRepository)
	curriculumsController := controllers.NewCurriculumsController(curriculumsService)

	academicsService := usecase.NewAcademicsUsecase(academicsRepository, studentsRepository, classroomsRepository, teachersRepository, curriculumsRepository)
	academicsController := controllers.NewAcademicsController(academicsService)

	academicStudentsRepository := repository.NewAcademicStudentsRepository(db)
	academicStudentsService := usecase.NewAcademicStudentsUsecase(academicStudentsRepository)
	academicStudentsController := controllers.NewAcademicStudentsController(academicStudentsService)

	subjSchedsRepository := repository.NewSubjSchedsRepository(db)
	subjSchedsService := usecase.NewSubjSchedsUsecase(subjSchedsRepository, academicsRepository, teachersRepository, studentsRepository)
	subjSchedsController := controllers.NewSubjSchedsController(subjSchedsService)

	studentAttsRepository := repository.NewStudentAttsRepository(db)
	studentAttsService := usecase.NewStudentAttsUsecase(studentAttsRepository, studentsRepository, academicsRepository, termsRepository)
	studentAttsController := controllers.NewStudentAttsController(studentAttsService)

	classNotesRepository := repository.NewClassNotesRepository(db)
	classNotesService := usecase.NewClassNotesUsecase(classNotesRepository, academicsRepository, studentAttsRepository, teachersRepository)
	classNotesController := controllers.NewClassNotesController(classNotesService)

	classNotesDetailsRepository := repository.NewClassNotesDetailsRepository(db)
	classNotesDetailsService := usecase.NewClassNotesDetailsUsecase(classNotesDetailsRepository, teachersRepository, studentAttsRepository)
	classNotesDetailsController := controllers.NewClassNotesDetailsController(classNotesDetailsService)

	studentBehaviourRepository := repository.NewStudentBehaviourActivitiesRepository(db)
	studentBehaviourService := usecase.NewStudentBehaviourActivitiesUsecase(studentBehaviourRepository, academicsRepository, termsRepository, studentsRepository)
	studentBehaviourController := controllers.NewStudentBehaviourActivitiesController(studentBehaviourService)

	studentGradesRepository := repository.NewStudentGradesRepository(db)
	studentGradesService := usecase.NewStudentGradesUsecase(
		studentGradesRepository,
		studentAttsRepository,
		academicsRepository,
		termsRepository,
		studentsRepository,
		subjectsRepository,
		studentBehaviourRepository,
		levelHistoriesRepository,
		curriculumSubjectsRepository,
		academicStudentsRepository,
	)
	studentGradesController := controllers.NewStudentGradesController(studentGradesService)

	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/register", usersController.Register)
	e.POST("/login", usersController.Login)

	eUser := e.Group("/users")
	eUser.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eUser.GET("", usersController.Browse)
	eUser.PUT("/:id", usersController.Update)
	eUser.PUT("/change-password/:id", usersController.ChangePassword)
	eUser.PUT("/undelete/:id", usersController.Undelete)
	eUser.DELETE("/:id", usersController.Delete)

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
	eApplicants.POST("/reject/:id", applicantsController.RejectRegistration)

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
	eLevelHistories.POST("", levelHistoriesController.Create)

	eCurriculums := e.Group("/curriculums")
	eCurriculums.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eCurriculums.GET("", curriculumsController.Browse)
	eCurriculums.GET("/:id", curriculumsController.Find)
	eCurriculums.POST("", curriculumsController.Create)
	eCurriculums.PUT("/:id", curriculumsController.Update)
	eCurriculums.DELETE("/:id", curriculumsController.Delete)
	eCurriculums.PUT("/undelete/:id", curriculumsController.UnDelete)

	eClassrooms := e.Group("/classrooms")
	eClassrooms.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eClassrooms.GET("", classroomsController.Browse)
	eClassrooms.GET("/:id", classroomsController.Find)
	eClassrooms.POST("", classroomsController.Create)
	eClassrooms.PUT("/:id", classroomsController.Update)
	eClassrooms.DELETE("/:id", classroomsController.Delete)
	eClassrooms.PUT("/undelete/:id", classroomsController.UnDelete)

	eStudents := e.Group("/students")
	eStudents.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eStudents.GET("", studentsController.Browse)
	eStudents.GET("/:id", studentsController.Find)
	eStudents.POST("", studentsController.Create)
	eStudents.PUT("/:id", studentsController.Update)
	eStudents.PUT("/update-current-academic/:academic_id", studentsController.UpdateStudentAcademicId)
	eStudents.DELETE("/:id", studentsController.Delete)

	eTeachers := e.Group("/teachers")
	eTeachers.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eTeachers.GET("", teachersController.Browse)
	eTeachers.GET("/my", teachersController.GetByToken)
	eTeachers.GET("/available-homeroom", teachersController.GetAvailableHomeroomTeachers)
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
	eTeacherAtts.GET("/report", teacherAttsController.GetReport)

	eTerms := e.Group("/terms")
	eTerms.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eTerms.PUT("/:id", termsController.Update)

	eAcademics := e.Group("/academics")
	eAcademics.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eAcademics.GET("", academicsController.Browse)
	eAcademics.GET("/:id", academicsController.Find)
	eAcademics.POST("/batch", academicsController.CreateBatch)
	eAcademics.POST("", academicsController.Create)
	eAcademics.PUT("/:id", academicsController.Update)
	eAcademics.DELETE("/:id", academicsController.Delete)

	eAcademicAtts := eAcademics.Group("/:academic_id/:term_id/attendances")
	eAcademicAtts.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eAcademicAtts.GET("", studentAttsController.BrowseByTermID)
	eAcademicAtts.PUT("", studentAttsController.UpdateByTermID)

	eAcademicHTNotes := eAcademics.Group("/notes")
	eAcademicHTNotes.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eAcademicHTNotes.PUT("", academicStudentsController.Update)

	eStudentAtts := eStudents.Group("/attendances")
	eStudentAtts.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eStudentAtts.POST("/batch", studentAttsController.CreateBatch)
	eStudentAtts.GET("/report", studentAttsController.GetReport)

	eStudentMarks := eStudents.Group("/marks")
	eStudentMarks.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eStudentMarks.GET("/report", studentGradesController.GetReport)

	eStudentBehaviour := eStudents.Group("/behaviour/:academic_id/:term_id")
	eStudentBehaviour.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eStudentBehaviour.GET("", studentBehaviourController.GetByAcademicIdAndTermId)
	eStudentBehaviour.POST("", studentBehaviourController.Create)
	eStudentBehaviour.PUT("", studentBehaviourController.Update)

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

	eAcademicGrades := eAcademics.Group("/:academic_id/:term_id/grades")
	eAcademicGrades.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	eAcademicGrades.GET("", studentGradesController.GetAll)
	eAcademicGrades.POST("", studentGradesController.Create)
	eAcademicGrades.PUT("", studentGradesController.UpdateByTermID)

	eAcademicReports := eAcademics.Group("/:academic_id/report")
	eAcademicReports.GET("/:term_id/:student_ids", studentGradesController.GetAllByStudent)
	eAcademicReports.GET("/monthly/:student_ids", studentGradesController.GetMonthlyReportByStudent)

	tSchedules := eTeachers.Group("/schedules")
	tSchedules.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	tSchedules.GET("", classNotesDetailsController.GetAllByTeacher)
	tSchedules.GET("/origin", subjSchedsController.GetAllByTeacher)
	tSchedules.GET("/:sched_id/classnotes", classNotesController.FindByTeacher)

	sStudents := eStudents.Group("/my")
	sStudents.Use(echojwt.JWT([]byte(constants.SECRET_KEY)))
	sStudents.GET("", studentsController.GetByToken)
	sStudents.GET("/academics", academicsController.GetAcademicsByStudent)
	sStudents.GET("/attendances", studentAttsController.GetAttendanceByStudent)
	sStudents.GET("/schedules", subjSchedsController.GetScheduleByStudent)
	sStudents.GET("/:academic_id/:term_id/scores", studentGradesController.GetStudentScoreByStudent)
}
