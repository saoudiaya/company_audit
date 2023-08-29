package membre

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Membre struct {
	ID            uint   `gorm:"primaryKey;autoIncrement;column:id;" json:"id"`
	UtilisateurID uint   `gorm:"column:utilisateur_id;primaryKey" json:"utilisateur_id"`
	AuditID       uint   `gorm:"column:audit_id;primaryKey" json:"audit_id"`
	Type          string `gorm:"column:type" json:"type"`
	Poste         string `gorm:"column:poste" json:"poste"`
	gorm.Model
}

// create new membre
func NewMembre(db *gorm.DB, membre Membre) (Membre, error) {
	return membre, db.Create(&membre).Error
}

// get all membres
func GetMembres(db *gorm.DB) (membres []Membre, err error) {
	return membres, db.Preload(clause.Associations).Find(&membres).Error
}

// check if membre exists
func CheckMembreExists(db *gorm.DB, id uint) bool {

	// init vars
	membre := &Membre{}

	// check if row exists
	check := db.Where("id=?", id).First(&membre)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get membre by id
func GetMembreById(db *gorm.DB, id uint) (membre Membre, err error) {
	return membre, db.Preload(clause.Associations).First(&membre, "id=?", id).Error
}

// search membres
func SearchMembres(db *gorm.DB, membre Membre) (membres []Membre, err error) {
	return membres, db.Where(&membre).Find(&membres).Error
}

// update membre
func UpdateMembre(db *gorm.DB, membre Membre) error {
	return db.Where("id=?", membre.ID).Updates(&membre).Error
}

// delete membre
func DeleteMembre(db *gorm.DB, membre_id uint) error {
	return db.Where("id=?", membre_id).Delete(&Membre{}).Error
}

// get membre by email
func GetMembreByEmail(db *gorm.DB, email string) (membre Membre, err error) {
	return membre, db.First(&membre, "email=?", email).Error
}

// compare two passwords
func ComparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}
