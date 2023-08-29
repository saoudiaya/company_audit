package observation

import (
	"net/http"
	"os"
	listecontrole "pfe/api/v1/listecontrole"
	"regexp"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new observation
func (db Database) NewObservation(ctx *gin.Context) {

	// init vars
	var observation ObservationRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&observation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(observation.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check liste controle exists
	if _, err := listecontrole.GetListeControleById(db.DB, observation.ListeControleID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this liste controle does not exist"})
		return
	}

	// init new observation
	new_observation := Observation{
		ID:          0,
		Nom:         observation.Nom,
		Observation: observation.Observation,
		Revue:       observation.Revue,
		//Entretien:             observation.Entretien,
		//VerificationTechnique: observation.VerificationTechnique,
		//Analyse:               observation.Analyse,
		ListeControleID: observation.ListeControleID,
	}

	// create observation
	if _, err := NewObservation(db.DB, new_observation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created"})

}

// get all observations from database
func (db Database) GetObservations(ctx *gin.Context) {

	// get observation
	observation, err := GetObservations(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, observation)
}

// get observation by id

func (db Database) GetObservationById(ctx *gin.Context) {

	// get id value from path
	observation_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get observation by id

	observation, err := GetObservationById(db.DB, uint(observation_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, observation)
}

// search observation from database
func (db Database) SearchObservation(ctx *gin.Context) {

	// init vars
	var observation Observation

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&observation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// check observation exists
	if exists := CheckObservationExists(db.DB, observation.ID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "observation  does not exist"})
		return
	}

	// search observation from database
	observations, err := SearchObservation(db.DB, observation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, observations)
}

// update observation

func (db Database) UpdateObservation(ctx *gin.Context) {

	// init vars
	var observation Observation
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&observation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	observation_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(observation.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	observation.ID = uint(observation_id)

	// update observation
	if err = UpdateObservation(db.DB, observation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteObservation(ctx *gin.Context) {

	// get id from path
	observation_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete observation
	if err = DeleteObservation(db.DB, uint(observation_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
