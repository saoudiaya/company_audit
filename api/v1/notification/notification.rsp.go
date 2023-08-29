package notification

type NotificationRsp struct {
	Nom           string `json:"nom"`
	Description   string `json:"description"`
	UtilisateurID uint   `json:"utilisateur_id"`
}
