package audit

import (
	"net/http"
	"os"
	"pfe/api/app/common"
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

// create new audit
func (db Database) NewAudit(ctx *gin.Context) {

	// init vars
	var audit AuditRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&audit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(audit.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check audit exists
	/*if _, err := audit.GetAuditById(db.DB, audit.EntrepriseID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this audit does not exist"})
		return
	}*/

	// init new audit
	new_audit := common.Audit{
		ID:                    0,
		Nom:                   audit.Nom,
		Description:           audit.Description,
		Type:                  audit.Type,
		Statut:                audit.Statut,
		Date_debut:            audit.Date_debut,
		Date_fin:              audit.Date_fin,
		Effacement:            audit.Effacement,
		Effacement_jours:      audit.Effacement_jours,
		Observation:           audit.Observation,
		UtilisateurPrincipale: audit.UtilisateurPrincipale,
		EntrepriseAuditie:     audit.EntrepriseAuditie,
		EntrepriseAuditrice:   audit.EntrepriseAuditrice,
		EntrepriseID:          audit.EntrepriseID,
	}

	// create audit
	if _, err := NewAudit(db.DB, new_audit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created"})

}

// get all audits from database
func (db Database) GetAudits(ctx *gin.Context) {

	// get audit
	audit, err := GetAudits(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, audit)
}

// get audit by id

func (db Database) GetAuditById(ctx *gin.Context) {

	// get id value from path
	audit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get audit by id

	audit, err := GetAuditById(db.DB, uint(audit_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, audit)
}

// search audit from database
func (db Database) SearchAudit(ctx *gin.Context) {

	// init vars
	var audit common.Audit

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&audit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// check audit exists
	if exists := CheckAuditExists(db.DB, audit.ID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "audit  does not exist"})
		return
	}

	// search audit from database
	audits, err := SearchAudit(db.DB, audit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, audits)
}

// update audit

func (db Database) UpdateAudit(ctx *gin.Context) {

	// init vars
	var audit common.Audit
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&audit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	audit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(audit.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	audit.ID = uint(audit_id)

	// update audit
	if err = UpdateAudit(db.DB, audit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteAudit(ctx *gin.Context) {

	// get id from path
	audit_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete audit
	if err = DeleteAudit(db.DB, uint(audit_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
