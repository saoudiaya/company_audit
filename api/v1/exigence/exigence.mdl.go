package exigence

import (
	"pfe/api/app/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Exigence struct {
	ID          uint         `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom         string       `gorm:"column:nom;not null" json:"nom"`
	Description string       `gorm:"column:description;" json:"description"`
	Niveau      string       `gorm:"column:niveau;" json:"niveau"`
	AuditID     uint         `gorm:"foreignKey:OwningAuditId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"audit_id"`
	Audit       common.Audit `gorm:"foreignKey:AuditID;references:ID" json:"audit"`
	gorm.Model
}

// create new exigence
func NewExigence(db *gorm.DB, exigence Exigence) (Exigence, error) {
	return exigence, db.Create(&exigence).Error
}

// get all exigence
func GetExigences(db *gorm.DB) (exigence []Exigence, err error) {
	return exigence, db.Preload(clause.Associations).Find(&exigence).Error
}

// check if exigence exists
func CheckExigenceExists(db *gorm.DB, id uint) bool {

	// init vars
	exigence := &Exigence{}

	// check if row exists
	check := db.Where("id=?", id).First(&exigence)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get exigence by id
func GetExigenceById(db *gorm.DB, id uint) (exigence Exigence, err error) {
	return exigence, db.Preload(clause.Associations).First(&exigence, "id=?", id).Error
}

// search exigence
func SearchExigence(db *gorm.DB, exigence Exigence) (exigences []Exigence, err error) {
	return exigences, db.Where(&exigence).Find(&exigences).Error
}

// update exigence
func UpdateExigence(db *gorm.DB, exigence Exigence) error {
	return db.Where("id=?", exigence.ID).Updates(&exigence).Error
}

// delete exigence
func DeleteExigence(db *gorm.DB, exigence_id uint) error {
	return db.Where("id=?", exigence_id).Delete(&Exigence{}).Error
}
