package dto

type CreateStudentGradesRequest struct {
	AcademicID uint                               `json:"academic_id"`
	Details    []CreateStudentGradesDetailRequest `json:"details" validate:"required,dive"`
}

type CreateStudentGradesDetailRequest struct {
	SubjectID uint                              `json:"subject_id"`
	Entries   []CreateStudentGradesEntryRequest `json:"entries" validate:"required,dive"`
}

type CreateStudentGradesEntryRequest struct {
	StudentID   uint    `json:"student_id"`
	Quiz        float64 `json:"quiz"`
	FirstMonth  float64 `json:"first_month"`
	SecondMonth float64 `json:"second_month"`
	Midterm     float64 `json:"midterm"`
	Finals      float64 `json:"finals"`
	Remarks     string  `json:"remarks"`
}
