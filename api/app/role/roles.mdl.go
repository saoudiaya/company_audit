package role

import (
	"pfe/api/app/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// create new role

func NewRole(db *gorm.DB, role common.Role) error {
	return db.Create(&role).Error
}

// get all roles

func GetRoles(db *gorm.DB) (role []common.Role, err error) {
	return role, db.Preload(clause.Associations).Find(&role).Error
}

// get role by nom

/*func GetRoleByNom(db *gorm.DB, nom string) (common.Role, error) {
	role := common.Role{}
	return role, db.Where("nom=?", nom).First(&role).Error
}*/

// get role by id

func GetRoleByID(db *gorm.DB, id uint) (common.Role, error) {
	role := common.Role{}
	return role, db.Preload(clause.Associations).Where("id=?", id).First(&role).Error
}

// search role

func SearchRoles(db *gorm.DB, role common.Role) (roles []common.Role, err error) {
	return roles, db.Where(&role).Find(&roles).Error
}

// update role

func UpdateRole(db *gorm.DB, role common.Role) error {
	return db.Where("id=?", role.ID).Updates(&role).Error
}

// delete role

func DeleteRole(db *gorm.DB, role_id uint) error {
	return db.Where("id=?", role_id).Delete(&common.Role{}).Error
}
