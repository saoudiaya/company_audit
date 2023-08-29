package reunion

import (
	"pfe/api/app/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Reunion struct {
	ID            uint               `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Titre         string             `gorm:"column:titre;not null" json:"titre"`
	UtilisateurID uint               `gorm:"foreignKey:OwningUtilisateurId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"utilisateur_id"`
	Utilisateur   common.Utilisateur `gorm:"foreignKey:UtilisateurID;references:ID" json:"utilisateur"`
	AuditID       uint               `gorm:"foreignKey:OwningAuditId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"audit_id"`
	Audit         common.Audit       `gorm:"foreignKey:AuditID;references:ID" json:"audit"`
	Datedebut     string             `gorm:"column:date_debut;not null" json:"date_debut"`
	Datefin       string             `gorm:"column:date_fin;not null" json:"date_fin"`
	gorm.Model
}

// create new reunion
func NewReunion(db *gorm.DB, reunion Reunion) (Reunion, error) {
	return reunion, db.Create(&reunion).Error
}

// get all reunion
func GetReunions(db *gorm.DB) (reunion []Reunion, err error) {
	return reunion, db.Preload(clause.Associations).Find(&reunion).Error
}

// check if reunion exists
func CheckReunionExists(db *gorm.DB, id uint) bool {

	// init vars
	reunion := &Reunion{}

	// check if row exists
	check := db.Where("id=?", id).First(&reunion)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get reunion by id
func GetReunionById(db *gorm.DB, id uint) (reunion Reunion, err error) {
	return reunion, db.Preload(clause.Associations).First(&reunion, "id=?", id).Error
}

// search reunion
func SearchReunion(db *gorm.DB, reunion Reunion) (reunions []Reunion, err error) {
	return reunions, db.Where(&reunion).Find(&reunions).Error
}

// update reunion
func UpdateReunion(db *gorm.DB, reunion Reunion) error {
	return db.Where("id=?", reunion.ID).Updates(&reunion).Error
}

// delete reunion
func DeleteReunion(db *gorm.DB, reunion_id uint) error {
	return db.Where("id=?", reunion_id).Delete(&Reunion{}).Error
}
