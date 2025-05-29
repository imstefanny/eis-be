package dto

type CreateTeacherAttsRequest struct {
	TeacherID         uint   `json:"teacher_id" validate:"required"`
	WorkingScheduleID uint   `json:"working_schedule_id" validate:"required"`
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
