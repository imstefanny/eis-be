package dto

type UpdateAcademicStudentsRequest struct {
	ID              int    `json:"id" validate:"required"`
	AcademicID      int    `json:"academic_id" validate:"required"`
	StudentID       int    `json:"student_id" validate:"required"`
	IsFirstTerm     bool   `json:"is_first_term"`
	FirstTermNotes  string `json:"first_term_notes"`
	SecondTermNotes string `json:"second_term_notes"`
}
