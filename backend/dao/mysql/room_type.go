package mysql

import (
	"chess-room-backend/model"
)

func GetRoomTypeList(status, page, pageSize int) ([]model.RoomType, int64, error) {
	var types []model.RoomType
	var total int64
	db := DB.Model(&model.RoomType{}).Order("sort_order ASC")
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB.Order("sort_order ASC")
		if status != 0 {
			db = db.Where("status = ?", status)
		}
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&types).Error
	return types, total, err
}

func GetRoomTypeByID(id int64) (*model.RoomType, error) {
	var roomType model.RoomType
	err := DB.Where("id = ?", id).First(&roomType).Error
	return &roomType, err
}

func CreateRoomType(roomType *model.RoomType) error {
	return DB.Create(roomType).Error
}

func UpdateRoomType(roomType *model.RoomType) error {
	return DB.Save(roomType).Error
}

func DeleteRoomType(id int64) error {
	return DB.Delete(&model.RoomType{}, id).Error
}

func BatchDeleteRoomType(ids []int64) error {
	return DB.Delete(&model.RoomType{}, ids).Error
}
