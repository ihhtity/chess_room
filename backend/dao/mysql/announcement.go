package mysql

import (
	"chess-room-backend/model"
)

func GetAnnouncementList() ([]model.Announcement, error) {
	var announcements []model.Announcement
	err := DB.Where("status = ?", 1).Order("sort_order ASC, created_at DESC").Find(&announcements).Error
	return announcements, err
}

func GetAnnouncementByID(id int64) (*model.Announcement, error) {
	var announcement model.Announcement
	err := DB.Where("id = ?", id).First(&announcement).Error
	return &announcement, err
}

func CreateAnnouncement(announcement *model.Announcement) error {
	return DB.Create(announcement).Error
}

func UpdateAnnouncement(announcement *model.Announcement) error {
	return DB.Save(announcement).Error
}

func DeleteAnnouncement(id int64) error {
	return DB.Delete(&model.Announcement{}, id).Error
}
