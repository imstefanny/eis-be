package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type StudentGradesUsecase interface {
	GetAll(termID int) (dto.GetStudentGradesResponse, error)
	Create(studentGrade dto.CreateStudentGradesRequest) error
	UpdateByTermID(termID int, studentGrade dto.UpdateStudentGradesRequest) (dto.GetStudentGradesResponse, error)
	GetAllByStudent(termID int, studentIDs []int) ([]dto.GetPrintReportByStudent, error)
	GetMonthlyReportByStudent(academicID int, studentsIDs []int) ([]dto.GetPrintMonthlyReportByStudent, error)
	GetReport(academicYear string, levelID, academicID, termID int) ([]dto.StudentGradesReport, error)

	// Students specific methods
	GetStudentScoreByStudent(userID, termID int) ([]dto.StudentScoreResponse, error)
}

type studentGradesUsecase struct {
	studentGradesRepository      repository.StudentGradesRepository
	studentAttsRepository        repository.StudentAttsRepository
	academicsRepository          repository.AcademicsRepository
	termsRepository              repository.TermsRepository
	studentsRepository           repository.StudentsRepository
	subjectsRepository           repository.SubjectsRepository
	studentBehaviourRepository   repository.StudentBehaviourActivitiesRepository
	levelHistoriesRepository     repository.LevelHistoriesRepository
	curriculumSubjectsRepository repository.CurriculumSubjectsRepository
	academicStudentsRepository   repository.AcademicStudentsRepository
}

func NewStudentGradesUsecase(
	studentGradesRepo repository.StudentGradesRepository,
	studentAttsRepo repository.StudentAttsRepository,
	academicsRepo repository.AcademicsRepository,
	termsRepo repository.TermsRepository,
	studentsRepo repository.StudentsRepository,
	subjectsRepo repository.SubjectsRepository,
	studentBehaviourRepo repository.StudentBehaviourActivitiesRepository,
	levelHistoriesRepo repository.LevelHistoriesRepository,
	curriculumSubjectsRepo repository.CurriculumSubjectsRepository,
	academicStudentsRepo repository.AcademicStudentsRepository) *studentGradesUsecase {
	return &studentGradesUsecase{
		studentGradesRepository:      studentGradesRepo,
		studentAttsRepository:        studentAttsRepo,
		academicsRepository:          academicsRepo,
		termsRepository:              termsRepo,
		studentsRepository:           studentsRepo,
		subjectsRepository:           subjectsRepo,
		studentBehaviourRepository:   studentBehaviourRepo,
		levelHistoriesRepository:     levelHistoriesRepo,
		curriculumSubjectsRepository: curriculumSubjectsRepo,
		academicStudentsRepository:   academicStudentsRepo,
	}
}

func validateCreateStudentGradesRequest(req dto.CreateStudentGradesRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func validateUpdateStudentGradesRequest(req dto.UpdateStudentGradesRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *studentGradesUsecase) GetAll(termID int) (dto.GetStudentGradesResponse, error) {
	term, err := s.termsRepository.Find(termID)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("term with ID %d not found: %w", termID, err)
	}

	studentGrades, err := s.studentGradesRepository.GetAll(termID)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("error browsing student grades: %w", err)
	}

	details := make(map[uint]*dto.GetStudentGradesDetailResponse)
	for _, grade := range studentGrades {
		if _, exists := details[grade.SubjectID]; !exists {
			details[grade.SubjectID] = &dto.GetStudentGradesDetailResponse{
				SubjectID: grade.SubjectID,
				Subject:   grade.Subject.Name,
				Students:  []dto.GetStudentGradesEntryResponse{},
			}
		}
		details[grade.SubjectID].Students = append(details[grade.SubjectID].Students, dto.GetStudentGradesEntryResponse{
			ID:          grade.ID,
			StudentID:   grade.StudentID,
			StudentName: grade.Student.FullName,
			DisplayName: grade.DisplayName,
			FirstQuiz:   grade.FirstQuiz,
			SecondQuiz:  grade.SecondQuiz,
			FirstMonth:  grade.FirstMonth,
			SecondMonth: grade.SecondMonth,
			Finals:      grade.Finals,
			FinalGrade:  grade.FinalGrade,
			Remarks:     grade.Remarks,
		})
	}
	detailsList := []dto.GetStudentGradesDetailResponse{}
	for _, detail := range details {
		detailsList = append(detailsList, *detail)
	}
	sort.Slice(detailsList, func(i, j int) bool {
		return detailsList[i].SubjectID < detailsList[j].SubjectID
	})
	notesList := []dto.GetTeacherNotesResponse{}
	teacherNotes, err := s.academicStudentsRepository.FindByAcademicID(term.AcademicID)
	for _, notes := range teacherNotes {
		termNotes := ""
		if term.Name == "Semester 1" {
			termNotes = notes.FirstTermNotes
		} else {
			termNotes = notes.SecondTermNotes
		}
		notesList = append(notesList, dto.GetTeacherNotesResponse{
			ID:        notes.ID,
			StudentID: notes.StudentsID,
			Student:   notes.Student.FullName,
			Notes:     termNotes,
		})
	}
	response := dto.GetStudentGradesResponse{
		AcademicID:   term.AcademicID,
		Academic:     term.Academic.DisplayName,
		TermID:       term.ID,
		Term:         term.Name,
		Details:      detailsList,
		TeacherNotes: notesList,
	}

	return response, nil
}

