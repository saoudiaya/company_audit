package listecontrole

import (
	"pfe/api/v1/audit"
	"pfe/api/v1/critere"

	"net/http"
	"os"
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

// create new liste controle
func (db Database) NewListeControle(ctx *gin.Context) {

	// init vars
	var listecontrole ListeControleRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&listecontrole); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(listecontrole.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check audit exists
	if _, err := audit.GetAuditById(db.DB, listecontrole.AuditID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this audit does not exist"})
		return
	}

	// check critere exists
	if exists := critere.CheckCritereExists(db.DB, listecontrole.CritereID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this critere does not exist"})
		return
	}

	// init new liste controle
	new_listecontrole := ListeControle{
		ID:                  0,
		Nom:                 listecontrole.Nom,
		Description:         listecontrole.Description,
		AuditID:             listecontrole.AuditID,
		CritereID:           listecontrole.CritereID,
		ApprobationAuditiee: listecontrole.ApprobationAuditiee,
	}

	// create liste controle
	if _, err := NewListeControle(db.DB, new_listecontrole); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created"})

}

// get all liste controles from database
func (db Database) GetListeControles(ctx *gin.Context) {

	// get liste controle
	listecontrole, err := GetListeControles(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, listecontrole)
}

// get liste controle by id

func (db Database) GetListeControleById(ctx *gin.Context) {

	// get id value from path
	listecontrole_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get liste controle by id

	listecontrole, err := GetListeControleById(db.DB, uint(listecontrole_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, listecontrole)
}

// search liste controle from database
func (db Database) SearchListeControle(ctx *gin.Context) {

	// init vars
	var listecontrole ListeControle

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&listecontrole); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// check liste controle exists
	if exists := CheckListeControleExists(db.DB, listecontrole.ID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "liste controle  does not exist"})
		return
	}

	// search liste controle from database
	listecontroles, err := SearchListeControle(db.DB, listecontrole)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, listecontroles)
}

// update listecontrole

func (db Database) UpdateListeControle(ctx *gin.Context) {

	// init vars
	var listecontrole ListeControle
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&listecontrole); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	listecontrole_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(listecontrole.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	listecontrole.ID = uint(listecontrole_id)

	// update liste controle
	if err = UpdateListeControle(db.DB, listecontrole); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteListeControle(ctx *gin.Context) {

	// get id from path
	listecontrole_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete listecontrole
	if err = DeleteListeControle(db.DB, uint(listecontrole_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
