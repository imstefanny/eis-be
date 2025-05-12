package dto

type CreateDocumentsRequest struct {
	TypeID       uint           `json:"type_id" validate:"required"`
	ApplicantID  uint           `json:"applicant_id"`
	StudentID    uint           `json:"student_id"`
	UploadedFile string         `json:"uploaded_file" validate:"required"`
	Description  string         `json:"description"`
}
