package mysql

import (
	"chess-room-backend/model"
)

func GetUserByOpenID(openID string) (*model.User, error) {
	var user model.User
	err := DB.Where("openid = ?", openID).First(&user).Error
	return &user, err
}

func GetUserByPhone(phone string) (*model.User, error) {
	var user model.User
	err := DB.Where("phone = ?", phone).First(&user).Error
	return &user, err
}

func GetUserByID(id int64) (*model.User, error) {
	var user model.User
	err := DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func CreateUser(user *model.User) error {
	return DB.Create(user).Error
}

func UpdateUser(user *model.User) error {
	return DB.Save(user).Error
}

func DeleteUser(id int64) error {
	return DB.Delete(&model.User{}, id).Error
}

func BatchDeleteUser(ids []int64) error {
	return DB.Where("id IN (?)", ids).Delete(&model.User{}).Error
}

func GetUserList(users *[]model.User) error {
	return DB.Order("created_at DESC").Find(users).Error
}

func GetUserListFiltered(nickname, phone string, status, page, pageSize int) ([]model.User, int64, error) {
	var users []model.User
	var total int64
	db := DB.Model(&model.User{}).Order("created_at DESC")
	if nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+nickname+"%")
	}
	if phone != "" {
		db = db.Where("phone LIKE ?", "%"+phone+"%")
	}
	if status >= 0 {
		db = db.Where("status = ?", status)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&users).Error
	return users, total, err
}
