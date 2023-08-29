package rapportobservation

import (
	"net/http"
	"os"
	"pfe/api/v1/observation"
	"pfe/api/v1/rapport"
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

// create new rapport observation
func (db Database) NewRapportObservation(ctx *gin.Context) {

	// init vars
	var rapportobservation RapportObservationRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&rapportobservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(rapportobservation.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check observation exists
	if _, err := observation.GetObservationById(db.DB, rapportobservation.ObservationID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this observation does not exist"})
		return
	}

	// check rapport exists
	if _, err := rapport.GetRapportById(db.DB, rapportobservation.RapportID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this rapport does not exist"})
		return
	}

	// init new rapport observation
	new_rapportobservation := RapportObservation{
		ID:            0,
		Nom:           rapportobservation.Nom,
		Description:   rapportobservation.Description,
		ObservationID: rapportobservation.ObservationID,
		RapportID:     rapportobservation.RapportID,
	}

	// create rapport observation
	if _, err := NewRapportObservation(db.DB, new_rapportobservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created"})

}

// get all rapports observations from database
func (db Database) GetRapportsObservations(ctx *gin.Context) {

	// get rapport observation
	rapportobservation, err := GetRapportsObservations(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, rapportobservation)
}

// get rapport observation by id

func (db Database) GetRapportObservationById(ctx *gin.Context) {

	// get id value from path
	rapportobservation_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get rapport observation by id

	rapportobservation, err := GetRapportObservationById(db.DB, uint(rapportobservation_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, rapportobservation)
}

// search rapport observation from database
func (db Database) SearchRapportObservation(ctx *gin.Context) {

	// init vars
	var rapportobservation RapportObservation

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&rapportobservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// check rapportobservation exists
	if exists := CheckRapportObservationExists(db.DB, rapportobservation.ID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "rapport observation  does not exist"})
		return
	}

	// search rapport observation from database
	rapportsobservations, err := SearchRapportObservation(db.DB, rapportobservation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, rapportsobservations)
}

// update rapport observation

func (db Database) UpdateRapportObservation(ctx *gin.Context) {

	// init vars
	var rapportobservation RapportObservation
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&rapportobservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	rapportobservation_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(rapportobservation.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	rapportobservation.ID = uint(rapportobservation_id)

	// update rapport observation
	if err = UpdateRapportObservation(db.DB, rapportobservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteRapportObservation(ctx *gin.Context) {

	// get id from path
	rapportobservation_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete rapport observation
	if err = DeleteRapportObservation(db.DB, uint(rapportobservation_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
