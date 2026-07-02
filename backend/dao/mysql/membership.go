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

func GetMembershipList(level, status int) ([]model.Membership, error) {
	var memberships []model.Membership
	db := DB.Preload("User")
	if level != 0 {
		db = db.Where("level = ?", level)
	}
	if status != 0 {
		db = db.Where("membership_status = ?", status)
	}
	err := db.Order("created_at DESC").Find(&memberships).Error
	return memberships, err
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
