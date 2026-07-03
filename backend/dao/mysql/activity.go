package mysql

import (
	"chess-room-backend/model"
	"time"
)

func GetActivityList() ([]model.Activity, error) {
	var activities []model.Activity
	err := DB.Where("status = ? AND valid_from <= ? AND valid_to >= ?", 1, time.Now(), time.Now()).
		Order("sort_order ASC").Find(&activities).Error
	return activities, err
}

func GetActivityListAdmin() ([]model.Activity, error) {
	var activities []model.Activity
	err := DB.Order("sort_order ASC, created_at DESC").Find(&activities).Error
	return activities, err
}

func GetActivityListAdminFiltered(name string, status int) ([]model.Activity, error) {
	var activities []model.Activity
	db := DB.Order("sort_order ASC, created_at DESC")
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if status >= 0 {
		db = db.Where("status = ?", status)
	}
	err := db.Find(&activities).Error
	return activities, err
}

func GetActivityByID(id int64) (*model.Activity, error) {
	var activity model.Activity
	err := DB.Where("id = ?", id).First(&activity).Error
	return &activity, err
}

func CreateActivity(activity *model.Activity) error {
	return DB.Create(activity).Error
}

func UpdateActivity(activity *model.Activity) error {
	return DB.Save(activity).Error
}

func DeleteActivity(id int64) error {
	return DB.Delete(&model.Activity{}, id).Error
}
