package mysql

import (
	"chess-room-backend/model"
)

func GetTimeSlotList(status, typeID, page, pageSize int) ([]model.TimeSlot, int64, error) {
	var slots []model.TimeSlot
	var total int64
	db := DB.Model(&model.TimeSlot{}).Order("sort_order ASC")
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	if typeID != 0 {
		db = db.Where("type_id = ?", typeID)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB.Order("sort_order ASC")
		if status != 0 {
			db = db.Where("status = ?", status)
		}
		if typeID != 0 {
			db = db.Where("type_id = ?", typeID)
		}
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&slots).Error
	return slots, total, err
}

func GetTimeSlotByID(id int64) (*model.TimeSlot, error) {
	var slot model.TimeSlot
	err := DB.Where("id = ?", id).First(&slot).Error
	return &slot, err
}

func CreateTimeSlot(slot *model.TimeSlot) error {
	return DB.Create(slot).Error
}

func UpdateTimeSlot(slot *model.TimeSlot) error {
	return DB.Save(slot).Error
}

func DeleteTimeSlot(id int64) error {
	return DB.Delete(&model.TimeSlot{}, id).Error
}

func BatchDeleteTimeSlot(ids []int64) error {
	return DB.Where("id IN (?)", ids).Delete(&model.TimeSlot{}).Error
}