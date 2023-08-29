package commentaire

import (
	"pfe/api/app/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Commentaire struct {
	ID            uint               `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom           string             `gorm:"column:nom;not null;" json:"nom"`
	Description   string             `gorm:"column:nom;not null;" json:"description"`
	UtilisateurID uint               `gorm:"foreignKey:OwningUtilisateurID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"utilisateur_id"`
	Utilisateur   common.Utilisateur `gorm:"foreignKey:UtilisateurID;references:ID" json:"utilisateur"`
	ManagedBy     uint               `gorm:"column:managed_by;" json:"managed_by"`
	CreatedBy     uint               `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// create new commentaire
func NewCommentaire(db *gorm.DB, commentaire Commentaire) (Commentaire, error) {
	return commentaire, db.Create(&commentaire).Error
}

// get all commentaires
func GetCommentaires(db *gorm.DB) (commentaire []Commentaire, err error) {
	return commentaire, db.Preload(clause.Associations).Find(&commentaire).Error
}

// check commentaire exists
func CheckCommentaireExists(db *gorm.DB, id uint) bool {

	// init vars
	commentaire := &Commentaire{}

	// check if row exists
	check := db.Where("id=?", id).First(&commentaire)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get commentaire by id
func GetCommentaireById(db *gorm.DB, id uint) (commentaire Commentaire, err error) {
	return commentaire, db.Preload(clause.Associations).First(&commentaire, "id=?", id).Error
}

// search commentaires
func SearchCommentaires(db *gorm.DB, commentaire Commentaire) (commentaires []Commentaire, err error) {
	return commentaires, db.Where(&commentaire).Find(&commentaires).Error
}

// update commentaire
func UpdateCommentaire(db *gorm.DB, commentaire Commentaire) error {
	return db.Where("id=?", commentaire.ID).Updates(&commentaire).Error
}

// delete commentaire
func DeleteCommentaire(db *gorm.DB, commentaire_id uint) error {
	return db.Where("id=?", commentaire_id).Delete(&Commentaire{}).Error
}
