package revuedocument

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

// create new revue document
func (db Database) NewRevueDocument(ctx *gin.Context) {

	// init vars
	var revuedocument RevueDocumentRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&revuedocument); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(revuedocument.Titre) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check audit exists
	if _, err := audit.GetAuditById(db.DB, revuedocument.AuditID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this audit does not exist"})
		return
	}

	// check utilisateur exists
	if _, err := utilisateur.GetUtilisateurById(db.DB, revuedocument.UtilisateurID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this utilisateur does not exist"})
		return
	}

	// init new revue document
	new_revuedocument := RevueDocument{
		ID:            0,
		Titre:         revuedocument.Titre,
		AuditID:       revuedocument.AuditID,
		UtilisateurID: revuedocument.UtilisateurID,
		Date:          revuedocument.Date,
	}

	// create revue document
	if _, err := NewRevueDocument(db.DB, new_revuedocument); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})

}

// get all revue documents from database
func (db Database) GetRevueDocuments(ctx *gin.Context) {

	// get revue document
	revuedocument, err := GetRevueDocuments(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, revuedocument)
}

// get revue document by id

func (db Database) GetRevueDocumentById(ctx *gin.Context) {

	// get id value from path
	revuedocument_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get revue document by id

	revuedocument, err := GetRevueDocumentById(db.DB, uint(revuedocument_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, revuedocument)
}

// search revue document from database
func (db Database) SearchRevueDocument(ctx *gin.Context) {

	// init vars
	var revuedocument RevueDocument

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&revuedocument); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check revue document exists
	if exists := CheckRevueDocumentExists(db.DB, revuedocument.ID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "revue document does not exist"})
		return
	}

	// search revue document from database
	revuedocuments, err := SearchRevueDocument(db.DB, revuedocument)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, revuedocuments)
}

func (db Database) UpdateRevueDocument(ctx *gin.Context) {

	// init vars
	var revuedocument RevueDocument
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&revuedocument); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	revuedocument_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(revuedocument.Titre) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	revuedocument.ID = uint(revuedocument_id)

	// update revuedocument
	if err = UpdateRevueDocument(db.DB, revuedocument); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteRevueDocument(ctx *gin.Context) {

	// get id from path
	revuedocument_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete revue document
	if err = DeleteRevueDocument(db.DB, uint(revuedocument_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
