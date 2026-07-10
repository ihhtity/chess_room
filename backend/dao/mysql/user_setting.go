package mysql

import (
	"chess-room-backend/model"

	"github.com/jinzhu/gorm"
)

func GetUserSettingByUserID(userID int64) (*model.UserSetting, error) {
	var setting model.UserSetting
	err := DB.Where("user_id = ?", userID).First(&setting).Error
	return &setting, err
}

func CreateUserSetting(setting *model.UserSetting) error {
	return DB.Create(setting).Error
}

func UpdateUserSetting(setting *model.UserSetting) error {
	return DB.Save(setting).Error
}

func UpsertUserSetting(userID int64, updates map[string]interface{}) error {
	var setting model.UserSetting
	err := DB.Where("user_id = ?", userID).First(&setting).Error

	if err == gorm.ErrRecordNotFound {
		setting.UserID = userID
		for k, v := range updates {
			switch k {
			case "notifications":
				setting.Notifications = v.(int)
			case "sound":
				setting.Sound = v.(int)
			case "vibrate":
				setting.Vibrate = v.(int)
			case "language":
				setting.Language = v.(string)
			case "theme":
				setting.Theme = v.(string)
			case "allow_push":
				setting.AllowPush = v.(int)
			case "allow_marketing":
				setting.AllowMarketing = v.(int)
			}
		}
		return DB.Create(&setting).Error
	}

	if err != nil {
		return err
	}

	for k, v := range updates {
		switch k {
		case "notifications":
			setting.Notifications = v.(int)
		case "sound":
			setting.Sound = v.(int)
		case "vibrate":
			setting.Vibrate = v.(int)
		case "language":
			setting.Language = v.(string)
		case "theme":
			setting.Theme = v.(string)
		case "allow_push":
			setting.AllowPush = v.(int)
		case "allow_marketing":
			setting.AllowMarketing = v.(int)
		}
	}

	return DB.Save(&setting).Error
}