func (s *studentGradesUsecase) Create(studentGrade dto.CreateStudentGradesRequest) error {
	e := validateCreateStudentGradesRequest(studentGrade)
	if e != nil {
		return e
	}

	term, err := s.termsRepository.Find(int(studentGrade.TermID))
	if err != nil {
		return fmt.Errorf("term with ID %d not found: %w", studentGrade.TermID, err)
	}

	studentGradesData := []models.StudentGrades{}
	for _, detail := range studentGrade.Details {
		subject, err := s.subjectsRepository.Find(int(detail.SubjectID))
		if err != nil {
			return fmt.Errorf("subject with ID %d not found: %w", detail.SubjectID, err)
		}
		for _, entry := range detail.Students {
			student, err := s.studentsRepository.Find(int(entry.StudentID))
			if err != nil {
				return fmt.Errorf("student with ID %d not found: %w", entry.StudentID, err)
			}
			finals := math.Round((((entry.FirstMonth+entry.SecondMonth)/2+(entry.FirstQuiz+entry.SecondQuiz)/2)/2 + entry.Finals) / 2)
			studentGradesData = append(studentGradesData, models.StudentGrades{
				DisplayName: term.Academic.DisplayName + " - " + subject.Name + " - " + student.FullName,
				AcademicID:  studentGrade.AcademicID,
				TermID:      term.ID,
				StudentID:   entry.StudentID,
				SubjectID:   detail.SubjectID,
				FirstQuiz:   entry.FirstQuiz,
				SecondQuiz:  entry.SecondQuiz,
				FirstMonth:  entry.FirstMonth,
				SecondMonth: entry.SecondMonth,
				Finals:      entry.Finals,
				FinalGrade:  finals,
				Remarks:     entry.Remarks,
			})
		}
	}

	err = s.studentGradesRepository.Create(studentGradesData)

	if err != nil {
		return err
	}

	return nil
}

