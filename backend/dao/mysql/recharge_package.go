package mysql

import (
	"chess-room-backend/model"
)

func GetRechargePackageList() ([]model.RechargePackage, error) {
	var packages []model.RechargePackage
	err := DB.Where("status = ?", 1).Order("sort_order ASC").Find(&packages).Error
	return packages, err
}

func GetRechargePackageListFiltered(name string, status int) ([]model.RechargePackage, error) {
	var packages []model.RechargePackage
	db := DB.Order("sort_order ASC")
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if status >= 0 {
		db = db.Where("status = ?", status)
	}
	err := db.Find(&packages).Error
	return packages, err
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
