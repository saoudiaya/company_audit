package critere

import (
	"net/http"
	"os"
	"pfe/api/v1/norme"
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

// create new critere
func (db Database) NewCritere(ctx *gin.Context) {

	// init vars
	var critere CritereRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&critere); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(critere.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check norme exists
	if _, err := norme.GetNormeById(db.DB, critere.NormeID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this norme does not exist"})
		return
	}

	// init new critere
	new_critere := Critere{
		ID:          0,
		Nom:         critere.Nom,
		Description: critere.Description,
		//Obseravtion:           critere.Obseravtion,
		Entretien:             critere.Entretien,
		VerificationTechnique: critere.VerificationTechnique,
		Analyse:               critere.Analyse,
		NormeID:               critere.NormeID,
	}

	// create critere
	if _, err := NewCritere(db.DB, new_critere); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created"})

}

// get all criteres from database
func (db Database) GetCriteres(ctx *gin.Context) {

	// get critere
	critere, err := GetCriteres(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, critere)
}

// get critere by id

func (db Database) GetCritereById(ctx *gin.Context) {

	// get id value from path
	critere_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get critere by id

	critere, err := GetCritereById(db.DB, uint(critere_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, critere)
}

// search critere from database
func (db Database) SearchCritere(ctx *gin.Context) {

	// init vars
	var critere Critere

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&critere); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// check critere exists
	if exists := CheckCritereExists(db.DB, critere.ID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "critere  does not exist"})
		return
	}

	// search critere from database
	criteres, err := SearchCritere(db.DB, critere)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, criteres)
}

// update critere

func (db Database) UpdateCritere(ctx *gin.Context) {

	// init vars
	var critere Critere
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&critere); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	critere_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(critere.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	critere.ID = uint(critere_id)

	// update critere
	if err = UpdateCritere(db.DB, critere); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteCritere(ctx *gin.Context) {

	// get id from path
	critere_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete critere
	if err = DeleteCritere(db.DB, uint(critere_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
