package audit

import (
	"pfe/api/app/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// create new audit
func NewAudit(db *gorm.DB, audit common.Audit) (common.Audit, error) {
	return audit, db.Create(&audit).Error
}

// get all audit
func GetAudits(db *gorm.DB) (audit []common.Audit, err error) {
	return audit, db.Preload(clause.Associations).Find(&audit).Error
}

// check if audit exists
func CheckAuditExists(db *gorm.DB, id uint) bool {

	// init vars
	audit := &common.Audit{}

	// check if row exists
	check := db.Where("id=?", id).First(&audit)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// get audit by id
func GetAuditById(db *gorm.DB, id uint) (audit common.Audit, err error) {
	return audit, db.Preload(clause.Associations).First(&audit, "id=?", id).Error
}

// search audit
func SearchAudit(db *gorm.DB, audit common.Audit) (audits []common.Audit, err error) {
	return audits, db.Where(&audit).Find(&audits).Error
}

// update audit
func UpdateAudit(db *gorm.DB, audit common.Audit) error {
	return db.Where("id=?", audit.ID).Updates(&audit).Error
}

// delete audit
func DeleteAudit(db *gorm.DB, audit_id uint) error {
	return db.Where("id=?", audit_id).Delete(&common.Audit{}).Error
}
