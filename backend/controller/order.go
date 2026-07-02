package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/pkg/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderCreateRequest struct {
	UserID    int64     `json:"user_id"`
	RoomID    int64     `json:"room_id" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
	Remark    string    `json:"remark"`
}

// @Summary 获取订单列表
// @Description 获取用户或管理员的订单列表
// @Tags 订单
// @Accept json
// @Produce json
// @Param room_id query int false "房间ID"
// @Param status query int false "订单状态"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.Order}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /orders [get]
func GetOrderList(c *gin.Context) {
	userID := c.GetInt64("user_id")
	roomID := c.Query("room_id")
	status := c.Query("status")

	var roomIDInt int
	var statusInt int
	var err error

	if roomID != "" {
		roomIDInt, err = strconv.Atoi(roomID)
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

	if userID == 0 {
		adminID := c.GetInt64("admin_id")
		if adminID != 0 {
			userID = -1
		}
	}

	orders, err := logic.GetOrderList(userID, roomIDInt, statusInt, time.Time{}, time.Time{})
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, orders)
}

// @Summary 获取订单详情
// @Description 根据订单ID获取订单详情
// @Tags 订单
// @Accept json
// @Produce json
// @Param id path int true "订单ID"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Order}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /orders/{id} [get]
func GetOrderDetail(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	order, err := logic.GetOrderByID(idInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, order)
}

// @Summary 创建订单
// @Description 用户创建新订单，管理员可指定用户ID
// @Tags 订单
// @Accept json
// @Produce json
// @Param body body OrderCreateRequest true "订单信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Order}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	var req OrderCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	userID := c.GetInt64("user_id")
	if userID == 0 {
		adminID := c.GetInt64("admin_id")
		if adminID != 0 && req.UserID != 0 {
			userID = req.UserID
		} else {
			response.Fail(c, 400, "用户ID不能为空")
			return
		}
	}

	order, err := logic.CreateOrder(userID, req.RoomID, req.StartTime, req.EndTime, req.Remark)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "创建成功", order)
}

// @Summary 取消订单
// @Description 用户取消未支付的订单
// @Tags 订单
// @Accept json
// @Produce json
// @Param id path int true "订单ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Router /orders/{id}/cancel [put]
func CancelOrder(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	userID := c.GetInt64("user_id")
	if err := logic.CancelOrder(userID, idInt); err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, nil)
}

// @Summary 确认订单
// @Description 用户或管理员确认订单
// @Tags 订单
// @Accept json
// @Produce json
// @Param id path int true "订单ID"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Order}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Router /orders/{id}/confirm [put]
func ConfirmOrder(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	userID := c.GetInt64("user_id")
	if userID != 0 {
		order, err := logic.GetOrderByID(idInt)
		if err != nil {
			response.HandleError(c, err)
			return
		}
		if order.UserID != userID {
			response.Fail(c, 403, "无权操作此订单")
			return
		}
	}

	order, err := logic.ConfirmOrder(idInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "确认成功", order)
}

// @Summary 完成订单
// @Description 用户或管理员完成订单
// @Tags 订单
// @Accept json
// @Produce json
// @Param id path int true "订单ID"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Order}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Router /orders/{id}/complete [put]
func CompleteOrder(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	userID := c.GetInt64("user_id")
	if userID != 0 {
		order, err := logic.GetOrderByID(idInt)
		if err != nil {
			response.HandleError(c, err)
			return
		}
		if order.UserID != userID {
			response.Fail(c, 403, "无权操作此订单")
			return
		}
	}

	order, err := logic.CompleteOrder(idInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "完成成功", order)
}

// @Summary 删除订单
// @Description 管理员删除订单
// @Tags 订单
// @Accept json
// @Produce json
// @Param id path int true "订单ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /orders/{id} [delete]
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.DeleteOrder(idInt); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "删除成功", nil)
}

type OrderUpdateRequest struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Duration  int    `json:"duration"`
	Status    *int   `json:"status"`
	Remark    string `json:"remark"`
}

// @Summary 更新订单
// @Description 管理员更新订单信息
// @Tags 订单
// @Accept json
// @Produce json
// @Param id path int true "订单ID"
// @Param body body OrderUpdateRequest true "订单更新信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Order}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /orders/{id} [put]
func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	var req OrderUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	var startTime, endTime *time.Time
	if req.StartTime != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", req.StartTime); err == nil {
			startTime = &t
		} else {
			response.Fail(c, 400, "开始时间格式错误，应为 YYYY-MM-DD HH:mm:ss")
			return
		}
	}

	if req.EndTime != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", req.EndTime); err == nil {
			endTime = &t
		} else {
			response.Fail(c, 400, "结束时间格式错误，应为 YYYY-MM-DD HH:mm:ss")
			return
		}
	}

	status := -1
	if req.Status != nil {
		status = *req.Status
	}

	order, err := logic.UpdateOrder(idInt, startTime, endTime, req.Duration, status, req.Remark)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", order)
}
