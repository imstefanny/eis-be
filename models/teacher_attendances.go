package models

import (
	"time"

	"gorm.io/gorm"
)

type TeacherAttendances struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	DisplayName       string         `json:"display_name" gorm:"size:255;unique"`
	TeacherID         uint           `json:"teacher_id" gorm:"index"`
	Teacher           Teachers       `json:"teacher" gorm:"foreignKey:TeacherID;references:ID"`
	WorkingScheduleID uint           `json:"working_schedule_id" gorm:"index"`
	WorkingSchedule   WorkScheds     `json:"working_schedule" gorm:"foreignKey:WorkingScheduleID;references:ID"`
	Date              time.Time      `json:"date"`
	LogInTime         time.Time      `json:"log_in_time"`
	LogOutTime        time.Time      `json:"log_out_time"`
	Remark            string         `json:"remark"`
	Note              string         `json:"note"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
