package mysql

import (
	"chess-room-backend/model"
	"time"
)

func GetOrderList(userID, roomID, status int, startTime, endTime time.Time) ([]model.Order, error) {
	var orders []model.Order
	db := DB.Preload("User").Preload("Room").Order("created_at DESC")
	if userID > 0 {
		db = db.Where("user_id = ?", userID)
	}
	if roomID != 0 {
		db = db.Where("room_id = ?", roomID)
	}
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	if !startTime.IsZero() {
		db = db.Where("start_time >= ?", startTime)
	}
	if !endTime.IsZero() {
		db = db.Where("end_time <= ?", endTime)
	}
	err := db.Find(&orders).Error
	return orders, err
}

func GetOrderByOrderNo(orderNo string) (*model.Order, error) {
	var order model.Order
	err := DB.Preload("User").Preload("Room").Where("order_no = ?", orderNo).First(&order).Error
	return &order, err
}

func GetOrderByID(id int64) (*model.Order, error) {
	var order model.Order
	err := DB.Preload("User").Preload("Room").Where("id = ?", id).First(&order).Error
	return &order, err
}

func CreateOrder(order *model.Order) error {
	return DB.Create(order).Error
}

func UpdateOrder(order *model.Order) error {
	return DB.Save(order).Error
}

func DeleteOrder(id int64) error {
	return DB.Delete(&model.Order{}, id).Error
}