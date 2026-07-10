package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateFeedbackRequest struct {
	Content string `json:"content" binding:"required"`
	Contact string `json:"contact"`
	Type    int    `json:"type"`
}

func CreateFeedback(c *gin.Context) {
	userID := c.GetInt64("user_id")
	
	var req CreateFeedbackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	
	if req.Content == "" {
		response.Fail(c, 400, "反馈内容不能为空")
		return
	}
	
	feedback, err := logic.CreateFeedback(userID, req.Content, req.Contact, req.Type)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	
	response.SuccessWithMsg(c, "提交成功", feedback)
}

func GetFeedback(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	
	feedback, err := logic.GetFeedback(idInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	
	response.Success(c, feedback)
}

func GetFeedbackList(c *gin.Context) {
	userID := c.GetInt64("user_id")
	
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	
	page := 1
	var err error
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			response.Fail(c, 400, "页码格式错误")
			return
		}
	}
	
	pageSize := 10
	if pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			response.Fail(c, 400, "每页数量格式错误")
			return
		}
	}
	
	feedbacks, total, err := logic.GetFeedbackList(userID, page, pageSize)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	
	response.Success(c, gin.H{
		"data":  feedbacks,
		"total": total,
	})
}

type ReplyFeedbackRequest struct {
	Reply string `json:"reply" binding:"required"`
}

func ReplyFeedback(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	
	var req ReplyFeedbackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	
	if req.Reply == "" {
		response.Fail(c, 400, "回复内容不能为空")
		return
	}
	
	feedback, err := logic.ReplyFeedback(idInt, req.Reply)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	
	response.SuccessWithMsg(c, "回复成功", feedback)
}
