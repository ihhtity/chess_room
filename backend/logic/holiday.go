package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetHolidayList(isHoliday string, page, pageSize int) ([]model.Holiday, int64, error) {
	isHolidayInt := -1
	if isHoliday != "" {
		var err error
		isHolidayInt, err = strconv.Atoi(isHoliday)
		if err != nil {
			return nil, 0, errno.New(errno.BadRequest)
		}
	}
	holidays, total, err := mysql.GetHolidayList(isHolidayInt, page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return holidays, total, nil
}

func GetHolidayByID(id string) (*model.Holiday, error) {
	holidayID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	holiday, err := mysql.GetHolidayByID(holidayID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.NotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return holiday, nil
}

func CreateHoliday(holiday *model.Holiday) error {
	if err := mysql.CreateHoliday(holiday); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func UpdateHoliday(id string, data map[string]interface{}) (*model.Holiday, error) {
	holidayID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}

	_, err = mysql.GetHolidayByID(holidayID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.NotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	if err := mysql.UpdateHolidayByID(holidayID, data); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return mysql.GetHolidayByID(holidayID)
}

func DeleteHoliday(id string) error {
	holidayID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	if err := mysql.DeleteHoliday(holidayID); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchDeleteHoliday(ids []string) error {
	var holidayIDs []int64
	for _, id := range ids {
		holidayID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errno.New(errno.BadRequest)
		}
		holidayIDs = append(holidayIDs, holidayID)
	}
	if err := mysql.BatchDeleteHoliday(holidayIDs); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchUpdateHoliday(reqs []struct {
	ID        int64 `json:"id"`
	IsHoliday int   `json:"is_holiday"`
}) error {
	for _, req := range reqs {
		holiday, err := mysql.GetHolidayByID(req.ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errno.New(errno.NotFound)
			}
			return errno.New(errno.InternalError)
		}

		holiday.IsHoliday = req.IsHoliday

		if err := mysql.UpdateHolidayByID(req.ID, map[string]interface{}{"is_holiday": req.IsHoliday}); err != nil {
			return errno.New(errno.InternalError)
		}
	}
	return nil
}
