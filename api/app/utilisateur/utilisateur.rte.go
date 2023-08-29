package utilisateur

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesUtilisateur(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("utilisateurs", "write", enforcer), baseInstance.NewUtilisateur)
	router.GET("/all", middleware.Authorize("utilisateurs", "read", enforcer), baseInstance.GetUtilisateurs)
	router.GET("/:id", middleware.Authorize("utilisateurs", "read", enforcer), baseInstance.GetUtilisateurById)
	router.POST("/search", middleware.Authorize("utilisateurs", "read", enforcer), baseInstance.SearchUtilisateurs)
	router.PUT("/:id", middleware.Authorize("utilisateurs", "write", enforcer), baseInstance.UpdateUtilisateur)
	router.DELETE("/:id", middleware.Authorize("utilisateurs", "write", enforcer), baseInstance.DeleteUtilisateur)

}
