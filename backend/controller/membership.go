package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RechargeRequest struct {
	Amount float64 `json:"amount"`
}

type MembershipUpdateRequest struct {
	Level            *int     `json:"level"`
	Balance          *float64 `json:"balance"`
	Points           *int     `json:"points"`
	MembershipStatus *int     `json:"membership_status"`
}

func GetMembership(c *gin.Context) {
	userID := c.GetInt64("user_id")
	membership, err := logic.GetMembership(userID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, membership)
}

func Recharge(c *gin.Context) {
	var req RechargeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if req.Amount <= 0 {
		response.Fail(c, 400, "金额必须大于0")
		return
	}

	userID := c.GetInt64("user_id")
	membership, err := logic.Recharge(userID, req.Amount)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "储值成功", membership)
}

func GetMembershipList(c *gin.Context) {
	level := c.Query("level")
	status := c.Query("status")

	var levelInt, statusInt int
	var err error

	if level != "" {
		levelInt, err = strconv.Atoi(level)
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

	memberships, err := logic.GetMembershipList(levelInt, statusInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, memberships)
}

func GetRechargeRecords(c *gin.Context) {
	userID := c.GetInt64("user_id")
	records, err := logic.GetRechargeRecords(userID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, records)
}

func GetMembershipDetail(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	membership, err := logic.GetMembershipByID(idInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, membership)
}

func UpdateMembership(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	var req MembershipUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	membership, err := logic.UpdateMembership(idInt, req.Level, req.Balance, req.Points, req.MembershipStatus)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", membership)
}

func DeleteMembership(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.DeleteMembership(idInt); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "删除成功", nil)
}
