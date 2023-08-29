package perimetre

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

// create new perimetre
func (db Database) NewPerimetre(ctx *gin.Context) {

	// init vars
	var perimetre PerimetreRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&perimetre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(perimetre.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check audit exists
	if _, err := audit.GetAuditById(db.DB, perimetre.AuditID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this audit does not exist"})
		return
	}

	// init new perimetre
	new_perimetre := Perimetre{
		ID:          0,
		Nom:         perimetre.Nom,
		Description: perimetre.Description,
		AuditID:     perimetre.AuditID,
	}

	// create perimetre
	if _, err := NewPerimetre(db.DB, new_perimetre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created"})

}

// get all perimetres from database
func (db Database) GetPerimetres(ctx *gin.Context) {

	// get perimetre
	perimetre, err := GetPerimetres(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, perimetre)
}

// get perimetre by id

func (db Database) GetPerimetreById(ctx *gin.Context) {

	// get id value from path
	perimetre_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get perimetre by id

	perimetre, err := GetPerimetreById(db.DB, uint(perimetre_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, perimetre)
}

// search perimetre from database
func (db Database) SearchPerimetre(ctx *gin.Context) {

	// init vars
	var perimetre Perimetre

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&perimetre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// check perimetre exists
	if exists := CheckPerimetreExists(db.DB, perimetre.ID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "perimetre  does not exist"})
		return
	}

	// search perimetre from database
	perimetres, err := SearchPerimetre(db.DB, perimetre)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, perimetres)
}

// update perimetre

func (db Database) UpdatePerimetre(ctx *gin.Context) {

	// init vars
	var perimetre Perimetre
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&perimetre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	perimetre_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(perimetre.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	perimetre.ID = uint(perimetre_id)

	// update perimetre
	if err = UpdatePerimetre(db.DB, perimetre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeletePerimetre(ctx *gin.Context) {

	// get id from path
	perimetre_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete perimetre
	if err = DeletePerimetre(db.DB, uint(perimetre_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
