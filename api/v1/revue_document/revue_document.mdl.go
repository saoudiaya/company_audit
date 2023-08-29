package revuedocument

import (
	"pfe/api/app/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RevueDocument struct {
	ID            uint               `gorm:"column:id;autoIncrement;primaryKey;unique" json:"id"`
	Titre         string             `gorm:"column:titre;not null" json:"titre"`
	Type          string             `gorm:"column:titre;not null" json:"type"`
	AuditID       uint               `gorm:"foreignKey:OwningAuditId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"audit_id"`
	Audit         common.Audit       `gorm:"foreignKey:AuditID;references:ID" json:"audit"`
	UtilisateurID uint               `gorm:"foreignKey:OwningUtilisateurId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"utilisateur_id"`
	Utilisateur   common.Utilisateur `gorm:"foreignKey:UtilisateurID;references:ID" json:"utilisateur"`
	Date          string             `gorm:"column:date;" json:"date"`
	gorm.Model
}

// create new revue document
func NewRevueDocument(db *gorm.DB, revuedocument RevueDocument) (RevueDocument, error) {
	return revuedocument, db.Create(&revuedocument).Error
}

// get all revue document
func GetRevueDocuments(db *gorm.DB) (revuedocument []RevueDocument, err error) {
	return revuedocument, db.Preload(clause.Associations).Find(&revuedocument).Error
}

// check if revue document exists
func CheckRevueDocumentExists(db *gorm.DB, id uint) bool {

	// init vars
	revuedocument := &RevueDocument{}

	// check if row exists
	check := db.Where("id=?", id).First(&revuedocument)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get revue document by id
func GetRevueDocumentById(db *gorm.DB, id uint) (revuedocument RevueDocument, err error) {
	return revuedocument, db.Preload(clause.Associations).First(&revuedocument, "id=?", id).Error
}

// search revue document
func SearchRevueDocument(db *gorm.DB, revuedocument RevueDocument) (revuedocuments []RevueDocument, err error) {
	return revuedocuments, db.Where(&revuedocument).Find(&revuedocuments).Error
}

// update revue document
func UpdateRevueDocument(db *gorm.DB, revuedocument RevueDocument) error {
	return db.Where("id=?", revuedocument.ID).Updates(&revuedocument).Error
}

// delete revue document
func DeleteRevueDocument(db *gorm.DB, revuedocument_id uint) error {
	return db.Where("id=?", revuedocument_id).Delete(&RevueDocument{}).Error
}
