package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"chess-room-backend/pkg/jwt"
	"chess-room-backend/pkg/utils"

	"github.com/jinzhu/gorm"
)

func AdminLogin(username, password string) (*model.Admin, string, error) {
	admin, err := mysql.GetAdminByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "", errno.New(errno.AdminNotFound)
		}
		return nil, "", errno.New(errno.InternalError)
	}

	if !utils.CheckPasswordHash(password, admin.Password) {
		return nil, "", errno.New(errno.PasswordError)
	}

	if admin.Status != 1 {
		return nil, "", errno.New(errno.AdminDisabled)
	}

	token, err := jwt.GenerateToken(admin.ID, admin.Username, admin.Role)
	if err != nil {
		return nil, "", errno.New(errno.InternalError)
	}

	return admin, token, nil
}

func GetAdminByID(id int64) (*model.Admin, error) {
	admin, err := mysql.GetAdminByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.AdminNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return admin, nil
}

func UpdateAdminProfile(id int64, username, realname string) (*model.Admin, error) {
	admin, err := mysql.GetAdminByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.AdminNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	if username != "" {
		admin.Username = username
	}
	if realname != "" {
		admin.Realname = realname
	}

	if err := mysql.UpdateAdmin(admin); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return admin, nil
}

func AdminChangePassword(id int64, oldPassword, newPassword string) error {
	admin, err := mysql.GetAdminByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errno.New(errno.AdminNotFound)
		}
		return errno.New(errno.InternalError)
	}

	if !utils.CheckPasswordHash(oldPassword, admin.Password) {
		return errno.New(errno.PasswordError)
	}

	newPasswordHash, err := utils.HashPassword(newPassword)
	if err != nil {
		return errno.New(errno.InternalError)
	}

	admin.Password = newPasswordHash

	if err := mysql.UpdateAdmin(admin); err != nil {
		return errno.New(errno.InternalError)
	}

	return nil
}
