package objectif

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesObjectif(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("objectifs", "write", enforcer), baseInstance.NewObjectif)
	router.GET("/all", middleware.Authorize("objectifs", "read", enforcer), baseInstance.GetObjectifs)
	router.GET("/:id", middleware.Authorize("objectifs", "read", enforcer), baseInstance.GetObjectifById)
	router.POST("/search", middleware.Authorize("objectifs", "read", enforcer), baseInstance.SearchObjectif)
	router.PUT("/:id", middleware.Authorize("objectifs", "write", enforcer), baseInstance.UpdateObjectif)
	router.DELETE("/:id", middleware.Authorize("objectifs", "write", enforcer), baseInstance.DeleteObjectif)

}
