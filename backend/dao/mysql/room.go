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

func GetRoomListFiltered(typeID int, floor string, status int, name string, page, pageSize int) ([]model.Room, int64, error) {
	var rooms []model.Room
	var total int64
	db := DB.Model(&model.Room{}).Order("sort_order ASC")
	if typeID != 0 {
		db = db.Where("type_id = ?", typeID)
	}
	if floor != "" {
		db = db.Where("floor = ?", floor)
	}
	if status >= 0 {
		db = db.Where("status = ?", status)
	}
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB.Preload("Type").Order("sort_order ASC")
		if typeID != 0 {
			db = db.Where("type_id = ?", typeID)
		}
		if floor != "" {
			db = db.Where("floor = ?", floor)
		}
		if status >= 0 {
			db = db.Where("status = ?", status)
		}
		if name != "" {
			db = db.Where("name LIKE ?", "%"+name+"%")
		}
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&rooms).Error
	return rooms, total, err
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

func BatchDeleteRoom(ids []int64) error {
	return DB.Where("id IN (?)", ids).Delete(&model.Room{}).Error
}
