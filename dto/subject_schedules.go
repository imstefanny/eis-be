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
	SubjectID uint   `json:"subject_id"`
	Subject   string `json:"subject"`
	Teacher   string `json:"teacher"`
	TeacherID uint   `json:"teacher_id"`
	StartHour string `json:"start_hour"`
	EndHour   string `json:"end_hour"`
}

// below is the response for the class notes list for the teacher view
type TeacherStudentGetSubjScheds struct {
	Day     string                               `json:"day"`
	Details []TeacherStudentGetSubjSchedsDetails `json:"details"`
}
type TeacherStudentGetSubjSchedsDetails struct {
	SubjSchedID uint   `json:"subj_sched_id"`
	ClassID     uint   `json:"class_id"`
	Class       string `json:"class"`
	SubjectID   uint   `json:"subject_id"`
	Subject     string `json:"subject"`
	TeacherID   uint   `json:"teacher_id"`
	StartHour   string `json:"start_hour"`
	EndHour     string `json:"end_hour"`
}
