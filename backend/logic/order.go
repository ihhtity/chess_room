package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"chess-room-backend/pkg/utils"
	"time"

	"github.com/jinzhu/gorm"
)

func GetOrderList(userID int64, roomID, status int, startTime, endTime time.Time) ([]model.Order, error) {
	orders, err := mysql.GetOrderList(int(userID), roomID, status, startTime, endTime)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return orders, nil
}

func GetOrderByID(id int64) (*model.Order, error) {
	order, err := mysql.GetOrderByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.OrderNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return order, nil
}

func CreateOrder(userID, roomID int64, startTime, endTime time.Time, remark string) (*model.Order, error) {
	room, err := mysql.GetRoomByID(roomID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.RoomNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	if room.Status != 1 {
		return nil, errno.New(errno.RoomDisabled)
	}

	orders, err := mysql.GetOrderList(0, int(roomID), 0, startTime, endTime)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errno.New(errno.InternalError)
	}

	if len(orders) > 0 {
		return nil, errno.New(errno.RoomOccupied)
	}

	duration := utils.CalculateDurationMinutes(startTime, endTime)
	price := room.Type.BasePrice * float64(duration) / 60

	order := &model.Order{
		OrderNo:     utils.GenerateOrderNo(),
		UserID:      userID,
		RoomID:      roomID,
		StartTime:   startTime,
		EndTime:     endTime,
		Duration:    duration,
		Status:      0,
		TotalAmount: price,
		PaidAmount:  0,
		Remark:      remark,
	}

	if err := mysql.CreateOrder(order); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return order, nil
}

func CancelOrder(userID, orderID int64) error {
	order, err := mysql.GetOrderByID(orderID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errno.New(errno.OrderNotFound)
		}
		return errno.New(errno.InternalError)
	}

	if order.UserID != userID {
		return errno.New(errno.Forbidden)
	}

	if order.Status != 0 {
		return errno.New(errno.OrderStatusError)
	}

	order.Status = 3
	now := time.Now()
	order.CancelTime = &now

	if err := mysql.UpdateOrder(order); err != nil {
		return errno.New(errno.OrderCancelFailed)
	}

	return nil
}

func ConfirmOrder(orderID int64) (*model.Order, error) {
	order, err := mysql.GetOrderByID(orderID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.OrderNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	if order.Status != 0 {
		return nil, errno.New(errno.OrderStatusError)
	}

	order.Status = 1

	if err := mysql.UpdateOrder(order); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return order, nil
}

func CompleteOrder(orderID int64) (*model.Order, error) {
	order, err := mysql.GetOrderByID(orderID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.OrderNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	if order.Status != 1 {
		return nil, errno.New(errno.OrderStatusError)
	}

	order.Status = 2
	now := time.Now()
	order.CompletedAt = &now

	if err := mysql.UpdateOrder(order); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return order, nil
}

func DeleteOrder(orderID int64) error {
	if err := mysql.DeleteOrder(orderID); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func UpdateOrder(orderID int64, startTime, endTime *time.Time, duration int, status int, remark string) (*model.Order, error) {
	order, err := mysql.GetOrderByID(orderID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.OrderNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	if startTime != nil {
		order.StartTime = *startTime
	}
	if endTime != nil {
		order.EndTime = *endTime
	}
	if duration > 0 {
		order.Duration = duration
	}
	if status >= 0 {
		order.Status = status
	}
	if remark != "" {
		order.Remark = remark
	}

	if err := mysql.UpdateOrder(order); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return order, nil
}
