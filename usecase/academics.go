package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type AcademicsUsecase interface {
	Browse(page, limit int, search, academicYear string, userId int) ([]dto.GetAcademicsResponse, int64, error)
	Create(academic dto.CreateAcademicsRequest) error
	CreateBatch(batch dto.CreateBatchAcademicsRequest) error
	Find(id int) (interface{}, error)
	Update(id int, academic dto.CreateAcademicsRequest) (models.Academics, error)
	Delete(id int) error

	// Students specific methods
	GetAcademicsByStudent(userID int) ([]dto.StudentGetAcademicsResponse, error)
}

type academicsUsecase struct {
	academicsRepository   repository.AcademicsRepository
	studentsRepository    repository.StudentsRepository
	classroomsRepository  repository.ClassroomsRepository
	teachersRepository    repository.TeachersRepository
	curriculumsRepository repository.CurriculumsRepository
}

func NewAcademicsUsecase(
	academicsRepo repository.AcademicsRepository,
	studentsRepo repository.StudentsRepository,
	classroomsRepo repository.ClassroomsRepository,
	teachersRepo repository.TeachersRepository,
	curriculumsRepo repository.CurriculumsRepository) *academicsUsecase {
	return &academicsUsecase{
		academicsRepository:   academicsRepo,
		studentsRepository:    studentsRepo,
		classroomsRepository:  classroomsRepo,
		teachersRepository:    teachersRepo,
		curriculumsRepository: curriculumsRepo,
	}
}

