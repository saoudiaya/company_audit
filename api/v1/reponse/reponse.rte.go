package reponse

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare reponse routes
func RoutesReponses(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("reponses", "write", enforcer), baseInstance.NewReponse)
	router.GET("/all", middleware.Authorize("reponses", "read", enforcer), baseInstance.GetReponses)
	router.GET("/:id", middleware.Authorize("reponses", "read", enforcer), baseInstance.GetReponseById)
	router.POST("/search", middleware.Authorize("reponses", "read", enforcer), baseInstance.SearchReponses)
	router.PUT("/:id", middleware.Authorize("reponses", "write", enforcer), baseInstance.UpdateReponse)
	router.DELETE("/:id", middleware.Authorize("reponses", "write", enforcer), baseInstance.DeleteReponse)
}
