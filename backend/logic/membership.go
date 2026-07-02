package logic

import (
	"chess-room-backend/dao/mysql"
	"chess-room-backend/model"
	"chess-room-backend/pkg/errno"
	"time"

	"github.com/jinzhu/gorm"
)

func GetMembership(userID int64) (*model.Membership, error) {
	membership, err := mysql.GetMembershipByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			membership = &model.Membership{
				UserID:           userID,
				Level:            0,
				Balance:          0,
				Points:           0,
				TotalConsumed:    0,
				TotalRecharged:   0,
				MembershipStatus: 1,
				JoinedAt:         time.Now(),
			}
			if err := mysql.CreateMembership(membership); err != nil {
				return nil, errno.New(errno.InternalError)
			}
		} else {
			return nil, errno.New(errno.InternalError)
		}
	}
	return membership, nil
}

func Recharge(userID int64, amount float64) (*model.Membership, error) {
	membership, err := GetMembership(userID)
	if err != nil {
		return nil, err
	}

	giftAmount := calculateGiftAmount(amount)
	actualAmount := amount + giftAmount

	membership.Balance += actualAmount
	membership.TotalRecharged += amount
	membership.Points += int(amount)
	updateLevel(membership)

	if err := mysql.UpdateMembership(membership); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return membership, nil
}

func calculateGiftAmount(amount float64) float64 {
	switch {
	case amount >= 1000:
		return 200
	case amount >= 500:
		return 80
	case amount >= 200:
		return 20
	default:
		return 0
	}
}

func updateLevel(membership *model.Membership) {
	total := membership.TotalConsumed + membership.TotalRecharged
	switch {
	case total >= 5000:
		membership.Level = 3
	case total >= 2000:
		membership.Level = 2
	case total >= 500:
		membership.Level = 1
	default:
		membership.Level = 0
	}
}

func Consume(userID int64, amount float64) (*model.Membership, error) {
	membership, err := GetMembership(userID)
	if err != nil {
		return nil, err
	}

	if membership.Balance < amount {
		return nil, errno.New(errno.BalanceInsufficient)
	}

	membership.Balance -= amount
	membership.TotalConsumed += amount
	membership.Points += int(amount * 0.5)
	updateLevel(membership)

	if err := mysql.UpdateMembership(membership); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return membership, nil
}

func GetMembershipList(level, status int) ([]model.Membership, error) {
	memberships, err := mysql.GetMembershipList(level, status)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return memberships, nil
}

func GetMembershipByID(id int64) (*model.Membership, error) {
	membership, err := mysql.GetMembershipByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.MembershipNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}
	return membership, nil
}

func UpdateMembership(id int64, level *int, balance *float64, points *int, membershipStatus *int, expiredAt *time.Time) (*model.Membership, error) {
	membership, err := mysql.GetMembershipByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.New(errno.MembershipNotFound)
		}
		return nil, errno.New(errno.InternalError)
	}

	if level != nil {
		membership.Level = *level
	}
	if balance != nil {
		membership.Balance = *balance
	}
	if points != nil {
		membership.Points = *points
	}
	if membershipStatus != nil {
		membership.MembershipStatus = *membershipStatus
	}
	if expiredAt != nil {
		membership.ExpiredAt = expiredAt
	}

	if err := mysql.UpdateMembership(membership); err != nil {
		return nil, errno.New(errno.InternalError)
	}

	return membership, nil
}

func DeleteMembership(id int64) error {
	if err := mysql.DeleteMembership(id); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func CreateMembership(membership *model.Membership) error {
	_, err := mysql.GetMembershipByUserID(membership.UserID)
	if err == nil {
		return errno.NewWithMessage(errno.BadRequest, "该用户已存在会员信息")
	}
	if err != gorm.ErrRecordNotFound {
		return errno.New(errno.InternalError)
	}

	if err := mysql.CreateMembership(membership); err != nil {
		return errno.New(errno.InternalError)
	}
	return nil
}

func GetRechargeRecords(userID int64) ([]model.RechargeRecord, error) {
	records, err := mysql.GetRechargeRecordList(userID, -1)
	if err != nil {
		return nil, errno.New(errno.InternalError)
	}
	return records, nil
}
