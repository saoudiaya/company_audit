package listecontrole

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesListeControle(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("listecontroles", "write", enforcer), baseInstance.NewListeControle)
	router.GET("/all", middleware.Authorize("listecontroles", "read", enforcer), baseInstance.GetListeControles)
	router.GET("/:id", middleware.Authorize("listecontroles", "read", enforcer), baseInstance.GetListeControleById)
	router.POST("/search", middleware.Authorize("listecontroles", "read", enforcer), baseInstance.SearchListeControle)
	router.PUT("/:id", middleware.Authorize("listecontroles", "write", enforcer), baseInstance.UpdateListeControle)
	router.DELETE("/:id", middleware.Authorize("listecontroles", "write", enforcer), baseInstance.DeleteListeControle)

}
