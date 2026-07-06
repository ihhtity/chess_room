package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetNotificationList(userID, notificationType, readStatus string, page, pageSize int) ([]model.Notification, int64, error) {
	userIDInt := int64(0)
	if userID != "" {
		var err error
		userIDInt, err = strconv.ParseInt(userID, 10, 64)
		if err != nil {
			return nil, 0, errno.New(errno.BadRequest)
		}
	}

	notificationTypeInt := 0
	if notificationType != "" {
		var err error
		notificationTypeInt, err = strconv.Atoi(notificationType)
		if err != nil {
			return nil, 0, errno.New(errno.BadRequest)
		}
	}

	readStatusInt := -1
	if readStatus != "" {
		var err error
		readStatusInt, err = strconv.Atoi(readStatus)
		if err != nil {
			return nil, 0, errno.New(errno.BadRequest)
		}
	}

	notifications, total, err := mysql.GetNotificationList(userIDInt, notificationTypeInt, readStatusInt, page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return notifications, total, nil
}

func GetNotificationByID(id string) (*model.Notification, error) {
	notificationID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	notification, err := mysql.GetNotificationByID(notificationID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.NotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return notification, nil
}

func CreateNotification(notification *model.Notification) error {
	if err := mysql.CreateNotification(notification); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func UpdateNotification(id string, data map[string]interface{}) (*model.Notification, error) {
	notificationID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}

	_, err = mysql.GetNotificationByID(notificationID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.NotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	if err := mysql.UpdateNotificationByID(notificationID, data); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return mysql.GetNotificationByID(notificationID)
}

func DeleteNotification(id string) error {
	notificationID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	if err := mysql.DeleteNotification(notificationID); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchDeleteNotification(ids []string) error {
	var idInts []int64
	for _, id := range ids {
		intID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errno.New(errno.BadRequest)
		}
		idInts = append(idInts, intID)
	}
	if err := mysql.BatchDeleteNotification(idInts); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchUpdateNotification(reqs []struct {
	ID         int64 `json:"id"`
	ReadStatus int   `json:"read_status"`
}) error {
	for _, req := range reqs {
		notification, err := mysql.GetNotificationByID(req.ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errno.New(errno.NotFound)
			}
			return errno.New(errno.InternalError)
		}

		notification.ReadStatus = req.ReadStatus

		if err := mysql.UpdateNotificationByID(req.ID, map[string]interface{}{"read_status": req.ReadStatus}); err != nil {
			return errno.New(errno.InternalError)
		}
	}
	return nil
}

func MarkAllNotificationAsRead(userID string) error {
	userIDInt, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	if err := mysql.BatchUpdateNotificationStatus(userIDInt, 1); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}
