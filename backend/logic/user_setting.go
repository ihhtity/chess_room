package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"

	"github.com/jinzhu/gorm"
)

func GetUserSetting(userID int64) (*model.UserSetting, error) {
	setting, err := mysql.GetUserSettingByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &model.UserSetting{
				UserID:         userID,
				Notifications:  1,
				Sound:          1,
				Vibrate:        1,
				Language:       "zh-CN",
				Theme:          "light",
				AllowPush:      1,
				AllowMarketing: 1,
			}, nil
		}
		return nil, errno.New(errno.InternalError)
	}
	return setting, nil
}

func UpdateUserSetting(userID int64, updates map[string]interface{}) (*model.UserSetting, error) {
	if err := mysql.UpsertUserSetting(userID, updates); err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return mysql.GetUserSettingByUserID(userID)
}

func ToggleSetting(userID int64, key string) (*model.UserSetting, error) {
	setting, err := mysql.GetUserSettingByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			setting = &model.UserSetting{
				UserID:         userID,
				Notifications:  1,
				Sound:          1,
				Vibrate:        1,
				Language:       "zh-CN",
				Theme:          "light",
				AllowPush:      1,
				AllowMarketing: 1,
			}
		} else {
			return nil, errno.New(errno.InternalError)
		}
	}

	switch key {
	case "notifications":
		setting.Notifications = 1 - setting.Notifications
	case "sound":
		setting.Sound = 1 - setting.Sound
	case "vibrate":
		setting.Vibrate = 1 - setting.Vibrate
	case "allow_push":
		setting.AllowPush = 1 - setting.AllowPush
	case "allow_marketing":
		setting.AllowMarketing = 1 - setting.AllowMarketing
	default:
		return nil, errno.NewWithMessage(errno.BadRequest, "无效的设置项")
	}

	if err := mysql.UpdateUserSetting(setting); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return setting, nil
}
