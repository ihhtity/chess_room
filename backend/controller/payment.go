package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 获取支付列表
// @Description 获取所有支付记录列表
// @Tags 支付管理
// @Accept json
// @Produce json
// @Param user_id query string false "用户ID"
// @Param payment_type query string false "支付类型"
// @Param status query string false "状态"
// @Success 200 {object} response.Response
// @Router /payments [get]
func GetPaymentList(c *gin.Context) {
	userID := c.Query("user_id")
	paymentType := c.Query("payment_type")
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

	payments, total, err := logic.GetPaymentList(userID, paymentType, status, page, pageSize)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, gin.H{
		"data":  payments,
		"total": total,
	})
}

// @Summary 获取支付详情
// @Description 根据ID获取支付记录详情
// @Tags 支付管理
// @Accept json
// @Produce json
// @Param id path string true "支付ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /payments/{id} [get]
func GetPaymentDetail(c *gin.Context) {
	id := c.Param("id")
	payment, err := logic.GetPaymentByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, payment)
}

// @Summary 更新支付记录
// @Description 更新支付记录信息
// @Tags 支付管理
// @Accept json
// @Produce json
// @Param id path string true "支付ID"
// @Param body body map[string]interface{} true "支付信息"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /payments/{id} [put]
func UpdatePayment(c *gin.Context) {
	id := c.Param("id")
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	payment, err := logic.UpdatePayment(id, data)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, payment)
}

// @Summary 删除支付记录
// @Description 删除支付记录
// @Tags 支付管理
// @Accept json
// @Produce json
// @Param id path string true "支付ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /payments/{id} [delete]
func DeletePayment(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeletePayment(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}

// @Summary 批量删除支付记录
// @Description 管理员批量删除支付记录
// @Tags 支付管理
// @Accept json
// @Produce json
// @Param body body object{ids=[]string} true "支付记录ID列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /payments/batch [delete]
func BatchDeletePayment(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	if len(req.IDs) == 0 {
		response.Fail(c, 400, "请选择要删除的支付记录")
		return
	}
	if err := logic.BatchDeletePayment(req.IDs); err != nil {
		response.HandleError(c, err)
		return
	}
	response.SuccessWithMsg(c, "批量删除成功", nil)
}

// @Summary 批量更新支付记录
// @Description 管理员批量更新支付记录状态
// @Tags 支付管理
// @Accept json
// @Produce json
// @Param body body []object{id=int64,status=int} true "支付记录更新信息列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /payments/batch [put]
func BatchUpdatePayment(c *gin.Context) {
	var reqs []struct {
		ID     int64 `json:"id"`
		Status int   `json:"status"`
	}
	if err := c.ShouldBindJSON(&reqs); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.BatchUpdatePayment(reqs); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "批量更新成功", nil)
}
