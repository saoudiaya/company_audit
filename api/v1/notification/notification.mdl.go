package notification

import (
	"pfe/api/app/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Notification struct {
	ID            uint               `gorm:"column:id;autoIncrement;primaryKey" json:"id"`
	Nom           string             `gorm:"column:nom;not null" json:"nom"`
	Description   string             `gorm:"column:description;" json:"description"`
	UtilisateurID uint               `gorm:"foreignKey:OwningUtilisateurId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"utilisateur_id"`
	Utilisateur   common.Utilisateur `gorm:"foreignKey:UtilisateurID" json:"utilisateur"`
	gorm.Model
}

// create new notification
func NewNotification(db *gorm.DB, notification Notification) (Notification, error) {
	return notification, db.Create(&notification).Error
}

// get all notification
func GetNotifications(db *gorm.DB) (notification []Notification, err error) {
	return notification, db.Preload(clause.Associations).Find(&notification).Error
}

// check if notification exists
func CheckNotificationExists(db *gorm.DB, id uint) bool {

	// init vars
	notification := &Notification{}

	// check if row exists
	check := db.Where("id=?", id).First(&notification)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get notification by id
func GetNotificationById(db *gorm.DB, id uint) (notification Notification, err error) {
	return notification, db.Preload(clause.Associations).First(&notification, "id=?", id).Error
}

// search notification
func SearchNotification(db *gorm.DB, notification Notification) (notifications []Notification, err error) {
	return notifications, db.Where(&notification).Find(&notifications).Error
}

// update notification
func UpdateNotification(db *gorm.DB, notification Notification) error {
	return db.Where("id=?", notification.ID).Updates(&notification).Error
}

// delete notification
func DeleteNotification(db *gorm.DB, notification_id uint) error {
	return db.Where("id=?", notification_id).Delete(&Notification{}).Error
}
