package utilisateur

type UtilisateurRsp struct {
	Nom          string `json:"nom"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	RoleNom      string `json:"role_nom"`
	RoleID       string `json:"role_id"`
	EntrepriseID uint   `json:"entreprise_id"`
}
