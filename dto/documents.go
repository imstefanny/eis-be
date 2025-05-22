package dto

type CreateDocumentsRequest struct {
	Name         string `json:"name" validate:"required"`
	TypeID       uint   `json:"type_id" validate:"required"`
	ApplicantID  uint   `json:"applicant_id"`
	StudentID    uint   `json:"student_id"`
	UploadedFile string `json:"uploaded_file" validate:"required"`
	Description  string `json:"description"`
}
