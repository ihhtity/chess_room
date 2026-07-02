package errno

import (
	"fmt"
)

type ErrCode int

const (
	Success       ErrCode = 200
	InternalError ErrCode = 500
	BadRequest    ErrCode = 400
	Unauthorized  ErrCode = 401
	Forbidden     ErrCode = 403
	NotFound      ErrCode = 404

	UserNotFound    ErrCode = 1001
	UserExists      ErrCode = 1002
	PasswordError   ErrCode = 1003
	InvalidPassword ErrCode = 1005
	UserDisabled    ErrCode = 1004

	AdminNotFound ErrCode = 1101
	AdminExists   ErrCode = 1102
	AdminDisabled ErrCode = 1103

	RoomNotFound ErrCode = 2001
	RoomDisabled ErrCode = 2002
	RoomOccupied ErrCode = 2003

	RoomTypeNotFound ErrCode = 2101
	RoomTypeExists   ErrCode = 2102

	OrderNotFound     ErrCode = 3001
	OrderStatusError  ErrCode = 3002
	OrderAlreadyPaid  ErrCode = 3003
	OrderCancelFailed ErrCode = 3004

	PaymentNotFound    ErrCode = 4001
	PaymentFailed      ErrCode = 4002
	PaymentAlreadyDone ErrCode = 4003

	MembershipNotFound  ErrCode = 5001
	BalanceInsufficient ErrCode = 5002

	ActivityNotFound ErrCode = 7001
	ActivityDisabled ErrCode = 7002

	AnnouncementNotFound ErrCode = 7101

	RechargePackageNotFound ErrCode = 7201

	ReviewNotFound    ErrCode = 7301
	ReviewAlreadyDone ErrCode = 7302

	TimeSlotNotFound ErrCode = 2201

	WechatError ErrCode = 6001
)

var errMessages = map[ErrCode]string{
	Success:                 "success",
	InternalError:           "内部服务器错误",
	BadRequest:              "请求参数错误",
	Unauthorized:            "未授权",
	Forbidden:               "禁止访问",
	NotFound:                "资源不存在",
	UserNotFound:            "用户不存在",
	UserExists:              "用户已存在",
	PasswordError:           "密码错误",
	InvalidPassword:         "密码不正确",
	UserDisabled:            "用户已禁用",
	AdminNotFound:           "管理员不存在",
	AdminExists:             "管理员已存在",
	AdminDisabled:           "管理员已禁用",
	RoomNotFound:            "包间不存在",
	RoomDisabled:            "包间不可用",
	RoomOccupied:            "包间已被占用",
	RoomTypeNotFound:        "包间类型不存在",
	RoomTypeExists:          "包间类型已存在",
	OrderNotFound:           "订单不存在",
	OrderStatusError:        "订单状态错误",
	OrderAlreadyPaid:        "订单已支付",
	OrderCancelFailed:       "取消订单失败",
	PaymentNotFound:         "支付记录不存在",
	PaymentFailed:           "支付失败",
	PaymentAlreadyDone:      "支付已完成",
	MembershipNotFound:      "会员信息不存在",
	BalanceInsufficient:     "余额不足",
	ActivityNotFound:        "活动不存在",
	ActivityDisabled:        "活动已结束",
	AnnouncementNotFound:    "公告不存在",
	RechargePackageNotFound: "充值套餐不存在",
	ReviewNotFound:          "评价不存在",
	ReviewAlreadyDone:       "该订单已评价",
	TimeSlotNotFound:        "时间槽不存在",
	WechatError:             "微信接口错误",
}

type Error struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code=%d, message=%s", e.Code, e.Message)
}

func New(code ErrCode) *Error {
	return &Error{
		Code:    code,
		Message: errMessages[code],
	}
}

func NewWithMessage(code ErrCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
