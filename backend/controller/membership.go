package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"strconv"
	"time"

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
	ExpiredAt        string   `json:"expired_at"`
}

type MembershipCreateRequest struct {
	UserID           int64   `json:"user_id" binding:"required"`
	Level            int     `json:"level" default:"0"`
	Balance          float64 `json:"balance" default:"0"`
	Points           int     `json:"points" default:"0"`
	TotalConsumed    float64 `json:"total_consumed" default:"0"`
	TotalRecharged   float64 `json:"total_recharged" default:"0"`
	MembershipStatus int     `json:"membership_status" default:"1"`
	ExpiredAt        string  `json:"expired_at"`
}

// @Summary 获取会员信息
// @Description 获取当前登录用户的会员信息
// @Tags 会员
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Membership}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /membership [get]
func GetMembership(c *gin.Context) {
	userID := c.GetInt64("user_id")
	membership, err := logic.GetMembership(userID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, membership)
}

// @Summary 会员储值
// @Description 用户进行会员储值
// @Tags 会员
// @Accept json
// @Produce json
// @Param body body RechargeRequest true "储值信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Membership}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /membership/recharge [post]
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

// @Summary 获取会员列表
// @Description 管理员获取会员列表，支持按等级、状态筛选
// @Tags 会员
// @Accept json
// @Produce json
// @Param user_id query string false "用户ID"
// @Param level query string false "会员等级"
// @Param membership_status query string false "会员状态"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.Membership}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /memberships [get]
func GetMembershipList(c *gin.Context) {
	userIDStr := c.Query("user_id")
	level := c.Query("level")
	status := c.Query("membership_status")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")

	var userID int64
	var levelInt int
	var statusInt int = -1
	var err error

	if userIDStr != "" {
		userID, err = strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			response.Fail(c, 400, "用户ID格式错误")
			return
		}
	}

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

	page := 1
	pageSize := 10
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 0 {
			page = 1
		}
	}

	if pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 0 {
			pageSize = 10
		}
	}

	memberships, total, err := logic.GetMembershipList(userID, levelInt, statusInt, page, pageSize)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, gin.H{
		"data":  memberships,
		"total": total,
	})
}

// @Summary 获取储值记录
// @Description 获取当前登录用户的储值记录
// @Tags 会员
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.RechargeRecord}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /membership/recharges [get]
func GetRechargeRecords(c *gin.Context) {
	userID := c.GetInt64("user_id")
	records, err := logic.GetRechargeRecords(userID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, records)
}

// @Summary 获取会员详情
// @Description 管理员根据会员ID获取会员详情
// @Tags 会员
// @Accept json
// @Produce json
// @Param id path string true "会员ID"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Membership}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /memberships/{id} [get]
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

// @Summary 创建会员
// @Description 管理员创建会员
// @Tags 会员
// @Accept json
// @Produce json
// @Param body body MembershipCreateRequest true "会员信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Membership}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /memberships [post]
func CreateMembership(c *gin.Context) {
	var req MembershipCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	var expiredAt *time.Time
	if req.ExpiredAt != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", req.ExpiredAt); err == nil {
			expiredAt = &t
		} else {
			response.Fail(c, 400, "过期时间格式错误，应为 YYYY-MM-DD HH:mm:ss")
			return
		}
	}

	membership := &model.Membership{
		UserID:           req.UserID,
		Level:            req.Level,
		Balance:          req.Balance,
		Points:           req.Points,
		TotalConsumed:    req.TotalConsumed,
		TotalRecharged:   req.TotalRecharged,
		MembershipStatus: req.MembershipStatus,
		JoinedAt:         time.Now(),
		ExpiredAt:        expiredAt,
	}

	if err := logic.CreateMembership(membership); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "创建成功", membership)
}

// @Summary 更新会员信息
// @Description 管理员更新会员信息
// @Tags 会员
// @Accept json
// @Produce json
// @Param id path string true "会员ID"
// @Param body body MembershipUpdateRequest true "会员更新信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Membership}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /memberships/{id} [put]
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

	var expiredAt *time.Time
	if req.ExpiredAt != "" {
		if t, err := time.Parse("2006-01-02 15:04:05", req.ExpiredAt); err == nil {
			expiredAt = &t
		} else {
			response.Fail(c, 400, "过期时间格式错误，应为 YYYY-MM-DD HH:mm:ss")
			return
		}
	}

	membership, err := logic.UpdateMembership(idInt, req.Level, req.Balance, req.Points, req.MembershipStatus, expiredAt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", membership)
}

// @Summary 删除会员
// @Description 管理员删除会员
// @Tags 会员
// @Accept json
// @Produce json
// @Param id path string true "会员ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /memberships/{id} [delete]
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

// @Summary 批量删除会员
// @Description 管理员批量删除会员
// @Tags 会员
// @Accept json
// @Produce json
// @Param body body object{ids=[]string} true "会员ID列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /memberships/batch [delete]
func BatchDeleteMembership(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	if len(req.IDs) == 0 {
		response.Fail(c, 400, "请选择要删除的会员")
		return
	}
	if err := logic.BatchDeleteMembership(req.IDs); err != nil {
		response.HandleError(c, err)
		return
	}
	response.SuccessWithMsg(c, "批量删除成功", nil)
}

// @Summary 批量更新会员信息
// @Description 管理员批量更新会员信息
// @Tags 会员
// @Accept json
// @Produce json
// @Param body body []object{id=int64,level=int,balance=float64,points=int,membership_status=int} true "会员更新信息列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /memberships/batch [put]
func BatchUpdateMembership(c *gin.Context) {
	var reqs []struct {
		ID               int64   `json:"id"`
		Level            int     `json:"level"`
		Balance          float64 `json:"balance"`
		Points           int     `json:"points"`
		MembershipStatus int     `json:"membership_status"`
	}
	if err := c.ShouldBindJSON(&reqs); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.BatchUpdateMembership(reqs); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "批量更新成功", nil)
}
