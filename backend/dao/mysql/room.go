package mysql

import (
	"chess-room-backend/model"
)

func GetRoomList(status, typeID int) ([]model.Room, error) {
	var rooms []model.Room
	db := DB.Preload("Type").Order("sort_order ASC")
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	if typeID != 0 {
		db = db.Where("type_id = ?", typeID)
	}
	err := db.Find(&rooms).Error
	return rooms, err
}

func GetRoomListByTypeID(typeID int) ([]model.Room, error) {
	var rooms []model.Room
	err := DB.Preload("Type").Where("type_id = ?", typeID).Order("sort_order ASC").Find(&rooms).Error
	return rooms, err
}

func GetRoomListByFloor(floor string) ([]model.Room, error) {
	var rooms []model.Room
	err := DB.Preload("Type").Where("floor = ?", floor).Order("sort_order ASC").Find(&rooms).Error
	return rooms, err
}

func GetRoomListByStatus(status int) ([]model.Room, error) {
	var rooms []model.Room
	err := DB.Preload("Type").Where("status = ?", status).Order("sort_order ASC").Find(&rooms).Error
	return rooms, err
}

func GetRoomByID(id int64) (*model.Room, error) {
	var room model.Room
	err := DB.Preload("Type").Where("id = ?", id).First(&room).Error
	return &room, err
}

func CreateRoom(room *model.Room) error {
	return DB.Create(room).Error
}

func UpdateRoom(room *model.Room) error {
	return DB.Save(room).Error
}

func DeleteRoom(id int64) error {
	return DB.Delete(&model.Room{}, id).Error
}
