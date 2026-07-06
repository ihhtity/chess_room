package mysql

import (
	"chess-room-backend/model"
)

func GetPaymentByOrderNo(orderNo string) (*model.Payment, error) {
	var payment model.Payment
	err := DB.Where("order_no = ?", orderNo).First(&payment).Error
	return &payment, err
}

func GetPaymentByID(id int64) (*model.Payment, error) {
	var payment model.Payment
	err := DB.Where("id = ?", id).First(&payment).Error
	return &payment, err
}

func GetPaymentByOrderID(orderID int64) (*model.Payment, error) {
	var payment model.Payment
	err := DB.Where("order_id = ?", orderID).First(&payment).Error
	return &payment, err
}

func GetPaymentList(userID int64, paymentType, status, page, pageSize int) ([]model.Payment, int64, error) {
	var payments []model.Payment
	var total int64
	db := DB.Model(&model.Payment{}).Order("created_at DESC")
	if userID != 0 {
		db = db.Where("user_id = ?", userID)
	}
	if paymentType != 0 {
		db = db.Where("payment_type = ?", paymentType)
	}
	if status != 0 {
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
		if paymentType != 0 {
			db = db.Where("payment_type = ?", paymentType)
		}
		if status != 0 {
			db = db.Where("status = ?", status)
		}
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&payments).Error
	return payments, total, err
}

func CreatePayment(payment *model.Payment) error {
	return DB.Create(payment).Error
}

func UpdatePayment(payment *model.Payment) error {
	return DB.Save(payment).Error
}

func DeletePayment(id int64) error {
	return DB.Delete(&model.Payment{}, id).Error
}

func BatchDeletePayment(ids []int64) error {
	return DB.Where("id IN (?)", ids).Delete(&model.Payment{}).Error
}