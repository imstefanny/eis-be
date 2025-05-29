package dto

type CreateClassNotesRequest struct {
	AcademicID uint                             `json:"academic_id" validate:"required"`
	Date       string                           `json:"date" validate:"required"`
	Details    []CreateClassNotesDetailsRequest `json:"details" validate:"required"`
}

type CreateClassNotesDetailsRequest struct {
	SubjSchedID uint   `json:"subj_sched_id" validate:"required"`
	TeacherID   uint   `json:"teacher_id" validate:"required"`
	Materials   string `json:"materials" validate:"required"`
	Notes       string `json:"notes" validate:"required"`
}

type CreateBatchClassNotesRequest struct {
	Date string `json:"date" validate:"required"`
}
