package mysql

import (
	"chess-room-backend/model"
)

func GetNotificationList(userID int64, notificationType, readStatus int) ([]model.Notification, error) {
	var notifications []model.Notification
	db := DB.Order("created_at DESC")
	if userID != 0 {
		db = db.Where("user_id = ?", userID)
	}
	if notificationType != 0 {
		db = db.Where("type = ?", notificationType)
	}
	if readStatus != -1 {
		db = db.Where("read_status = ?", readStatus)
	}
	err := db.Find(&notifications).Error
	return notifications, err
}

func GetNotificationByID(id int64) (*model.Notification, error) {
	var notification model.Notification
	err := DB.Where("id = ?", id).First(&notification).Error
	return &notification, err
}

func CreateNotification(notification *model.Notification) error {
	return DB.Create(notification).Error
}

func UpdateNotificationByID(id int64, data map[string]interface{}) error {
	return DB.Model(&model.Notification{}).Where("id = ?", id).Updates(data).Error
}

func DeleteNotification(id int64) error {
	return DB.Delete(&model.Notification{}, id).Error
}

func BatchUpdateNotificationStatus(userID int64, readStatus int) error {
	return DB.Model(&model.Notification{}).Where("user_id = ?", userID).Update("read_status", readStatus).Error
}
