package reponse

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

// create new reponse
func (db Database) NewReponse(ctx *gin.Context) {

	// init vars
	var reponse ReponseRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	log.Println("ok")

	// check json validity
	if err := ctx.ShouldBindJSON(&reponse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	log.Println("ok")
	// check fields
	if empty_reg.MatchString(reponse.Nom) || empty_reg.MatchString(reponse.Description) || reponse.ManagedBy < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	log.Println("ok")
	// check utilisateur exists
	if exists := common.CheckUtilisateurExists(db.DB, reponse.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "manager does not exist"})
		return
	}
	log.Println("ok")

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new reponse
	new_reponse := Reponse{
		Nom:           reponse.Nom,
		Description:   reponse.Description,
		CommentaireID: reponse.CommentaireID,
		ManagedBy:     reponse.ManagedBy,
		CreatedBy:     session.UtilisateurID,
	}

	// create new reponse
	_, err := NewReponse(db.DB, new_reponse)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all reponses from database
func (db Database) GetReponses(ctx *gin.Context) {

	// get reponses
	reponses, err := GetReponses(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reponses)
}

// get reponse by id
func (db Database) GetReponseById(ctx *gin.Context) {

	// get id value from path
	reponse_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get reponse by id

	reponse, err := GetReponseById(db.DB, uint(reponse_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reponse)
}

// search reponses from database
func (db Database) SearchReponses(ctx *gin.Context) {

	// init vars
	var reponse Reponse

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&reponse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search reponses
	reponses, err := SearchReponses(db.DB, reponse)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reponses)
}

// update reponse
func (db Database) UpdateReponse(ctx *gin.Context) {

	// init vars
	var reponse Reponse
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&reponse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(reponse.Nom) || empty_reg.MatchString(reponse.Description) || reponse.ManagedBy < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check subsidiary_of exists
	if exists := CheckReponseExists(db.DB, reponse.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid manager id"})
		return
	}

	// check utilisateur exists
	if exists := common.CheckUtilisateurExists(db.DB, reponse.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "manager does not exist"})
		return
	}

	if reponse.ManagedBy != 0 {

		// check utilisateur exists
		if exists := common.CheckUtilisateurExists(db.DB, reponse.ManagedBy); !exists {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid utilisateur id"})
			return
		}
	}

	// get id value from path
	reponse_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// init update utilisateur
	update_reponse := Reponse{
		ID:          uint(reponse_id),
		Nom:         reponse.Nom,
		Description: reponse.Description,
		ManagedBy:   reponse.ManagedBy,
	}

	// update reponse
	if err = UpdateReponse(db.DB, update_reponse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// delete reponse from database
func (db Database) DeleteReponse(ctx *gin.Context) {

	// get id value from path
	reponse_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete reponse
	if err := DeleteReponse(db.DB, uint(reponse_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
