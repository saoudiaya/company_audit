package revuedocument

type RevueDocumentRsp struct {
	Titre         string `json:"titre"`
	AuditID       uint   `json:"audit_id"`
	UtilisateurID uint   `json:"utilisateur_id"`
	Date          string `json:"date"`
}
