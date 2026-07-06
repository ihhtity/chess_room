package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewCreateRequest struct {
	OrderID int64  `json:"order_id" binding:"required"`
	UserID  int64  `json:"user_id"`
	RoomID  int64  `json:"room_id" binding:"required"`
	Rating  int    `json:"rating" binding:"required"`
	Content string `json:"content"`
	Images  string `json:"images"`
	Status  int    `json:"status" default:"1"`
}

// @Summary 获取评价列表
// @Description 获取评价列表，支持按房间ID、用户ID、评分、状态筛选
// @Tags 评价
// @Accept json
// @Produce json
// @Param room_id query int false "房间ID"
// @Param user_id query int false "用户ID"
// @Param rating query int false "评分"
// @Param status query int false "评价状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} response.Response{data=[]model.Review}
// @Failure 400 {object} response.Response
// @Router /reviews [get]
func GetReviewList(c *gin.Context) {
	roomID := c.Query("room_id")
	userID := c.Query("user_id")
	rating := c.Query("rating")
	statusStr := c.Query("status")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")

	var roomIDInt, userIDInt, ratingInt int
	var statusInt int = -1
	var err error

	if roomID != "" {
		roomIDInt, err = strconv.Atoi(roomID)
		if err != nil {
			response.Fail(c, 400, "参数错误")
			return
		}
	}

	if userID != "" {
		userIDInt, err = strconv.Atoi(userID)
		if err != nil {
			response.Fail(c, 400, "参数错误")
			return
		}
	}

	if rating != "" {
		ratingInt, err = strconv.Atoi(rating)
		if err != nil || ratingInt < 1 || ratingInt > 5 {
			response.Fail(c, 400, "评分参数错误")
			return
		}
	}

	if statusStr != "" {
		statusInt, err = strconv.Atoi(statusStr)
		if err != nil {
			response.Fail(c, 400, "参数错误")
			return
		}
	}

	page := 1
	pageSize := 10
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

	reviews, total, err := logic.GetReviewListFiltered(roomIDInt, userIDInt, ratingInt, statusInt, page, pageSize)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, gin.H{
		"data":  reviews,
		"total": total,
	})
}

// @Summary 获取评价详情
// @Description 根据评价ID获取评价详情
// @Tags 评价
// @Accept json
// @Produce json
// @Param id path string true "评价ID"
// @Success 200 {object} response.Response{data=model.Review}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /reviews/{id} [get]
func GetReviewDetail(c *gin.Context) {
	id := c.Param("id")
	review, err := logic.GetReviewByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, review)
}

// @Summary 创建评价
// @Description 用户创建评价，管理员可指定用户ID和状态
// @Tags 评价
// @Accept json
// @Produce json
// @Param body body ReviewCreateRequest true "评价信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Review}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /reviews [post]
func CreateReview(c *gin.Context) {
	var req ReviewCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	userID := c.GetInt64("user_id")
	if userID == 0 {
		adminID := c.GetInt64("admin_id")
		if adminID != 0 && req.UserID != 0 {
			userID = req.UserID
		} else {
			response.Fail(c, 400, "用户ID不能为空")
			return
		}
	}

	status := req.Status
	if status == 0 {
		status = 1
	}

	review := &model.Review{
		OrderID: req.OrderID,
		UserID:  userID,
		RoomID:  req.RoomID,
		Rating:  req.Rating,
		Content: req.Content,
		Images:  req.Images,
		Status:  status,
	}

	if err := logic.CreateReview(review); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "评价成功", review)
}

// @Summary 更新评价
// @Description 管理员更新评价
// @Tags 评价
// @Accept json
// @Produce json
// @Param id path string true "评价ID"
// @Param body body map[string]interface{} true "评价信息"
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.Review}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /reviews/{id} [put]
func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	review, err := logic.UpdateReview(id, data)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, review)
}

// @Summary 删除评价
// @Description 管理员删除评价
// @Tags 评价
// @Accept json
// @Produce json
// @Param id path string true "评价ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /reviews/{id} [delete]
func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteReview(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}

// @Summary 批量删除评价
// @Description 管理员批量删除评价
// @Tags 评价
// @Accept json
// @Produce json
// @Param body body object{ids=[]string} true "评价ID列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /reviews/batch [delete]
func BatchDeleteReview(c *gin.Context) {
	var req struct {
		IDs []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	if len(req.IDs) == 0 {
		response.Fail(c, 400, "请选择要删除的评价")
		return
	}
	if err := logic.BatchDeleteReview(req.IDs); err != nil {
		response.HandleError(c, err)
		return
	}
	response.SuccessWithMsg(c, "批量删除成功", nil)
}

// @Summary 批量更新评价
// @Description 管理员批量更新评价状态
// @Tags 评价
// @Accept json
// @Produce json
// @Param body body []object{id=int64,status=int} true "评价更新信息列表"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /reviews/batch [put]
func BatchUpdateReview(c *gin.Context) {
	var reqs []struct {
		ID      int64  `json:"id"`
		Rating  int    `json:"rating"`
		Content string `json:"content"`
		Status  int    `json:"status"`
	}
	if err := c.ShouldBindJSON(&reqs); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := logic.BatchUpdateReview(reqs); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "批量更新成功", nil)
}
