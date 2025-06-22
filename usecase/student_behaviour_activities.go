package usecase

import (
	"eis-be/dto"
	"eis-be/models"
	"eis-be/repository"
	"fmt"
)

type StudentBehaviourActivitiesUsecase interface {
	GetByAcademicIdAndTermId(academicID, termID int) ([]dto.StudentBehaviourActivityRequest, error)
	Create(studentGrade []dto.StudentBehaviourActivityRequest) error
	Update(studentGrade []dto.StudentBehaviourActivityRequest) error
}

type studentBehaviourActivitiesUsecase struct {
	studentBehaviourRepository repository.StudentBehaviourActivitiesRepository
	academicsRepository        repository.AcademicsRepository
	termsRepository            repository.TermsRepository
	studentsRepository         repository.StudentsRepository
}

func NewStudentBehaviourActivitiesUsecase(studentBehaviourRepo repository.StudentBehaviourActivitiesRepository, academicsRepo repository.AcademicsRepository, termsRepo repository.TermsRepository, studentsRepo repository.StudentsRepository) *studentBehaviourActivitiesUsecase {
	return &studentBehaviourActivitiesUsecase{
		studentBehaviourRepository: studentBehaviourRepo,
		academicsRepository:        academicsRepo,
		termsRepository:            termsRepo,
		studentsRepository:         studentsRepo,
	}
}

func (s *studentBehaviourActivitiesUsecase) GetByAcademicIdAndTermId(academicId, termID int) ([]dto.StudentBehaviourActivityRequest, error) {
	studentBehaviour, err := s.studentBehaviourRepository.GetByAcademicIdAndTermId(academicId, termID)
	if err != nil {
		return []dto.StudentBehaviourActivityRequest{}, fmt.Errorf("error browsing student grades: %w", err)
	}

	response := []dto.StudentBehaviourActivityRequest{}
	for _, studentBehaviour := range studentBehaviour {
		response = append(response, dto.StudentBehaviourActivityRequest{
			ID:                                    studentBehaviour.ID,
			AcademicID:                            studentBehaviour.AcademicID,
			TermID:                                studentBehaviour.TermID,
			StudentID:                             studentBehaviour.StudentID,
			StudentNIS:                            studentBehaviour.Student.NIS,
			StudentName:                           studentBehaviour.Student.FullName,
			FirstBehaviour:                        studentBehaviour.FirstBehaviour,
			FirstNeatness:                         studentBehaviour.FirstNeatness,
			FirstCrafts:                           studentBehaviour.FirstCrafts,
			FirstMonthExtracurricularFirst:        studentBehaviour.FirstMonthExtracurricularFirst,
			FirstMonthExtracurricularScoreFirst:   studentBehaviour.FirstMonthExtracurricularScoreFirst,
			FirstMonthExtracurricularSecond:       studentBehaviour.FirstMonthExtracurricularSecond,
			FirstMonthExtracurricularScoreSecond:  studentBehaviour.FirstMonthExtracurricularScoreSecond,
			SecondNeatness:                        studentBehaviour.SecondNeatness,
			SecondCrafts:                          studentBehaviour.SecondCrafts,
			SecondBehaviour:                       studentBehaviour.SecondBehaviour,
			SecondMonthExtracurricularFirst:       studentBehaviour.SecondMonthExtracurricularFirst,
			SecondMonthExtracurricularScoreFirst:  studentBehaviour.SecondMonthExtracurricularScoreFirst,
			SecondMonthExtracurricularSecond:      studentBehaviour.SecondMonthExtracurricularSecond,
			SecondMonthExtracurricularScoreSecond: studentBehaviour.SecondMonthExtracurricularScoreSecond,
		})
	}

	return response, nil
}

func (s *studentBehaviourActivitiesUsecase) Create(studentBehaviour []dto.StudentBehaviourActivityRequest) error {

	studentBehaviourData := []models.StudentBehaviourActivities{}
	for _, detail := range studentBehaviour {
		studentBehaviourData = append(studentBehaviourData, models.StudentBehaviourActivities{
			AcademicID:                            detail.AcademicID,
			TermID:                                detail.TermID,
			StudentID:                             detail.StudentID,
			FirstBehaviour:                        detail.FirstBehaviour,
			FirstNeatness:                         detail.FirstNeatness,
			FirstCrafts:                           detail.FirstCrafts,
			FirstMonthExtracurricularFirst:        detail.FirstMonthExtracurricularFirst,
			FirstMonthExtracurricularScoreFirst:   detail.FirstMonthExtracurricularScoreFirst,
			FirstMonthExtracurricularSecond:       detail.FirstMonthExtracurricularSecond,
			FirstMonthExtracurricularScoreSecond:  detail.FirstMonthExtracurricularScoreSecond,
			SecondNeatness:                        detail.SecondNeatness,
			SecondCrafts:                          detail.SecondCrafts,
			SecondBehaviour:                       detail.SecondBehaviour,
			SecondMonthExtracurricularFirst:       detail.SecondMonthExtracurricularFirst,
			SecondMonthExtracurricularScoreFirst:  detail.SecondMonthExtracurricularScoreFirst,
			SecondMonthExtracurricularSecond:      detail.SecondMonthExtracurricularSecond,
			SecondMonthExtracurricularScoreSecond: detail.SecondMonthExtracurricularScoreSecond,
		})
	}

	err := s.studentBehaviourRepository.Create(studentBehaviourData)

	if err != nil {
		return err
	}

	return nil
}

func (s *studentBehaviourActivitiesUsecase) Update(studentBehaviour []dto.StudentBehaviourActivityRequest) error {
	for _, detail := range studentBehaviour {
		result, _ := s.studentBehaviourRepository.FindByStudentIDAndAcademicIDAndTermID(detail.StudentID, detail.AcademicID, detail.TermID)

		entry := models.StudentBehaviourActivities{
			ID:                                    detail.ID,
			AcademicID:                            detail.AcademicID,
			TermID:                                detail.TermID,
			StudentID:                             detail.StudentID,
			FirstBehaviour:                        detail.FirstBehaviour,
			FirstNeatness:                         detail.FirstNeatness,
			FirstCrafts:                           detail.FirstCrafts,
			FirstMonthExtracurricularFirst:        detail.FirstMonthExtracurricularFirst,
			FirstMonthExtracurricularScoreFirst:   detail.FirstMonthExtracurricularScoreFirst,
			FirstMonthExtracurricularSecond:       detail.FirstMonthExtracurricularSecond,
			FirstMonthExtracurricularScoreSecond:  detail.FirstMonthExtracurricularScoreSecond,
			SecondNeatness:                        detail.SecondNeatness,
			SecondCrafts:                          detail.SecondCrafts,
			SecondBehaviour:                       detail.SecondBehaviour,
			SecondMonthExtracurricularFirst:       detail.SecondMonthExtracurricularFirst,
			SecondMonthExtracurricularScoreFirst:  detail.SecondMonthExtracurricularScoreFirst,
			SecondMonthExtracurricularSecond:      detail.SecondMonthExtracurricularSecond,
			SecondMonthExtracurricularScoreSecond: detail.SecondMonthExtracurricularScoreSecond,
		}

		if result != nil {
			entry.CreatedAt = result.CreatedAt
		}

		var err error
		if detail.ID == 0 {
			err = s.studentBehaviourRepository.Create([]models.StudentBehaviourActivities{entry})
		} else {
			err = s.studentBehaviourRepository.Update(entry)
		}

		if err != nil {
			return err
		}
	}
	return nil
}
