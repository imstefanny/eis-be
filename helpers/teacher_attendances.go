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

		parseInTime, _ := time.Parse("15:04", schedule.WorkStart)
		parseOutTime, _ := time.Parse("15:04", schedule.WorkEnd)
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

func CountWorkdays(startDate, endDate string, workSched models.WorkScheds) int {
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)
	count := 0

	days := map[string]bool{}
	for _, schedule := range workSched.Details {
		days[schedule.Day] = true
	}
	for d := start; d.Before(end) || d.Equal(end); d = d.AddDate(0, 0, 1) {
		dayOfWeek := d.Weekday()
		if days[dayOfWeek.String()] {
			count++
		}
	}

	return count
}
