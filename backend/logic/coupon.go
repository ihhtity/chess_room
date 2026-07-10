package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"errors"

	"github.com/jinzhu/gorm"
)

func GetCouponList(userID int64) ([]model.Coupon, int64, error) {
	var userCoupons []model.UserCoupon
	var total int64

	err := mysql.DB.Preload("Coupon").Where("user_id = ?", userID).Order("created_at DESC").Find(&userCoupons).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var coupons []model.Coupon
	for _, uc := range userCoupons {
		coupon := uc.Coupon
		coupon.Status = uc.Status
		coupons = append(coupons, coupon)
	}

	return coupons, total, nil
}

func GetCouponDetail(userID, couponID int64) (*model.Coupon, error) {
	var userCoupon model.UserCoupon
	err := mysql.DB.Preload("Coupon").Where("user_id = ? AND id = ?", userID, couponID).First(&userCoupon).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	coupon := userCoupon.Coupon
	coupon.Status = userCoupon.Status
	return &coupon, nil
}

func UseCoupon(userID, couponID, orderID int64) error {
	var userCoupon model.UserCoupon
	err := mysql.DB.Preload("Coupon").Where("user_id = ? AND id = ? AND status = 0", userID, couponID).First(&userCoupon).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("优惠券不存在或已使用")
		}
		return err
	}

	userCoupon.Status = 1
	userCoupon.OrderID = orderID
	return mysql.DB.Save(&userCoupon).Error
}
