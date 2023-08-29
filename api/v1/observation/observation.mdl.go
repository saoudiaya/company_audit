package observation

import (
	listecontrole "pfe/api/v1/listecontrole"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Observation struct {
	ID                    uint                        `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom                   string                      `gorm:"column:nom;not null" json:"nom"`
	Observation           string                      `gorm:"column:observation;" json:"observation"`
	Revue                 string                      `gorm:"column:revue;" json:"revue"`
	Entretien             string                      `gorm:"column:entretien;" json:"entretien"`
	VerificationTechnique string                      `gorm:"column:verification_technique;" json:"verification_technique"`
	Analyse               string                      `gorm:"column:analyse;" json:"analyse"`
	ListeControleID       uint                        `gorm:"foreignKey:OwningListecontroleId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"listecontrole_id"`
	ListesControle        listecontrole.ListeControle `gorm:"foreignKey:ListeControleID;references:ID" json:"listecontrole"`
	gorm.Model
}

// create new observation
func NewObservation(db *gorm.DB, observation Observation) (Observation, error) {
	return observation, db.Create(&observation).Error
}

// get all observation
func GetObservations(db *gorm.DB) (observation []Observation, err error) {
	return observation, db.Preload(clause.Associations).Find(&observation).Error
}

// check if observation exists
func CheckObservationExists(db *gorm.DB, id uint) bool {

	// init vars
	observation := &Observation{}

	// check if row exists
	check := db.Where("id=?", id).First(&observation)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get observation by id
func GetObservationById(db *gorm.DB, id uint) (observation Observation, err error) {
	return observation, db.Preload(clause.Associations).First(&observation, "id=?", id).Error
}

// search observation
func SearchObservation(db *gorm.DB, observation Observation) (observations []Observation, err error) {
	return observations, db.Where(&observation).Find(&observations).Error
}

// update observation
func UpdateObservation(db *gorm.DB, observation Observation) error {
	return db.Where("id=?", observation.ID).Updates(&observation).Error
}

// delete observation
func DeleteObservation(db *gorm.DB, observation_id uint) error {
	return db.Where("id=?", observation_id).Delete(&Observation{}).Error
}
