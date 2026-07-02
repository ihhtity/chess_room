package mysql

import (
	"chess-room-backend/model"
)

func GetRoomTypeList(status int) ([]model.RoomType, error) {
	var types []model.RoomType
	db := DB.Order("sort_order ASC")
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	err := db.Find(&types).Error
	return types, err
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