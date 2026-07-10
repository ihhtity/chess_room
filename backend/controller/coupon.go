package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UseCouponRequest struct {
	OrderID int64 `json:"order_id"`
}

// @Summary 获取用户优惠券列表
// @Description 获取当前登录用户的优惠券列表
// @Tags 优惠券
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.Coupon}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /coupons [get]
func GetCouponList(c *gin.Context) {
	userID := c.GetInt64("user_id")
	coupons, total, err := logic.GetCouponList(userID)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, gin.H{
		"data":  coupons,
		"total": total,
	})
}

// @Summary 获取优惠券详情
// @Description 获取指定优惠券详情
// @Tags 优惠券
// @Accept json
// @Produce json
// @Param id path string true "优惠券ID"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Coupon}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /coupons/{id} [get]
func GetCouponDetail(c *gin.Context) {
	userID := c.GetInt64("user_id")
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	coupon, err := logic.GetCouponDetail(userID, idInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	if coupon == nil {
		response.Fail(c, 404, "优惠券不存在")
		return
	}

	response.Success(c, coupon)
}

// @Summary 使用优惠券
// @Description 用户使用优惠券
// @Tags 优惠券
// @Accept json
// @Produce json
// @Param id path string true "优惠券ID"
// @Param body body UseCouponRequest true "使用信息"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /coupons/{id}/use [post]
func UseCoupon(c *gin.Context) {
	userID := c.GetInt64("user_id")
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	var req UseCouponRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.UseCoupon(userID, idInt, req.OrderID); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "使用成功", nil)
}
