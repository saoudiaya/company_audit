package exigence

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

// create new exigence
func (db Database) NewExigence(ctx *gin.Context) {

	// init vars
	var exigence ExigenceRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&exigence); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(exigence.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check audit exists
	if _, err := audit.GetAuditById(db.DB, exigence.AuditID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this audit does not exist"})
		return
	}

	// init new exigence
	new_exigence := Exigence{
		ID:          0,
		Nom:         exigence.Nom,
		Description: exigence.Description,
		Niveau:      exigence.Niveau,
		AuditID:     exigence.AuditID,
	}

	// create exigence
	if _, err := NewExigence(db.DB, new_exigence); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created"})

}

// get all exigences from database
func (db Database) GetExigences(ctx *gin.Context) {

	// get exigence
	exigence, err := GetExigences(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exigence)
}

// get exigence by id

func (db Database) GetExigenceById(ctx *gin.Context) {

	// get id value from path
	exigence_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get exigence by id

	exigence, err := GetExigenceById(db.DB, uint(exigence_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exigence)
}

// search exigence from database
func (db Database) SearchExigence(ctx *gin.Context) {

	// init vars
	var exigence Exigence

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&exigence); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// check exigence exists
	if exists := CheckExigenceExists(db.DB, exigence.ID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "exigence  does not exist"})
		return
	}

	// search exigence from database
	exigences, err := SearchExigence(db.DB, exigence)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exigences)
}

// update exigence

func (db Database) UpdateExigence(ctx *gin.Context) {

	// init vars
	var exigence Exigence
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&exigence); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	exigence_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(exigence.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	exigence.ID = uint(exigence_id)

	// update exigence
	if err = UpdateExigence(db.DB, exigence); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteExigence(ctx *gin.Context) {

	// get id from path
	exigence_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete exigence
	if err = DeleteExigence(db.DB, uint(exigence_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
