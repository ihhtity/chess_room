package mysql

import (
	"chess-room-backend/model"
	"errors"

	"github.com/jinzhu/gorm"
)

func GetPermissionByID(id int64) (*model.Permission, error) {
	var permission model.Permission
	if err := DB.Where("id = ?", id).First(&permission).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, errors.New("获取权限失败")
	}
	return &permission, nil
}

func GetPermissionByCode(code string) (*model.Permission, error) {
	var permission model.Permission
	if err := DB.Where("code = ?", code).First(&permission).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, errors.New("获取权限失败")
	}
	return &permission, nil
}

func GetPermissionList() ([]model.Permission, error) {
	var permissions []model.Permission
	if err := DB.Order("sort_order ASC").Find(&permissions).Error; err != nil {
		return nil, errors.New("获取权限列表失败")
	}
	return permissions, nil
}

func GetPermissionListByGroup() ([]model.Permission, error) {
	var permissions []model.Permission
	if err := DB.Order("group ASC, sort_order ASC").Find(&permissions).Error; err != nil {
		return nil, errors.New("获取权限列表失败")
	}
	return permissions, nil
}

func CreatePermission(permission *model.Permission) error {
	if err := DB.Create(permission).Error; err != nil {
		return errors.New("创建权限失败")
	}
	return nil
}

func UpdatePermission(id int64, permission *model.Permission) error {
	if err := DB.Model(&model.Permission{}).Where("id = ?", id).Updates(permission).Error; err != nil {
		return errors.New("更新权限失败")
	}
	return nil
}

func DeletePermission(id int64) error {
	if err := DB.Delete(&model.Permission{}, id).Error; err != nil {
		return errors.New("删除权限失败")
	}
	return nil
}
