package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type ClassNotesUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	BrowseByAcademicID(academicID, page, limit int, search string) (interface{}, int64, error)
	Create(classNote dto.CreateClassNotesRequest) error
	CreateBatch(classNote dto.CreateBatchClassNotesRequest) error
	Find(id int) (interface{}, error)
	// Update(id int, classNote dto.CreateClassNotesRequest) (models.ClassNotes, error)
	// Delete(id int) error
}

type classNotesUsecase struct {
	classNotesRepository repository.ClassNotesRepository
	academicsRepository  repository.AcademicsRepository
}

func NewClassNotesUsecase(classNotesRepo repository.ClassNotesRepository, academicsRepo repository.AcademicsRepository) *classNotesUsecase {
	return &classNotesUsecase{
		classNotesRepository: classNotesRepo,
		academicsRepository:  academicsRepo,
	}
}

func validateCreateClassNotesRequest(req dto.CreateClassNotesRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func validateCreateBatchClassNotesRequest(req dto.CreateBatchClassNotesRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}

func (s *classNotesUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	classNotes, total, err := s.classNotesRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return classNotes, total, nil
}

func (s *classNotesUsecase) BrowseByAcademicID(academicID, page, limit int, search string) (interface{}, int64, error) {
	classNotes, total, err := s.classNotesRepository.BrowseByAcademicID(academicID, page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return classNotes, total, nil
}

func (s *classNotesUsecase) Create(classNote dto.CreateClassNotesRequest) error {
	e := validateCreateClassNotesRequest(classNote)

	if e != nil {
		return e
	}

	details := []models.ClassNotesDetails{}
	if len(classNote.Details) > 0 {
		for _, detail := range classNote.Details {
			detailData := models.ClassNotesDetails{
				SubjSchedID: detail.SubjSchedID,
				TeacherID:   detail.TeacherID,
				Materials:   detail.Materials,
				Notes:       detail.Notes,
			}
			details = append(details, detailData)
		}
	}

	parsedDate, edate := time.Parse("2006-01-02", classNote.Date)
	if edate != nil {
		return edate
	}

	academic, eAcademic := s.academicsRepository.Find(int(classNote.AcademicID))
	if eAcademic != nil {
		return eAcademic
	}

	classNoteData := models.ClassNotes{
		DisplayName: academic.DisplayName + " - " + classNote.Date,
		AcademicID:  classNote.AcademicID,
		Date:        parsedDate,
		Details:     details,
	}

	err := s.classNotesRepository.Create(classNoteData)

	if err != nil {
		return err
	}

	return nil
}

func (s *classNotesUsecase) CreateBatch(classNote dto.CreateBatchClassNotesRequest) error {
	e := validateCreateBatchClassNotesRequest(classNote)

	if e != nil {
		return e
	}

	parsedDate, edate := time.Parse("2006-01-02", classNote.Date)
	if edate != nil {
		return edate
	}

	academics, err := s.academicsRepository.GetAll()
	if err != nil {
		return err
	}

	classNoteData := []models.ClassNotes{}
	for _, academic := range academics {
		details := []models.ClassNotesDetails{}
		for _, subjSched := range academic.SubjScheds {
			detailData := models.ClassNotesDetails{
				SubjSchedID: subjSched.ID,
				TeacherID:   subjSched.TeacherID,
			}
			details = append(details, detailData)
		}
		classNoteData = append(classNoteData, models.ClassNotes{
			DisplayName: academic.DisplayName + " - " + classNote.Date,
			AcademicID:  academic.ID,
			Date:        parsedDate,
			Details:     details,
		})
	}

	err = s.classNotesRepository.CreateBatch(classNoteData)

	if err != nil {
		return err
	}

	return nil
}

func (s *classNotesUsecase) Find(id int) (interface{}, error) {
	classNote, err := s.classNotesRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return classNote, nil
}

// func (s *classNotesUsecase) Update(id int, classNote dto.CreateClassNotesRequest) (models.ClassNotes, error) {
// 	classNoteData, err := s.classNotesRepository.Find(id)

// 	if err != nil {
// 		return models.ClassNotes{}, err
// 	}

// 	students := []models.Students{}
// 	if len(classNote.Students) > 0 {
// 		studentsData, e := s.studentsRepository.GetByIds(classNote.Students)
// 		if e != nil {
// 			return models.ClassNotes{}, e
// 		}
// 		if len(studentsData) == 0 {
// 			return models.ClassNotes{}, fmt.Errorf("Students not found")
// 		}
// 		students = studentsData
// 	}

// 	classNoteData.DisplayName = classNote.DisplayName
// 	classNoteData.StartYear = classNote.StartYear
// 	classNoteData.EndYear = classNote.EndYear
// 	classNoteData.ClassroomID = classNote.ClassroomID
// 	classNoteData.Major = classNote.Major
// 	classNoteData.HomeroomTeacherID = classNote.HomeroomTeacherID
// 	classNoteData.Students = students

// 	e := s.classNotesRepository.Update(id, classNoteData)

// 	if e != nil {
// 		return models.ClassNotes{}, e
// 	}

// 	classNoteUpdated, err := s.classNotesRepository.Find(id)

// 	if err != nil {
// 		return models.ClassNotes{}, err
// 	}

// 	return classNoteUpdated, nil
// }

// func (s *classNotesUsecase) Delete(id int) error {
// 	err := s.classNotesRepository.Delete(id)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
