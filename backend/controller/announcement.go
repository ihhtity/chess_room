package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type AnnouncementCreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    int    `json:"type"`
}

type AnnouncementUpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    int    `json:"type"`
	Status  int    `json:"status"`
}

func GetAnnouncementList(c *gin.Context) {
	announcements, err := logic.GetAnnouncementList()
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, announcements)
}

func GetAnnouncementDetail(c *gin.Context) {
	id := c.Param("id")
	announcement, err := logic.GetAnnouncementByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, announcement)
}

func CreateAnnouncement(c *gin.Context) {
	var req AnnouncementCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	announcement := &model.Announcement{
		Title:   req.Title,
		Content: req.Content,
		Type:    req.Type,
		Status:  1,
	}

	if err := logic.CreateAnnouncement(announcement); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "创建成功", announcement)
}

func UpdateAnnouncement(c *gin.Context) {
	id := c.Param("id")
	var req AnnouncementUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	announcement := &model.Announcement{
		Title:   req.Title,
		Content: req.Content,
		Type:    req.Type,
		Status:  req.Status,
	}

	if err := logic.UpdateAnnouncement(id, announcement); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", nil)
}

func DeleteAnnouncement(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteAnnouncement(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}
