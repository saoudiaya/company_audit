package rapport

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesRapport(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("rapports", "write", enforcer), baseInstance.NewRapport)
	router.GET("/all", middleware.Authorize("rapports", "read", enforcer), baseInstance.GetRapports)
	router.GET("/:id", middleware.Authorize("rapports", "read", enforcer), baseInstance.GetRapportById)
	router.POST("/search", middleware.Authorize("rapports", "read", enforcer), baseInstance.SearchRapport)
	router.PUT("/:id", middleware.Authorize("rapports", "write", enforcer), baseInstance.UpdateRapport)
	router.DELETE("/:id", middleware.Authorize("rapports", "write", enforcer), baseInstance.DeleteRapport)

}
