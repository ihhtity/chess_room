package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomCreateRequest struct {
	Name     string `json:"name"`
	TypeID   int64  `json:"type_id"`
	Floor    string `json:"floor"`
	Capacity int    `json:"capacity"`
	Images   string `json:"images"`
}

type RoomUpdateRequest struct {
	Name     string `json:"name"`
	TypeID   int64  `json:"type_id"`
	Floor    string `json:"floor"`
	Capacity int    `json:"capacity"`
	Images   string `json:"images"`
	Status   *int   `json:"status"`
}

func GetRoomList(c *gin.Context) {
	typeID := c.Query("type_id")
	floor := c.Query("floor")
	status := c.Query("status")

	var rooms []model.Room
	var err error

	if typeID != "" {
		rooms, err = logic.GetRoomListByTypeID(typeID)
	} else if floor != "" {
		rooms, err = logic.GetRoomListByFloor(floor)
	} else if status != "" {
		rooms, err = logic.GetRoomListByStatus(status)
	} else {
		rooms, err = logic.GetRoomList()
	}

	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, rooms)
}

func GetRoomDetail(c *gin.Context) {
	id := c.Param("id")
	room, err := logic.GetRoomByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, room)
}

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
	if req.TypeID != 0 {
		room.TypeID = req.TypeID
	}
	if req.Floor != "" {
		room.Floor = req.Floor
	}
	if req.Capacity != 0 {
		room.Capacity = req.Capacity
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

func DeleteRoom(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteRoom(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}

func GetRoomTypeList(c *gin.Context) {
	types, err := logic.GetRoomTypeList()
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, types)
}

func GetRoomTypeDetail(c *gin.Context) {
	id := c.Param("id")
	roomType, err := logic.GetRoomTypeByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, roomType)
}

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

func DeleteRoomType(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteRoomType(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
