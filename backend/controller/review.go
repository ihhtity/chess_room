package controller

import (
	"chess-room-backend/logic"
	"chess-room-backend/model"
	"chess-room-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewCreateRequest struct {
	OrderID int64  `json:"order_id"`
	Rating  int    `json:"rating"`
	Content string `json:"content"`
	Images  string `json:"images"`
}

func GetReviewList(c *gin.Context) {
	roomID := c.Query("room_id")
	var roomIDInt int
	var err error

	if roomID != "" {
		roomIDInt, err = strconv.Atoi(roomID)
		if err != nil {
			response.Fail(c, 400, "参数错误")
			return
		}
	}

	reviews, err := logic.GetReviewList(roomIDInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, reviews)
}

func GetReviewDetail(c *gin.Context) {
	id := c.Param("id")
	review, err := logic.GetReviewByID(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, review)
}

func CreateReview(c *gin.Context) {
	var req ReviewCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	userID := c.GetInt64("user_id")

	review := &model.Review{
		OrderID: req.OrderID,
		UserID:  userID,
		Rating:  req.Rating,
		Content: req.Content,
		Images:  req.Images,
		Status:  1,
	}

	if err := logic.CreateReview(review); err != nil {
		response.HandleError(c, err)
		return
	}

	response.SuccessWithMsg(c, "评价成功", review)
}

func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	if err := logic.DeleteReview(id); err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, nil)
}
