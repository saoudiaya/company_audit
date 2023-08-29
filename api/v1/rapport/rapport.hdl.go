package rapport

import (
	"net/http"
	"os"
	"pfe/api/v1/observation"
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

// create new rapport
func (db Database) NewRapport(ctx *gin.Context) {

	// init vars
	var rapport RapportRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&rapport); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(rapport.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check observation exists
	if _, err := observation.GetObservationById(db.DB, rapport.ObservationID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this observation does not exist"})
		return
	}

	// init new rapport
	new_rapport := Rapport{
		ID:                 0,
		Nom:                rapport.Nom,
		Nombreofconformite: rapport.Nombreofconformite,
		Process:            rapport.Process,
		Nombreofarticle:    rapport.Nombreofarticle,
		Description:        rapport.Description,
		ObservationID:      rapport.ObservationID,
	}

	// create rapport
	if _, err := NewRapport(db.DB, new_rapport); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created"})

}

// get all rapports from database
func (db Database) GetRapports(ctx *gin.Context) {

	// get rapport
	rapport, err := GetRapports(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, rapport)
}

// get rapport by id

func (db Database) GetRapportById(ctx *gin.Context) {

	// get id value from path
	rapport_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get rapport by id

	rapport, err := GetRapportById(db.DB, uint(rapport_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, rapport)
}

// search rapport from database
func (db Database) SearchRapport(ctx *gin.Context) {

	// init vars
	var rapport Rapport

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&rapport); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// check rapport exists
	if exists := CheckRapportExists(db.DB, rapport.ID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "rapport  does not exist"})
		return
	}

	// search rapport from database
	rapports, err := SearchRapport(db.DB, rapport)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, rapports)
}

// update rapport

func (db Database) UpdateRapport(ctx *gin.Context) {

	// init vars
	var rapport Rapport
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&rapport); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	rapport_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(rapport.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	rapport.ID = uint(rapport_id)

	// update rapport
	if err = UpdateRapport(db.DB, rapport); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteRapport(ctx *gin.Context) {

	// get id from path
	rapport_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete rapport
	if err = DeleteRapport(db.DB, uint(rapport_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
