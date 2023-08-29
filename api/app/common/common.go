package common

import (
	"time"

	"gorm.io/gorm"
)

type Utilisateur struct {
	ID           uint       `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom          string     `gorm:"column:nom;not null" json:"nom"`
	Address      string     `gorm:"column:address" json:"address"`
	Email        string     `gorm:"column:email;not null;unique" json:"email"`
	Phone        string     `gorm:"column:phone" json:"phone"`
	Password     string     `gorm:"column:password;not null" json:"password"`
	LastLogin    time.Time  `gorm:"column:last_login" json:"last_login"`
	RoleNom      string     `gorm:"column:role_nom;not null" json:"role_nom"`
	EntrepriseID uint       `gorm:"foreignKey:OwningEntrepriseId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"entreprise_id"`
	Entreprise   Entreprise `gorm:"foreignKey:EntrepriseID" json:"entreprise"`
	RoleID       uint       `gorm:"foreignKey:OwningRoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role_id"`
	Role         Role       `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	Audits       []Audit    `gorm:"many2many:membres" json:"audit"`
	CreatedBy    uint       `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

type Entreprise struct {
	ID        uint   `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom       string `gorm:"column:nom;not null;" json:"nom"`
	Email     string `gorm:"column:email;not null;unique;" json:"email"`
	Phone     string `gorm:"column:phone;not null;" json:"phone"`
	Address   string `gorm:"column:address;" json:"address"`
	ManagedBy uint   `gorm:"column:managed_by;" json:"managed_by"`
	CreatedBy uint   `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

type Role struct {
	ID           uint       `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom          string     `gorm:"column:nom;not null;" json:"nom"`
	EntrepriseID uint       `gorm:"foreignKey:OwningEntrepriseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"entreprise_id"`
	Entreprise   Entreprise `gorm:"foreignKey:EntrepriseID;references:ID" json:"entreprise"`
	CreatedBy    uint       `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}
type Audit struct {
	ID                    uint       `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom                   string     `gorm:"column:nom;not null" json:"nom"`
	Description           string     `gorm:"column:description;not null" json:"description"`
	Type                  string     `gorm:"column:type;not null" json:"type"`
	Statut                string     `gorm:"column:statut;not null" json:"statut"`
	Date_debut            string     `gorm:"column:date_debut;not null" json:"date_debut"`
	Date_fin              string     `gorm:"column:date_fin;not null" json:"date_fin"`
	Effacement            bool       `gorm:"column:effacement;not null" json:"effacement"`
	Effacement_jours      uint       `gorm:"column:effacement_jours;not null" json:"effacement_jours"`
	Observation           string     `gorm:"column:observation;not null" json:"observation"`
	UtilisateurPrincipale uint       `gorm:"column:utilisateur_principale;" json:"utilisateur_principale"`
	EntrepriseAuditie     uint       `gorm:"column:entreprise_auditie;" json:"entreprise_auditie"`
	EntrepriseAuditrice   uint       `gorm:"column:entreprise_auditrice;" json:"entreprise_auditrice"`
	EntrepriseID          uint       `gorm:"foreignKey:OwningEntrepriseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"entreprise_id"`
	Entreprise            Entreprise `gorm:"foreignKey:EntrepriseID;references:ID" json:"entreprise"`
	CreatedBy             uint       `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}
type Norme struct {
	ID          uint   `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom         string `gorm:"column:nom;not null" json:"nom"`
	Description string `gorm:"column:description;" json:"description"`
	gorm.Model
}

// check utilisateur exists
func CheckUtilisateurExists(db *gorm.DB, id uint) bool {

	// init vars
	utilisateur := &Utilisateur{}

	// check if row exists
	check := db.Where("id=?", id).First(&utilisateur)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
