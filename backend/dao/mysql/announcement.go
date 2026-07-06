package mysql

import (
	"chess-room-backend/model"
)

func GetAnnouncementList(page, pageSize int) ([]model.Announcement, int64, error) {
	var announcements []model.Announcement
	var total int64
	db := DB.Model(&model.Announcement{}).Where("status = ?", 1)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB.Where("status = ?", 1).Order("sort_order ASC, created_at DESC").Offset((page - 1) * pageSize).Limit(pageSize)
	} else {
		db = DB.Where("status = ?", 1).Order("sort_order ASC, created_at DESC")
	}
	err := db.Find(&announcements).Error
	return announcements, total, err
}

func GetAnnouncementListFiltered(title string, typeInt, statusInt, page, pageSize int) ([]model.Announcement, int64, error) {
	var announcements []model.Announcement
	var total int64
	db := DB.Model(&model.Announcement{}).Order("sort_order ASC, created_at DESC")
	if title != "" {
		db = db.Where("title LIKE ?", "%"+title+"%")
	}
	if typeInt >= 0 {
		db = db.Where("type = ?", typeInt)
	}
	if statusInt >= 0 {
		db = db.Where("status = ?", statusInt)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB.Order("sort_order ASC, created_at DESC")
		if title != "" {
			db = db.Where("title LIKE ?", "%"+title+"%")
		}
		if typeInt >= 0 {
			db = db.Where("type = ?", typeInt)
		}
		if statusInt >= 0 {
			db = db.Where("status = ?", statusInt)
		}
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&announcements).Error
	return announcements, total, err
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

func BatchDeleteAnnouncement(ids []int64) error {
	return DB.Where("id IN (?)", ids).Delete(&model.Announcement{}).Error
}
