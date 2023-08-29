package app_auth

import (
	"net/http"
	"os"
	"pfe/api/app/common"
	"pfe/api/app/entreprise"
	"pfe/api/app/role"
	utilisateur "pfe/api/app/utilisateur"
	"pfe/middleware"
	"regexp"
	"strconv"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// signup utilisateur
func (db Database) SignUpUtilisateur(ctx *gin.Context) {

	// init vars
	var account InsertUtilisateurRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// upmarshal sent json
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check field validity
	if empty_reg.MatchString(account.Nom) || empty_reg.MatchString(account.Email) || empty_reg.MatchString(account.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check role exists
	if _, err := role.GetRoleByID(db.DB, account.RoleID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this role does not exist"})
		return
	}

	// check entreprise exists
	if _, err := entreprise.GetEntrepriseById(db.DB, account.EntrepriseID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this entreprise does not exist"})
		return
	}

	// hash password
	utilisateur.HashPassword(&account.Password)

	// create new utilisateur
	new_utilisateur := common.Utilisateur{
		Nom:          account.Nom,
		Email:        account.Email,
		Password:     account.Password,
		RoleNom:      account.RoleNom,
		RoleID:       account.RoleID,
		EntrepriseID: account.EntrepriseID,
	}

	// create utilisateur
	saved_utilisateur, err := utilisateur.NewUtilisateur(db.DB, new_utilisateur)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// super add will add role
	db.Enforcer.AddGroupingPolicy(strconv.FormatUint(uint64(saved_utilisateur.ID), 10), saved_utilisateur.RoleNom)

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// signin utilisateur
func (db Database) SignInUtilisateur(ctx *gin.Context) {

	// init cars
	var utilisateur_login UtilisateurLogIn
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&utilisateur_login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	// check field validity
	if empty_reg.MatchString(utilisateur_login.Email) || empty_reg.MatchString(utilisateur_login.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check if email exists
	dbUtilisateur, err := utilisateur.GetUtilisateurByEmail(db.DB, utilisateur_login.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "no Such Utilisateur Found"})
		return
	}

	// update last login
	dbUtilisateur.LastLogin = time.Now()

	// update utilisateur
	if err := utilisateur.UpdateUtilisateur(db.DB, dbUtilisateur); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// compare password
	if isTrue := utilisateur.ComparePassword(dbUtilisateur.Password, utilisateur_login.Password); isTrue {

		// generate token
		token := middleware.GenerateToken(dbUtilisateur.ID, dbUtilisateur.EntrepriseID, dbUtilisateur.Role.Nom)
		ctx.JSON(http.StatusOK, UtilisateurLogedIn{Token: token})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"message": "password not matched"})
}
