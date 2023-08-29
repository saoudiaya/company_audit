package app_permission

import (
	"pfe/api/app/common"

	"gorm.io/gorm"
)

type CasbinRule struct {
	ID uint   `json:"id"`
	V0 string `json:"role"`
	V1 string `json:"object"`
	V2 string `json:"action"`
}

// get permission

func GetPermissions(db *gorm.DB) (permissions []CasbinRule, err error) {
	return permissions, db.Table("casbin_rule").Find(&permissions, "ptype=?", "p").Error
}

// get permission by id

func GetPermissionByID(db *gorm.DB, id uint) (permission CasbinRule, err error) {
	return permission, db.Table("casbin_rule").Where("id = ? AND ptype = 'p'", id).Find(&permission).Error
}

// search permissions
func SearchPermissions(db *gorm.DB, permission CasbinRule) (permissions []CasbinRule, err error) {
	return permissions, db.Where(&permission).Find(&permissions).Error
}

func CheckRolePermissionExists(db *gorm.DB, role string) (role_exists common.Role, err error) {
	return role_exists, db.Table("roles").Where("nom=?", role).Error
}

// check if the role exists

func CheckRoleExists(db *gorm.DB, nom string) (role common.Role, err error) {
	return role, db.Table("roles").Where("nom = ?", nom).Find(&role).Error
}
