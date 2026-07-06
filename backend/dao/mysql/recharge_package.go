package mysql

import (
	"chess-room-backend/model"
)

func GetRechargePackageList(page, pageSize int) ([]model.RechargePackage, int64, error) {
	var packages []model.RechargePackage
	var total int64
	db := DB.Model(&model.RechargePackage{}).Where("status = ?", 1)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB.Where("status = ?", 1).Order("sort_order ASC").Offset((page - 1) * pageSize).Limit(pageSize)
	} else {
		db = DB.Where("status = ?", 1).Order("sort_order ASC")
	}
	err := db.Find(&packages).Error
	return packages, total, err
}

func GetRechargePackageListFiltered(name string, status, page, pageSize int) ([]model.RechargePackage, int64, error) {
	var packages []model.RechargePackage
	var total int64
	db := DB.Model(&model.RechargePackage{}).Order("sort_order ASC")
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if status >= 0 {
		db = db.Where("status = ?", status)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if page > 0 && pageSize > 0 {
		db = DB.Order("sort_order ASC")
		if name != "" {
			db = db.Where("name LIKE ?", "%"+name+"%")
		}
		if status >= 0 {
			db = db.Where("status = ?", status)
		}
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	err := db.Find(&packages).Error
	return packages, total, err
}

func GetRechargePackageByID(id int64) (*model.RechargePackage, error) {
	var pkg model.RechargePackage
	err := DB.Where("id = ?", id).First(&pkg).Error
	return &pkg, err
}

func CreateRechargePackage(pkg *model.RechargePackage) error {
	return DB.Create(pkg).Error
}

func UpdateRechargePackage(pkg *model.RechargePackage) error {
	return DB.Save(pkg).Error
}

func DeleteRechargePackage(id int64) error {
	return DB.Delete(&model.RechargePackage{}, id).Error
}

func BatchDeleteRechargePackage(ids []int64) error {
	return DB.Where("id IN (?)", ids).Delete(&model.RechargePackage{}).Error
}
