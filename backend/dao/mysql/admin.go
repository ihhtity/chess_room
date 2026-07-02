package mysql

import (
	"chess-room-backend/model"
)

func GetAdminByUsername(username string) (*model.Admin, error) {
	var admin model.Admin
	err := DB.Where("username = ?", username).First(&admin).Error
	return &admin, err
}

func GetAdminByID(id int64) (*model.Admin, error) {
	var admin model.Admin
	err := DB.Where("id = ?", id).First(&admin).Error
	return &admin, err
}

func CreateAdmin(admin *model.Admin) error {
	return DB.Create(admin).Error
}

func UpdateAdmin(admin *model.Admin) error {
	return DB.Save(admin).Error
}

func DeleteAdmin(id int64) error {
	return DB.Delete(&model.Admin{}, id).Error
}