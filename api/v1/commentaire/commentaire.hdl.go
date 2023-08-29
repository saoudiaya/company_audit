package commentaire

import (
	"log"
	"net/http"
	"os"
	"pfe/api/app/common"
	"pfe/middleware"
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

// create new commentaire
func (db Database) NewCommentaire(ctx *gin.Context) {

	// init vars
	var commentaire CommentaireRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	log.Println("ok")

	// check json validity
	if err := ctx.ShouldBindJSON(&commentaire); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	log.Println("ok")
	// check fields
	if empty_reg.MatchString(commentaire.Nom) || empty_reg.MatchString(commentaire.Description) || commentaire.ManagedBy < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	log.Println("ok")
	// check utilisateur exists
	if exists := common.CheckUtilisateurExists(db.DB, commentaire.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "manager does not exist"})
		return
	}
	log.Println("ok")

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new commentaire
	new_commentaire := Commentaire{
		Nom:           commentaire.Nom,
		Description:   commentaire.Description,
		UtilisateurID: commentaire.UtilisateurID,
		ManagedBy:     commentaire.ManagedBy,
		CreatedBy:     session.UtilisateurID,
	}

	// create new commentaire
	_, err := NewCommentaire(db.DB, new_commentaire)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all commentaires from database
func (db Database) GetCommentaires(ctx *gin.Context) {

	// get commentaires
	commentaires, err := GetCommentaires(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, commentaires)
}

// get commentaire by id
func (db Database) GetCommentaireById(ctx *gin.Context) {

	// get id value from path
	commentaire_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get commentaire by id

	commentaire, err := GetCommentaireById(db.DB, uint(commentaire_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, commentaire)
}

// search commentaires from database
func (db Database) SearchCommentaires(ctx *gin.Context) {

	// init vars
	var commentaire Commentaire

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&commentaire); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search commentaires
	commentaires, err := SearchCommentaires(db.DB, commentaire)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, commentaires)
}

// update commentaire
func (db Database) UpdateCommentaire(ctx *gin.Context) {

	// init vars
	var commentaire Commentaire
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&commentaire); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(commentaire.Nom) || empty_reg.MatchString(commentaire.Description) || commentaire.ManagedBy < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check subsidiary_of exists
	if exists := CheckCommentaireExists(db.DB, commentaire.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid manager id"})
		return
	}

	// check utilisateur exists
	if exists := common.CheckUtilisateurExists(db.DB, commentaire.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "manager does not exist"})
		return
	}

	if commentaire.ManagedBy != 0 {

		// check utilisateur exists
		if exists := common.CheckUtilisateurExists(db.DB, commentaire.ManagedBy); !exists {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid utilisateur id"})
			return
		}
	}

	// get id value from path
	commentaire_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// init update utilisateur
	update_commentaire := Commentaire{
		ID:          uint(commentaire_id),
		Nom:         commentaire.Nom,
		Description: commentaire.Description,
		ManagedBy:   commentaire.ManagedBy,
	}

	// update commentaire
	if err = UpdateCommentaire(db.DB, update_commentaire); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// delete commentaire from database
func (db Database) DeleteCommentaire(ctx *gin.Context) {

	// get id value from path
	commentaire_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete commentaire
	if err := DeleteCommentaire(db.DB, uint(commentaire_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
