package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetReviewList(roomID int) ([]model.Review, error) {
	reviews, err := mysql.GetReviewList(roomID)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return reviews, nil
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
