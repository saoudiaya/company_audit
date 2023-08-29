package audit

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesAudit(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("audits", "write", enforcer), baseInstance.NewAudit)
	router.GET("/all", middleware.Authorize("audits", "read", enforcer), baseInstance.GetAudits)
	router.GET("/:id", middleware.Authorize("audits", "read", enforcer), baseInstance.GetAuditById)
	router.POST("/search", middleware.Authorize("audits", "read", enforcer), baseInstance.SearchAudit)
	router.PUT("/:id", middleware.Authorize("audits", "write", enforcer), baseInstance.UpdateAudit)
	router.DELETE("/:id", middleware.Authorize("audits", "write", enforcer), baseInstance.DeleteAudit)

}
