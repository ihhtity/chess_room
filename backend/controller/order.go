package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/pkg/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderCreateRequest struct {
	RoomID    int64     `json:"room_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Remark    string    `json:"remark"`
}

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

func CreateOrder(c *gin.Context) {
	var req OrderCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	userID := c.GetInt64("user_id")
	order, err := logic.CreateOrder(userID, req.RoomID, req.StartTime, req.EndTime, req.Remark)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "创建成功", order)
}

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

func ConfirmOrder(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	order, err := logic.ConfirmOrder(idInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "确认成功", order)
}

func CompleteOrder(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	order, err := logic.CompleteOrder(idInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "完成成功", order)
}

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
