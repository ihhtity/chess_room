package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/pkg/errno"
	"chess-room-backend/pkg/response"
	"chess-room-backend/pkg/wechat"
	"fmt"

	"github.com/gin-gonic/gin"
)

type WechatLoginRequest struct {
	Code     string `json:"code"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
}

func WechatLogin(c *gin.Context) {
	var req WechatLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	session, err := wechat.GetSession(req.Code)
	if err != nil {
		response.HandleError(c, errno.New(errno.WechatError))
		return
	}

	user, token, err := logic.UserLogin(session.OpenID, "", "", req.Nickname, req.Avatar, req.Gender)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, gin.H{
		"user":  user,
		"token": token,
	})
}

type WechatPayRequest struct {
	OrderID int64   `json:"order_id"`
	Amount  float64 `json:"amount"`
	OpenID  string  `json:"open_id"`
}

func WechatPay(c *gin.Context) {
	var req WechatPayRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	userID := c.GetInt64("user_id")

	payment, err := logic.CreatePayment(req.OrderID, userID, req.Amount, 1)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	payResult, err := wechat.CreateOrder(payment.TransactionNo, "棋牌室预约", fmt.Sprintf("%.2f", req.Amount), req.OpenID)
	if err != nil {
		response.HandleError(c, errno.New(errno.PaymentFailed))
		return
	}

	response.Success(c, gin.H{
		"payment":  payment,
		"pay_info": payResult,
	})
}
