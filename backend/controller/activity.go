package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ActivityCreateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Discount    float64 `json:"discount"`
	ValidFrom   string  `json:"valid_from"`
	ValidTo     string  `json:"valid_to"`
	Status      int     `json:"status"`
}

type ActivityUpdateRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Discount    *float64 `json:"discount"`
	ValidFrom   string   `json:"valid_from"`
	ValidTo     string   `json:"valid_to"`
	Status      *int     `json:"status"`
	SortOrder   *int     `json:"sort_order"`
}

// @Summary 获取活动列表
// @Description 获取活动列表，管理员查看全部，普通用户只看启用的活动，支持按名称、状态筛选
// @Tags 活动
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name query string false "活动名称"
// @Param status query string false "活动状态"
// @Success 200 {object} response.Response{data=[]model.Activity}
// @Failure 400 {object} response.Response
// @Router /activities [get]
func GetActivityList(c *gin.Context) {
	adminID := c.GetInt64("admin_id")
	name := c.Query("name")
	statusStr := c.Query("status")

	status := -1
	var err error
	if statusStr != "" {
		status, err = strconv.Atoi(statusStr)
		if err != nil {
			response.Fail(c, 400, "状态格式错误")
			return
		}
	}

	var activities []model.Activity
	if adminID != 0 {
		activities, err = logic.GetActivityListAdminFiltered(name, status)
	} else {
		activities, err = logic.GetActivityList()
	}
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, activities)
}

// @Summary 获取活动详情
// @Description 根据活动ID获取活动详情
// @Tags 活动
// @Accept json
// @Produce json
// @Param id path string true "活动ID"
// @Success 200 {object} response.Response{data=model.Activity}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /activities/{id} [get]
func GetActivityDetail(c *gin.Context) {
	id := c.Param("id")
	activity, err := logic.GetActivityByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, activity)
}

// @Summary 创建活动
// @Description 管理员创建新活动
// @Tags 活动
// @Accept json
// @Produce json
// @Param body body ActivityCreateRequest true "活动信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Activity}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /activities [post]
func CreateActivity(c *gin.Context) {
	var req ActivityCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	activity := &model.Activity{
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
		Discount:    req.Discount,
		Status:      req.Status,
	}

	if req.ValidFrom != "" {
		if validFrom, err := time.Parse("2006-01-02 15:04:05", req.ValidFrom); err == nil {
			activity.ValidFrom = validFrom
		} else {
			response.Fail(c, 400, "开始时间格式错误，应为 YYYY-MM-DD HH:mm:ss")
			return
		}
	}

	if req.ValidTo != "" {
		if validTo, err := time.Parse("2006-01-02 15:04:05", req.ValidTo); err == nil {
			activity.ValidTo = validTo
		} else {
			response.Fail(c, 400, "结束时间格式错误，应为 YYYY-MM-DD HH:mm:ss")
			return
		}
	}

	if err := logic.CreateActivity(activity); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "创建成功", activity)
}

// @Summary 更新活动
// @Description 管理员更新活动信息
// @Tags 活动
// @Accept json
// @Produce json
// @Param id path string true "活动ID"
// @Param body body ActivityUpdateRequest true "活动更新信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Activity}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /activities/{id} [put]
func UpdateActivity(c *gin.Context) {
	id := c.Param("id")
	var req ActivityUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	activity, err := logic.GetActivityByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	if req.Name != "" {
		activity.Name = req.Name
	}
	if req.Description != "" {
		activity.Description = req.Description
	}
	if req.Image != "" {
		activity.Image = req.Image
	}
	if req.Discount != nil {
		activity.Discount = *req.Discount
	}
	if req.ValidFrom != "" {
		if validFrom, err := time.Parse("2006-01-02 15:04:05", req.ValidFrom); err == nil {
			activity.ValidFrom = validFrom
		} else {
			response.Fail(c, 400, "开始时间格式错误，应为 YYYY-MM-DD HH:mm:ss")
			return
		}
	}
	if req.ValidTo != "" {
		if validTo, err := time.Parse("2006-01-02 15:04:05", req.ValidTo); err == nil {
			activity.ValidTo = validTo
		} else {
			response.Fail(c, 400, "结束时间格式错误，应为 YYYY-MM-DD HH:mm:ss")
			return
		}
	}
	if req.Status != nil {
		activity.Status = *req.Status
	}
	if req.SortOrder != nil {
		activity.SortOrder = *req.SortOrder
	}

	if err := logic.UpdateActivity(id, activity); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", activity)
}

// @Summary 删除活动
// @Description 管理员删除活动
// @Tags 活动
// @Accept json
// @Produce json
// @Param id path string true "活动ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /activities/{id} [delete]
func DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteActivity(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.SuccessWithMsg(c, "删除成功", nil)
}
