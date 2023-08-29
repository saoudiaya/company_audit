package norme

import (
	"pfe/api/app/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// create new norme
func NewNorme(db *gorm.DB, norme common.Norme) (common.Norme, error) {
	return norme, db.Create(&norme).Error
}

// get all norme
func GetNormes(db *gorm.DB) (norme []common.Norme, err error) {
	return norme, db.Preload(clause.Associations).Find(&norme).Error
}

// check if norme exists
func CheckNormeExists(db *gorm.DB, id uint) bool {

	// init vars
	norme := &common.Norme{}

	// check if row exists
	check := db.Where("id=?", id).First(&norme)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get norme by id
func GetNormeById(db *gorm.DB, id uint) (norme common.Norme, err error) {
	return norme, db.First(&norme, "id=?", id).Error
}

// search norme
func SearchNorme(db *gorm.DB, norme common.Norme) (normes []common.Norme, err error) {
	return normes, db.Where(&norme).Find(&normes).Error
}

// update norme
func UpdateNorme(db *gorm.DB, norme common.Norme) error {
	return db.Where("id=?", norme.ID).Updates(&norme).Error
}

// delete norme
func DeleteNorme(db *gorm.DB, norme_id uint) error {
	return db.Where("id=?", norme_id).Delete(&common.Norme{}).Error
}