func (s *studentGradesUsecase) UpdateByTermID(termID int, studentGrade dto.UpdateStudentGradesRequest) (dto.GetStudentGradesResponse, error) {
	e := validateUpdateStudentGradesRequest(studentGrade)
	if e != nil {
		return dto.GetStudentGradesResponse{}, e
	}
	term, err := s.termsRepository.Find(termID)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("term with ID %d not found: %w", termID, err)
	}

	studentGradeData, err := s.studentGradesRepository.GetAll(termID)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("error browsing student grades: %w", err)
	}

	existingGrades := make(map[uint]models.StudentGrades)
	for _, grade := range studentGradeData {
		existingGrades[grade.ID] = grade
	}
	studentGradesData := []models.StudentGrades{}
	newStudents := []models.StudentGrades{}
	for _, detail := range studentGrade.Details {
		subject, err := s.subjectsRepository.Find(int(detail.SubjectID))
		if err != nil {
			return dto.GetStudentGradesResponse{}, fmt.Errorf("subject with ID %d not found: %w", detail.SubjectID, err)
		}
		for _, grade := range detail.Students {
			finals := math.Round((((grade.FirstMonth+grade.SecondMonth)/2+(grade.FirstQuiz+grade.SecondQuiz)/2)/2 + grade.Finals) / 2)
			if grade.ID != 0 {
				studentGradesData = append(studentGradesData, models.StudentGrades{
					ID:          grade.ID,
					DisplayName: existingGrades[grade.ID].DisplayName,
					AcademicID:  existingGrades[grade.ID].AcademicID,
					TermID:      existingGrades[grade.ID].TermID,
					StudentID:   existingGrades[grade.ID].StudentID,
					SubjectID:   detail.SubjectID,
					FirstQuiz:   grade.FirstQuiz,
					SecondQuiz:  grade.SecondQuiz,
					FirstMonth:  grade.FirstMonth,
					SecondMonth: grade.SecondMonth,
					Finals:      grade.Finals,
					FinalGrade:  finals,
					Remarks:     grade.Remarks,
				})
			} else {
				student, _ := s.studentsRepository.Find(int(grade.StudentID))
				newStudents = append(newStudents, models.StudentGrades{
					DisplayName: term.Academic.DisplayName + " - " + subject.Name + " - " + student.FullName,
					AcademicID:  studentGrade.AcademicID,
					TermID:      term.ID,
					StudentID:   grade.StudentID,
					SubjectID:   detail.SubjectID,
					FirstQuiz:   grade.FirstQuiz,
					SecondQuiz:  grade.SecondQuiz,
					FirstMonth:  grade.FirstMonth,
					SecondMonth: grade.SecondMonth,
					Finals:      grade.Finals,
					FinalGrade:  finals,
					Remarks:     grade.Remarks,
				})
			}
		}
	}

	err = s.studentGradesRepository.UpdateByTermID(studentGradesData, newStudents)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("error updating student grades for term ID %d: %w", termID, err)
	}

	studentGradesUpdated, err := s.studentGradesRepository.GetAll(termID)
	if err != nil {
		return dto.GetStudentGradesResponse{}, fmt.Errorf("error browsing student grades: %w", err)
	}

	details := make(map[uint]*dto.GetStudentGradesDetailResponse)
	for _, grade := range studentGradesUpdated {
		if _, exists := details[grade.SubjectID]; !exists {
			details[grade.SubjectID] = &dto.GetStudentGradesDetailResponse{
				SubjectID: grade.SubjectID,
				Subject:   grade.Subject.Name,
				Students:  []dto.GetStudentGradesEntryResponse{},
			}
		}
		details[grade.SubjectID].Students = append(details[grade.SubjectID].Students, dto.GetStudentGradesEntryResponse{
			ID:          grade.ID,
			StudentID:   grade.StudentID,
			StudentName: grade.Student.FullName,
			DisplayName: grade.DisplayName,
			FirstQuiz:   grade.FirstQuiz,
			SecondQuiz:  grade.SecondQuiz,
			FirstMonth:  grade.FirstMonth,
			SecondMonth: grade.SecondMonth,
			Finals:      grade.Finals,
			FinalGrade:  grade.FinalGrade,
			Remarks:     grade.Remarks,
		})
	}
	var detailsList []dto.GetStudentGradesDetailResponse
	for _, detail := range details {
		detailsList = append(detailsList, *detail)
	}
	response := dto.GetStudentGradesResponse{
		AcademicID: term.AcademicID,
		Academic:   term.Academic.DisplayName,
		TermID:     term.ID,
		Term:       term.Name,
		Details:    detailsList,
	}

	return response, nil
}

