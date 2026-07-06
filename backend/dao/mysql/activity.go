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

func GetActivityListAdminFiltered(name string, status, page, pageSize int) ([]model.Activity, int64, error) {
	var activities []model.Activity
	var total int64
	db := DB.Model(&model.Activity{}).Order("sort_order ASC, created_at DESC")
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if status >= 0 {
		db = db.Where("status = ?", status)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB.Order("sort_order ASC, created_at DESC")
		if name != "" {
			db = db.Where("name LIKE ?", "%"+name+"%")
		}
		if status >= 0 {
			db = db.Where("status = ?", status)
		}
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&activities).Error
	return activities, total, err
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

func BatchDeleteActivity(ids []int64) error {
	return DB.Where("id IN (?)", ids).Delete(&model.Activity{}).Error
}
