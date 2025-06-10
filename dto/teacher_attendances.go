package dto

type CreateTeacherAttsRequest struct {
	TeacherID         uint   `json:"teacher_id" validate:"required"`
	WorkingScheduleID uint   `json:"working_schedule_id"`
	Date              string `json:"date" validate:"required"`
	LogInTime         string `json:"log_in_time" validate:"required"`
	LogOutTime        string `json:"log_out_time"`
	Remark            string `json:"remark"`
	Note              string `json:"note"`
}

type CreateBatchTeacherAttsRequest struct {
	Entries []CreateTeacherAttsRequest `json:"entries" validate:"required,dive"`
}

type GetTeacherAttsRequest struct {
	ID         uint   `json:"id"`
	Teacher    string `json:"teacher"`
	WorkSched  string `json:"work_sched"`
	Date       string `json:"date"`
	LogInTime  string `json:"log_in_time"`
	LogOutTime string `json:"log_out_time"`
	Remark     string `json:"remark"`
	Note       string `json:"note"`
}

type GetTeacherAttsReport struct {
	Teacher         string `json:"teacher"`
	LateCount       int    `json:"late"`
	EarlyLeaveCount int    `json:"early_leave"`
	AbsenceCount    int    `json:"absence"`
	PresentCount    int    `json:"present"`
}
