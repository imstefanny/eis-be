package dto

type CreateStudentGradesRequest struct {
	AcademicID uint                               `json:"academic_id"`
	Details    []CreateStudentGradesDetailRequest `json:"details" validate:"required,dive"`
}

type CreateStudentGradesDetailRequest struct {
	SubjectID uint                              `json:"subject_id"`
	Students  []CreateStudentGradesEntryRequest `json:"students" validate:"required,dive"`
}

type CreateStudentGradesEntryRequest struct {
	StudentID   uint    `json:"student_id"`
	Quiz        float64 `json:"quiz"`
	FirstMonth  float64 `json:"first_month"`
	SecondMonth float64 `json:"second_month"`
	Finals      float64 `json:"finals"`
	Remarks     string  `json:"remarks"`
}

type GetStudentGradesResponse struct {
	AcademicID uint                             `json:"academic_id"`
	Academic   string                           `json:"academic"`
	Details    []GetStudentGradesDetailResponse `json:"details"`
}
type GetStudentGradesDetailResponse struct {
	SubjectID uint                            `json:"subject_id"`
	Subject   string                          `json:"subject"`
	Students  []GetStudentGradesEntryResponse `json:"students"`
}
type GetStudentGradesEntryResponse struct {
	ID          uint    `json:"id"`
	StudentID   uint    `json:"student_id"`
	StudentName string  `json:"student_name"`
	DisplayName string  `json:"display_name"`
	Quiz        float64 `json:"quiz"`
	FirstMonth  float64 `json:"first_month"`
	SecondMonth float64 `json:"second_month"`
	Finals      float64 `json:"finals"`
	Remarks     string  `json:"remarks"`
}
