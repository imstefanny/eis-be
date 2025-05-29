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

type UpdateBatchSubjSchedsRequest struct {
	AcademicID uint                      `json:"academic_id" validate:"required"`
	Entries    []UpdateSubjSchedsRequest `json:"entries" validate:"required,dive"`
}

type UpdateSubjSchedsRequest struct {
	ID          uint   `json:"id"`
	DisplayName string `json:"display_name"`
	SubjectID   uint   `json:"subject_id" validate:"required"`
	TeacherID   uint   `json:"teacher_id"`
	Day         string `json:"day" validate:"required"`
	StartHour   string `json:"start_hour" validate:"required"`
	EndHour     string `json:"end_hour" validate:"required"`
}

type UpdateSubjSchedsResponse struct {
	AcademicID uint                         `json:"academic_id"`
	Entries    []GetSubjectScheduleResponse `json:"entries"`
}

type GetSubjectScheduleResponse struct {
	Day     string                            `json:"day"`
	Entries []GetSubjectScheduleEntryResponse `json:"entries"`
}
type GetSubjectScheduleEntryResponse struct {
	ID        uint   `json:"id"`
	Subject   string `json:"subject"`
	Teacher   string `json:"teacher"`
	StartHour string `json:"start_hour"`
	EndHour   string `json:"end_hour"`
}
