package entreprise

import (
	"pfe/api/app/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// create new entreprise
func NewEntreprise(db *gorm.DB, entreprise common.Entreprise) (common.Entreprise, error) {
	return entreprise, db.Create(&entreprise).Error
}

// get all entreprises
func GetEntreprises(db *gorm.DB) (entreprise []common.Entreprise, err error) {
	return entreprise, db.Preload(clause.Associations).Find(&entreprise).Error
}

// check entreprise exists
func CheckEntrepriseExists(db *gorm.DB, id uint) bool {

	// init vars
	entreprise := &common.Entreprise{}

	// check if row exists
	check := db.Where("id=?", id).First(&entreprise)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get entreprise by id
func GetEntrepriseById(db *gorm.DB, id uint) (entreprise common.Entreprise, err error) {
	return entreprise, db.Preload(clause.Associations).First(&entreprise, "id=?", id).Error
}

// search entreprises
func SearchEntreprises(db *gorm.DB, entreprise common.Entreprise) (entreprises []common.Entreprise, err error) {
	return entreprises, db.Where(&entreprise).Find(&entreprises).Error
}

// update entreprise
func UpdateEntreprise(db *gorm.DB, entreprise common.Entreprise) error {
	return db.Where("id=?", entreprise.ID).Updates(&entreprise).Error
}

// delete entreprise
func DeleteEntreprise(db *gorm.DB, entreprise_id uint) error {
	return db.Where("id=?", entreprise_id).Delete(&common.Entreprise{}).Error
}
