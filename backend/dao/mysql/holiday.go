package mysql

import (
	"chess-room-backend/model"
)

func GetHolidayList(isHoliday int) ([]model.Holiday, error) {
	var holidays []model.Holiday
	db := DB
	if isHoliday != -1 {
		db = db.Where("is_holiday = ?", isHoliday)
	}
	err := db.Order("date DESC").Find(&holidays).Error
	return holidays, err
}

func GetHolidayByID(id int64) (*model.Holiday, error) {
	var holiday model.Holiday
	err := DB.Where("id = ?", id).First(&holiday).Error
	return &holiday, err
}

func CreateHoliday(holiday *model.Holiday) error {
	return DB.Create(holiday).Error
}

func UpdateHolidayByID(id int64, data map[string]interface{}) error {
	return DB.Model(&model.Holiday{}).Where("id = ?", id).Updates(data).Error
}

func DeleteHoliday(id int64) error {
	return DB.Delete(&model.Holiday{}, id).Error
}
