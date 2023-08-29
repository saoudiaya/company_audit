package entreprise

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

// create new entreprise
func (db Database) NewEntreprise(ctx *gin.Context) {

	// init vars
	var entreprise EntrepriseRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	log.Println("ok")

	// check json validity
	if err := ctx.ShouldBindJSON(&entreprise); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	log.Println("ok")
	// check fields
	if empty_reg.MatchString(entreprise.Nom) || empty_reg.MatchString(entreprise.Email) || empty_reg.MatchString(entreprise.Phone) || empty_reg.MatchString(entreprise.Address) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	log.Println("ok")
	// check utilisateur exists
	if exists := common.CheckUtilisateurExists(db.DB, entreprise.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "manager does not exist"})
		return
	}
	log.Println("ok")

	// get values from session
	session := middleware.ExtractTokenValues(ctx)

	// init new entreprise
	new_entreprise := common.Entreprise{
		Nom:       entreprise.Nom,
		Email:     entreprise.Email,
		Phone:     entreprise.Phone,
		Address:   entreprise.Address,
		ManagedBy: entreprise.ManagedBy,
		CreatedBy: session.UtilisateurID,
	}

	// create new entreprise
	_, err := NewEntreprise(db.DB, new_entreprise)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created"})
}

// get all entreprises from database
func (db Database) GetEntreprises(ctx *gin.Context) {

	// get entreprises
	entreprises, err := GetEntreprises(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, entreprises)
}

// get entreprise by id
func (db Database) GetEntrepriseById(ctx *gin.Context) {

	// get id value from path
	entreprise_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get entreprise by id

	entreprise, err := GetEntrepriseById(db.DB, uint(entreprise_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, entreprise)
}

// @Summary     Rechercher des entreprises
// @Description Cette méthode permet de rechercher des entreprises en utilisant un objet JSON envoyé dans le corps de la requête. Les paramètres de recherche disponibles incluent le nom, l'adresse et le numéro de téléphone de l'entreprise etc.
// @Tags        Entreprise
// @Accept      json
// @Produce     json
// @Param       request body InsertEntreprise true "Entreprise required fields"
// @Schemes
// @Success     200 {array}     app_common.Entreprises
// @Failure     400 {object}    responses.Response
// @Failure     401 {object}    responses.Response
// @Failure     403 {object}    responses.Response
// @Failure     500 {object}    responses.Response
// @Router      /app/entreprises/search [post]
// search entreprises from database
func (db Database) SearchEntreprises(ctx *gin.Context) {

	// init vars
	var entreprise common.Entreprise

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&entreprise); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// search entreprises
	entreprises, err := SearchEntreprises(db.DB, entreprise)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, entreprises)
}

// update entreprise
func (db Database) UpdateEntreprise(ctx *gin.Context) {

	// init vars
	var entreprise common.Entreprise
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&entreprise); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check fields
	if empty_reg.MatchString(entreprise.Nom) || empty_reg.MatchString(entreprise.Email) || empty_reg.MatchString(entreprise.Phone) || empty_reg.MatchString(entreprise.Address) || entreprise.ManagedBy < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid fields"})
		return
	}

	// check subsidiary_of exists
	if exists := CheckEntrepriseExists(db.DB, entreprise.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid manager id"})
		return
	}

	// check utilisateur exists
	if exists := common.CheckUtilisateurExists(db.DB, entreprise.ManagedBy); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "manager does not exist"})
		return
	}

	if entreprise.ManagedBy != 0 {

		// check utilisateur exists
		if exists := common.CheckUtilisateurExists(db.DB, entreprise.ManagedBy); !exists {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid utilisateur id"})
			return
		}
	}

	// get id value from path
	entreprise_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// init update utilisateur
	update_entreprise := common.Entreprise{
		ID:        uint(entreprise_id),
		Nom:       entreprise.Nom,
		Email:     entreprise.Email,
		Phone:     entreprise.Phone,
		Address:   entreprise.Address,
		ManagedBy: entreprise.ManagedBy,
	}

	// update entreprise
	if err = UpdateEntreprise(db.DB, update_entreprise); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// delete entreprise from database
func (db Database) DeleteEntreprise(ctx *gin.Context) {

	// get id value from path
	entreprise_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete entreprise
	if err := DeleteEntreprise(db.DB, uint(entreprise_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
