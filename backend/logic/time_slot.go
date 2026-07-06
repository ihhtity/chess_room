package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetTimeSlotList(typeID string, page, pageSize int) ([]model.TimeSlot, int64, error) {
	id := 0
	if typeID != "" {
		var err error
		id, err = strconv.Atoi(typeID)
		if err != nil {
			return nil, 0, errno.New(errno.BadRequest)
		}
	}
	slots, total, err := mysql.GetTimeSlotList(0, id, page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return slots, total, nil
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

func BatchDeleteTimeSlot(ids []string) error {
	var slotIDs []int64
	for _, id := range ids {
		slotID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errno.New(errno.BadRequest)
		}
		slotIDs = append(slotIDs, slotID)
	}
	if err := mysql.BatchDeleteTimeSlot(slotIDs); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchUpdateTimeSlot(reqs []struct {
	ID     int64 `json:"id"`
	Status int   `json:"status"`
}) error {
	for _, req := range reqs {
		slot, err := mysql.GetTimeSlotByID(req.ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errno.New(errno.TimeSlotNotFound)
			}
			return errno.New(errno.InternalError)
		}

		slot.Status = req.Status

		if err := mysql.UpdateTimeSlot(slot); err != nil {
			return errno.New(errno.InternalError)
		}
	}
	return nil
}
