package api

import (
	"pfe/api/app"
	auth "pfe/api/auth"
	v_one "pfe/api/v1"
	"pfe/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare api routes
func RoutesApis(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	// auth routes
	auth.RoutesAuth(router.Group("/auth"), db, enforcer)

	// app routes
	app.RoutesApps(router.Group("/app", middleware.AuthorizeJWT()), db, enforcer)

	// v1 routes
	v_one.RoutesV1(router.Group("/v1", middleware.AuthorizeJWT()), db, enforcer)
}
