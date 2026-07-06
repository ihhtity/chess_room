package mysql

import (
	"chess-room-backend/model"
)

func GetReviewList(roomID, page, pageSize int) ([]model.Review, int64, error) {
	var reviews []model.Review
	var total int64
	db := DB.Model(&model.Review{}).Where("status = ?", 1)
	if roomID != 0 {
		db = db.Where("room_id = ?", roomID)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB.Preload("User").Where("status = ?", 1)
		if roomID != 0 {
			db = db.Where("room_id = ?", roomID)
		}
		db = db.Order("created_at DESC").Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&reviews).Error
	return reviews, total, err
}

func GetReviewListFiltered(roomID, userID, rating, status, page, pageSize int) ([]model.Review, int64, error) {
	var reviews []model.Review
	var total int64
	db := DB.Model(&model.Review{}).Order("created_at DESC")
	if roomID != 0 {
		db = db.Where("room_id = ?", roomID)
	}
	if userID != 0 {
		db = db.Where("user_id = ?", userID)
	}
	if rating != 0 {
		db = db.Where("rating = ?", rating)
	}
	if status >= 0 {
		db = db.Where("status = ?", status)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB.Preload("User").Order("created_at DESC")
		if roomID != 0 {
			db = db.Where("room_id = ?", roomID)
		}
		if userID != 0 {
			db = db.Where("user_id = ?", userID)
		}
		if rating != 0 {
			db = db.Where("rating = ?", rating)
		}
		if status >= 0 {
			db = db.Where("status = ?", status)
		}
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&reviews).Error
	return reviews, total, err
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

func UpdateReviewByID(id int64, data map[string]interface{}) error {
	return DB.Model(&model.Review{}).Where("id = ?", id).Updates(data).Error
}

func DeleteReview(id int64) error {
	return DB.Delete(&model.Review{}, id).Error
}

func BatchDeleteReview(ids []int64) error {
	return DB.Where("id IN (?)", ids).Delete(&model.Review{}).Error
}
