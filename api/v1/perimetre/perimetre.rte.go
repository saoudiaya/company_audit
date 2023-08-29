package perimetre

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesPerimetre(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("perimetres", "write", enforcer), baseInstance.NewPerimetre)
	router.GET("/all", middleware.Authorize("perimetres", "read", enforcer), baseInstance.GetPerimetres)
	router.GET("/:id", middleware.Authorize("perimetres", "read", enforcer), baseInstance.GetPerimetreById)
	router.POST("/search", middleware.Authorize("perimetres", "read", enforcer), baseInstance.SearchPerimetre)
	router.PUT("/:id", middleware.Authorize("perimetres", "write", enforcer), baseInstance.UpdatePerimetre)
	router.DELETE("/:id", middleware.Authorize("perimetres", "write", enforcer), baseInstance.DeletePerimetre)

}
