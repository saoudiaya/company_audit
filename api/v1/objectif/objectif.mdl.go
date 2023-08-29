package objectif

import (
	"pfe/api/app/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Objectif struct {
	ID          uint         `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom         string       `gorm:"column:nom;not null" json:"nom"`
	Description string       `gorm:"column:description;" json:"description"`
	AuditID     uint         `gorm:"foreignKey:OwningAuditId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"audit_id"`
	Audit       common.Audit `gorm:"foreignKey:AuditID" json:"audit"`
	gorm.Model
}

// create new objectif
func NewObjectif(db *gorm.DB, objectif Objectif) (Objectif, error) {
	return objectif, db.Create(&objectif).Error
}

// get all objectif
func GetObjectifs(db *gorm.DB) (objectif []Objectif, err error) {
	return objectif, db.Preload(clause.Associations).Find(&objectif).Error
}

// check if objectif exists
func CheckObjectifExists(db *gorm.DB, id uint) bool {

	// init vars
	objectif := &Objectif{}

	// check if row exists
	check := db.Where("id=?", id).First(&objectif)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get objectif by id
func GetObjectifById(db *gorm.DB, id uint) (objectif Objectif, err error) {
	return objectif, db.Preload(clause.Associations).First(&objectif, "id=?", id).Error
}

// search objectif
func SearchObjectif(db *gorm.DB, objectif Objectif) (objectifs []Objectif, err error) {
	return objectifs, db.Where(&objectif).Find(&objectifs).Error
}

// update objectif
func UpdateObjectif(db *gorm.DB, objectif Objectif) error {
	return db.Where("id=?", objectif.ID).Updates(&objectif).Error
}

// delete objectif
func DeleteObjectif(db *gorm.DB, objectif_id uint) error {
	return db.Where("id=?", objectif_id).Delete(&Objectif{}).Error
}
