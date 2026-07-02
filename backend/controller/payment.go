package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentCreateRequest struct {
	OrderID     int64   `json:"order_id"`
	Amount      float64 `json:"amount"`
	PaymentType int     `json:"payment_type"`
}

func CreatePayment(c *gin.Context) {
	var req PaymentCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	userID := c.GetInt64("user_id")
	payment, err := logic.CreatePayment(req.OrderID, userID, req.Amount, req.PaymentType)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, payment)
}

func GetPaymentDetail(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	payment, err := logic.GetPaymentByID(idInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, payment)
}

func GetPaymentList(c *gin.Context) {
	userID := c.GetInt64("user_id")
	paymentType := c.Query("payment_type")
	status := c.Query("status")

	var paymentTypeInt, statusInt int
	var err error

	if paymentType != "" {
		paymentTypeInt, err = strconv.Atoi(paymentType)
		if err != nil {
			response.Fail(c, 400, "参数错误")
			return
		}
	}

	if status != "" {
		statusInt, err = strconv.Atoi(status)
		if err != nil {
			response.Fail(c, 400, "参数错误")
			return
		}
	}

	payments, err := logic.GetPaymentList(userID, paymentTypeInt, statusInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, payments)
}

func WechatPayNotify(c *gin.Context) {
	var req struct {
		PaymentID     int64  `json:"payment_id"`
		TransactionNo string `json:"transaction_no"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.UpdatePaymentStatus(req.PaymentID, 1, req.TransactionNo); err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, nil)
}

func GetPaymentByOrderNo(c *gin.Context) {
	orderNo := c.Param("orderNo")
	payment, err := logic.GetPaymentByOrderNo(orderNo)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, payment)
}
