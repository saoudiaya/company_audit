package exigence

type ExigenceRsp struct {
	Nom         string `json:"nom"`
	Description string `json:"description"`
	Niveau      string `json:"niveau"`
	AuditID     uint   `json:"audit_id"`
}
