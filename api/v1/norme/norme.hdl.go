package norme

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

// create new norme
func (db Database) NewNorme(ctx *gin.Context) {

	// init vars
	var norme NormeRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&norme); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(norme.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// init new norme
	new_norme := common.Norme{
		ID:          0,
		Nom:         norme.Nom,
		Description: norme.Description,
	}

	// create norme
	if _, err := NewNorme(db.DB, new_norme); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created"})

}

// get all normes from database
func (db Database) GetNormes(ctx *gin.Context) {

	// get norme
	norme, err := GetNormes(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, norme)
}

// get norme by id

func (db Database) GetNormeById(ctx *gin.Context) {

	// get id value from path
	norme_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get norme by id

	norme, err := GetNormeById(db.DB, uint(norme_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, norme)
}

// search norme from database
func (db Database) SearchNorme(ctx *gin.Context) {

	// init vars
	var norme common.Norme

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&norme); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// check norme exists
	if exists := CheckNormeExists(db.DB, norme.ID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "norme  does not exist"})
		return
	}

	// search norme from database
	normes, err := SearchNorme(db.DB, norme)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, normes)
}

// update norme

func (db Database) UpdateNorme(ctx *gin.Context) {

	// init vars
	var norme common.Norme
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&norme); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	norme_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(norme.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	norme.ID = uint(norme_id)

	// update norme
	if err = UpdateNorme(db.DB, norme); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteNorme(ctx *gin.Context) {

	// get id from path
	norme_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete norme
	if err = DeleteNorme(db.DB, uint(norme_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
