package mysql

import (
	"chess-room-backend/model"
)

func CreateFeedback(feedback *model.Feedback) error {
	return DB.Create(feedback).Error
}

func GetFeedbackByID(id int64) (*model.Feedback, error) {
	var feedback model.Feedback
	err := DB.Preload("User").Where("id = ?", id).First(&feedback).Error
	return &feedback, err
}

func GetFeedbackList(userID int64, page, pageSize int) ([]model.Feedback, int64, error) {
	var feedbacks []model.Feedback
	var total int64

	query := DB.Preload("User")
	if userID > 0 {
		query = query.Where("user_id = ?", userID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&feedbacks).Error
	return feedbacks, total, err
}

func UpdateFeedback(feedback *model.Feedback) error {
	return DB.Save(feedback).Error
}
