package objectif

import (
	"net/http"
	"os"
	"pfe/api/v1/audit"
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

// create new objectif
func (db Database) NewObjectif(ctx *gin.Context) {

	// init vars
	var objectif ObjectifRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&objectif); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(objectif.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check audit exists
	if _, err := audit.GetAuditById(db.DB, objectif.AuditID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this audit does not exist"})
		return
	}

	// init new objectif
	new_objectif := Objectif{
		ID:          0,
		Nom:         objectif.Nom,
		Description: objectif.Description,
		AuditID:     objectif.AuditID,
	}

	// create objectif
	if _, err := NewObjectif(db.DB, new_objectif); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created"})

}

// get all objectifs from database
func (db Database) GetObjectifs(ctx *gin.Context) {

	// get objectif
	objectif, err := GetObjectifs(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, objectif)
}

// get objectif by id

func (db Database) GetObjectifById(ctx *gin.Context) {

	// get id value from path
	objectif_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get objectif by id

	objectif, err := GetObjectifById(db.DB, uint(objectif_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, objectif)
}

// search objectif from database
func (db Database) SearchObjectif(ctx *gin.Context) {

	// init vars
	var objectif Objectif

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&objectif); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// check objectif exists
	if exists := CheckObjectifExists(db.DB, objectif.ID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "objectif  does not exist"})
		return
	}

	// search objectif from database
	objectifs, err := SearchObjectif(db.DB, objectif)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, objectifs)
}

// update objectif

func (db Database) UpdateObjectif(ctx *gin.Context) {

	// init vars
	var objectif Objectif
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&objectif); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	objectif_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(objectif.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	objectif.ID = uint(objectif_id)

	// update objectif
	if err = UpdateObjectif(db.DB, objectif); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteObjectif(ctx *gin.Context) {

	// get id from path
	objectif_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete objectif
	if err = DeleteObjectif(db.DB, uint(objectif_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
