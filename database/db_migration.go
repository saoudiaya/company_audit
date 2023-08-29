package database

import (
	"fmt"
	"os"
	"pfe/api/app/common"
	permission "pfe/api/app/permission"
	"pfe/api/app/utilisateur"
	"pfe/api/v1/commentaire"
	"pfe/api/v1/critere"
	"pfe/api/v1/exigence"
	listecontrole "pfe/api/v1/listecontrole"
	"pfe/api/v1/membre"
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
	"strconv"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

// auto migrate datbles
func _auto_migrate_tables(db *gorm.DB) {

	// auto migrate casbin table
	if err := db.Table("casbin_rule").AutoMigrate(&permission.CasbinRule{}); err != nil {
		panic(fmt.Sprintf("Error while creating casbin table: %v", err))
	}
	// auto migrate tables
	//if err := db.AutoMigrate(
	//	&common.Audit{},
	//&common.Utilisateur{},
	//&critere.Critere{},
	//); err != nil {
	//panic(err)
	//}

	// auto migrate tables
	if err := db.AutoMigrate(
		&common.Audit{},
		&common.Utilisateur{},
		&critere.Critere{},
	); err != nil {
		panic(err)
	}
	if err :=
		db.Table("liste_controles").AutoMigrate(&listecontrole.ListeControle{}); err != nil {
		panic(err)
	}
	db.SetupJoinTable(&critere.Critere{}, "liste_controles", &common.Audit{})

	// auto migrate utilisateur,membre, role & entreprise tables
	if err := db.AutoMigrate(
		&common.Role{},
		&common.Entreprise{},
		&common.Norme{},
		&exigence.Exigence{},
		&objectif.Objectif{},
		&reunion.Reunion{},
		&observation.Observation{},
		&notification.Notification{},
		&perimetre.Perimetre{},
		&rapport.Rapport{},
		&commentaire.Commentaire{},
		&reponse.Reponse{},
		&rapportobservation.RapportObservation{},
		&tache.Tache{},
		&revuedocument.RevueDocument{},
	); err != nil {
		panic(err)
	}

	if err :=
		db.Table("membres").AutoMigrate(&membre.Membre{}); err != nil {
		panic(err)
	}
	db.SetupJoinTable(&common.Utilisateur{}, "Audits", "membres")
}

// auto create root utilisateur
func _create_root_utilisateur(db *gorm.DB, enforcer *casbin.Enforcer) {

	// init vars:
	// root
	var utilisateur_id uint
	root_utilisateur := common.Utilisateur{}
	// default role
	utilisateur_role := common.Role{}

	// create entreprise
	// check entreprise exists
	var db_entreprise common.Entreprise
	if check := db.Where("nom=?", os.Getenv("DEFAULT_COMPANY_NAME")).Find(&db_entreprise); check.RowsAffected == 0 && check.Error == nil {

		// create entreprise
		db_entreprise = common.Entreprise{Nom: os.Getenv("DEFAULT_COMPANY_NAME"), Email: os.Getenv("DEFAULT_COMPANY_EMAIL"), Phone: os.Getenv("DEFAULT_COMPANY_PHONE"), Address: os.Getenv("DEFAULT_COMPANY_ADDRESS"), ManagedBy: utilisateur_id, CreatedBy: utilisateur_id}
		err := db.Create(&db_entreprise).Error
		if err != nil {
			panic(fmt.Sprintf("[WARNING] error while creating the root entreprise: %v", err))
		}

		// edit utilisateur to add entreprise id
		if check := db.Where("email=?", os.Getenv("DEFAULT_EMAIL")).Find(&root_utilisateur); check.RowsAffected == 1 && check.Error == nil {
			root_utilisateur.EntrepriseID = db_entreprise.ID
			if update := db.Where("id=?", root_utilisateur.ID).Updates(&root_utilisateur); update.Error != nil {
				panic(fmt.Sprintf("[WARNING] error while updating the root utilisateur: %v", update.Error))
			}
		}
	}
	// create root role
	// check root role exists
	db_role := common.Role{}
	if check := db.Where("nom=?", os.Getenv("DEFAULT_ROOT")).Find(&db_role); check.RowsAffected == 0 && check.Error == nil {

		// create role utilisateur
		db_role = common.Role{Nom: os.Getenv("DEFAULT_ROOT"), EntrepriseID: db_entreprise.ID}

		if err := db.Create(&db_role).Error; err != nil {
			panic(fmt.Sprintf("[WARNING] error while creating the root role: %v", err))
		}
	}
	// create root utilisateur
	// check root utilisateur exists
	if check := db.Where("email=?", os.Getenv("DEFAULT_EMAIL")).Find(&root_utilisateur); check.RowsAffected == 0 && check.Error == nil {

		// create utilisateur
		db_utilisateur := common.Utilisateur{Nom: os.Getenv("DEFAULT_NAME"), Email: os.Getenv("DEFAULT_EMAIL"), Password: os.Getenv("DEFAULT_PASSWORD"), RoleID: db_role.ID, EntrepriseID: db_entreprise.ID}
		utilisateur.HashPassword(&db_utilisateur.Password)

		if err := db.Create(&db_utilisateur).Error; err != nil {
			panic(fmt.Sprintf("[WARNING] error while creating the root utilisateur: %v", err))
		}

		// used to save utilisateur id to create entreprise
		utilisateur_id = db_utilisateur.ID
	} else {

		// used to save utilisateur id to create entreprise
		utilisateur_id = root_utilisateur.ID
	}

	// add policy
	enforcer.AddGroupingPolicy(strconv.FormatUint(uint64(utilisateur_id), 10), os.Getenv("DEFAULT_ROOT"))

	// create utilisateur
	if check := db.Where("nom=?", os.Getenv("DEFAULT_USER")).Find(&utilisateur_role); check.RowsAffected == 0 && check.Error == nil {

		// create role utilisateur
		db_role := common.Role{Nom: os.Getenv("DEFAULT_USER"), EntrepriseID: db_entreprise.ID}

		if err := db.Create(&db_role).Error; err != nil {
			panic(fmt.Sprintf("[WARNING] error while creating the utilisateur role: %v", err))
		}
	}

	// add policy
	enforcer.AddGroupingPolicy(strconv.FormatUint(uint64(0), 10), os.Getenv("DEFAULT_USER"))

}

func AutoMigrateDatabase(db *gorm.DB, enforcer *casbin.Enforcer) {

	// create tables
	_auto_migrate_tables(db)

	// create root
	_create_root_utilisateur(db, enforcer)

}
