package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetRoomList() ([]model.Room, error) {
	rooms, err := mysql.GetRoomList(0, 0)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return rooms, nil
}

func GetRoomListByTypeID(typeID string) ([]model.Room, error) {
	id, err := strconv.Atoi(typeID)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	rooms, err := mysql.GetRoomListByTypeID(id)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return rooms, nil
}

func GetRoomListByFloor(floor string) ([]model.Room, error) {
	rooms, err := mysql.GetRoomListByFloor(floor)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return rooms, nil
}

func GetRoomListByStatus(status string) ([]model.Room, error) {
	s, err := strconv.Atoi(status)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	rooms, err := mysql.GetRoomListByStatus(s)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return rooms, nil
}

func GetRoomByID(id string) (*model.Room, error) {
	roomID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	room, err := mysql.GetRoomByID(roomID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.RoomNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return room, nil
}

func CreateRoom(room *model.Room) error {
	if err := mysql.CreateRoom(room); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func UpdateRoom(room *model.Room) error {
	if err := mysql.UpdateRoom(room); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func DeleteRoom(id string) error {
	roomID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	if err := mysql.DeleteRoom(roomID); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func GetRoomTypeList() ([]model.RoomType, error) {
	types, err := mysql.GetRoomTypeList(0)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return types, nil
}

func GetRoomTypeByID(id string) (*model.RoomType, error) {
	typeID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	roomType, err := mysql.GetRoomTypeByID(typeID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.RoomTypeNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return roomType, nil
}

func CreateRoomType(roomType *model.RoomType) error {
	if err := mysql.CreateRoomType(roomType); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func UpdateRoomType(roomType *model.RoomType) error {
	if err := mysql.UpdateRoomType(roomType); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func DeleteRoomType(id string) error {
	typeID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	if err := mysql.DeleteRoomType(typeID); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}
