package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

func CreatePayment(orderID, userID int64, amount float64, paymentType int) (*model.Payment, error) {
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

	if order.TotalAmount != amount {
		return nil, errno.New(errno.PaymentFailed)
	}

	payment := &model.Payment{
		OrderID:     orderID,
		UserID:      userID,
		Amount:      amount,
		PaymentType: paymentType,
		Status:      0,
		TransactionNo: "TXN" + time.Now().Format("20060102150405") +
			fmt.Sprintf("%d", time.Now().Nanosecond()%1000000000/100000000),
	}

	if err := mysql.CreatePayment(payment); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return payment, nil
}

func GetPaymentByID(id int64) (*model.Payment, error) {
	payment, err := mysql.GetPaymentByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.PaymentNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return payment, nil
}

func GetPaymentByOrderID(orderID int64) (*model.Payment, error) {
	payment, err := mysql.GetPaymentByOrderID(orderID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.PaymentNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return payment, nil
}

func UpdatePaymentStatus(id int64, status int, transactionNo string) error {
	payment, err := mysql.GetPaymentByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errno.New(errno.PaymentNotFound)
		}
		return errno.New(errno.InternalError)
	}

	if payment.Status != 0 {
		return errno.New(errno.PaymentAlreadyDone)
	}

	payment.Status = status
	if transactionNo != "" {
		payment.TransactionNo = transactionNo
	}
	now := time.Now()
	payment.PaidAt = &now

	if err := mysql.UpdatePayment(payment); err != nil {
		return errno.New(errno.InternalError)
	}

	if status == 1 {
		order, _ := mysql.GetOrderByID(payment.OrderID)
		if order != nil && order.Status == 0 {
			order.Status = 1
			order.PaidAmount = payment.Amount
			order.PaidAt = &now
			mysql.UpdateOrder(order)
		}
	}

	return nil
}

func GetPaymentList(userID int64, paymentType, status int) ([]model.Payment, error) {
	payments, err := mysql.GetPaymentList(userID, paymentType, status)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return payments, nil
}

func GetPaymentByOrderNo(orderNo string) (*model.Payment, error) {
	payment, err := mysql.GetPaymentByOrderNo(orderNo)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.PaymentNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return payment, nil
}
