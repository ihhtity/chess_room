package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	TypeID   int64  `json:"type_id" binding:"required"`
	Floor    string `json:"floor"`
	Capacity int    `json:"capacity" binding:"required"`
	Images   string `json:"images"`
}

type RoomUpdateRequest struct {
	Name     string `json:"name"`
	TypeID   *int64 `json:"type_id"`
	Floor    string `json:"floor"`
	Capacity *int   `json:"capacity"`
	Images   string `json:"images"`
	Status   *int   `json:"status"`
}

// @Summary 获取房间列表
// @Description 获取房间列表，支持按类型、楼层、状态、名称筛选
// @Tags 房间
// @Accept json
// @Produce json
// @Param type_id query string false "房间类型ID"
// @Param floor query string false "楼层"
// @Param status query string false "房间状态"
// @Param name query string false "房间名称"
// @Success 200 {object} response.Response{data=[]model.Room}
// @Failure 400 {object} response.Response
// @Router /rooms [get]
func GetRoomList(c *gin.Context) {
	typeID := c.Query("type_id")
	floor := c.Query("floor")
	status := c.Query("status")
	name := c.Query("name")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")

	typeIDInt := 0
	statusInt := -1
	var err error

	if typeID != "" {
		typeIDInt, err = strconv.Atoi(typeID)
		if err != nil {
			response.Fail(c, 400, "类型ID格式错误")
			return
		}
	}

	if status != "" {
		statusInt, err = strconv.Atoi(status)
		if err != nil {
			response.Fail(c, 400, "状态格式错误")
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

	rooms, total, err := logic.GetRoomListFiltered(typeIDInt, floor, statusInt, name, page, pageSize)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, gin.H{
		"data":  rooms,
		"total": total,
	})
}

// @Summary 获取房间详情
// @Description 根据房间ID获取房间详情
// @Tags 房间
// @Accept json
// @Produce json
// @Param id path string true "房间ID"
// @Success 200 {object} response.Response{data=model.Room}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /rooms/{id} [get]
func GetRoomDetail(c *gin.Context) {
	id := c.Param("id")
	room, err := logic.GetRoomByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, room)
}

// @Summary 创建房间
// @Description 管理员创建新房间
// @Tags 房间
// @Accept json
// @Produce json
// @Param body body RoomCreateRequest true "房间信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Room}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /room [post]
func CreateRoom(c *gin.Context) {
	var req RoomCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	room := &model.Room{
		Name:     req.Name,
		TypeID:   req.TypeID,
		Floor:    req.Floor,
		Capacity: req.Capacity,
		Images:   req.Images,
		Status:   1,
	}

	if err := logic.CreateRoom(room); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "创建成功", room)
}

// @Summary 更新房间
// @Description 管理员更新房间信息
// @Tags 房间
// @Accept json
// @Produce json
// @Param id path string true "房间ID"
// @Param body body RoomUpdateRequest true "房间更新信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Room}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /rooms/{id} [put]
func UpdateRoom(c *gin.Context) {
	id := c.Param("id")
	var req RoomUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	room, err := logic.GetRoomByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	if req.Name != "" {
		room.Name = req.Name
	}
	if req.TypeID != nil {
		room.TypeID = *req.TypeID
	}
	if req.Floor != "" {
		room.Floor = req.Floor
	}
	if req.Capacity != nil {
		room.Capacity = *req.Capacity
	}
	if req.Images != "" {
		room.Images = req.Images
	}
	if req.Status != nil {
		room.Status = *req.Status
	}

	if err := logic.UpdateRoom(room); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", room)
}

// @Summary 删除房间
// @Description 管理员删除房间
// @Tags 房间
// @Accept json
// @Produce json
// @Param id path string true "房间ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /rooms/{id} [delete]
func DeleteRoom(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteRoom(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}

// @Summary 批量删除房间
// @Description 管理员批量删除房间
// @Tags 房间
// @Accept json
// @Produce json
// @Param body body object{ids=[]string} true "房间ID列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /rooms/batch [delete]
func BatchDeleteRoom(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	if len(req.IDs) == 0 {
		response.Fail(c, 400, "请选择要删除的房间")
		return
	}
	if err := logic.BatchDeleteRoom(req.IDs); err != nil {
		response.HandleError(c, err)
		return
	}
	response.SuccessWithMsg(c, "批量删除成功", nil)
}

// @Summary 批量更新房间
// @Description 管理员批量更新房间信息
// @Tags 房间
// @Accept json
// @Produce json
// @Param body body []object{id=int64,name=string,type_id=int64,floor=string,capacity=int,images=string,status=int} true "房间更新信息列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /room/batch [put]
func BatchUpdateRoom(c *gin.Context) {
	var reqs []struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		TypeID   int64  `json:"type_id"`
		Floor    string `json:"floor"`
		Capacity int    `json:"capacity"`
		Images   string `json:"images"`
		Status   int    `json:"status"`
	}
	if err := c.ShouldBindJSON(&reqs); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.BatchUpdateRoom(reqs); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "批量更新成功", nil)
}

