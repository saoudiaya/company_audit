package rapport

import (
	"pfe/api/v1/observation"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Rapport struct {
	ID                 uint                    `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom                string                  `gorm:"column:nom;not null" json:"nom"`
	Nombreofconformite uint                    `gorm:"column:nombreofconformite;" json:"nombreofconformite"`
	Process            string                  `gorm:"column:process;" json:"process"`
	Nombreofarticle    uint                    `gorm:"column:nombreofarticle;" json:"nombreofarticle"`
	Description        string                  `gorm:"column:description;" json:"description"`
	ObservationID      uint                    `gorm:"foreignKey:OwningObservationId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"observation_id"`
	Observation        observation.Observation `gorm:"foreignKey:ObservationID" json:"observation"`
	gorm.Model
}

// create new rapport
func NewRapport(db *gorm.DB, rapport Rapport) (Rapport, error) {
	return rapport, db.Create(&rapport).Error
}

// get all rapport
func GetRapports(db *gorm.DB) (rapport []Rapport, err error) {
	return rapport, db.Preload(clause.Associations).Find(&rapport).Error
}

// check if rapport exists
func CheckRapportExists(db *gorm.DB, id uint) bool {

	// init vars
	rapport := &Rapport{}

	// check if row exists
	check := db.Where("id=?", id).First(&rapport)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get rapport by id
func GetRapportById(db *gorm.DB, id uint) (rapport Rapport, err error) {
	return rapport, db.Preload(clause.Associations).First(&rapport, "id=?", id).Error
}

// search rapport
func SearchRapport(db *gorm.DB, rapport Rapport) (rapports []Rapport, err error) {
	return rapports, db.Where(&rapport).Find(&rapports).Error
}

// update rapport
func UpdateRapport(db *gorm.DB, rapport Rapport) error {
	return db.Where("id=?", rapport.ID).Updates(&rapport).Error
}

// delete rapport
func DeleteRapport(db *gorm.DB, rapport_id uint) error {
	return db.Where("id=?", rapport_id).Delete(&Rapport{}).Error
}
