package dto

type CreateLevelsRequest struct {
	Name string `json:"name"`
}

type CreateLevelHistoriesRequest struct {
	LevelID       uint   `json:"level_id"`
	OpCertNum     string `json:"op_cert_num"`
	Accreditation string `json:"accreditation"`
	Curriculum    string `json:"curriculum"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	PrincipleID   uint   `json:"principle_id"`
	OperatorID    uint   `json:"operator_id"`
	State         bool   `json:"state"`
}
