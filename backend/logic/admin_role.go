package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"

	"github.com/jinzhu/gorm"
)

func GetAdminRoleList() ([]model.AdminRole, error) {
	roles, err := mysql.GetAdminRoleList()
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return roles, nil
}

func GetAdminRoleByID(id int64) (*model.AdminRole, error) {
	role, err := mysql.GetAdminRoleByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.NotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return role, nil
}

func CreateAdminRole(currentAdminRoleID int64, role *model.AdminRole) error {
	currentRole, err := mysql.GetAdminRoleByID(currentAdminRoleID)
	if err != nil {
		return errno.New(errno.InternalError)
	}

	if role.Level <= currentRole.Level {
		return errno.NewWithMessage(errno.Forbidden, "只能创建层级低于自己的角色")
	}

	if err := mysql.CreateAdminRole(role); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func UpdateAdminRole(currentAdminRoleID int64, id int64, role *model.AdminRole) error {
	targetRole, err := mysql.GetAdminRoleByID(id)
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
		return errno.NewWithMessage(errno.Forbidden, "只能修改层级低于自己的角色")
	}

	if role.Level != 0 && role.Level <= currentRole.Level {
		return errno.NewWithMessage(errno.Forbidden, "只能将角色层级设置为低于自己的层级")
	}

	if err := mysql.UpdateAdminRole(id, role); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func DeleteAdminRole(currentAdminRoleID int64, id int64) error {
	targetRole, err := mysql.GetAdminRoleByID(id)
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
		return errno.NewWithMessage(errno.Forbidden, "只能删除层级低于自己的角色")
	}

	if err := mysql.DeleteAdminRole(id); err != nil {
		return errno.New(errno.InternalError)
	}

	if err := mysql.DeleteRolePermissionsByRoleID(id); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func GetRolesByLevel(level int) ([]model.AdminRole, error) {
	roles, err := mysql.GetRolesByLevel(level)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return roles, nil
}
