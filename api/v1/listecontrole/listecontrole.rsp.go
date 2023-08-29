package listecontrole

type ListeControleRsp struct {
	Nom                 string `json:"nom"`
	Description         string `json:"description"`
	AuditID             uint   `json:"audit_id"`
	CritereID           uint   `json:"critere_id"`
	ApprobationAuditiee bool   `json:"approbationauditiee"`
}
