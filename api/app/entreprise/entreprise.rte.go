package entreprise

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare entreprise routes
func RoutesEntreprises(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("entreprises", "write", enforcer), baseInstance.NewEntreprise)
	router.GET("/all", middleware.Authorize("entreprises", "read", enforcer), baseInstance.GetEntreprises)
	router.GET("/:id", middleware.Authorize("entreprises", "read", enforcer), baseInstance.GetEntrepriseById)
	router.POST("/search", middleware.Authorize("entreprises", "read", enforcer), baseInstance.SearchEntreprises)
	router.PUT("/:id", middleware.Authorize("entreprises", "write", enforcer), baseInstance.UpdateEntreprise)
	router.DELETE("/:id", middleware.Authorize("entreprises", "write", enforcer), baseInstance.DeleteEntreprise)
}
