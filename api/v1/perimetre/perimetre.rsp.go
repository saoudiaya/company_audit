package perimetre

type PerimetreRsp struct {
	Nom         string `json:"nom"`
	Description string `json:"description"`
	AuditID     uint   `json:"audit_id"`
}
