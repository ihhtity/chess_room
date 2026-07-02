package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type NotificationCreateRequest struct {
	UserID     int    `json:"user_id" comment:"用户ID"`
	Type       int    `json:"type" comment:"类型"`
	Title      string `json:"title" binding:"required" comment:"标题"`
	Content    string `json:"content" comment:"内容"`
	ReadStatus int    `json:"read_status" comment:"读取状态"`
	Link       string `json:"link" comment:"链接"`
}

// @Summary 获取通知列表
// @Description 获取所有通知列表
// @Tags 通知管理
// @Accept json
// @Produce json
// @Param user_id query string false "用户ID"
// @Param type query string false "类型"
// @Param read_status query string false "读取状态"
// @Success 200 {object} response.Response
// @Router /notifications [get]
func GetNotificationList(c *gin.Context) {
	userID := c.Query("user_id")
	notificationType := c.Query("type")
	readStatus := c.Query("read_status")
	notifications, err := logic.GetNotificationList(userID, notificationType, readStatus)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, notifications)
}

// @Summary 获取通知详情
// @Description 根据ID获取通知详情
// @Tags 通知管理
// @Accept json
// @Produce json
// @Param id path string true "通知ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /notifications/{id} [get]
func GetNotificationDetail(c *gin.Context) {
	id := c.Param("id")
	notification, err := logic.GetNotificationByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, notification)
}

// @Summary 创建通知
// @Description 创建新的通知
// @Tags 通知管理
// @Accept json
// @Produce json
// @Param body body NotificationCreateRequest true "通知信息"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /notifications [post]
func CreateNotification(c *gin.Context) {
	var req NotificationCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	notification := &model.Notification{
		UserID:     int64(req.UserID),
		Type:       req.Type,
		Title:      req.Title,
		Content:    req.Content,
		ReadStatus: req.ReadStatus,
		Link:       req.Link,
	}

	if err := logic.CreateNotification(notification); err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, notification)
}

// @Summary 更新通知
// @Description 更新通知信息
// @Tags 通知管理
// @Accept json
// @Produce json
// @Param id path string true "通知ID"
// @Param body body map[string]interface{} true "通知信息"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /notifications/{id} [put]
func UpdateNotification(c *gin.Context) {
	id := c.Param("id")
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	notification, err := logic.UpdateNotification(id, data)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, notification)
}

// @Summary 删除通知
// @Description 删除通知
// @Tags 通知管理
// @Accept json
// @Produce json
// @Param id path string true "通知ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /notifications/{id} [delete]
func DeleteNotification(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteNotification(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}

// @Summary 标记全部已读
// @Description 标记用户的所有通知为已读
// @Tags 通知管理
// @Accept json
// @Produce json
// @Param user_id query string true "用户ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /notifications/mark-all-read [post]
func MarkAllNotificationAsRead(c *gin.Context) {
	userID := c.Query("user_id")
	if err := logic.MarkAllNotificationAsRead(userID); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}
