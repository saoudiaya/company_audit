package objectif

type ObjectifRsp struct {
	Nom         string `json:"nom"`
	Description string `json:"description"`
	AuditID     uint   `json:"audit_id"`
}
