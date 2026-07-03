package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetAnnouncementList() ([]model.Announcement, error) {
	announcements, err := mysql.GetAnnouncementList()
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return announcements, nil
}

func GetAnnouncementListFiltered(title string, typeInt, statusInt int) ([]model.Announcement, error) {
	announcements, err := mysql.GetAnnouncementListFiltered(title, typeInt, statusInt)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return announcements, nil
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
