package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"chess-room-backend/pkg/jwt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(openID, phone, password, nickname, avatar string, gender int) (*model.User, string, error) {
	var user *model.User
	var err error

	if openID != "" {
		user, err = mysql.GetUserByOpenID(openID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, "", errno.New(errno.InternalError)
		}
	} else if phone != "" {
		user, err = mysql.GetUserByPhone(phone)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, "", errno.New(errno.InternalError)
		}
	}

	if err == gorm.ErrRecordNotFound || user == nil {
		var hashedPassword string
		if password != "" {
			hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				return nil, "", errno.New(errno.InternalError)
			}
			hashedPassword = string(hashed)
		}
		user = &model.User{
			OpenID:   openID,
			Phone:    phone,
			Password: hashedPassword,
			Nickname: nickname,
			Avatar:   avatar,
			Gender:   gender,
			Status:   1,
		}
		if err := mysql.CreateUser(user); err != nil {
			return nil, "", errno.New(errno.InternalError)
		}
	} else {
		if password != "" && user.Password != "" {
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
				return nil, "", errno.New(errno.InvalidPassword)
			}
		}

		if nickname != "" {
			user.Nickname = nickname
		}
		if avatar != "" {
			user.Avatar = avatar
		}
		if gender != 0 {
			user.Gender = gender
		}
		if err := mysql.UpdateUser(user); err != nil {
			return nil, "", errno.New(errno.InternalError)
		}
	}

	token, err := jwt.GenerateToken(user.ID, user.Nickname, 0, 0)
	if err != nil {
		return nil, "", errno.New(errno.InternalError)
	}

	return user, token, nil
}

func ChangePassword(userID int64, oldPassword, newPassword string) error {
	user, err := mysql.GetUserByID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errno.New(errno.UserNotFound)
		}
		return errno.New(errno.InternalError)
	}

	if user.Password == "" {
		return errno.New(errno.InvalidPassword)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errno.New(errno.InvalidPassword)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errno.New(errno.InternalError)
	}

	user.Password = string(hashedPassword)
	if err := mysql.UpdateUser(user); err != nil {
		return errno.New(errno.InternalError)
	}

	return nil
}

func GetUserByID(id int64) (*model.User, error) {
	user, err := mysql.GetUserByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.UserNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return user, nil
}

func UpdateUser(user *model.User) error {
	if err := mysql.UpdateUser(user); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func GetUserList(users *[]model.User) error {
	err := mysql.GetUserList(users)
	if err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func UpdateUserStatus(userID int64, status int) (*model.User, error) {
	user, err := mysql.GetUserByID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.UserNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	user.Status = status
	if err := mysql.UpdateUser(user); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return user, nil
}
