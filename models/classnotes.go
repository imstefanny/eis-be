package models

import (
	"time"

	"gorm.io/gorm"
)

type ClassNotes struct {
	ID          uint                `json:"id" gorm:"primaryKey"`
	DisplayName string              `json:"display_name" gorm:"default:null"`
	AcademicID  uint                `json:"academic_id" gorm:"default:null"`
	Academic    Academics           `json:"academic" gorm:"foreignKey:AcademicID"`
	TermID      uint                `json:"term_id" gorm:"default:null"`
	Term        Terms               `json:"term" gorm:"foreignKey:TermID"`
	Date        time.Time           `json:"date"`
	Details     []ClassNotesDetails `json:"details" gorm:"foreignKey:NoteID"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	DeletedAt   gorm.DeletedAt      `json:"deleted_at" gorm:"index"`
}

type ClassNotesDetails struct {
	ID          uint             `json:"id" gorm:"primaryKey"`
	NoteID      uint             `json:"note_id" gorm:"default:null"`
	Note        ClassNotes       `json:"note" gorm:"foreignKey:NoteID"`
	SubjSchedID uint             `json:"subj_sched_id" gorm:"default:null"`
	SubjSched   SubjectSchedules `json:"subj_sched" gorm:"foreignKey:SubjSchedID"`
	TeacherID   uint             `json:"teacher_id" gorm:"default:null"`
	Teacher     Teachers         `json:"teacher" gorm:"foreignKey:TeacherID"`
	Materials   string           `json:"materials" gorm:"default:null"`
	Notes       string           `json:"notes" gorm:"default:null"`
}
