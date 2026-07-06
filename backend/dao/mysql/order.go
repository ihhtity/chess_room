package mysql

import (
	"chess-room-backend/model"
	"time"
)

func GetOrderList(userID, roomID, status int, orderNo string, startTime, endTime time.Time, page, pageSize int) ([]model.Order, int64, error) {
	var orders []model.Order
	var total int64
	db := DB.Model(&model.Order{}).Preload("User").Preload("Room").Order("created_at DESC")
	if userID > 0 {
		db = db.Where("user_id = ?", userID)
	}
	if roomID != 0 {
		db = db.Where("room_id = ?", roomID)
	}
	if status >= 0 {
		db = db.Where("status = ?", status)
	}
	if orderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+orderNo+"%")
	}
	if !startTime.IsZero() {
		db = db.Where("start_time >= ?", startTime)
	}
	if !endTime.IsZero() {
		db = db.Where("end_time <= ?", endTime)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&orders).Error
	return orders, total, err
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

func BatchDeleteOrder(ids []int64) error {
	return DB.Where("id IN (?)", ids).Delete(&model.Order{}).Error
}
