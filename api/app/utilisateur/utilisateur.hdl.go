package utilisateur

import (
	"net/http"
	"os"
	"pfe/api/app/common"
	"pfe/api/app/entreprise"
	"pfe/api/app/role"
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

// create new utilisateur
func (db Database) NewUtilisateur(ctx *gin.Context) {

	// init vars
	var utilisateur common.Utilisateur
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))
	email_reg, _ := regexp.Compile(os.Getenv("EMAIL_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&utilisateur); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(utilisateur.Nom) || empty_reg.MatchString(utilisateur.Email) || empty_reg.MatchString(utilisateur.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check email validity
	if !email_reg.MatchString(utilisateur.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please enter a valid email address"})
		return
	}

	// check role exists
	if _, err := role.GetRoleByID(db.DB, utilisateur.RoleID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this role does not exist"})
		return
	}

	// check entreprise exists
	if _, err := entreprise.GetEntrepriseById(db.DB, utilisateur.EntrepriseID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this entreprise does not exist"})
		return
	}

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// hash password
	HashPassword(&utilisateur.Password)

	// init new utilisateur
	new_utilisateur := common.Utilisateur{
		ID:           0,
		Nom:          utilisateur.Nom,
		Address:      utilisateur.Address,
		Email:        utilisateur.Email,
		Phone:        utilisateur.Phone,
		Password:     utilisateur.Password,
		RoleNom:      utilisateur.RoleNom,
		RoleID:       utilisateur.RoleID,
		EntrepriseID: utilisateur.EntrepriseID,
		CreatedBy:    session.UtilisateurID,
	}

	// create utilisateur
	if _, err := NewUtilisateur(db.DB, new_utilisateur); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// permission
	db.Enforcer.AddGroupingPolicy(strconv.FormatUint(uint64(utilisateur.ID), 10), utilisateur.Role)

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all utilisateurs from database
func (db Database) GetUtilisateurs(ctx *gin.Context) {

	// get utilisateurs
	utilisateurs, err := GetUtilisateurs(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, utilisateurs)
}

// get utilisateur by id

func (db Database) GetUtilisateurById(ctx *gin.Context) {

	// get id value from path
	utilisateur_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get utilisateur by id

	utilisateur, err := GetUtilisateurById(db.DB, uint(utilisateur_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, utilisateur)
}

// search utilisateurs from database
func (db Database) SearchUtilisateurs(ctx *gin.Context) {

	// init vars
	var utilisateur common.Utilisateur

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&utilisateur); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search utilisateurs from database
	utilisateurs, err := SearchUtilisateurs(db.DB, utilisateur)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, utilisateurs)
}

func (db Database) UpdateUtilisateur(ctx *gin.Context) {

	// init vars
	var utilisateur common.Utilisateur
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&utilisateur); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	utilisateur_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(utilisateur.Nom) || empty_reg.MatchString(utilisateur.Email) || empty_reg.MatchString(utilisateur.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check role exists
	if _, err := role.GetRoleByID(db.DB, utilisateur.RoleID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this role does not exist"})
		return
	}

	// check entreprise exists
	if _, err := entreprise.GetEntrepriseById(db.DB, utilisateur.EntrepriseID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this entreprise does not exist"})
		return
	}

	// hash password
	HashPassword(&utilisateur.Password)

	// ignore key attributs
	utilisateur.ID = uint(utilisateur_id)
	utilisateur.CreatedBy = 0
	utilisateur.EntrepriseID = 0

	// update utilisateur
	if err = UpdateUtilisateur(db.DB, utilisateur); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteUtilisateur(ctx *gin.Context) {

	// get id from path
	utilisateur_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete utilisateur
	if err = DeleteUtilisateur(db.DB, uint(utilisateur_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
