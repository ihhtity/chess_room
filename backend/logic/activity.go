package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/dao/redis"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"encoding/json"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetActivityList() ([]model.Activity, error) {
	cacheKey := "activity:list"
	if cacheData, err := redis.Get(cacheKey); err == nil && cacheData != "" {
		var activities []model.Activity
		if err := json.Unmarshal([]byte(cacheData), &activities); err == nil {
			return activities, nil
		}
	}

	activities, err := mysql.GetActivityList()
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}

	if data, err := json.Marshal(activities); err == nil {
		redis.Set(cacheKey, string(data), 300)
	}

	return activities, nil
}

func GetActivityListAdmin() ([]model.Activity, error) {
	activities, err := mysql.GetActivityListAdmin()
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return activities, nil
}

func GetActivityListAdminFiltered(name string, status int) ([]model.Activity, error) {
	activities, err := mysql.GetActivityListAdminFiltered(name, status)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return activities, nil
}

func GetActivityByID(id string) (*model.Activity, error) {
	activityID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	activity, err := mysql.GetActivityByID(activityID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.ActivityNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return activity, nil
}

func CreateActivity(activity *model.Activity) error {
	if err := mysql.CreateActivity(activity); err != nil {
		return errno.New(errno.InternalError)
	}
	redis.Del("activity:list")
	return nil
}

func UpdateActivity(id string, activity *model.Activity) error {
	activityID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	existing, err := mysql.GetActivityByID(activityID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errno.New(errno.ActivityNotFound)
		}
		return errno.New(errno.InternalError)
	}
	existing.Name = activity.Name
	existing.Description = activity.Description
	existing.Image = activity.Image
	existing.Discount = activity.Discount
	existing.ValidFrom = activity.ValidFrom
	existing.ValidTo = activity.ValidTo
	existing.Status = activity.Status
	existing.SortOrder = activity.SortOrder
	if err := mysql.UpdateActivity(existing); err != nil {
		return errno.New(errno.InternalError)
	}
	redis.Del("activity:list")
	return nil
}

func DeleteActivity(id string) error {
	activityID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	if err := mysql.DeleteActivity(activityID); err != nil {
		return errno.New(errno.InternalError)
	}
	redis.Del("activity:list")
	return nil
}
