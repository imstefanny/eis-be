package dto

type CreateStudentGradesRequest struct {
	AcademicID uint                               `json:"academic_id" validate:"required"`
	Details    []CreateStudentGradesDetailRequest `json:"details" validate:"required,dive"`
}

type CreateStudentGradesDetailRequest struct {
	SubjectID uint                              `json:"subject_id" validate:"required"`
	Students  []CreateStudentGradesEntryRequest `json:"students" validate:"required,dive"`
}

type CreateStudentGradesEntryRequest struct {
	StudentID   uint    `json:"student_id" validate:"required"`
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
	FinalGrade  float64 `json:"final_grade"`
	Remarks     string  `json:"remarks"`
}

type UpdateStudentGradesRequest struct {
	AcademicID uint                               `json:"academic_id" validate:"required"`
	Details    []UpdateStudentGradesDetailRequest `json:"details" validate:"required,dive"`
}
type UpdateStudentGradesDetailRequest struct {
	SubjectID uint                              `json:"subject_id" validate:"required"`
	Students  []UpdateStudentGradesEntryRequest `json:"students" validate:"required,dive"`
}
type UpdateStudentGradesEntryRequest struct {
	ID          uint    `json:"id"`
	StudentID   uint    `json:"student_id" validate:"required"`
	Quiz        float64 `json:"quiz"`
	FirstMonth  float64 `json:"first_month"`
	SecondMonth float64 `json:"second_month"`
	Finals      float64 `json:"finals"`
	Remarks     string  `json:"remarks"`
}

type StudentScoreResponse struct {
	SubjectName string  `json:"subject_name"`
	FirstMonth  float64 `json:"first_month"`
	SecondMonth float64 `json:"second_month"`
	Quiz        float64 `json:"quiz"`
	Finals      float64 `json:"finals"`
}

type StudentGradesReport struct {
	Class    string                          `json:"class"`
	Average  float64                         `json:"average"`
	Students []StudentGradesReportTopStudent `json:"students"`
}
type StudentGradesReportTopStudent struct {
	Rank    int     `json:"rank"`
	Student string  `json:"student"`
	NIS     string  `json:"nis"`
	Class   string  `json:"class"`
	Finals  float64 `json:"finals"`
}
