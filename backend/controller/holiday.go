package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type HolidayCreateRequest struct {
	Name        string `json:"name" binding:"required" comment:"节假日名称"`
	Date        string `json:"date" binding:"required" comment:"日期"`
	IsHoliday   int    `json:"is_holiday" comment:"是否节假日"`
	Description string `json:"description" comment:"描述"`
}

// @Summary 获取节假日列表
// @Description 获取所有节假日列表
// @Tags 节假日
// @Accept json
// @Produce json
// @Param is_holiday query string false "是否节假日"
// @Success 200 {object} response.Response
// @Router /holidays [get]
func GetHolidayList(c *gin.Context) {
	isHoliday := c.Query("is_holiday")
	holidays, err := logic.GetHolidayList(isHoliday)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, holidays)
}

// @Summary 获取节假日详情
// @Description 根据ID获取节假日详情
// @Tags 节假日
// @Accept json
// @Produce json
// @Param id path string true "节假日ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /holidays/{id} [get]
func GetHolidayDetail(c *gin.Context) {
	id := c.Param("id")
	holiday, err := logic.GetHolidayByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, holiday)
}

// @Summary 创建节假日
// @Description 创建新的节假日
// @Tags 节假日
// @Accept json
// @Produce json
// @Param body body HolidayCreateRequest true "节假日信息"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /holidays [post]
func CreateHoliday(c *gin.Context) {
	var req HolidayCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	holiday := &model.Holiday{
		Name:        req.Name,
		Date:        req.Date,
		IsHoliday:   req.IsHoliday,
		Description: req.Description,
	}

	if err := logic.CreateHoliday(holiday); err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, holiday)
}

// @Summary 更新节假日
// @Description 更新节假日信息
// @Tags 节假日
// @Accept json
// @Produce json
// @Param id path string true "节假日ID"
// @Param body body map[string]interface{} true "节假日信息"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /holidays/{id} [put]
func UpdateHoliday(c *gin.Context) {
	id := c.Param("id")
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	holiday, err := logic.UpdateHoliday(id, data)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, holiday)
}

// @Summary 删除节假日
// @Description 删除节假日
// @Tags 节假日
// @Accept json
// @Produce json
// @Param id path string true "节假日ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /holidays/{id} [delete]
func DeleteHoliday(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteHoliday(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}
