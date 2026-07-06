package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetPaymentList(userID, paymentType, status string, page, pageSize int) ([]model.Payment, int64, error) {
	userIDInt := int64(0)
	if userID != "" {
		var err error
		userIDInt, err = strconv.ParseInt(userID, 10, 64)
		if err != nil {
			return nil, 0, errno.New(errno.BadRequest)
		}
	}

	paymentTypeInt := 0
	if paymentType != "" {
		var err error
		paymentTypeInt, err = strconv.Atoi(paymentType)
		if err != nil {
			return nil, 0, errno.New(errno.BadRequest)
		}
	}

	statusInt := 0
	if status != "" {
		var err error
		statusInt, err = strconv.Atoi(status)
		if err != nil {
			return nil, 0, errno.New(errno.BadRequest)
		}
	}

	payments, total, err := mysql.GetPaymentList(userIDInt, paymentTypeInt, statusInt, page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return payments, total, nil
}

func GetPaymentByID(id string) (*model.Payment, error) {
	paymentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	payment, err := mysql.GetPaymentByID(paymentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.NotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return payment, nil
}

func UpdatePayment(id string, data map[string]interface{}) (*model.Payment, error) {
	paymentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}

	payment, err := mysql.GetPaymentByID(paymentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.NotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	if err := mysql.DB.Model(payment).Updates(data).Error; err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return mysql.GetPaymentByID(paymentID)
}

func DeletePayment(id string) error {
	paymentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	if err := mysql.DeletePayment(paymentID); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchDeletePayment(ids []string) error {
	var paymentIDs []int64
	for _, id := range ids {
		paymentID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errno.New(errno.BadRequest)
		}
		paymentIDs = append(paymentIDs, paymentID)
	}
	if err := mysql.BatchDeletePayment(paymentIDs); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchUpdatePayment(reqs []struct {
	ID     int64 `json:"id"`
	Status int   `json:"status"`
}) error {
	for _, req := range reqs {
		payment, err := mysql.GetPaymentByID(req.ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errno.New(errno.NotFound)
			}
			return errno.New(errno.InternalError)
		}

		payment.Status = req.Status

		if err := mysql.DB.Model(payment).Update("status", req.Status).Error; err != nil {
			return errno.New(errno.InternalError)
		}
	}
	return nil
}

func CreatePayment(orderID, userID int64, amount float64, paymentType int) (*model.Payment, error) {
	payment := &model.Payment{
		OrderID:       orderID,
		UserID:        userID,
		Amount:        amount,
		PaymentType:   paymentType,
		Status:        0,
		TransactionNo: "TXN" + strconv.FormatInt(orderID, 10) + strconv.FormatInt(userID, 10),
	}

	if err := mysql.DB.Create(payment).Error; err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return payment, nil
}
