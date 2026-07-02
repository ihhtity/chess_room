package mysql

import (
	"chess-room-backend/model"
)

func GetTimeSlotList(status, typeID int) ([]model.TimeSlot, error) {
	var slots []model.TimeSlot
	db := DB.Order("sort_order ASC")
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	if typeID != 0 {
		db = db.Where("type_id = ?", typeID)
	}
	err := db.Find(&slots).Error
	return slots, err
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