package membre

import (
	"net/http"
	"os"
	"pfe/api/app/common"
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

// create new membre
func (db Database) NewMembre(ctx *gin.Context) {

	// init vars
	var membre MembreRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&membre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(membre.Poste) || empty_reg.MatchString(membre.Type) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check audit exists
	if exists := audit.CheckAuditExists(db.DB, membre.AuditID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this audit does not exist"})
		return
	}

	// check utilisateur exists
	if exists := common.CheckUtilisateurExists(db.DB, membre.UtilisateurID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this utilisateur not exist"})
		return
	}

	// init new membre
	new_membre := Membre{
		AuditID:       membre.AuditID,
		UtilisateurID: membre.UtilisateurID,
		Type:          membre.Type,
		Poste:         membre.Poste,
	}

	// create membre
	if _, err := NewMembre(db.DB, new_membre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// permission
	//db.Enforcer.AddGroupingPolicy(strconv.FormatUint(uint64(membre.ID), 10))

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all membres from database
func (db Database) GetMembres(ctx *gin.Context) {

	// get membres
	membres, err := GetMembres(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, membres)
}

// get membre by id

func (db Database) GetMembreById(ctx *gin.Context) {

	// get id value from path
	membre_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get membre by id

	membre, err := GetMembreById(db.DB, uint(membre_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, membre)
}

// search membres from database
func (db Database) SearchMembres(ctx *gin.Context) {

	// init vars
	var membre Membre

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&membre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search membres from database
	membres, err := SearchMembres(db.DB, membre)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, membres)
}

func (db Database) UpdateMembre(ctx *gin.Context) {

	// init vars
	var membre Membre
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&membre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	membre_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(membre.Poste) || empty_reg.MatchString(membre.Type) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	membre.ID = uint(membre_id)

	// update membre
	if err = UpdateMembre(db.DB, membre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteMembre(ctx *gin.Context) {

	// get id from path
	membre_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete membre
	if err = DeleteMembre(db.DB, uint(membre_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
