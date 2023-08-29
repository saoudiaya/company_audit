package membre

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesMembre(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("membres", "write", enforcer), baseInstance.NewMembre)
	router.GET("/all", middleware.Authorize("membres", "read", enforcer), baseInstance.GetMembres)
	router.GET("/:id", middleware.Authorize("membres", "read", enforcer), baseInstance.GetMembreById)
	router.POST("/search", middleware.Authorize("membres", "read", enforcer), baseInstance.SearchMembres)
	router.PUT("/:id", middleware.Authorize("membres", "write", enforcer), baseInstance.UpdateMembre)
	router.DELETE("/:id", middleware.Authorize("membres", "write", enforcer), baseInstance.DeleteMembre)

}
