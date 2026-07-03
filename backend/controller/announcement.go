package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"strconv"

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
	Type    *int   `json:"type"`
	Status  *int   `json:"status"`
}

// @Summary 获取公告列表
// @Description 获取公告列表，支持按标题、类型、状态筛选
// @Tags 公告
// @Accept json
// @Produce json
// @Param title query string false "公告标题"
// @Param type query string false "公告类型"
// @Param status query string false "公告状态"
// @Success 200 {object} response.Response{data=[]model.Announcement}
// @Failure 400 {object} response.Response
// @Router /announcements [get]
func GetAnnouncementList(c *gin.Context) {
	title := c.Query("title")
	typeStr := c.Query("type")
	statusStr := c.Query("status")

	typeInt := -1
	statusInt := -1
	var err error

	if typeStr != "" {
		typeInt, err = strconv.Atoi(typeStr)
		if err != nil {
			response.Fail(c, 400, "类型格式错误")
			return
		}
	}

	if statusStr != "" {
		statusInt, err = strconv.Atoi(statusStr)
		if err != nil {
			response.Fail(c, 400, "状态格式错误")
			return
		}
	}

	announcements, err := logic.GetAnnouncementListFiltered(title, typeInt, statusInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, announcements)
}

// @Summary 获取公告详情
// @Description 根据公告ID获取公告详情
// @Tags 公告
// @Accept json
// @Produce json
// @Param id path string true "公告ID"
// @Success 200 {object} response.Response{data=model.Announcement}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /announcements/{id} [get]
func GetAnnouncementDetail(c *gin.Context) {
	id := c.Param("id")
	announcement, err := logic.GetAnnouncementByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, announcement)
}

// @Summary 创建公告
// @Description 管理员创建新公告
// @Tags 公告
// @Accept json
// @Produce json
// @Param body body AnnouncementCreateRequest true "公告信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Announcement}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /announcements [post]
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

// @Summary 更新公告
// @Description 管理员更新公告信息
// @Tags 公告
// @Accept json
// @Produce json
// @Param id path string true "公告ID"
// @Param body body AnnouncementUpdateRequest true "公告更新信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Announcement}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /announcements/{id} [put]
func UpdateAnnouncement(c *gin.Context) {
	id := c.Param("id")
	var req AnnouncementUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	announcement, err := logic.GetAnnouncementByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	if req.Title != "" {
		announcement.Title = req.Title
	}
	if req.Content != "" {
		announcement.Content = req.Content
	}
	if req.Type != nil {
		announcement.Type = *req.Type
	}
	if req.Status != nil {
		announcement.Status = *req.Status
	}

	if err := logic.UpdateAnnouncement(id, announcement); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "更新成功", announcement)
}

// @Summary 删除公告
// @Description 管理员删除公告
// @Tags 公告
// @Accept json
// @Produce json
// @Param id path string true "公告ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /announcements/{id} [delete]
func DeleteAnnouncement(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteAnnouncement(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}
