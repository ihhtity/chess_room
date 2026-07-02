package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
)

type ActivityCreateRequest struct {
	Name        string  `json:"name"`
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

func GetActivityList(c *gin.Context) {
	adminID := c.GetInt64("admin_id")
	var activities []model.Activity
	var err error
	if adminID != 0 {
		activities, err = logic.GetActivityListAdmin()
	} else {
		activities, err = logic.GetActivityList()
	}
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, activities)
}

func GetActivityDetail(c *gin.Context) {
	id := c.Param("id")
	activity, err := logic.GetActivityByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, activity)
}

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
		if validFrom, err := time.Parse(time.RFC3339, req.ValidFrom); err == nil {
			activity.ValidFrom = validFrom
		}
	}

	if req.ValidTo != "" {
		if validTo, err := time.Parse(time.RFC3339, req.ValidTo); err == nil {
			activity.ValidTo = validTo
		}
	}

	if err := logic.CreateActivity(activity); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "创建成功", activity)
}

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
		if validFrom, err := time.Parse(time.RFC3339, req.ValidFrom); err == nil {
			activity.ValidFrom = validFrom
		}
	}
	if req.ValidTo != "" {
		if validTo, err := time.Parse(time.RFC3339, req.ValidTo); err == nil {
			activity.ValidTo = validTo
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

func DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteActivity(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.SuccessWithMsg(c, "删除成功", nil)
}
