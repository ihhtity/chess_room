package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"strconv"

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

	holidays, total, err := logic.GetHolidayList(isHoliday, page, pageSize)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, gin.H{
		"data":  holidays,
		"total": total,
	})
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

// @Summary 批量删除节假日
// @Description 管理员批量删除节假日
// @Tags 节假日
// @Accept json
// @Produce json
// @Param body body object{ids=[]string} true "节假日ID列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /holidays/batch [delete]
func BatchDeleteHoliday(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	if len(req.IDs) == 0 {
		response.Fail(c, 400, "请选择要删除的节假日")
		return
	}
	if err := logic.BatchDeleteHoliday(req.IDs); err != nil {
		response.HandleError(c, err)
		return
	}
	response.SuccessWithMsg(c, "批量删除成功", nil)
}

// @Summary 批量更新节假日
// @Description 管理员批量更新节假日是否为假期
// @Tags 节假日
// @Accept json
// @Produce json
// @Param body body []object{id=int64,is_holiday=int} true "节假日更新信息列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /holidays/batch [put]
func BatchUpdateHoliday(c *gin.Context) {
	var reqs []struct {
		ID        int64 `json:"id"`
		IsHoliday int   `json:"is_holiday"`
	}
	if err := c.ShouldBindJSON(&reqs); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.BatchUpdateHoliday(reqs); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "批量更新成功", nil)
}
