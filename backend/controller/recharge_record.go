package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type RechargeRecordCreateRequest struct {
	UserID     int64   `json:"user_id" binding:"required" comment:"用户ID"`
	Amount     float64 `json:"amount" binding:"required" comment:"金额"`
	GiftAmount float64 `json:"gift_amount" comment:"赠送金额"`
	PaymentID  int64   `json:"payment_id" comment:"支付ID"`
	Status     int     `json:"status" comment:"状态"`
}

// @Summary 获取充值记录列表
// @Description 获取所有充值记录列表
// @Tags 充值记录管理
// @Accept json
// @Produce json
// @Param user_id query string false "用户ID"
// @Param status query string false "状态"
// @Success 200 {object} response.Response
// @Router /recharge-records [get]
func GetRechargeRecordList(c *gin.Context) {
	userID := c.Query("user_id")
	status := c.Query("status")
	records, err := logic.GetRechargeRecordList(userID, status)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, records)
}

// @Summary 获取充值记录详情
// @Description 根据ID获取充值记录详情
// @Tags 充值记录管理
// @Accept json
// @Produce json
// @Param id path string true "充值记录ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /recharge-records/{id} [get]
func GetRechargeRecordDetail(c *gin.Context) {
	id := c.Param("id")
	record, err := logic.GetRechargeRecordByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, record)
}

// @Summary 创建充值记录
// @Description 创建新的充值记录
// @Tags 充值记录管理
// @Accept json
// @Produce json
// @Param body body RechargeRecordCreateRequest true "充值记录信息"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /recharge-records [post]
func CreateRechargeRecord(c *gin.Context) {
	var req RechargeRecordCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	record := &model.RechargeRecord{
		UserID:     req.UserID,
		Amount:     req.Amount,
		GiftAmount: req.GiftAmount,
		PaymentID:  req.PaymentID,
		Status:     req.Status,
	}

	if err := logic.CreateRechargeRecord(record); err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, record)
}

// @Summary 更新充值记录
// @Description 更新充值记录信息
// @Tags 充值记录管理
// @Accept json
// @Produce json
// @Param id path string true "充值记录ID"
// @Param body body map[string]interface{} true "充值记录信息"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /recharge-records/{id} [put]
func UpdateRechargeRecord(c *gin.Context) {
	id := c.Param("id")
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	record, err := logic.UpdateRechargeRecord(id, data)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, record)
}

// @Summary 删除充值记录
// @Description 删除充值记录
// @Tags 充值记录管理
// @Accept json
// @Produce json
// @Param id path string true "充值记录ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /recharge-records/{id} [delete]
func DeleteRechargeRecord(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteRechargeRecord(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}
