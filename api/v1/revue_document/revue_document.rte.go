package revuedocument

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesRevueDocument(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("revuedocuments", "write", enforcer), baseInstance.NewRevueDocument)
	router.GET("/all", middleware.Authorize("revuedocuments", "read", enforcer), baseInstance.GetRevueDocuments)
	router.GET("/:id", middleware.Authorize("revuedocuments", "read", enforcer), baseInstance.GetRevueDocumentById)
	router.POST("/search", middleware.Authorize("revuedocuments", "read", enforcer), baseInstance.SearchRevueDocument)
	router.PUT("/:id", middleware.Authorize("revuedocuments", "write", enforcer), baseInstance.UpdateRevueDocument)
	router.DELETE("/:id", middleware.Authorize("revuedocuments", "write", enforcer), baseInstance.DeleteRevueDocument)

}
