package reponse

import (
	"pfe/api/v1/commentaire"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Reponse struct {
	ID            uint                    `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom           string                  `gorm:"column:nom;not null;" json:"nom"`
	Description   string                  `gorm:"column:nom;not null;" json:"description"`
	CommentaireID uint                    `gorm:"foreignKey:OwningCommentaireID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"commentaire_id"`
	Commentaire   commentaire.Commentaire `gorm:"foreignKey:CommentaireID;references:ID" json:"commentaire"`
	ManagedBy     uint                    `gorm:"column:managed_by;" json:"managed_by"`
	CreatedBy     uint                    `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

// create new reponse
func NewReponse(db *gorm.DB, reponse Reponse) (Reponse, error) {
	return reponse, db.Create(&reponse).Error
}

// get all reponses
func GetReponses(db *gorm.DB) (reponse []Reponse, err error) {
	return reponse, db.Preload(clause.Associations).Find(&reponse).Error
}

// check reponse exists
func CheckReponseExists(db *gorm.DB, id uint) bool {

	// init vars
	reponse := &Reponse{}

	// check if row exists
	check := db.Where("id=?", id).First(&reponse)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get reponse by id
func GetReponseById(db *gorm.DB, id uint) (reponse Reponse, err error) {
	return reponse, db.Preload(clause.Associations).First(&reponse, "id=?", id).Error
}

// search reponses
func SearchReponses(db *gorm.DB, reponse Reponse) (reponses []Reponse, err error) {
	return reponses, db.Where(&reponse).Find(&reponses).Error
}

// update reponse
func UpdateReponse(db *gorm.DB, reponse Reponse) error {
	return db.Where("id=?", reponse.ID).Updates(&reponse).Error
}

// delete reponse
func DeleteReponse(db *gorm.DB, reponse_id uint) error {
	return db.Where("id=?", reponse_id).Delete(&Reponse{}).Error
}
