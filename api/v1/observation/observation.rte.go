package observation

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesObservation(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("observations", "write", enforcer), baseInstance.NewObservation)
	router.GET("/all", middleware.Authorize("observations", "read", enforcer), baseInstance.GetObservations)
	router.GET("/:id", middleware.Authorize("observations", "read", enforcer), baseInstance.GetObservationById)
	router.POST("/search", middleware.Authorize("observations", "read", enforcer), baseInstance.SearchObservation)
	router.PUT("/:id", middleware.Authorize("observations", "write", enforcer), baseInstance.UpdateObservation)
	router.DELETE("/:id", middleware.Authorize("observations", "write", enforcer), baseInstance.DeleteObservation)

}
