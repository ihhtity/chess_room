package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetAnnouncementList(page, pageSize int) ([]model.Announcement, int64, error) {
	announcements, total, err := mysql.GetAnnouncementList(page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return announcements, total, nil
}

func GetAnnouncementListFiltered(title string, typeInt, statusInt, page, pageSize int) ([]model.Announcement, int64, error) {
	announcements, total, err := mysql.GetAnnouncementListFiltered(title, typeInt, statusInt, page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return announcements, total, nil
}

func GetAnnouncementByID(id string) (*model.Announcement, error) {
	announcementID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	announcement, err := mysql.GetAnnouncementByID(announcementID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.AnnouncementNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return announcement, nil
}

func CreateAnnouncement(announcement *model.Announcement) error {
	if err := mysql.CreateAnnouncement(announcement); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func UpdateAnnouncement(id string, announcement *model.Announcement) error {
	announcementID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	existing, err := mysql.GetAnnouncementByID(announcementID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errno.New(errno.AnnouncementNotFound)
		}
		return errno.New(errno.InternalError)
	}
	existing.Title = announcement.Title
	existing.Content = announcement.Content
	existing.Type = announcement.Type
	existing.Status = announcement.Status
	if err := mysql.UpdateAnnouncement(existing); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func DeleteAnnouncement(id string) error {
	announcementID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	if err := mysql.DeleteAnnouncement(announcementID); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchDeleteAnnouncement(ids []string) error {
	var announcementIDs []int64
	for _, id := range ids {
		announcementID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errno.New(errno.BadRequest)
		}
		announcementIDs = append(announcementIDs, announcementID)
	}
	if err := mysql.BatchDeleteAnnouncement(announcementIDs); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchUpdateAnnouncement(reqs []struct {
	ID     int64 `json:"id"`
	Status int   `json:"status"`
}) error {
	for _, req := range reqs {
		announcement, err := mysql.GetAnnouncementByID(req.ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errno.New(errno.AnnouncementNotFound)
			}
			return errno.New(errno.InternalError)
		}

		announcement.Status = req.Status

		if err := mysql.UpdateAnnouncement(announcement); err != nil {
			return errno.New(errno.InternalError)
		}
	}
	return nil
}
