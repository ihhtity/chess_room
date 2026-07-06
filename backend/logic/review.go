package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetReviewList(roomID, page, pageSize int) ([]model.Review, int64, error) {
	reviews, total, err := mysql.GetReviewList(roomID, page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return reviews, total, nil
}

func GetReviewListFiltered(roomID, userID, rating, status, page, pageSize int) ([]model.Review, int64, error) {
	reviews, total, err := mysql.GetReviewListFiltered(roomID, userID, rating, status, page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return reviews, total, nil
}

func GetReviewByID(id string) (*model.Review, error) {
	reviewID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	review, err := mysql.GetReviewByID(reviewID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.ReviewNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return review, nil
}

func CreateReview(review *model.Review) error {
	_, err := mysql.GetReviewByOrderID(review.OrderID)
	if err == nil {
		return errno.New(errno.ReviewAlreadyDone)
	}
	if err != gorm.ErrRecordNotFound {
		return errno.New(errno.InternalError)
	}

	if review.Rating < 1 || review.Rating > 5 {
		return errno.NewWithMessage(errno.BadRequest, "评分必须在1-5之间")
	}

	if err := mysql.CreateReview(review); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func DeleteReview(id string) error {
	reviewID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	if err := mysql.DeleteReview(reviewID); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchDeleteReview(ids []string) error {
	var reviewIDs []int64
	for _, id := range ids {
		reviewID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errno.New(errno.BadRequest)
		}
		reviewIDs = append(reviewIDs, reviewID)
	}
	if err := mysql.BatchDeleteReview(reviewIDs); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchUpdateReview(reqs []struct {
	ID      int64  `json:"id"`
	Rating  int    `json:"rating"`
	Content string `json:"content"`
	Status  int    `json:"status"`
}) error {
	for _, req := range reqs {
		_, err := mysql.GetReviewByID(req.ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errno.New(errno.ReviewNotFound)
			}
			return errno.New(errno.InternalError)
		}

		data := make(map[string]interface{})
		if req.Rating != 0 {
			if req.Rating < 1 || req.Rating > 5 {
				return errno.NewWithMessage(errno.BadRequest, "评分必须在1-5之间")
			}
			data["rating"] = req.Rating
		}
		if req.Content != "" {
			data["content"] = req.Content
		}
		if req.Status == 0 || req.Status == 1 {
			data["status"] = req.Status
		}

		if len(data) == 0 {
			continue
		}

		if err := mysql.UpdateReviewByID(req.ID, data); err != nil {
			return errno.New(errno.InternalError)
		}
	}
	return nil
}

func UpdateReview(id string, data map[string]interface{}) (*model.Review, error) {
	reviewID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}

	_, err = mysql.GetReviewByID(reviewID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.ReviewNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	if rating, ok := data["rating"]; ok {
		if r, ok := rating.(float64); ok && (r < 1 || r > 5) {
			return nil, errno.NewWithMessage(errno.BadRequest, "评分必须在1-5之间")
		}
	}

	if err := mysql.UpdateReviewByID(reviewID, data); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return mysql.GetReviewByID(reviewID)
}
