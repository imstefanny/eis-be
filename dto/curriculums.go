package dto

type CreateCurriculumsRequest struct {
	Name               string                            `json:"name" validate:"required"`
	LevelID            uint                              `json:"level_id" validate:"required"`
	Grade              string                            `json:"grade" validate:"required"`
	CurriculumSubjects []CreateCurriculumSubjectsRequest `json:"curriculum_subjects" validate:"required,dive,required"`
}

type CreateCurriculumSubjectsRequest struct {
	ID         uint   `json:"id"`
	SubjectID  uint   `json:"subject_id" validate:"required"`
	Competence string `json:"competence" validate:"required"`
}

type GetCurriculumsResponse struct {
	ID                 uint                            `json:"id"`
	DisplayName        string                          `json:"display_name"`
	Name               string                          `json:"name"`
	LevelID            uint                            `json:"level_id"`
	Level              string                          `json:"level"`
	Grade              string                          `json:"grade"`
	CurriculumSubjects []GetCurriculumSubjectsResponse `json:"curriculum_subjects"`
}
type GetCurriculumSubjectsResponse struct {
	ID         uint   `json:"id"`
	SubjectID  uint   `json:"subject_id"`
	Subject    string `json:"subject"`
	Competence string `json:"competence"`
}
