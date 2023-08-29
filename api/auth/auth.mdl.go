package app_auth

import "pfe/api/app/common"

type InsertUtilisateur struct {
	Nom          string            `gorm:"column:nom;not null" json:"nom"`
	Email        string            `gorm:"column:email;not null;unique" json:"email"`
	Password     string            `gorm:"column:password;not null" json:"password"`
	RoleNom      string            `gorm:"column:role_nom;not null" json:"role_nom"`
	RoleID       uint              `gorm:"column:role_id" json:"role_id"`
	Role         common.Role       `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	EntrepriseID uint              `gorm:"column:entreprise_id" json:"entreprise_id"`
	Entreprise   common.Entreprise `gorm:"foreignKey:EntrepriseID;references:ID" json:"entreprise"`
}

type UtilisateurLogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UtilisateurLogedIn struct {
	Token string `json:"token"`
}
