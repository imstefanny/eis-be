package helpers

import (
	"eis-be/models"
	"strings"
	"time"
)

func TeacherAttsRemark(att models.TeacherAttendances, workSched models.WorkScheds) string {
	dayOfWeek := att.Date.Weekday()
	remark := []string{}
	for _, schedule := range workSched.Details {
		date := att.Date
		loc, _ := time.LoadLocation("Asia/Jakarta")
		parseInTime, _ := time.Parse("15:04:05", schedule.WorkStart)
		parseOutTime, _ := time.Parse("15:04:05", schedule.WorkEnd)
		work_start := time.Date(date.Year(), date.Month(), date.Day(), parseInTime.Hour(), parseInTime.Minute(), parseInTime.Second(), 0, loc)
		work_end := time.Date(date.Year(), date.Month(), date.Day(), parseOutTime.Hour(), parseOutTime.Minute(), parseOutTime.Second(), 0, loc)
		if schedule.Day == dayOfWeek.String() {
			if att.LogInTime.After(work_start) {
				remark = append(remark, "Terlambat")
			}
			if att.LogOutTime.Before(work_end) {
				remark = append(remark, "Pulang Cepat")
			}
		}
	}
	return strings.Join(remark, ", ")
}
