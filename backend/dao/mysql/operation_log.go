package mysql

import (
	"chess-room-backend/model"
)

func GetOperationLogList(adminID int64, module string) ([]model.OperationLog, error) {
	var logs []model.OperationLog
	db := DB.Order("created_at DESC")
	if adminID != 0 {
		db = db.Where("admin_id = ?", adminID)
	}
	if module != "" {
		db = db.Where("module = ?", module)
	}
	err := db.Find(&logs).Error
	return logs, err
}

func GetOperationLogByID(id int64) (*model.OperationLog, error) {
	var log model.OperationLog
	err := DB.Where("id = ?", id).First(&log).Error
	return &log, err
}

func CreateOperationLog(log *model.OperationLog) error {
	return DB.Create(log).Error
}

func DeleteOperationLog(id int64) error {
	return DB.Delete(&model.OperationLog{}, id).Error
}
