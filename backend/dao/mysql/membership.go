package mysql

import (
	"chess-room-backend/model"
)

func GetMembershipByUserID(userID int64) (*model.Membership, error) {
	var membership model.Membership
	err := DB.Preload("User").Where("user_id = ?", userID).First(&membership).Error
	return &membership, err
}

func GetMembershipByID(id int64) (*model.Membership, error) {
	var membership model.Membership
	err := DB.Preload("User").Where("id = ?", id).First(&membership).Error
	return &membership, err
}

func GetMembershipList(userID int64, level, status, page, pageSize int) ([]model.Membership, int64, error) {
	var memberships []model.Membership
	var total int64
	db := DB.Model(&model.Membership{})
	if userID != 0 {
		db = db.Where("user_id = ?", userID)
	}
	if level != 0 {
		db = db.Where("level = ?", level)
	}
	if status >= 0 {
		db = db.Where("membership_status = ?", status)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB.Preload("User")
		if userID != 0 {
			db = db.Where("user_id = ?", userID)
		}
		if level != 0 {
			db = db.Where("level = ?", level)
		}
		if status >= 0 {
			db = db.Where("membership_status = ?", status)
		}
		db = db.Order("created_at DESC").Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&memberships).Error
	return memberships, total, err
}

func CreateMembership(membership *model.Membership) error {
	return DB.Create(membership).Error
}

func UpdateMembership(membership *model.Membership) error {
	return DB.Save(membership).Error
}

func DeleteMembership(id int64) error {
	return DB.Delete(&model.Membership{}, id).Error
}

func BatchDeleteMembership(ids []int64) error {
	return DB.Where("id IN (?)", ids).Delete(&model.Membership{}).Error
}
