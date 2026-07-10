package mysql

import (
	"chess-room-backend/model"
)

func GetHolidayList(isHoliday, page, pageSize int) ([]model.Holiday, int64, error) {
	var holidays []model.Holiday
	var total int64
	db := DB.Model(&model.Holiday{})
	if isHoliday != -1 {
		db = db.Where("is_holiday = ?", isHoliday)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB
		if isHoliday != -1 {
			db = db.Where("is_holiday = ?", isHoliday)
		}
		db = db.Order("date DESC").Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&holidays).Error
	return holidays, total, err
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

func BatchDeleteHoliday(ids []int64) error {
	return DB.Where("id IN (?)", ids).Delete(&model.Holiday{}).Error
}

func GetHolidayByDate(date string) (*model.Holiday, error) {
	var holiday model.Holiday
	err := DB.Where("date = ?", date).First(&holiday).Error
	return &holiday, err
}
