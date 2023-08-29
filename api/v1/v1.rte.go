package v1

import (
	"pfe/api/v1/audit"
	"pfe/api/v1/commentaire"
	"pfe/api/v1/critere"
	"pfe/api/v1/exigence"
	listecontrole "pfe/api/v1/listecontrole"
	"pfe/api/v1/membre"
	"pfe/api/v1/norme"
	"pfe/api/v1/notification"
	"pfe/api/v1/objectif"
	"pfe/api/v1/observation"
	"pfe/api/v1/perimetre"
	"pfe/api/v1/rapport"
	rapportobservation "pfe/api/v1/rapport_observation"
	"pfe/api/v1/reponse"
	"pfe/api/v1/reunion"
	revuedocument "pfe/api/v1/revue_document"
	"pfe/api/v1/tache"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesV1(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {

	//project routes
	// audit routes
	audit.RoutesAudit(router.Group("/audit"), db, enforcer)

	// critere routes
	critere.RoutesCritere(router.Group("/critere"), db, enforcer)

	// exigence routes
	exigence.RoutesExigence(router.Group("/exigence"), db, enforcer)

	// liste controle routes
	listecontrole.RoutesListeControle(router.Group("/listecontrole"), db, enforcer)

	// membre routes
	membre.RoutesMembre(router.Group("/membre"), db, enforcer)

	// norme routes
	norme.RoutesNorme(router.Group("/norme"), db, enforcer)

	// notification routes
	notification.RoutesNotification(router.Group("/notification"), db, enforcer)

	// objectif routes
	objectif.RoutesObjectif(router.Group("/objectif"), db, enforcer)

	// observation routes
	observation.RoutesObservation(router.Group("/observation"), db, enforcer)

	// perimetre routes
	perimetre.RoutesPerimetre(router.Group("/perimetre"), db, enforcer)

	// rapport routes
	rapport.RoutesRapport(router.Group("/rapport"), db, enforcer)

	// rapport observation routes
	rapportobservation.RoutesRapportObservation(router.Group("/rapportobservation"), db, enforcer)

	// reunion routes
	reunion.RoutesReunion(router.Group("/reunion"), db, enforcer)

	// revue document routes
	revuedocument.RoutesRevueDocument(router.Group("/revuedocument"), db, enforcer)

	// tache document routes
	tache.RoutesTache(router.Group("/tache"), db, enforcer)

	// commentaire document routes
	commentaire.RoutesCommentaires(router.Group("/commentaire"), db, enforcer)

	// reponse document routes
	reponse.RoutesReponses(router.Group("/reponse"), db, enforcer)

}
