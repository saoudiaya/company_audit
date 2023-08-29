package commentaire

type CommentaireRsp struct {
	Nom           string `json:"nom"`
	Description   string `json:"description"`
	UtilisateurID uint   `json:"utilisateur_id"`
	ManagedBy     uint   `json:"managed_by"`
}
