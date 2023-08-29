package reponse

type ReponseRsp struct {
	Nom           string `json:"nom"`
	Description   string `json:"description"`
	CommentaireID uint   `json:"commentaire_id"`
	ManagedBy     uint   `json:"managed_by"`
}
