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

func GetPaymentList(userID int64, paymentType, status int) ([]model.Payment, error) {
	var payments []model.Payment
	db := DB.Order("created_at DESC")
	if userID != 0 {
		db = db.Where("user_id = ?", userID)
	}
	if paymentType != 0 {
		db = db.Where("payment_type = ?", paymentType)
	}
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	err := db.Find(&payments).Error
	return payments, err
}

func CreatePayment(payment *model.Payment) error {
	return DB.Create(payment).Error
}

func UpdatePayment(payment *model.Payment) error {
	return DB.Save(payment).Error
}