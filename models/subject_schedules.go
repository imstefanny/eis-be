package models

import (
	"time"

	"gorm.io/gorm"
)

type SubjectSchedules struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	DisplayName string    `json:"display_name"`
	AcademicID  uint      `json:"academic_id"`
	Academic    Academics `json:"academic" gorm:"foreignKey:AcademicID"`
	SubjectID   uint      `json:"subject_id"`
	Subject     Subjects  `json:"subject" gorm:"foreignKey:SubjectID"`
	TeacherID   uint      `json:"teacher_id" gorm:"default:null"`
	Teacher     Users     `json:"teacher" gorm:"foreignKey:TeacherID"`
	Day         string    `json:"day"`
	StartHour   string    `json:"start_hour"`
	EndHour     string    `json:"end_hour"`
	// Histories []LevelHistories `json:"histories" gorm:"foreignKey:LevelID;references:ID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// type LevelHistories struct {
// 	ID            uint      `json:"id" gorm:"primaryKey"`
// 	LevelID       uint      `json:"level_id"`
// 	OpCertNum     string    `json:"op_cert_num"`
// 	Accreditation string    `json:"accreditation"`
// 	Curriculum    string    `json:"curriculum"`
// 	Email         string    `json:"email"`
// 	Phone         string    `json:"phone"`
// 	PrincipleID   *uint     `json:"principle_id"`
// 	Principle     Users     `json:"principle" gorm:"foreignKey:PrincipleID;references:ID"`
// 	OperatorID    *uint     `json:"operator_id"`
// 	Operator      Users     `json:"operator" gorm:"foreignKey:OperatorID;references:ID"`
// 	State         bool      `json:"state"`
// 	CreatedAt     time.Time `json:"created_at"`
// 	UpdatedAt     time.Time `json:"updated_at"`
// }
