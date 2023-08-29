package rapportobservation

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesRapportObservation(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("rapportsobservations", "write", enforcer), baseInstance.NewRapportObservation)
	router.GET("/all", middleware.Authorize("rapportsobservations", "read", enforcer), baseInstance.GetRapportsObservations)
	router.GET("/:id", middleware.Authorize("rapportsobservations", "read", enforcer), baseInstance.GetRapportObservationById)
	router.POST("/search", middleware.Authorize("rapportsobservations", "read", enforcer), baseInstance.SearchRapportObservation)
	router.PUT("/:id", middleware.Authorize("rapportsobservations", "write", enforcer), baseInstance.UpdateRapportObservation)
	router.DELETE("/:id", middleware.Authorize("rapportsobservations", "write", enforcer), baseInstance.DeleteRapportObservation)

}
