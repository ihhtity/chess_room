package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetTimeSlotList(typeID string) ([]model.TimeSlot, error) {
	id := 0
	if typeID != "" {
		var err error
		id, err = strconv.Atoi(typeID)
		if err != nil {
			return nil, errno.New(errno.BadRequest)
		}
	}
	slots, err := mysql.GetTimeSlotList(0, id)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return slots, nil
}

func GetTimeSlotByID(id string) (*model.TimeSlot, error) {
	slotID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	slot, err := mysql.GetTimeSlotByID(slotID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.TimeSlotNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return slot, nil
}

func CreateTimeSlot(slot *model.TimeSlot) error {
	if err := mysql.CreateTimeSlot(slot); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func UpdateTimeSlot(slot *model.TimeSlot) error {
	if err := mysql.UpdateTimeSlot(slot); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func DeleteTimeSlot(id string) error {
	slotID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	if err := mysql.DeleteTimeSlot(slotID); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}
