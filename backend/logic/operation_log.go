package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetOperationLogList(adminID, module string, page, pageSize int) ([]model.OperationLog, int64, error) {
	adminIDInt := int64(0)
	if adminID != "" {
		var err error
		adminIDInt, err = strconv.ParseInt(adminID, 10, 64)
		if err != nil {
			return nil, 0, errno.New(errno.BadRequest)
		}
	}

	logs, total, err := mysql.GetOperationLogList(adminIDInt, module, page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return logs, total, nil
}

func GetOperationLogByID(id string) (*model.OperationLog, error) {
	logID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	log, err := mysql.GetOperationLogByID(logID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.NotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return log, nil
}

func CreateOperationLog(log *model.OperationLog) error {
	if err := mysql.CreateOperationLog(log); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func DeleteOperationLog(id string) error {
	logID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	if err := mysql.DeleteOperationLog(logID); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchDeleteOperationLog(ids []string) error {
	var idInts []int64
	for _, id := range ids {
		intID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errno.New(errno.BadRequest)
		}
		idInts = append(idInts, intID)
	}
	if err := mysql.BatchDeleteOperationLog(idInts); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchUpdateOperationLog(logs []model.OperationLog) error {
	if err := mysql.BatchUpdateOperationLog(logs); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}
