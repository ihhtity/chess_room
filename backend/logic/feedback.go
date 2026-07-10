package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"time"

	"github.com/jinzhu/gorm"
)

func CreateFeedback(userID int64, content, contact string, feedbackType int) (*model.Feedback, error) {
	feedback := &model.Feedback{
		UserID:  userID,
		Content: content,
		Contact: contact,
		Type:    feedbackType,
		Status:  0,
	}
	
	if err := mysql.CreateFeedback(feedback); err != nil {
		return nil, errno.New(errno.InternalError)
	}
	
	return feedback, nil
}

func GetFeedback(id int64) (*model.Feedback, error) {
	feedback, err := mysql.GetFeedbackByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.UserNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return feedback, nil
}

func GetFeedbackList(userID int64, page, pageSize int) ([]model.Feedback, int64, error) {
	feedbacks, total, err := mysql.GetFeedbackList(userID, page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return feedbacks, total, nil
}

func ReplyFeedback(id int64, reply string) (*model.Feedback, error) {
	feedback, err := mysql.GetFeedbackByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.UserNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	
	feedback.Reply = reply
	feedback.Status = 1
	now := time.Now()
	feedback.RepliedAt = &now
	
	if err := mysql.UpdateFeedback(feedback); err != nil {
		return nil, errno.New(errno.InternalError)
	}
	
	return feedback, nil
}
