package app_auth

type InsertUtilisateurRsp struct {
	Nom          string `json:"nom"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RoleNom      string `json:"role_nom"`
	RoleID       uint   `json:"role_id"`
	EntrepriseID uint   `json:"entreprise_id"`
}
