package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"

	"github.com/jinzhu/gorm"
)

func GetAdminRoleList(name string, status int, page, pageSize int) ([]model.AdminRole, int64, error) {
	roles, total, err := mysql.GetAdminRoleListFiltered(name, status, page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return roles, total, nil
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

func BatchUpdateAdminRole(currentAdminRoleID int64, reqs []struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Level       int    `json:"level"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}) error {
	for _, req := range reqs {
		targetRole, err := mysql.GetAdminRoleByID(req.ID)
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

		if req.Name != "" {
			targetRole.Name = req.Name
		}
		if req.Level != 0 && req.Level <= currentRole.Level {
			return errno.NewWithMessage(errno.Forbidden, "只能将角色层级设置为低于自己的层级")
		}
		if req.Level != 0 {
			targetRole.Level = req.Level
		}
		if req.Description != "" {
			targetRole.Description = req.Description
		}
		if req.Status >= 0 {
			targetRole.Status = req.Status
		}

		if err := mysql.UpdateAdminRole(req.ID, targetRole); err != nil {
			return errno.New(errno.InternalError)
		}
	}

	return nil
}
