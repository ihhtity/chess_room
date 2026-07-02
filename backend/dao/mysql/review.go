package mysql

import (
	"chess-room-backend/model"
)

func GetReviewList(roomID int) ([]model.Review, error) {
	var reviews []model.Review
	db := DB.Preload("User").Where("status = ?", 1)
	if roomID != 0 {
		db = db.Where("room_id = ?", roomID)
	}
	err := db.Order("created_at DESC").Find(&reviews).Error
	return reviews, err
}

func GetReviewByOrderID(orderID int64) (*model.Review, error) {
	var review model.Review
	err := DB.Where("order_id = ?", orderID).First(&review).Error
	return &review, err
}

func GetReviewByID(id int64) (*model.Review, error) {
	var review model.Review
	err := DB.Preload("User").Where("id = ?", id).First(&review).Error
	return &review, err
}

func CreateReview(review *model.Review) error {
	return DB.Create(review).Error
}

func UpdateReview(review *model.Review) error {
	return DB.Save(review).Error
}

func DeleteReview(id int64) error {
	return DB.Delete(&model.Review{}, id).Error
}
