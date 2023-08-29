package critere

import (
	"pfe/api/app/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Critere struct {
	ID                    uint           `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom                   string         `gorm:"column:nom;not null" json:"nom"`
	Description           string         `gorm:"column:description;" json:"description"`
	Entretien             string         `gorm:"column:entretien;" json:"entretien"`
	VerificationTechnique string         `gorm:"column:verification_technique;" json:"verification_technique"`
	Analyse               string         `gorm:"column:analyse;" json:"analyse"`
	NormeID               uint           `gorm:"foreignKey:OwningNormeId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"norme_id"`
	Norme                 common.Norme   `gorm:"foreignKey:NormeID" json:"norme"`
	Audit                 []common.Audit `gorm:"many2many:liste_controles" json:"audit"`
	gorm.Model
}

// create new critere
func NewCritere(db *gorm.DB, critere Critere) (Critere, error) {
	return critere, db.Create(&critere).Error
}

// get all critere
func GetCriteres(db *gorm.DB) (critere []Critere, err error) {
	return critere, db.Preload(clause.Associations).Find(&critere).Error
}

// check if critere exists
func CheckCritereExists(db *gorm.DB, id uint) bool {

	// init vars
	critere := &Critere{}

	// check if row exists
	check := db.Where("id=?", id).First(&critere)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get critere by id
func GetCritereById(db *gorm.DB, id uint) (critere Critere, err error) {
	return critere, db.Preload(clause.Associations).First(&critere, "id=?", id).Error
}

// search critere
func SearchCritere(db *gorm.DB, critere Critere) (criteres []Critere, err error) {
	return criteres, db.Where(&critere).Find(&criteres).Error
}

// update critere
func UpdateCritere(db *gorm.DB, critere Critere) error {
	return db.Where("id=?", critere.ID).Updates(&critere).Error
}

// delete critere
func DeleteCritere(db *gorm.DB, critere_id uint) error {
	return db.Where("id=?", critere_id).Delete(&Critere{}).Error
}
