package audit

type AuditRsp struct {
	ID                    uint   `json:"ID"`
	Nom                   string ` json:"nom"`
	Description           string ` json:"description"`
	Type                  string ` json:"type"`
	Statut                string ` json:"statut"`
	Date_debut            string ` json:"date_debut"`
	Date_fin              string `json:"date_fin"`
	Effacement            bool   `json:"effacement"`
	Effacement_jours      uint   ` json:"effacement_jours"`
	Observation           string ` json:"observation"`
	UtilisateurPrincipale uint   ` json:"utilisateur_principale"`
	EntrepriseAuditie     uint   ` json:"entreprise_auditie"`
	EntrepriseAuditrice   uint   `json:"entreprise_auditrice"`
	EntrepriseID          uint   ` json:"entreprise_id"`
}
