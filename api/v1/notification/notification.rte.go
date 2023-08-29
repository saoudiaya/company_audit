package notification

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesNotification(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("notifications", "write", enforcer), baseInstance.NewNotification)
	router.GET("/all", middleware.Authorize("notifications", "read", enforcer), baseInstance.GetNotifications)
	router.GET("/:id", middleware.Authorize("notifications", "read", enforcer), baseInstance.GetNotificationById)
	router.POST("/search", middleware.Authorize("notifications", "read", enforcer), baseInstance.SearchNotification)
	router.PUT("/:id", middleware.Authorize("notifications", "write", enforcer), baseInstance.UpdateNotification)
	router.DELETE("/:id", middleware.Authorize("notifications", "write", enforcer), baseInstance.DeleteNotification)

}
