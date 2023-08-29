package tache

type TacheRsp struct {
	Nom           string `json:"nom"`
	Description   string `json:"description"`
	UtilisateurID uint   `json:"utilisateur_id"`
	AuditID       uint   `json:"audit_id"`
	Detefin       string `json:"datefin"`
}
