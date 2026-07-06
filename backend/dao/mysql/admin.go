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

func GetAdminList(username, realname string, roleID int64, status, maxLevel, page, pageSize int) ([]model.Admin, int64, error) {
	var admins []model.Admin
	var total int64
	query := DB.Model(&model.Admin{}).Order("id DESC")

	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}
	if realname != "" {
		query = query.Where("realname LIKE ?", "%"+realname+"%")
	}
	if roleID != 0 {
		query = query.Where("role_id = ?", roleID)
	}
	if status >= 0 {
		query = query.Where("status = ?", status)
	}

	query = query.Joins("LEFT JOIN admin_roles ar ON admins.role_id = ar.id").
		Where("ar.level >= ? OR admins.role_id = 0", maxLevel)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if page > 0 && pageSize > 0 {
		query = DB.Preload("Role").Order("id DESC")

		if username != "" {
			query = query.Where("username LIKE ?", "%"+username+"%")
		}
		if realname != "" {
			query = query.Where("realname LIKE ?", "%"+realname+"%")
		}
		if roleID != 0 {
			query = query.Where("role_id = ?", roleID)
		}
		if status >= 0 {
			query = query.Where("status = ?", status)
		}

		query = query.Joins("LEFT JOIN admin_roles ar ON admins.role_id = ar.id").
			Where("ar.level >= ? OR admins.role_id = 0", maxLevel).
			Offset((page - 1) * pageSize).Limit(pageSize)
	}

	err := query.Find(&admins).Error
	return admins, total, err
}
