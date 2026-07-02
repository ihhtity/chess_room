package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetRechargeRecordList(userID, status string) ([]model.RechargeRecord, error) {
	userIDInt := int64(0)
	if userID != "" {
		var err error
		userIDInt, err = strconv.ParseInt(userID, 10, 64)
		if err != nil {
			return nil, errno.New(errno.BadRequest)
		}
	}

	statusInt := -1
	if status != "" {
		var err error
		statusInt, err = strconv.Atoi(status)
		if err != nil {
			return nil, errno.New(errno.BadRequest)
		}
	}

	records, err := mysql.GetRechargeRecordList(userIDInt, statusInt)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return records, nil
}

func GetRechargeRecordByID(id string) (*model.RechargeRecord, error) {
	recordID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	record, err := mysql.GetRechargeRecordByID(recordID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.NotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return record, nil
}

func CreateRechargeRecord(record *model.RechargeRecord) error {
	if err := mysql.CreateRechargeRecord(record); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func UpdateRechargeRecord(id string, data map[string]interface{}) (*model.RechargeRecord, error) {
	recordID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}

	_, err = mysql.GetRechargeRecordByID(recordID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.NotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	if err := mysql.UpdateRechargeRecordByID(recordID, data); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return mysql.GetRechargeRecordByID(recordID)
}

func DeleteRechargeRecord(id string) error {
	recordID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	if err := mysql.DeleteRechargeRecord(recordID); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}
