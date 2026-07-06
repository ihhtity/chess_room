package mysql

import (
	"chess-room-backend/model"
)

func GetRechargeRecordList(userID int64, status, page, pageSize int) ([]model.RechargeRecord, int64, error) {
	var records []model.RechargeRecord
	var total int64
	db := DB.Model(&model.RechargeRecord{}).Order("created_at DESC")
	if userID != 0 {
		db = db.Where("user_id = ?", userID)
	}
	if status != -1 {
		db = db.Where("status = ?", status)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB.Order("created_at DESC")
		if userID != 0 {
			db = db.Where("user_id = ?", userID)
		}
		if status != -1 {
			db = db.Where("status = ?", status)
		}
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&records).Error
	return records, total, err
}

func GetRechargeRecordByID(id int64) (*model.RechargeRecord, error) {
	var record model.RechargeRecord
	err := DB.Where("id = ?", id).First(&record).Error
	return &record, err
}

func CreateRechargeRecord(record *model.RechargeRecord) error {
	return DB.Create(record).Error
}

func UpdateRechargeRecordByID(id int64, data map[string]interface{}) error {
	return DB.Model(&model.RechargeRecord{}).Where("id = ?", id).Updates(data).Error
}

func DeleteRechargeRecord(id int64) error {
	return DB.Delete(&model.RechargeRecord{}, id).Error
}

func BatchDeleteRechargeRecord(ids []int64) error {
	return DB.Where("id IN (?)", ids).Delete(&model.RechargeRecord{}).Error
}
