package tache

import (
	"net/http"
	"os"
	"pfe/api/app/utilisateur"
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

// create new tache
func (db Database) NewTache(ctx *gin.Context) {

	// init vars
	var tache TacheRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&tache); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(tache.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check utilisateur exists
	if _, err := utilisateur.GetUtilisateurById(db.DB, tache.UtilisateurID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this utilisateur does not exist"})
		return
	}

	// check audit exists
	if _, err := audit.GetAuditById(db.DB, tache.AuditID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this audit does not exist"})
		return
	}

	// init new tache
	new_tache := Tache{
		ID:            0,
		Nom:           tache.Nom,
		Description:   tache.Description,
		UtilisateurID: tache.UtilisateurID,
		AuditID:       tache.AuditID,
		Datefin:       tache.Detefin,
	}

	// create tache
	if _, err := NewTache(db.DB, new_tache); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created"})

}

// get all taches from database
func (db Database) GetTaches(ctx *gin.Context) {

	// get tache
	tache, err := GetTaches(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tache)
}

// get tache by id

func (db Database) GetTacheById(ctx *gin.Context) {

	// get id value from path
	tache_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get tache by id

	tache, err := GetTacheById(db.DB, uint(tache_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tache)
}

// search tache from database
func (db Database) SearchTache(ctx *gin.Context) {

	// init vars
	var tache Tache

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&tache); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// check tache exists
	if exists := CheckTacheExists(db.DB, tache.ID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "tache  does not exist"})
		return
	}

	// search tache from database
	taches, err := SearchTache(db.DB, tache)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, taches)
}

// update tache

func (db Database) UpdateTache(ctx *gin.Context) {

	// init vars
	var tache Tache
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&tache); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	tache_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(tache.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	tache.ID = uint(tache_id)

	// update tache
	if err = UpdateTache(db.DB, tache); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteTache(ctx *gin.Context) {

	// get id from path
	tache_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete tache
	if err = DeleteTache(db.DB, uint(tache_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
