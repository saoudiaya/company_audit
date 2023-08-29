package tache

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesTache(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("taches", "write", enforcer), baseInstance.NewTache)
	router.GET("/all", middleware.Authorize("taches", "read", enforcer), baseInstance.GetTaches)
	router.GET("/:id", middleware.Authorize("taches", "read", enforcer), baseInstance.GetTacheById)
	router.POST("/search", middleware.Authorize("taches", "read", enforcer), baseInstance.SearchTache)
	router.PUT("/:id", middleware.Authorize("taches", "write", enforcer), baseInstance.UpdateTache)
	router.DELETE("/:id", middleware.Authorize("taches", "write", enforcer), baseInstance.DeleteTache)

}
