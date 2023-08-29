package listecontrole

import (
	//"pfe/api/v1/audit"
	//"pfe/api/v1/critere"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ListeControle struct {
	ID                  uint   `gorm:"primaryKey;autoIncrement;column:id;unique;" json:"id"`
	Nom                 string `gorm:"column:nom;not null" json:"nom"`
	Description         string `gorm:"column:description;" json:"description"`
	AuditID             uint   `gorm:"column:audit_id;primaryKey" json:"audit_id"`
	CritereID           uint   `gorm:"column:critere_id;primaryKey" json:"critere_id"`
	ApprobationAuditiee bool   `gorm:"column:approbationauditiee" json:"approbationauditiee"`
	gorm.Model
}

// create new liste controle
func NewListeControle(db *gorm.DB, listecontrole ListeControle) (ListeControle, error) {
	return listecontrole, db.Create(&listecontrole).Error
}

// get all liste controle
func GetListeControles(db *gorm.DB) (listecontrole []ListeControle, err error) {
	return listecontrole, db.Preload(clause.Associations).Find(&listecontrole).Error
}

// check if liste controle exists
func CheckListeControleExists(db *gorm.DB, id uint) bool {

	// init vars
	listecontrole := &ListeControle{}

	// check if row exists
	check := db.Where("id=?", id).First(&listecontrole)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get liste controle by id
func GetListeControleById(db *gorm.DB, id uint) (listecontrole ListeControle, err error) {
	return listecontrole, db.Preload(clause.Associations).First(&listecontrole, "id=?", id).Error
}

// search liste controle
func SearchListeControle(db *gorm.DB, listecontrole ListeControle) (listecontroles []ListeControle, err error) {
	return listecontroles, db.Where(&listecontrole).Find(&listecontroles).Error
}

// update liste controle
func UpdateListeControle(db *gorm.DB, listecontrole ListeControle) error {
	return db.Where("id=?", listecontrole.ID).Updates(&listecontrole).Error
}

// delete liste controle
func DeleteListeControle(db *gorm.DB, listecontrole_id uint) error {
	return db.Where("id=?", listecontrole_id).Delete(&ListeControle{}).Error
}
