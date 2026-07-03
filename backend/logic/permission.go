package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"

	"github.com/jinzhu/gorm"
)

func GetPermissionList() ([]model.Permission, error) {
	permissions, err := mysql.GetPermissionList()
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return permissions, nil
}

func GetPermissionListByGroup() ([]model.Permission, error) {
	permissions, err := mysql.GetPermissionListByGroup()
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return permissions, nil
}

func GetPermissionByID(id int64) (*model.Permission, error) {
	permission, err := mysql.GetPermissionByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.NotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return permission, nil
}

func CreatePermission(permission *model.Permission) error {
	if err := mysql.CreatePermission(permission); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func UpdatePermission(id int64, permission *model.Permission) error {
	_, err := mysql.GetPermissionByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errno.New(errno.NotFound)
		}
		return errno.New(errno.InternalError)
	}

	if err := mysql.UpdatePermission(id, permission); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func DeletePermission(id int64) error {
	if err := mysql.DeletePermission(id); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func GetPermissionsByRoleID(roleID int64) ([]model.Permission, error) {
	permissions, err := mysql.GetPermissionsByRoleID(roleID)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return permissions, nil
}

func GetPermissionCodesByRoleID(roleID int64) ([]string, error) {
	codes, err := mysql.GetPermissionCodesByRoleID(roleID)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return codes, nil
}

func CheckPermission(roleID int64, permissionCode string) (bool, error) {
	codes, err := GetPermissionCodesByRoleID(roleID)
	if err != nil {
		return false, err
	}

	for _, code := range codes {
		if code == permissionCode {
			return true, nil
		}
	}
	return false, nil
}

func AssignPermissions(currentAdminRoleID int64, roleID int64, permissionIDs []int64) error {
	targetRole, err := mysql.GetAdminRoleByID(roleID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errno.New(errno.NotFound)
		}
		return errno.New(errno.InternalError)
	}

	currentRole, err := mysql.GetAdminRoleByID(currentAdminRoleID)
	if err != nil {
		return errno.New(errno.InternalError)
	}

	if targetRole.Level <= currentRole.Level {
		return errno.NewWithMessage(errno.Forbidden, "只能为层级低于自己的角色分配权限")
	}

	currentPermissionCodes, err := mysql.GetPermissionCodesByRoleID(currentAdminRoleID)
	if err != nil {
		return errno.New(errno.InternalError)
	}

	currentPermissionIDs := make(map[int64]bool)
	for _, code := range currentPermissionCodes {
		perm, err := mysql.GetPermissionByCode(code)
		if err == nil {
			currentPermissionIDs[perm.ID] = true
		}
	}

	for _, pid := range permissionIDs {
		if !currentPermissionIDs[pid] {
			return errno.NewWithMessage(errno.Forbidden, "只能分配自己拥有的权限")
		}
	}

	if err := mysql.BatchCreateRolePermissions(roleID, permissionIDs); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}