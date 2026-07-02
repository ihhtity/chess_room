package mysql

import (
	"chess-room-backend/model"
)

func CreateRechargeRecord(record *model.RechargeRecord) error {
	return DB.Create(record).Error
}

func GetRechargeRecords(userID int64) ([]model.RechargeRecord, error) {
	var records []model.RechargeRecord
	err := DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&records).Error
	return records, err
}