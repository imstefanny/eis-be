package dto

type CreateSubjSchedsRequest struct {
	AcademicID uint                     `json:"academic_id" validate:"required"`
	Schedules  []CreateSubjSchedRequest `json:"schedules" validate:"required"`
}

type CreateSubjSchedRequest struct {
	Day     string                          `json:"day" validate:"required"`
	Entries []CreateSubjSchedDetailsRequest `json:"entries" validate:"required"`
}

type CreateSubjSchedDetailsRequest struct {
	SubjectID uint   `json:"subject_id" validate:"required"`
	TeacherID uint   `json:"teacher_id"`
	StartHour string `json:"start_hour" validate:"required"`
	EndHour   string `json:"end_hour" validate:"required"`
}

type UpdateSubjSchedsRequest struct {
	DisplayName string `json:"display_name" validate:"required"`
	AcademicID  uint   `json:"academic_id" validate:"required"`
	SubjectID   uint   `json:"subject_id" validate:"required"`
	TeacherID   uint   `json:"teacher_id"`
	Day         string `json:"day" validate:"required"`
	StartHour   string `json:"start_hour" validate:"required"`
	EndHour     string `json:"end_hour" validate:"required"`
}
