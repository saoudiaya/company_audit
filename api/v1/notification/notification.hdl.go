package notification

import (
	"net/http"
	"os"
	"pfe/api/app/utilisateur"
	"regexp"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

// create new notification
func (db Database) NewNotification(ctx *gin.Context) {

	// init vars
	var notification NotificationRsp
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(notification.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check utilisateur exists
	if _, err := utilisateur.GetUtilisateurById(db.DB, notification.UtilisateurID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this utilisateur does not exist"})
		return
	}

	// init new notification
	new_notification := Notification{
		ID:            0,
		Nom:           notification.Nom,
		Description:   notification.Description,
		UtilisateurID: notification.UtilisateurID,
	}

	// create notification
	if _, err := NewNotification(db.DB, new_notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created"})

}

// get all notifications from database
func (db Database) GetNotifications(ctx *gin.Context) {

	// get notification
	notification, err := GetNotifications(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notification)
}

// get notification by id

func (db Database) GetNotificationById(ctx *gin.Context) {

	// get id value from path
	notification_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get notification by id

	notification, err := GetNotificationById(db.DB, uint(notification_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notification)
}

// search notification from database
func (db Database) SearchNotification(ctx *gin.Context) {

	// init vars
	var notification Notification

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// check notification exists
	if exists := CheckNotificationExists(db.DB, notification.ID); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "notification  does not exist"})
		return
	}

	// search notification from database
	notifications, err := SearchNotification(db.DB, notification)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notifications)
}

// update notification

func (db Database) UpdateNotification(ctx *gin.Context) {

	// init vars
	var notification Notification
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))

	// unmarshal sent json
	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id value from path
	notification_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check values validity
	if empty_reg.MatchString(notification.Nom) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// ignore key attributs
	notification.ID = uint(notification_id)

	// update notification
	if err = UpdateNotification(db.DB, notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (db Database) DeleteNotification(ctx *gin.Context) {

	// get id from path
	notification_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete notification
	if err = DeleteNotification(db.DB, uint(notification_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
