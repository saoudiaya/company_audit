package reunion

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

// create new reunion
func (db Database) NewReunion(ctx *gin.Context) {

	// init vars
	var reunion Reunion
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&reunion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(reunion.Titre) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check audit exists
	if _, err := audit.GetAuditById(db.DB, reunion.AuditID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this audit does not exist"})
		return
	}

	//check members exists
	//membre := reunion.Membre
	//for i := 0; i < len(membre); i++ {
	//if exists := common.CheckUtilisateurExists(db.DB, uint(membre[i])); !exists {
	//ctx.JSON(http.StatusBadRequest, gin.H{"message": "verifie member ID"})
	//return
	//}
	//}
	// check utilisateur exists
	if _, err := audit.GetAuditById(db.DB, reunion.UtilisateurID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this user does not exist"})
		return
	}
	// init new reunion
	new_reunion := Reunion{
		ID:            0,
		Titre:         reunion.Titre,
		UtilisateurID: reunion.UtilisateurID,
		AuditID:       reunion.AuditID,
		Datedebut:     reunion.Datedebut,
		Datefin:       reunion.Datefin,
	}

	// create reunion
	if _, err := NewReunion(db.DB, new_reunion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})

}

// get all reunions from database
func (db Database) GetReunions(ctx *gin.Context) {

	// get reunion
	reunion, err := GetReunions(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reunion)
}

// get reunion by id

func (db Database) GetReunionById(ctx *gin.Context) {

	// get id value from path
	reunion_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get reunion by id

	reunion, err := GetReunionById(db.DB, uint(reunion_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reunion)
}

// search reunion from database
func (db Database) SearchReunion(ctx *gin.Context) {

	// init vars
	var reunion Reunion

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&reunion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check reunion exists
	if exists := CheckReunionExists(db.DB, reunion.ID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "reunion does not exist"})
		return
	}

	// search reunion from database
	reunions, err := SearchReunion(db.DB, reunion)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reunions)
}

func (db Database) UpdateReunion(ctx *gin.Context) {

	// init vars
	var reunion Reunion
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&reunion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	reunion_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(reunion.Titre) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	reunion.ID = uint(reunion_id)
	//reunion.CreatedBy = 0
	//reunion.EntrepriseId = 0

	// update reunion
	if err = UpdateReunion(db.DB, reunion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteReunion(ctx *gin.Context) {

	// get id from path
	reunion_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete reunion
	if err = DeleteReunion(db.DB, uint(reunion_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
