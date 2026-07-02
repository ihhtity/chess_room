package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type TimeSlotCreateRequest struct {
	TypeID       int64   `json:"type_id" binding:"required" comment:"房间类型ID"`
	Name         string  `json:"name" binding:"required" comment:"时间槽名称"`
	StartTime    string  `json:"start_time" binding:"required" comment:"开始时间"`
	EndTime      string  `json:"end_time" binding:"required" comment:"结束时间"`
	Price        float64 `json:"price" comment:"价格"`
	WeekdayPrice float64 `json:"weekday_price" comment:"工作日价格"`
	WeekendPrice float64 `json:"weekend_price" comment:"周末价格"`
	HolidayPrice float64 `json:"holiday_price" comment:"节假日价格"`
	SortOrder    int     `json:"sort_order" comment:"排序顺序"`
	Status       int     `json:"status" comment:"状态"`
}

type TimeSlotUpdateRequest struct {
	TypeID       *int64  `json:"type_id" comment:"房间类型ID"`
	Name         string  `json:"name" comment:"时间槽名称"`
	StartTime    string  `json:"start_time" comment:"开始时间"`
	EndTime      string  `json:"end_time" comment:"结束时间"`
	Price        float64 `json:"price" comment:"价格"`
	WeekdayPrice float64 `json:"weekday_price" comment:"工作日价格"`
	WeekendPrice float64 `json:"weekend_price" comment:"周末价格"`
	HolidayPrice float64 `json:"holiday_price" comment:"节假日价格"`
	SortOrder    *int    `json:"sort_order" comment:"排序顺序"`
	Status       *int    `json:"status" comment:"状态"`
}

// @Summary 获取时间槽列表
// @Description 获取所有时间槽列表
// @Tags 时间槽
// @Accept json
// @Produce json
// @Param type_id query string false "房间类型ID"
// @Success 200 {object} response.Response
// @Router /time-slots [get]
func GetTimeSlotList(c *gin.Context) {
	typeID := c.Query("type_id")
	slots, err := logic.GetTimeSlotList(typeID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, slots)
}

// @Summary 获取时间槽详情
// @Description 根据ID获取时间槽详情
// @Tags 时间槽
// @Accept json
// @Produce json
// @Param id path string true "时间槽ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /time-slots/{id} [get]
func GetTimeSlotDetail(c *gin.Context) {
	id := c.Param("id")
	slot, err := logic.GetTimeSlotByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, slot)
}

// @Summary 创建时间槽
// @Description 创建新的时间槽
// @Tags 时间槽
// @Accept json
// @Produce json
// @Param body body TimeSlotCreateRequest true "时间槽信息"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /time-slots [post]
func CreateTimeSlot(c *gin.Context) {
	var req TimeSlotCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	slot := &model.TimeSlot{
		TypeID:       req.TypeID,
		Name:         req.Name,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		Price:        req.Price,
		WeekdayPrice: req.WeekdayPrice,
		WeekendPrice: req.WeekendPrice,
		HolidayPrice: req.HolidayPrice,
		SortOrder:    req.SortOrder,
		Status:       req.Status,
	}

	if err := logic.CreateTimeSlot(slot); err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, slot)
}

// @Summary 更新时间槽
// @Description 更新时间槽信息
// @Tags 时间槽
// @Accept json
// @Produce json
// @Param id path string true "时间槽ID"
// @Param body body TimeSlotUpdateRequest true "时间槽信息"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /time-slots/{id} [put]
func UpdateTimeSlot(c *gin.Context) {
	id := c.Param("id")
	var req TimeSlotUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	slot, err := logic.GetTimeSlotByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	if req.TypeID != nil {
		slot.TypeID = *req.TypeID
	}
	if req.Name != "" {
		slot.Name = req.Name
	}
	if req.StartTime != "" {
		slot.StartTime = req.StartTime
	}
	if req.EndTime != "" {
		slot.EndTime = req.EndTime
	}
	slot.Price = req.Price
	slot.WeekdayPrice = req.WeekdayPrice
	slot.WeekendPrice = req.WeekendPrice
	slot.HolidayPrice = req.HolidayPrice
	if req.SortOrder != nil {
		slot.SortOrder = *req.SortOrder
	}
	if req.Status != nil {
		slot.Status = *req.Status
	}

	if err := logic.UpdateTimeSlot(slot); err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, slot)
}

// @Summary 删除时间槽
// @Description 删除时间槽
// @Tags 时间槽
// @Accept json
// @Produce json
// @Param id path string true "时间槽ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /time-slots/{id} [delete]
func DeleteTimeSlot(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteTimeSlot(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}