// @Summary 获取房间类型列表
// @Description 获取房间类型列表
// @Tags 房间类型
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]model.RoomType}
// @Failure 400 {object} response.Response
// @Router /room-types [get]
func GetRoomTypeList(c *gin.Context) {
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

	types, total, err := logic.GetRoomTypeList(page, pageSize)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, gin.H{
		"data":  types,
		"total": total,
	})
}

// @Summary 获取房间类型详情
// @Description 根据房间类型ID获取房间类型详情
// @Tags 房间类型
// @Accept json
// @Produce json
// @Param id path string true "房间类型ID"
// @Success 200 {object} response.Response{data=model.RoomType}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /room-types/{id} [get]
func GetRoomTypeDetail(c *gin.Context) {
	id := c.Param("id")
	roomType, err := logic.GetRoomTypeByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, roomType)
}

// @Summary 创建房间类型
// @Description 管理员创建房间类型
// @Tags 房间类型
// @Accept json
// @Produce json
// @Param body body object true "房间类型信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.RoomType}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /room-types [post]
func CreateRoomType(c *gin.Context) {
	var req struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		BasePrice   float64 `json:"base_price"`
		MaxPeople   int     `json:"max_people"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	roomType := &model.RoomType{
		Name:        req.Name,
		Description: req.Description,
		BasePrice:   req.BasePrice,
		MaxPeople:   req.MaxPeople,
	}

	if err := logic.CreateRoomType(roomType); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "创建成功", roomType)
}

// @Summary 更新房间类型
// @Description 管理员更新房间类型信息
// @Tags 房间类型
// @Accept json
// @Produce json
// @Param id path string true "房间类型ID"
// @Param body body object true "房间类型更新信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.RoomType}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /room-types/{id} [put]
func UpdateRoomType(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		BasePrice   *float64 `json:"base_price"`
		MaxPeople   *int     `json:"max_people"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	roomType, err := logic.GetRoomTypeByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	if req.Name != "" {
		roomType.Name = req.Name
	}
	if req.Description != "" {
		roomType.Description = req.Description
	}
	if req.BasePrice != nil {
		roomType.BasePrice = *req.BasePrice
	}
	if req.MaxPeople != nil {
		roomType.MaxPeople = *req.MaxPeople
	}

	if err := logic.UpdateRoomType(roomType); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", roomType)
}

// @Summary 删除房间类型
// @Description 管理员删除房间类型
// @Tags 房间类型
// @Accept json
// @Produce json
// @Param id path string true "房间类型ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /room-types/{id} [delete]
func DeleteRoomType(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteRoomType(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}

// @Summary 批量删除房间类型
// @Description 管理员批量删除房间类型
// @Tags 房间类型
// @Accept json
// @Produce json
// @Param body body object{ids=[]string} true "房间类型ID列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /room-types/batch [delete]
func BatchDeleteRoomType(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	if len(req.IDs) == 0 {
		response.Fail(c, 400, "请选择要删除的房间类型")
		return
	}
	if err := logic.BatchDeleteRoomType(req.IDs); err != nil {
		response.HandleError(c, err)
		return
	}
	response.SuccessWithMsg(c, "批量删除成功", nil)
}

// @Summary 批量更新房间类型
// @Description 管理员批量更新房间类型信息
// @Tags 房间类型
// @Accept json
// @Produce json
// @Param body body []object{id=int64,name=string,description=string,base_price=float64,max_people=int} true "房间类型更新信息列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /room-types/batch [put]
func BatchUpdateRoomType(c *gin.Context) {
	var reqs []struct {
		ID          int64   `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		BasePrice   float64 `json:"base_price"`
		MaxPeople   int     `json:"max_people"`
	}
	if err := c.ShouldBindJSON(&reqs); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.BatchUpdateRoomType(reqs); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "批量更新成功", nil)
}

// @Summary 健康检查
// @Description 服务健康检查接口
// @Tags 系统
// @Accept json
// @Produce json
// @Success 200 {object} object{status=string}
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
