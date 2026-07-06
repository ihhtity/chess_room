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

func GetRoomList() ([]model.Room, error) {
	cacheKey := "room:list"
	if cacheData, err := redis.Get(cacheKey); err == nil && cacheData != "" {
		var rooms []model.Room
		if err := json.Unmarshal([]byte(cacheData), &rooms); err == nil {
			return rooms, nil
		}
	}

	rooms, err := mysql.GetRoomList(0, 0)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}

	if data, err := json.Marshal(rooms); err == nil {
		redis.Set(cacheKey, string(data), 300)
	}

	return rooms, nil
}

func GetRoomListFiltered(typeID int, floor string, status int, name string, page, pageSize int) ([]model.Room, int64, error) {
	rooms, total, err := mysql.GetRoomListFiltered(typeID, floor, status, name, page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return rooms, total, nil
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
	redis.Del("room:list")
	return nil
}

func UpdateRoom(room *model.Room) error {
	if err := mysql.UpdateRoom(room); err != nil {
		return errno.New(errno.InternalError)
	}
	redis.Del("room:list")
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
	redis.Del("room:list")
	return nil
}

func BatchDeleteRoom(ids []string) error {
	var roomIDs []int64
	for _, id := range ids {
		roomID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errno.New(errno.BadRequest)
		}
		roomIDs = append(roomIDs, roomID)
	}
	if err := mysql.BatchDeleteRoom(roomIDs); err != nil {
		return errno.New(errno.InternalError)
	}
	redis.Del("room:list")
	return nil
}

func BatchUpdateRoom(reqs []struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	TypeID   int64  `json:"type_id"`
	Floor    string `json:"floor"`
	Capacity int    `json:"capacity"`
	Images   string `json:"images"`
	Status   int    `json:"status"`
}) error {
	for _, req := range reqs {
		room, err := mysql.GetRoomByID(req.ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errno.New(errno.RoomNotFound)
			}
			return errno.New(errno.InternalError)
		}

		if req.Name != "" {
			room.Name = req.Name
		}
		if req.TypeID != 0 {
			room.TypeID = req.TypeID
		}
		if req.Floor != "" {
			room.Floor = req.Floor
		}
		if req.Capacity != 0 {
			room.Capacity = req.Capacity
		}
		if req.Images != "" {
			room.Images = req.Images
		}
		if req.Status >= 0 {
			room.Status = req.Status
		}

		if err := mysql.UpdateRoom(room); err != nil {
			return errno.New(errno.InternalError)
		}
	}

	redis.Del("room:list")
	return nil
}

func GetRoomTypeList(page, pageSize int) ([]model.RoomType, int64, error) {
	types, total, err := mysql.GetRoomTypeList(0, page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return types, total, nil
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
	redis.Del("room:type:list")
	return nil
}

func UpdateRoomType(roomType *model.RoomType) error {
	if err := mysql.UpdateRoomType(roomType); err != nil {
		return errno.New(errno.InternalError)
	}
	redis.Del("room:type:list")
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
	redis.Del("room:type:list")
	return nil
}

func BatchDeleteRoomType(ids []string) error {
	var typeIDs []int64
	for _, id := range ids {
		typeID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errno.New(errno.BadRequest)
		}
		typeIDs = append(typeIDs, typeID)
	}
	if err := mysql.BatchDeleteRoomType(typeIDs); err != nil {
		return errno.New(errno.InternalError)
	}
	redis.Del("room:type:list")
	return nil
}

func BatchUpdateRoomType(reqs []struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	BasePrice   float64 `json:"base_price"`
	MaxPeople   int     `json:"max_people"`
}) error {
	for _, req := range reqs {
		roomType, err := mysql.GetRoomTypeByID(req.ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errno.New(errno.RoomTypeNotFound)
			}
			return errno.New(errno.InternalError)
		}

		if req.Name != "" {
			roomType.Name = req.Name
		}
		if req.Description != "" {
			roomType.Description = req.Description
		}
		if req.BasePrice > 0 {
			roomType.BasePrice = req.BasePrice
		}
		if req.MaxPeople > 0 {
			roomType.MaxPeople = req.MaxPeople
		}

		if err := mysql.UpdateRoomType(roomType); err != nil {
			return errno.New(errno.InternalError)
		}
	}

	redis.Del("room:type:list")
	return nil
}
