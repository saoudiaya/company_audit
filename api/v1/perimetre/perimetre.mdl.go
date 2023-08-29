package perimetre

import (
	"pfe/api/app/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Perimetre struct {
	ID          uint         `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom         string       `gorm:"column:nom;not null" json:"nom"`
	Description string       `gorm:"column:description;" json:"description"`
	AuditID     uint         `gorm:"foreignKey:OwningAuditId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"audit_id"`
	Audit       common.Audit `gorm:"foreignKey:AuditID" json:"audit"`
	gorm.Model
}

// create new perimetre
func NewPerimetre(db *gorm.DB, perimetre Perimetre) (Perimetre, error) {
	return perimetre, db.Create(&perimetre).Error
}

// get all perimetre
func GetPerimetres(db *gorm.DB) (perimetre []Perimetre, err error) {
	return perimetre, db.Preload(clause.Associations).Find(&perimetre).Error
}

// check if perimetre exists
func CheckPerimetreExists(db *gorm.DB, id uint) bool {

	// init vars
	perimetre := &Perimetre{}

	// check if row exists
	check := db.Where("id=?", id).First(&perimetre)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get perimetre by id
func GetPerimetreById(db *gorm.DB, id uint) (perimetre Perimetre, err error) {
	return perimetre, db.Preload(clause.Associations).First(&perimetre, "id=?", id).Error
}

// search perimetre
func SearchPerimetre(db *gorm.DB, perimetre Perimetre) (perimetres []Perimetre, err error) {
	return perimetres, db.Where(&perimetre).Find(&perimetres).Error
}

// update perimetre
func UpdatePerimetre(db *gorm.DB, perimetre Perimetre) error {
	return db.Where("id=?", perimetre.ID).Updates(&perimetre).Error
}

// delete perimetre
func DeletePerimetre(db *gorm.DB, perimetre_id uint) error {
	return db.Where("id=?", perimetre_id).Delete(&Perimetre{}).Error
}
