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

func GetUserList(users *[]model.User) error {
	return DB.Order("created_at DESC").Find(users).Error
}