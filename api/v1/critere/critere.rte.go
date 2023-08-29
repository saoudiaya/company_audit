package critere

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesCritere(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("criteres", "write", enforcer), baseInstance.NewCritere)
	router.GET("/all", middleware.Authorize("criteres", "read", enforcer), baseInstance.GetCriteres)
	router.GET("/:id", middleware.Authorize("criteres", "read", enforcer), baseInstance.GetCritereById)
	router.POST("/search", middleware.Authorize("criteres", "read", enforcer), baseInstance.SearchCritere)
	router.PUT("/:id", middleware.Authorize("criteres", "write", enforcer), baseInstance.UpdateCritere)
	router.DELETE("/:id", middleware.Authorize("criteres", "write", enforcer), baseInstance.DeleteCritere)

}