func validateCreateAcademicsRequest(req dto.CreateAcademicsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func validateBatchCreateAcademicsRequest(req dto.CreateBatchAcademicsRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *academicsUsecase) Browse(page, limit int, search, academicYear string, userId int) ([]dto.GetAcademicsResponse, int64, error) {
	startYear, endYear := "", ""
	if academicYear != "" {
		startYear, endYear = academicYear[:4], academicYear[5:9]
	}

	teacher, _ := s.teachersRepository.GetByToken(userId)
	if teacher.User.Role.Name != "Homeroom Teacher" {
		teacher.ID = 0
	}

	academics, total, err := s.academicsRepository.Browse(page, limit, search, startYear, endYear, int(teacher.ID))

	if err != nil {
		return nil, total, err
	}

	responses := []dto.GetAcademicsResponse{}
	for _, academic := range academics {
		response := dto.GetAcademicsResponse{
			ID:              academic.ID,
			DisplayName:     academic.DisplayName,
			Classroom:       academic.Classroom.DisplayName,
			LevelName:       academic.Classroom.Level.Name,
			Major:           academic.Major,
			Curriculum:      academic.Curriculum.DisplayName,
			HomeroomTeacher: academic.HomeroomTeacher.Name,
			Students:        academic.Students,
			Terms:           academic.Terms,
			CreatedAt:       academic.CreatedAt,
			UpdatedAt:       academic.UpdatedAt,
			DeletedAt:       academic.DeletedAt,
		}
		responses = append(responses, response)
	}

	return responses, total, nil
}

func (s *academicsUsecase) Create(academic dto.CreateAcademicsRequest) error {
	e := validateCreateAcademicsRequest(academic)

	if e != nil {
		return e
	}

	students := []models.Students{}
	if len(academic.Students) > 0 {
		studentsData, err := s.studentsRepository.GetByIds(academic.Students)
		if err != nil {
			return err
		}
		if len(studentsData) == 0 {
			return fmt.Errorf("students not found")
		}
		students = studentsData
	}

	terms := []models.Terms{
		{
			Name: "Semester 1",
		},
		{
			Name: "Semester 2",
		},
	}

	academicData := models.Academics{
		DisplayName:       academic.DisplayName,
		StartYear:         academic.StartYear,
		EndYear:           academic.EndYear,
		ClassroomID:       academic.ClassroomID,
		Major:             academic.Major,
		CurriculumID:      academic.CurriculumID,
		HomeroomTeacherID: academic.HomeroomTeacherID,
		Students:          students,
		Terms:             terms,
	}

	err := s.academicsRepository.Create(academicData)
	if err != nil {
		return err
	}

	return nil
}

func (s *academicsUsecase) CreateBatch(batch dto.CreateBatchAcademicsRequest) error {
	e := validateBatchCreateAcademicsRequest(batch)
	if e != nil {
		return e
	}

	classrooms, eClass := s.classroomsRepository.GetAll()
	if eClass != nil {
		return eClass
	}

	academicsData := []models.Academics{}

	for _, classroom := range classrooms {
		curriculum, err := s.curriculumsRepository.GetCurriculumnsByLevelIDandGrade(int(classroom.LevelID), classroom.Grade)
		if err != nil || len(curriculum) == 0 {
			curriculum = []models.Curriculums{
				{
					ID: 0,
				},
			}
		}
		terms := []models.Terms{
			{
				Name: "Semester 1",
			},
			{
				Name: "Semester 2",
			},
		}
		academic := models.Academics{
			DisplayName:  "T.A." + batch.StartYear + "/" + batch.EndYear + " - " + classroom.DisplayName,
			StartYear:    batch.StartYear,
			EndYear:      batch.EndYear,
			ClassroomID:  classroom.ID,
			CurriculumID: curriculum[0].ID,
			Major:        "General",
			Terms:        terms,
		}
		academicsData = append(academicsData, academic)
	}

	err := s.academicsRepository.CreateBatch(academicsData)
	if err != nil {
		return err
	}

	return nil
}

func (s *academicsUsecase) Find(id int) (interface{}, error) {
	academic, err := s.academicsRepository.Find(id)

	if err != nil {
		return nil, err
	}

	terms := []dto.GetTermResponse{}
	for _, term := range academic.Terms {
		termResponse := dto.GetTermResponse{
			ID:              term.ID,
			Name:            term.Name,
			FirstStartDate:  term.FirstStartDate,
			FirstEndDate:    term.FirstEndDate,
			SecondStartDate: term.SecondStartDate,
			SecondEndDate:   term.SecondEndDate,
		}
		terms = append(terms, termResponse)
	}
	students := []dto.GetStudentResponse{}
	for _, student := range academic.Students {
		response := dto.GetStudentResponse{
			ID:       student.ID,
			FullName: student.FullName,
			NIS:      student.NIS,
			NISN:     student.NISN,
		}
		students = append(students, response)
	}
	scheduleDays := map[string][]dto.GetSubjectScheduleEntryResponse{}
	for _, schedule := range academic.SubjScheds {
		day := schedule.Day
		if _, exists := scheduleDays[day]; !exists {
			scheduleDays[day] = []dto.GetSubjectScheduleEntryResponse{}
		}
		scheduleEntry := dto.GetSubjectScheduleEntryResponse{
			ID:        schedule.ID,
			SubjectID: schedule.Subject.ID,
			Subject:   schedule.Subject.Name,
			TeacherID: schedule.Teacher.ID,
			Teacher:   schedule.Teacher.Name,
			StartHour: schedule.StartHour,
			EndHour:   schedule.EndHour,
		}
		scheduleDays[day] = append(scheduleDays[day], scheduleEntry)
	}
	schedules := []dto.GetSubjectScheduleResponse{}
	for day, entries := range scheduleDays {
		schedule := dto.GetSubjectScheduleResponse{
			Day:     day,
			Entries: entries,
		}
		schedules = append(schedules, schedule)
	}
	classnotes := []dto.GetClassNoteResponse{}
	for _, note := range academic.ClassNotes {
		entries := []dto.GetClassNoteEntryResponse{}
		for _, entry := range note.Details {
			classNoteEntry := dto.GetClassNoteEntryResponse{
				ID:                entry.ID,
				Subject:           entry.SubjSched.Subject.Name,
				SubjectScheduleId: entry.SubjSched.ID,
				Teacher:           entry.SubjSched.Teacher.Name,
				TeacherAct:        entry.Teacher.Name,
				TeacherID:         entry.SubjSched.Teacher.ID,
				TeacherActID:      entry.TeacherID,
				Materials:         entry.Materials,
				Notes:             entry.Notes,
			}
			entries = append(entries, classNoteEntry)
		}
		classNote := dto.GetClassNoteResponse{
			ID:      note.ID,
			Date:    note.Date,
			Entries: entries,
		}
		classnotes = append(classnotes, classNote)
	}

	curriculumSubjects := []dto.GetCurriculumSubjectResponse{}
	for _, subject := range academic.Curriculum.CurriculumSubjects {
		subjectResponse := dto.GetCurriculumSubjectResponse{
			ID:          subject.ID,
			SubjectID:   subject.Subject.ID,
			SubjectName: subject.Subject.Name,
		}
		curriculumSubjects = append(curriculumSubjects, subjectResponse)
	}

	response := dto.GetAcademicDetailResponse{
		ID:                 academic.ID,
		DisplayName:        academic.DisplayName,
		StartYear:          academic.StartYear,
		EndYear:            academic.EndYear,
		Classroom:          academic.Classroom.DisplayName,
		LevelName:          academic.Classroom.Level.Name,
		Major:              academic.Major,
		HomeroomTeacherId:  academic.HomeroomTeacherID,
		HomeroomTeacher:    academic.HomeroomTeacher.Name,
		CurriculumID:       academic.CurriculumID,
		Curriculum:         academic.Curriculum.DisplayName,
		CurriculumSubjects: curriculumSubjects,
		Terms:              terms,
		Students:           students,
		SubjScheds:         schedules,
		ClassNotes:         classnotes,
		CreatedAt:          academic.CreatedAt,
		UpdatedAt:          academic.UpdatedAt,
		DeletedAt:          academic.DeletedAt,
	}

	return response, nil
}

func (s *academicsUsecase) Update(id int, academic dto.CreateAcademicsRequest) (models.Academics, error) {
	academicData, err := s.academicsRepository.Find(id)

	if err != nil {
		return models.Academics{}, err
	}

	if len(academicData.Students) > 0 {
		for _, student := range academicData.Students {
			s.studentsRepository.UpdateStudentCurrentAcademicIdToNull(int(student.ID))
		}
	}

	students := []models.Students{}
	if len(academic.Students) > 0 {
		studentsData, e := s.studentsRepository.GetByIds(academic.Students)
		if e != nil {
			return models.Academics{}, e
		}
		if len(studentsData) == 0 {
			return models.Academics{}, fmt.Errorf("students not found")
		}
		for _, student := range studentsData {
			student.CurrentAcademicID = academicData.ID
			s.studentsRepository.Update(int(student.ID), student)
		}
		students = studentsData
	}

	academicData.DisplayName = academic.DisplayName
	academicData.StartYear = academic.StartYear
	academicData.EndYear = academic.EndYear
	academicData.ClassroomID = academic.ClassroomID
	academicData.CurriculumID = academic.CurriculumID
	academicData.Major = academic.Major
	academicData.HomeroomTeacherID = academic.HomeroomTeacherID
	academicData.Students = students

	e := s.academicsRepository.Update(id, academicData)

	if e != nil {
		return models.Academics{}, e
	}

	academicUpdated, err := s.academicsRepository.Find(id)

	if err != nil {
		return models.Academics{}, err
	}

	return academicUpdated, nil
}

func (s *academicsUsecase) Delete(id int) error {
	err := s.academicsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

// Students specific methods
func (s *academicsUsecase) GetAcademicsByStudent(userID int) ([]dto.StudentGetAcademicsResponse, error) {
	student, err := s.studentsRepository.GetByToken(userID)
	if err != nil {
		return nil, err
	}

	academics, err := s.academicsRepository.GetAcademicsByStudent(int(student.ID))
	if err != nil {
		return nil, err
	}

	responses := []dto.StudentGetAcademicsResponse{}
	for _, academic := range academics {
		terms := []dto.StudentGetTermResponse{}
		for _, term := range academic.Terms {
			terms = append(terms, dto.StudentGetTermResponse{
				ID:          term.ID,
				DisplayName: academic.DisplayName + " - " + term.Name,
				Name:        term.Name,
			})
		}
		response := dto.StudentGetAcademicsResponse{
			ID:          academic.ID,
			DisplayName: academic.DisplayName,
			StartYear:   academic.StartYear,
			EndYear:     academic.EndYear,
			Terms:       terms,
		}
		responses = append(responses, response)
	}

	return responses, nil
}
