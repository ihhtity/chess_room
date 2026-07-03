package mysql

import (
	"chess-room-backend/model"
)

func GetAnnouncementList() ([]model.Announcement, error) {
	var announcements []model.Announcement
	err := DB.Where("status = ?", 1).Order("sort_order ASC, created_at DESC").Find(&announcements).Error
	return announcements, err
}

func GetAnnouncementListFiltered(title string, typeInt, statusInt int) ([]model.Announcement, error) {
	var announcements []model.Announcement
	db := DB.Order("sort_order ASC, created_at DESC")
	if title != "" {
		db = db.Where("title LIKE ?", "%"+title+"%")
	}
	if typeInt >= 0 {
		db = db.Where("type = ?", typeInt)
	}
	if statusInt >= 0 {
		db = db.Where("status = ?", statusInt)
	}
	err := db.Find(&announcements).Error
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
