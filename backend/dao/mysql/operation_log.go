package mysql

import (
	"chess-room-backend/model"
)

func GetOperationLogList(adminID int64, module string, page, pageSize int) ([]model.OperationLog, int64, error) {
	var logs []model.OperationLog
	var total int64
	db := DB.Model(&model.OperationLog{}).Order("created_at DESC")
	if adminID != 0 {
		db = db.Where("admin_id = ?", adminID)
	}
	if module != "" {
		db = db.Where("module = ?", module)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB.Order("created_at DESC")
		if adminID != 0 {
			db = db.Where("admin_id = ?", adminID)
		}
		if module != "" {
			db = db.Where("module = ?", module)
		}
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&logs).Error
	return logs, total, err
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

func BatchDeleteOperationLog(ids []int64) error {
	return DB.Where("id IN (?)", ids).Delete(&model.OperationLog{}).Error
}

func BatchUpdateOperationLog(logs []model.OperationLog) error {
	for _, log := range logs {
		if err := DB.Model(&model.OperationLog{}).Where("id = ?", log.ID).Updates(log).Error; err != nil {
			return err
		}
	}
	return nil
}
