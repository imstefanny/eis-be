package dto

import "time"

type ClassNotesRepoRes struct {
	ID          uint      `json:"id"`
	NoteID      uint      `json:"note_id"`
	Day         string    `json:"day"`
	Date        time.Time `json:"date"`
	Class       string    `json:"class"`
	Subject     string    `json:"subject"`
	SubjSchedID uint      `json:"subj_sched_id"`
	AcademicID  uint      `json:"academic_id"`
	Teacher     string    `json:"teacher"`
	StartHour   string    `json:"start_hour"`
	EndHour     string    `json:"end_hour"`
	Materials   string    `json:"materials"`
	Notes       string    `json:"notes"`
}

type GetTeacherSchedsResponse struct {
	ID             uint                          `json:"id"`
	NoteID         uint                          `json:"note_id"`
	Day            string                        `json:"day"`
	Date           time.Time                     `json:"date"`
	Class          string                        `json:"class"`
	Subject        string                        `json:"subject"`
	SubjSchedID    uint                          `json:"subj_sched_id"`
	AcademicID     uint                          `json:"academic_id"`
	Teacher        string                        `json:"teacher"`
	StartHour      string                        `json:"start_hour"`
	EndHour        string                        `json:"end_hour"`
	Materials      string                        `json:"materials"`
	Notes          string                        `json:"notes"`
	AbsenceCount   []GetClassNoteAbsenceResponse `json:"absence_count"`
	AbsenceDetails []GetClassNoteAbsenceDetails  `json:"absence_details"`
}
