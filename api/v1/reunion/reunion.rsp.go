package reunion

type ReunionRsp struct {
	Titre         string `json:"titre"`
	UtilisateurID uint   `json:"utilisateur_id"`
	AuditID       uint   `json:"audit_id"`
	Datedebut     string `json:"date_debut"`
	Datefin       string `json:"date_fin"`
}
