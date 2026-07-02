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
