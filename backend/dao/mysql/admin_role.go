package mysql

import (
	"chess-room-backend/model"
	"errors"

	"github.com/jinzhu/gorm"
)

func GetAdminRoleByID(id int64) (*model.AdminRole, error) {
	var role model.AdminRole
	if err := DB.Where("id = ?", id).First(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, errors.New("获取角色失败")
	}
	return &role, nil
}

func GetAdminRoleList() ([]model.AdminRole, error) {
	var roles []model.AdminRole
	if err := DB.Where("status = 1").Order("level ASC").Find(&roles).Error; err != nil {
		return nil, errors.New("获取角色列表失败")
	}
	return roles, nil
}

func CreateAdminRole(role *model.AdminRole) error {
	if err := DB.Create(role).Error; err != nil {
		return errors.New("创建角色失败")
	}
	return nil
}

func UpdateAdminRole(id int64, role *model.AdminRole) error {
	if err := DB.Model(&model.AdminRole{}).Where("id = ?", id).Updates(role).Error; err != nil {
		return errors.New("更新角色失败")
	}
	return nil
}

func DeleteAdminRole(id int64) error {
	if err := DB.Model(&model.AdminRole{}).Where("id = ?", id).Update("status", 0).Error; err != nil {
		return errors.New("删除角色失败")
	}
	return nil
}

func GetRolesByLevel(level int) ([]model.AdminRole, error) {
	var roles []model.AdminRole
	if err := DB.Where("level >= ? AND status = 1", level).Order("level ASC").Find(&roles).Error; err != nil {
		return nil, errors.New("获取角色列表失败")
	}
	return roles, nil
}
