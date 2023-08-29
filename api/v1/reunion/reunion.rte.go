package reunion

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesReunion(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("reunions", "write", enforcer), baseInstance.NewReunion)
	router.GET("/all", middleware.Authorize("reunions", "read", enforcer), baseInstance.GetReunions)
	router.GET("/:id", middleware.Authorize("reunions", "read", enforcer), baseInstance.GetReunionById)
	router.POST("/search", middleware.Authorize("reunions", "read", enforcer), baseInstance.SearchReunion)
	router.PUT("/:id", middleware.Authorize("reunions", "write", enforcer), baseInstance.UpdateReunion)
	router.DELETE("/:id", middleware.Authorize("reunions", "write", enforcer), baseInstance.DeleteReunion)

}
