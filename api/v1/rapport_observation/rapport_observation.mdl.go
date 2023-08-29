package rapportobservation

import (
	"pfe/api/v1/observation"
	"pfe/api/v1/rapport"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RapportObservation struct {
	ID            uint                    `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom           string                  `gorm:"column:nom;not null" json:"nom"`
	Description   string                  `gorm:"column:description;" json:"description"`
	ObservationID uint                    `gorm:"foreignKey:OwningObservationId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"observation_id"`
	Observation   observation.Observation `gorm:"foreignKey:ObservationID" json:"observation"`
	RapportID     uint                    `gorm:"foreignKey:OwningRapportId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"rapport_id"`
	Rapport       rapport.Rapport         `gorm:"foreignKey:RapportID" json:"rapport"`
	gorm.Model
}

// create new rapport observation
func NewRapportObservation(db *gorm.DB, rapportobservation RapportObservation) (RapportObservation, error) {
	return rapportobservation, db.Create(&rapportobservation).Error
}

// get all rapport observation
func GetRapportsObservations(db *gorm.DB) (rapportobservation []RapportObservation, err error) {
	return rapportobservation, db.Preload(clause.Associations).Find(&rapportobservation).Error
}

// check if rapport observation exists
func CheckRapportObservationExists(db *gorm.DB, id uint) bool {

	// init vars
	rapportobservation := &RapportObservation{}

	// check if row exists
	check := db.Where("id=?", id).First(&rapportobservation)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get rapport observation by id
func GetRapportObservationById(db *gorm.DB, id uint) (rapportobservation RapportObservation, err error) {
	return rapportobservation, db.Preload(clause.Associations).First(&rapportobservation, "id=?", id).Error
}

// search rapport observation
func SearchRapportObservation(db *gorm.DB, rapportobservation RapportObservation) (rapportsobservations []RapportObservation, err error) {
	return rapportsobservations, db.Where(&rapportobservation).Find(&rapportsobservations).Error
}

// update rapport observation
func UpdateRapportObservation(db *gorm.DB, rapportobservation RapportObservation) error {
	return db.Where("id=?", rapportobservation.ID).Updates(&rapportobservation).Error
}

// delete rapport observation observation
func DeleteRapportObservation(db *gorm.DB, rapportobservation_id uint) error {
	return db.Where("id=?", rapportobservation_id).Delete(&RapportObservation{}).Error
}
