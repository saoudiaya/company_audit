package app

import (
	"pfe/api/app/entreprise"
	permission "pfe/api/app/permission"
	"pfe/api/app/role"
	"pfe/api/app/utilisateur"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// declare app routes
func RoutesApps(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	// utilisateur routes
	utilisateur.RoutesUtilisateur(router.Group("/utilisateur"), db, enforcer)

	// role routes
	role.RoutesRoles(router.Group("/role"), db, enforcer)

	// permission routes
	permission.RoutesPermissions(router.Group("/permission"), db, enforcer)

	// entreprise routes
	entreprise.RoutesEntreprises(router.Group("/entreprise"), db, enforcer)

}
