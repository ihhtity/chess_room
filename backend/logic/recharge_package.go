package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetRechargePackageList(page, pageSize int) ([]model.RechargePackage, int64, error) {
	packages, total, err := mysql.GetRechargePackageList(page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return packages, total, nil
}

func GetRechargePackageListFiltered(name string, status, page, pageSize int) ([]model.RechargePackage, int64, error) {
	packages, total, err := mysql.GetRechargePackageListFiltered(name, status, page, pageSize)
	if err != nil {
		return nil, 0, errno.New(errno.InternalError)
	}
	return packages, total, nil
}

func GetRechargePackageByID(id string) (*model.RechargePackage, error) {
	packageID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errno.New(errno.BadRequest)
	}
	pkg, err := mysql.GetRechargePackageByID(packageID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.RechargePackageNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return pkg, nil
}

func CreateRechargePackage(pkg *model.RechargePackage) error {
	if err := mysql.CreateRechargePackage(pkg); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func UpdateRechargePackage(id string, pkg *model.RechargePackage) error {
	packageID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	existing, err := mysql.GetRechargePackageByID(packageID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errno.New(errno.RechargePackageNotFound)
		}
		return errno.New(errno.InternalError)
	}
	existing.Name = pkg.Name
	existing.Amount = pkg.Amount
	existing.GiftAmount = pkg.GiftAmount
	existing.GiftPoints = pkg.GiftPoints
	existing.Description = pkg.Description
	existing.Status = pkg.Status
	if err := mysql.UpdateRechargePackage(existing); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func DeleteRechargePackage(id string) error {
	packageID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errno.New(errno.BadRequest)
	}
	if err := mysql.DeleteRechargePackage(packageID); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchDeleteRechargePackage(ids []string) error {
	var packageIDs []int64
	for _, id := range ids {
		packageID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errno.New(errno.BadRequest)
		}
		packageIDs = append(packageIDs, packageID)
	}
	if err := mysql.BatchDeleteRechargePackage(packageIDs); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func BatchUpdateRechargePackage(reqs []struct {
	ID     int64 `json:"id"`
	Status int   `json:"status"`
}) error {
	for _, req := range reqs {
		pkg, err := mysql.GetRechargePackageByID(req.ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errno.New(errno.RechargePackageNotFound)
			}
			return errno.New(errno.InternalError)
		}

		pkg.Status = req.Status

		if err := mysql.UpdateRechargePackage(pkg); err != nil {
			return errno.New(errno.InternalError)
		}
	}
	return nil
}
