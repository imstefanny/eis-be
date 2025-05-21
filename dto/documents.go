package dto

import "eis-be/models"

type CreateDocumentsRequest struct {
	TypeID       uint   `json:"type_id" validate:"required"`
	ApplicantID  uint   `json:"applicant_id"`
	StudentID    uint   `json:"student_id"`
	UploadedFile string `json:"uploaded_file" validate:"required"`
	Description  string `json:"description"`
}

type DocumentsResponse struct {
	models.Documents
	TypeName string `json:"type_name"`
}

func (DocumentsResponse) TableName() string {
	return "documents"
}
