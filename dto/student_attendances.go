package dto

// type CreateStudentAttsRequest struct {
// 	AcademicID uint   `json:"academic_id" validate:"required"`
// 	StudentID  uint   `json:"student_id" validate:"required"`
// 	Date       string `json:"date" validate:"required"`
// 	Status     string `json:"status" validate:"required"`
// 	Remarks    string `json:"remarks" validate:"required"`
// }

type CreateBatchStudentAttsRequest struct {
	Date string `json:"date" validate:"required"`
}

type UpdateStudentAttsRequest struct {
	AcademicID uint                            `json:"academic_id" validate:"required"`
	Date       string                          `json:"date" validate:"required"`
	Students   []UpdateStudentAttsEntryRequest `json:"students" validate:"required,dive"`
}
type UpdateStudentAttsEntryRequest struct {
	StudentID uint   `json:"student_id" validate:"required"`
	Status    string `json:"status" validate:"required"`
	Remarks   string `json:"remarks" validate:"required"`
}
