package norme

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesNorme(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("normes", "write", enforcer), baseInstance.NewNorme)
	router.GET("/all", middleware.Authorize("normes", "read", enforcer), baseInstance.GetNormes)
	router.GET("/:id", middleware.Authorize("normes", "read", enforcer), baseInstance.GetNormeById)
	router.POST("/search", middleware.Authorize("normes", "read", enforcer), baseInstance.SearchNorme)
	router.PUT("/:id", middleware.Authorize("normes", "write", enforcer), baseInstance.UpdateNorme)
	router.DELETE("/:id", middleware.Authorize("normes", "write", enforcer), baseInstance.DeleteNorme)

}
