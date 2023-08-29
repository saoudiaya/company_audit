package membre

type MembreRsp struct {
	UtilisateurID uint   `json:"utilisateur_id"`
	AuditID       uint   `json:"audit_id"`
	Type          string `json:"type"`
	Poste         string `json:"poste"`
}
