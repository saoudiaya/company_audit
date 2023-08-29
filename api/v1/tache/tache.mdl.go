package tache

import (
	"pfe/api/app/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Tache struct {
	ID            uint               `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom           string             `gorm:"column:nom;not null" json:"nom"`
	Description   string             `gorm:"column:description;" json:"description"`
	UtilisateurID uint               `gorm:"foreignKey:OwningUtilisateurId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"utilisateur_id"`
	Utilisateur   common.Utilisateur `gorm:"foreignKey:UtilisateurID;references:ID" json:"utilisateur"`
	AuditID       uint               `gorm:"foreignKey:OwningAuditId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"audit_id"`
	Audit         common.Audit       `gorm:"foreignKey:AuditID;references:ID" json:"audit"`
	Datefin       string             `gorm:"column:datefin;" json:"datefin"`
	gorm.Model
}

// create new tache
func NewTache(db *gorm.DB, tache Tache) (Tache, error) {
	return tache, db.Create(&tache).Error
}

// get all tache
func GetTaches(db *gorm.DB) (tache []Tache, err error) {
	return tache, db.Preload(clause.Associations).Find(&tache).Error
}

// check if tache exists
func CheckTacheExists(db *gorm.DB, id uint) bool {

	// init vars
	tache := &Tache{}

	// check if row exists
	check := db.Where("id=?", id).First(&tache)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get tache by id
func GetTacheById(db *gorm.DB, id uint) (tache Tache, err error) {
	return tache, db.Preload(clause.Associations).First(&tache, "id=?", id).Error
}

// search tache
func SearchTache(db *gorm.DB, tache Tache) (taches []Tache, err error) {
	return taches, db.Where(&tache).Find(&taches).Error
}

// update tache
func UpdateTache(db *gorm.DB, tache Tache) error {
	return db.Where("id=?", tache.ID).Updates(&tache).Error
}

// delete tache
func DeleteTache(db *gorm.DB, tache_id uint) error {
	return db.Where("id=?", tache_id).Delete(&Tache{}).Error
}