func (s *studentGradesUsecase) GetAllByStudent(termID int, studentIDs []int) ([]dto.GetPrintReportByStudent, error) {
	term, err := s.termsRepository.Find(termID)
	if err != nil {
		return []dto.GetPrintReportByStudent{}, fmt.Errorf("term with ID %d not found: %w", termID, err)
	}
	responses := []dto.GetPrintReportByStudent{}
	for _, studentID := range studentIDs {
		studentScores, err := s.studentGradesRepository.GetStudentScoreByStudent(studentID, termID)
		if err != nil {
			return []dto.GetPrintReportByStudent{}, fmt.Errorf("error getting student scores: %w", err)
		}
		studentAtts, err := s.studentAttsRepository.GetByTermStudent(termID, studentID)
		if err != nil {
			return []dto.GetPrintReportByStudent{}, fmt.Errorf("error getting student attendance: %w", err)
		}
		student, err := s.studentsRepository.Find(studentID)
		if err != nil {
			return []dto.GetPrintReportByStudent{}, fmt.Errorf("student with ID %d not found: %w", studentID, err)
		}
		grades := []dto.GetPrintReportGrade{}
		for _, score := range studentScores {
			remarks, _ := s.curriculumSubjectsRepository.GetByCurriculumSubjectID(term.Academic.CurriculumID, score.SubjectID)
			grades = append(grades, dto.GetPrintReportGrade{
				Subject: score.Subject.Name,
				Finals:  score.FinalGrade,
				Remarks: remarks.Competence,
			})
		}
		extracurriculars := []dto.GetPrintReportExtracurricular{}
		extracurricularsData, err := s.studentBehaviourRepository.FindByStudentIDAndAcademicIDAndTermID(student.ID, term.AcademicID, term.ID)
		if err != nil {
			extracurricularsData = &models.StudentBehaviourActivities{
				FirstMonthExtracurricularFirst:        "",
				FirstMonthExtracurricularScoreFirst:   "",
				FirstMonthExtracurricularSecond:       "",
				FirstMonthExtracurricularScoreSecond:  "",
				SecondMonthExtracurricularFirst:       "",
				SecondMonthExtracurricularScoreFirst:  "",
				SecondMonthExtracurricularSecond:      "",
				SecondMonthExtracurricularScoreSecond: "",
			}
		}
		extMap := map[string]int{
			"A": 4,
			"B": 3,
			"C": 2,
			"D": 1,
			"E": 0,
		}
		stFirstExtScore := extMap[extracurricularsData.FirstMonthExtracurricularScoreFirst]
		stSecondExtScore := extMap[extracurricularsData.FirstMonthExtracurricularScoreSecond]
		ndFirstExtScore := extMap[extracurricularsData.SecondMonthExtracurricularScoreFirst]
		ndSecondExtScore := extMap[extracurricularsData.SecondMonthExtracurricularScoreSecond]
		firstScore := ""
		if stFirstExtScore > ndFirstExtScore {
			firstScore = extracurricularsData.FirstMonthExtracurricularScoreFirst
		} else {
			firstScore = extracurricularsData.SecondMonthExtracurricularScoreFirst
		}
		secondScore := ""
		if stSecondExtScore > ndSecondExtScore {
			secondScore = extracurricularsData.FirstMonthExtracurricularScoreSecond
		} else {
			secondScore = extracurricularsData.SecondMonthExtracurricularScoreSecond
		}
		extracurriculars = append(extracurriculars, []dto.GetPrintReportExtracurricular{
			{
				Name:  extracurricularsData.FirstMonthExtracurricularFirst,
				Score: firstScore,
			},
			{
				Name:  extracurricularsData.FirstMonthExtracurricularSecond,
				Score: secondScore,
			},
			{
				Name:  "",
				Score: "",
			},
		}...)
		attsMap := make(map[string]int)
		for _, att := range studentAtts {
			if _, exists := attsMap[att.Status]; !exists {
				attsMap[att.Status] = 0
			}
			attsMap[att.Status]++
		}
		academicStudent, _ := s.academicStudentsRepository.FindByAcademicIDAndStudentID(term.AcademicID, student.ID)
		teacherNotes := ""
		principal := ""
		termDate := term.EndDate.Format("2006-01-02")
		if term.Name == "Semester 1" {
			teacherNotes = academicStudent.FirstTermNotes
		} else {
			teacherNotes = academicStudent.SecondTermNotes
		}
		levelHistories, _ := s.levelHistoriesRepository.GetAllByLevelID(term.Academic.Classroom.LevelID)
		rangeMap := []map[string]string{}
		if len(levelHistories) != 0 {
			for _, levelHistory := range levelHistories {
				start := levelHistory.CreatedAt
				end := time.Time{}
				if levelHistory.DeletedAt.Valid {
					end = levelHistory.DeletedAt.Time
				} else {
					end, _ = time.Parse("2006-01-02", termDate)
				}
				rangeMap = append(rangeMap, map[string]string{
					"start":     start.Format("2006-01-02"),
					"end":       end.Format("2006-01-02"),
					"principal": levelHistory.Principle.Name,
				})
			}
		}
		for _, r := range rangeMap {
			if r["start"] <= termDate && r["end"] >= termDate {
				principal = r["principal"]
				break
			}
		}
		termName, _ := strconv.Atoi(term.Name[9:])
		response := dto.GetPrintReportByStudent{
			Name:             student.FullName,
			NIS:              student.NIS,
			NISN:             student.NISN,
			Level:            term.Academic.Classroom.Level.Name,
			Class:            term.Academic.Classroom.Grade,
			Fase:             term.Academic.Classroom.Name,
			Term:             termName,
			AcademicYear:     term.Academic.StartYear + "/" + term.Academic.EndYear,
			Grades:           grades,
			Extracurriculars: extracurriculars,
			Sick:             attsMap["Sick"],
			Absent:           attsMap["Absent"],
			Permission:       attsMap["Permission"],
			HomeRoomTeacher:  term.Academic.HomeroomTeacher.Name,
			Principal:        principal,
			EndDate:          termDate,
			TeacherNotes:     teacherNotes,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (s *studentGradesUsecase) GetMonthlyReportByStudent(academicID int, studentsIDs []int) ([]dto.GetPrintMonthlyReportByStudent, error) {
	academic, err := s.academicsRepository.Find(academicID)
	if err != nil {
		return []dto.GetPrintMonthlyReportByStudent{}, fmt.Errorf("academic year with ID %d not found: %w", academicID, err)
	}
	stTerm := models.Terms{}
	ndTerm := models.Terms{}
	for _, term := range academic.Terms {
		if term.Name == "Semester 1" {
			stTerm = term
		} else if term.Name == "Semester 2" {
			ndTerm = term
		}
	}
	stFirstDate := stTerm.FirstEndDate
	stSecondDate := stTerm.SecondEndDate
	ndFirstDate := ndTerm.FirstEndDate
	ndSecondDate := ndTerm.SecondEndDate
	responses := []dto.GetPrintMonthlyReportByStudent{}
	for _, studentID := range studentsIDs {
		student, err := s.studentsRepository.Find(studentID)
		if err != nil {
			return []dto.GetPrintMonthlyReportByStudent{}, fmt.Errorf("student with ID %d not found: %w", studentID, err)
		}
		monthlyScores, err := s.studentGradesRepository.GetMonthlyReportByStudent(academicID, studentID)
		if err != nil {
			return []dto.GetPrintMonthlyReportByStudent{}, fmt.Errorf("error getting monthly scores for student ID %d: %w", studentID, err)
		}
		stFirstAtts, _ := s.studentAttsRepository.GetAttendanceByStudent(studentID, stTerm.FirstStartDate.Format("2006-01-02"), stTerm.FirstEndDate.Format("2006-01-02"))
		stSecondAtts, _ := s.studentAttsRepository.GetAttendanceByStudent(studentID, stTerm.SecondStartDate.Format("2006-01-02"), stTerm.SecondEndDate.Format("2006-01-02"))
		ndFirstAtts, _ := s.studentAttsRepository.GetAttendanceByStudent(studentID, ndTerm.FirstStartDate.Format("2006-01-02"), ndTerm.FirstEndDate.Format("2006-01-02"))
		ndSecondAtts, _ := s.studentAttsRepository.GetAttendanceByStudent(studentID, ndTerm.SecondStartDate.Format("2006-01-02"), ndTerm.SecondEndDate.Format("2006-01-02"))
		stFirstAttsMap := make(map[string]int)
		stSecondAttsMap := make(map[string]int)
		ndFirstAttsMap := make(map[string]int)
		ndSecondAttsMap := make(map[string]int)
		for _, att := range stFirstAtts {
			if _, exists := stFirstAttsMap[att.Status]; !exists {
				stFirstAttsMap[att.Status] = 0
			}
			stFirstAttsMap[att.Status]++
		}
		for _, att := range stSecondAtts {
			if _, exists := stSecondAttsMap[att.Status]; !exists {
				stSecondAttsMap[att.Status] = 0
			}
			stSecondAttsMap[att.Status]++
		}
		for _, att := range ndFirstAtts {
			if _, exists := ndFirstAttsMap[att.Status]; !exists {
				ndFirstAttsMap[att.Status] = 0
			}
			ndFirstAttsMap[att.Status]++
		}
		for _, att := range ndSecondAtts {
			if _, exists := ndSecondAttsMap[att.Status]; !exists {
				ndSecondAttsMap[att.Status] = 0
			}
			ndSecondAttsMap[att.Status]++
		}
		response := dto.GetPrintMonthlyReportByStudent{
			Name:               student.FullName,
			NIS:                student.NIS,
			Class:              academic.Classroom.DisplayName,
			AcademicYear:       academic.StartYear + "/" + academic.EndYear,
			Grades:             monthlyScores,
			HomeRoomTeacher:    academic.HomeroomTeacher.Name,
			StFirstSick:        stFirstAttsMap["Sick"],
			StSecondSick:       stSecondAttsMap["Sick"],
			StFirstPermission:  stFirstAttsMap["Permission"],
			StSecondPermission: stSecondAttsMap["Permission"],
			StFirstAbsent:      stFirstAttsMap["Absent"],
			StSecondAbsent:     stSecondAttsMap["Absent"],
			NdFirstSick:        ndFirstAttsMap["Sick"],
			NdSecondSick:       ndSecondAttsMap["Sick"],
			NdFirstPermission:  ndFirstAttsMap["Permission"],
			NdSecondPermission: ndSecondAttsMap["Permission"],
			NdFirstAbsent:      ndFirstAttsMap["Absent"],
			NdSecondAbsent:     ndSecondAttsMap["Absent"],
			StFirstDate:        stFirstDate.Format("02 January 2006"),
			StSecondDate:       stSecondDate.Format("02 January 2006"),
			NdFirstDate:        ndFirstDate.Format("02 January 2006"),
			NdSecondDate:       ndSecondDate.Format("02 January 2006"),
		}
		stTermBehavior, _ := s.studentBehaviourRepository.FindByStudentIDAndAcademicIDAndTermID(student.ID, academic.ID, stTerm.ID)
		ndTermBehavior, _ := s.studentBehaviourRepository.FindByStudentIDAndAcademicIDAndTermID(student.ID, academic.ID, ndTerm.ID)
		if stTermBehavior != nil {
			response.StFirstBehavior = stTermBehavior.FirstBehaviour
			response.StSecondBehavior = stTermBehavior.SecondBehaviour
			response.StFirstCraft = stTermBehavior.FirstCrafts
			response.StSecondCraft = stTermBehavior.SecondCrafts
			response.StFirstTidiness = stTermBehavior.FirstNeatness
			response.StSecondTidiness = stTermBehavior.SecondNeatness
			response.StFirstExtracurricularFirst = stTermBehavior.FirstMonthExtracurricularFirst
			response.StFirstExtracurricularScoreFirst = stTermBehavior.FirstMonthExtracurricularScoreFirst
			response.StFirstExtracurricularSecond = stTermBehavior.FirstMonthExtracurricularSecond
			response.StFirstExtracurricularScoreSecond = stTermBehavior.FirstMonthExtracurricularScoreSecond
			response.StSecondExtracurricularFirst = stTermBehavior.SecondMonthExtracurricularFirst
			response.StSecondExtracurricularScoreFirst = stTermBehavior.SecondMonthExtracurricularScoreFirst
			response.StSecondExtracurricularSecond = stTermBehavior.SecondMonthExtracurricularSecond
			response.StSecondExtracurricularScoreSecond = stTermBehavior.SecondMonthExtracurricularScoreSecond
			response.StFirstNotes = stTermBehavior.FirstNotes
			response.StSecondNotes = stTermBehavior.SecondNotes
		}
		if ndTermBehavior != nil {
			response.NdFirstBehavior = ndTermBehavior.FirstBehaviour
			response.NdSecondBehavior = ndTermBehavior.SecondBehaviour
			response.NdFirstCraft = ndTermBehavior.FirstCrafts
			response.NdSecondCraft = ndTermBehavior.SecondCrafts
			response.NdFirstExtracurricularFirst = ndTermBehavior.FirstMonthExtracurricularFirst
			response.NdFirstExtracurricularScoreFirst = ndTermBehavior.FirstMonthExtracurricularScoreFirst
			response.NdFirstExtracurricularSecond = ndTermBehavior.FirstMonthExtracurricularSecond
			response.NdFirstExtracurricularScoreSecond = ndTermBehavior.FirstMonthExtracurricularScoreSecond
			response.NdSecondExtracurricularFirst = stTermBehavior.SecondMonthExtracurricularFirst
			response.NdSecondExtracurricularScoreFirst = stTermBehavior.SecondMonthExtracurricularScoreFirst
			response.NdSecondExtracurricularSecond = stTermBehavior.SecondMonthExtracurricularSecond
			response.NdSecondExtracurricularScoreSecond = stTermBehavior.SecondMonthExtracurricularScoreSecond
			response.NdFirstNotes = ndTermBehavior.FirstNotes
			response.NdSecondNotes = ndTermBehavior.SecondNotes
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (s *studentGradesUsecase) GetReport(academicYear string, levelID, academicID, termID int) ([]dto.StudentGradesReport, error) {
	startYear, endYear := "", ""
	if academicYear != "" {
		startYear, endYear = academicYear[:4], academicYear[5:9]
	}

	studentGrades, err := s.studentGradesRepository.GetReport(startYear, endYear, levelID, academicID, termID)
	if err != nil {
		return []dto.StudentGradesReport{}, err
	}
	responses := []dto.StudentGradesReport{}
	classMap := make(map[string][]dto.StudentGradesReportTopStudent)
	for _, grade := range studentGrades {
		if _, exists := classMap[grade.Class]; !exists {
			classMap[grade.Class] = []dto.StudentGradesReportTopStudent{}
		}
		top := dto.StudentGradesReportTopStudent{
			Rank:    0,
			Student: grade.Student,
			NIS:     grade.NIS,
			Class:   grade.Class,
			Finals:  grade.Finals,
		}
		classMap[grade.Class] = append(classMap[grade.Class], top)
	}
	for class, students := range classMap {
		average := 0.0
		for idx, student := range students {
			average += student.Finals
			students[idx].Rank = idx + 1
		}
		if len(students) > 0 {
			average /= float64(len(students))
		}
		responses = append(responses, dto.StudentGradesReport{
			Class:    class,
			Average:  math.Round(average*100) / 100,
			Students: students,
		})
	}
	return responses, nil
}

// Students specific methods
func (s *studentGradesUsecase) GetStudentScoreByStudent(userID, termID int) ([]dto.StudentScoreResponse, error) {
	student, err := s.studentsRepository.GetByToken(userID)
	if err != nil {
		return []dto.StudentScoreResponse{}, fmt.Errorf("student with user ID %d not found", userID)
	}

	studentScores, err := s.studentGradesRepository.GetStudentScoreByStudent(int(student.ID), termID)
	if err != nil {
		return nil, err
	}

	responses := []dto.StudentScoreResponse{}
	for _, score := range studentScores {
		responses = append(responses, dto.StudentScoreResponse{
			SubjectName: score.Subject.Name,
			FirstQuiz:   score.FirstQuiz,
			SecondQuiz:  score.SecondQuiz,
			FirstMonth:  score.FirstMonth,
			SecondMonth: score.SecondMonth,
			Finals:      score.Finals,
			FinalGrade:  score.FinalGrade,
			Remarks:     score.Remarks,
		})
	}
	return responses, nil
}
