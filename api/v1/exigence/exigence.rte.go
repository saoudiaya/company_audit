package exigence

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesExigence(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("exigences", "write", enforcer), baseInstance.NewExigence)
	router.GET("/all", middleware.Authorize("exigences", "read", enforcer), baseInstance.GetExigences)
	router.GET("/:id", middleware.Authorize("exigences", "read", enforcer), baseInstance.GetExigenceById)
	router.POST("/search", middleware.Authorize("exigences", "read", enforcer), baseInstance.SearchExigence)
	router.PUT("/:id", middleware.Authorize("exigences", "write", enforcer), baseInstance.UpdateExigence)
	router.DELETE("/:id", middleware.Authorize("exigences", "write", enforcer), baseInstance.DeleteExigence)

}
