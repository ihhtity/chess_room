package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"strconv"

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
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")

	page := 1
	pageSize := 10
	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
	}

	if pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 {
			pageSize = 10
		}
	}

	records, total, err := logic.GetRechargeRecordList(userID, status, page, pageSize)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, gin.H{
		"data":  records,
		"total": total,
	})
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

// @Summary 批量删除充值记录
// @Description 批量删除充值记录
// @Tags 充值记录管理
// @Accept json
// @Produce json
// @Param body body struct{ IDs []string `json:"ids"` } true "充值记录ID列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /recharge-records/batch [delete]
func BatchDeleteRechargeRecord(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	if err := logic.BatchDeleteRechargeRecord(req.IDs); err != nil {
		response.HandleError(c, err)
		return
	}
	response.SuccessWithMsg(c, "批量删除成功", nil)
}

// @Summary 批量更新充值记录
// @Description 管理员批量更新充值记录状态
// @Tags 充值记录管理
// @Accept json
// @Produce json
// @Param body body []object{id=int64,status=int} true "充值记录更新信息列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /recharge-records/batch [put]
func BatchUpdateRechargeRecord(c *gin.Context) {
	var reqs []struct {
		ID     int64 `json:"id"`
		Status int   `json:"status"`
	}
	if err := c.ShouldBindJSON(&reqs); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.BatchUpdateRechargeRecord(reqs); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "批量更新成功", nil)
}
