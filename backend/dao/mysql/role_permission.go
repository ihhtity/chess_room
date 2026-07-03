package mysql

import (
	"chess-room-backend/model"
	"errors"
)

func GetRolePermissions(roleID int64) ([]model.RolePermission, error) {
	var rolePermissions []model.RolePermission
	if err := DB.Where("role_id = ?", roleID).Find(&rolePermissions).Error; err != nil {
		return nil, errors.New("获取角色权限失败")
	}
	return rolePermissions, nil
}

func GetPermissionsByRoleID(roleID int64) ([]model.Permission, error) {
	var permissions []model.Permission
	if err := DB.Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Where("role_permissions.role_id = ?", roleID).
		Order("permissions.sort_order ASC").
		Find(&permissions).Error; err != nil {
		return nil, errors.New("获取角色权限失败")
	}
	return permissions, nil
}

func GetPermissionCodesByRoleID(roleID int64) ([]string, error) {
	var codes []string
	if err := DB.Model(&model.RolePermission{}).
		Select("permissions.code").
		Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
		Where("role_permissions.role_id = ?", roleID).
		Pluck("code", &codes).Error; err != nil {
		return nil, errors.New("获取角色权限编码失败")
	}
	return codes, nil
}

func BatchCreateRolePermissions(roleID int64, permissionIDs []int64) error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return errors.New("开启事务失败")
	}

	if err := tx.Where("role_id = ?", roleID).Delete(&model.RolePermission{}).Error; err != nil {
		tx.Rollback()
		return errors.New("清除角色权限失败")
	}

	for _, permissionID := range permissionIDs {
		rolePermission := model.RolePermission{
			RoleID:       roleID,
			PermissionID: permissionID,
		}
		if err := tx.Create(&rolePermission).Error; err != nil {
			tx.Rollback()
			return errors.New("创建角色权限关联失败")
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.New("提交事务失败")
	}

	return nil
}

func DeleteRolePermissionsByRoleID(roleID int64) error {
	if err := DB.Where("role_id = ?", roleID).Delete(&model.RolePermission{}).Error; err != nil {
		return errors.New("删除角色权限失败")
	}
	return nil
}
