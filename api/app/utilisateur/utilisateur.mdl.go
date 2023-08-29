package utilisateur

import (
	"pfe/api/app/common"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// hash password
func HashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

// create new utilisateur
func NewUtilisateur(db *gorm.DB, utilisateur common.Utilisateur) (common.Utilisateur, error) {
	return utilisateur, db.Create(&utilisateur).Error
}

// get all utilisateurs
func GetUtilisateurs(db *gorm.DB) (utilisateurs []common.Utilisateur, err error) {
	return utilisateurs, db.Preload(clause.Associations).Find(&utilisateurs).Error
}

// check if utilisateur exists

// get utilisateur by id
func GetUtilisateurById(db *gorm.DB, id uint) (utilisateur common.Utilisateur, err error) {
	return utilisateur, db.Preload(clause.Associations).First(&utilisateur, "id=?", id).Error
}

// search utilisateurs
func SearchUtilisateurs(db *gorm.DB, utilisateur common.Utilisateur) (utilisateurs []common.Utilisateur, err error) {
	return utilisateurs, db.Where(&utilisateur).Find(&utilisateurs).Error
}

// update utilisateur
func UpdateUtilisateur(db *gorm.DB, utilisateur common.Utilisateur) error {
	return db.Where("id=?", utilisateur.ID).Updates(&utilisateur).Error
}

// delete utilisateur
func DeleteUtilisateur(db *gorm.DB, utilisateur_id uint) error {
	return db.Where("id=?", utilisateur_id).Delete(&common.Utilisateur{}).Error
}

// get utilisateur by email
func GetUtilisateurByEmail(db *gorm.DB, email string) (utilisateur common.Utilisateur, err error) {
	return utilisateur, db.Preload(clause.Associations).First(&utilisateur, "email=?", email).Error
}

// compare two passwords
func ComparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}
