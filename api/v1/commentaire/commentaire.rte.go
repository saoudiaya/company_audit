package commentaire

import (
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare commentaire routes
func RoutesCommentaires(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	baseInstance := Database{DB: db, Enforcer: enforcer}

	router.POST("/new", middleware.Authorize("commentaires", "write", enforcer), baseInstance.NewCommentaire)
	router.GET("/all", middleware.Authorize("commentaires", "read", enforcer), baseInstance.GetCommentaires)
	router.GET("/:id", middleware.Authorize("commentaires", "read", enforcer), baseInstance.GetCommentaireById)
	router.POST("/search", middleware.Authorize("commentaires", "read", enforcer), baseInstance.SearchCommentaires)
	router.PUT("/:id", middleware.Authorize("commentaires", "write", enforcer), baseInstance.UpdateCommentaire)
	router.DELETE("/:id", middleware.Authorize("commentaires", "write", enforcer), baseInstance.DeleteCommentaire)
}
