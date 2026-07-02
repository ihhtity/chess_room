package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"chess-room-backend/pkg/log"
)

func InitDefaultData(db *gorm.DB) {
	initDefaultAdmin(db)
	initDefaultRoomTypes(db)
	initDefaultTimeSlots(db)
	initDefaultRooms(db)
	initTestUsers(db)
	initTestOrders(db)
	initTestMemberships(db)
	initTestPayments(db)
	initTestRechargeRecords(db)
	log.Logger.Info("Default data initialized successfully")
}

func initDefaultAdmin(db *gorm.DB) {
	var count int
	db.Model(&Admin{}).Count(&count)
	if count > 0 {
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.Error("Failed to hash default admin password:", err)
		return
	}

	admin := Admin{
		Username: "admin",
		Password: string(passwordHash),
		Realname: "超级管理员",
		Role:     0,
		Status:   1,
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Logger.Error("Failed to create default admin:", err)
	} else {
		log.Logger.Info("Default admin created: username=admin, password=123456")
	}
}

func initDefaultRoomTypes(db *gorm.DB) {
	var count int
	db.Model(&RoomType{}).Count(&count)
	if count > 0 {
		return
	}

	roomTypes := []RoomType{
		{
			Name:        "普通棋牌室",
			Description: "适合2-4人使用，配备基础棋牌设施",
			BasePrice:   50.00,
			MaxPeople:   4,
			SortOrder:   1,
			Status:      1,
		},
		{
			Name:        "豪华棋牌室",
			Description: "适合4-6人使用，配备高档棋牌设施和独立卫生间",
			BasePrice:   80.00,
			MaxPeople:   6,
			SortOrder:   2,
			Status:      1,
		},
		{
			Name:        "VIP包间",
			Description: "适合6-8人使用，豪华装修，配备麻将机和茶水服务",
			BasePrice:   120.00,
			MaxPeople:   8,
			SortOrder:   3,
			Status:      1,
		},
	}

	for _, rt := range roomTypes {
		if err := db.Create(&rt).Error; err != nil {
			log.Logger.Error("Failed to create room type:", err)
		}
	}

	log.Logger.Info(fmt.Sprintf("Default room types created: %d", len(roomTypes)))
}

func initDefaultTimeSlots(db *gorm.DB) {
	var count int
	db.Model(&TimeSlot{}).Count(&count)
	if count > 0 {
		return
	}

	timeSlots := []TimeSlot{
		{
			TypeID:       1,
			Name:         "上午场",
			StartTime:    "09:00:00",
			EndTime:      "12:00:00",
			Price:        50.00,
			WeekdayPrice: 50.00,
			WeekendPrice: 60.00,
			HolidayPrice: 70.00,
			SortOrder:    1,
			Status:       1,
		},
		{
			TypeID:       1,
			Name:         "下午场",
			StartTime:    "13:00:00",
			EndTime:      "18:00:00",
			Price:        80.00,
			WeekdayPrice: 80.00,
			WeekendPrice: 100.00,
			HolidayPrice: 120.00,
			SortOrder:    2,
			Status:       1,
		},
		{
			TypeID:       1,
			Name:         "晚场",
			StartTime:    "18:00:00",
			EndTime:      "23:00:00",
			Price:        100.00,
			WeekdayPrice: 100.00,
			WeekendPrice: 120.00,
			HolidayPrice: 150.00,
			SortOrder:    3,
			Status:       1,
		},
	}

	for _, ts := range timeSlots {
		if err := db.Create(&ts).Error; err != nil {
			log.Logger.Error("Failed to create time slot:", err)
		}
	}

	log.Logger.Info(fmt.Sprintf("Default time slots created: %d", len(timeSlots)))
}

func initDefaultRooms(db *gorm.DB) {
	var count int
	db.Model(&Room{}).Count(&count)
	if count > 0 {
		return
	}

	rooms := []Room{
		{
			Name:        "101室",
			TypeID:      1,
			Floor:       "1",
			Capacity:    4,
			Equipment:   `["麻将机","扑克牌","茶几","沙发"]`,
			Images:      `["https://example.com/room101-1.jpg","https://example.com/room101-2.jpg"]`,
			Description: "阳光充足，通风良好",
			Status:      1,
			SortOrder:   1,
		},
		{
			Name:        "102室",
			TypeID:      1,
			Floor:       "1",
			Capacity:    4,
			Equipment:   `["麻将机","扑克牌","茶几","沙发"]`,
			Images:      `["https://example.com/room102-1.jpg","https://example.com/room102-2.jpg"]`,
			Description: "安静舒适，适合朋友聚会",
			Status:      1,
			SortOrder:   2,
		},
		{
			Name:        "201室",
			TypeID:      2,
			Floor:       "2",
			Capacity:    6,
			Equipment:   `["麻将机","扑克牌","茶几","沙发","独立卫生间"]`,
			Images:      `["https://example.com/room201-1.jpg","https://example.com/room201-2.jpg"]`,
			Description: "豪华装修，空间宽敞",
			Status:      1,
			SortOrder:   3,
		},
		{
			Name:        "202室",
			TypeID:      2,
			Floor:       "2",
			Capacity:    6,
			Equipment:   `["麻将机","扑克牌","茶几","沙发","独立卫生间"]`,
			Images:      `["https://example.com/room202-1.jpg","https://example.com/room202-2.jpg"]`,
			Description: "景观优美，视野开阔",
			Status:      1,
			SortOrder:   4,
		},
		{
			Name:        "301室",
			TypeID:      3,
			Floor:       "3",
			Capacity:    8,
			Equipment:   `["麻将机","扑克牌","茶几","沙发","独立卫生间","茶水服务"]`,
			Images:      `["https://example.com/room301-1.jpg","https://example.com/room301-2.jpg"]`,
			Description: "VIP专属包间，豪华配置",
			Status:      1,
			SortOrder:   5,
		},
	}

	for _, room := range rooms {
		if err := db.Create(&room).Error; err != nil {
			log.Logger.Error("Failed to create room:", err)
		}
	}

	log.Logger.Info(fmt.Sprintf("Default rooms created: %d", len(rooms)))
}

func initTestUsers(db *gorm.DB) {
	var count int
	db.Model(&User{}).Count(&count)
	if count > 0 {
		return
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

	users := []User{
		{
			OpenID:   "test_openid_1",
			Phone:    "13800138001",
			Password: string(passwordHash),
			Nickname: "张三",
			Realname: "张三",
			Avatar:   "https://example.com/avatar1.jpg",
			Gender:   1,
			Status:   1,
		},
		{
			OpenID:   "test_openid_2",
			Phone:    "13800138002",
			Password: string(passwordHash),
			Nickname: "李四",
			Realname: "李四",
			Avatar:   "https://example.com/avatar2.jpg",
			Gender:   1,
			Status:   1,
		},
		{
			OpenID:   "test_openid_3",
			Phone:    "13800138003",
			Password: string(passwordHash),
			Nickname: "王五",
			Realname: "王五",
			Avatar:   "https://example.com/avatar3.jpg",
			Gender:   2,
			Status:   1,
		},
		{
			OpenID:   "test_openid_4",
			Phone:    "13800138004",
			Password: string(passwordHash),
			Nickname: "赵六",
			Realname: "赵六",
			Avatar:   "https://example.com/avatar4.jpg",
			Gender:   1,
			Status:   1,
		},
		{
			OpenID:   "test_openid_5",
			Phone:    "13800138005",
			Password: string(passwordHash),
			Nickname: "钱七",
			Realname: "钱七",
			Avatar:   "https://example.com/avatar5.jpg",
			Gender:   2,
			Status:   1,
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Logger.Error("Failed to create test user:", err)
		}
	}

	log.Logger.Info(fmt.Sprintf("Test users created: %d", len(users)))
}

func initTestOrders(db *gorm.DB) {
	var count int
	db.Model(&Order{}).Count(&count)
	if count > 0 {
		return
	}

	now := time.Now()
	orders := []Order{
		{
			OrderNo:     "ORD" + now.Format("20060102150405") + "001",
			UserID:      1,
			RoomID:      1,
			StartTime:   now.Add(time.Hour * 2),
			EndTime:     now.Add(time.Hour * 5),
			Duration:    180,
			Status:      0,
			TotalAmount: 150.00,
			PaidAmount:  0,
			Remark:      "朋友聚会",
		},
		{
			OrderNo:     "ORD" + now.Format("20060102150405") + "002",
			UserID:      2,
			RoomID:      2,
			StartTime:   now.Add(time.Hour * 3),
			EndTime:     now.Add(time.Hour * 8),
			Duration:    300,
			Status:      1,
			TotalAmount: 250.00,
			PaidAmount:  250.00,
			PaidAt:      &now,
		},
		{
			OrderNo:     "ORD" + now.Add(-time.Hour*24).Format("20060102150405") + "003",
			UserID:      3,
			RoomID:      3,
			StartTime:   now.Add(-time.Hour * 24),
			EndTime:     now.Add(-time.Hour*24 + time.Hour*5),
			Duration:    300,
			Status:      2,
			TotalAmount: 400.00,
			PaidAmount:  400.00,
			PaidAt:      func() *time.Time { t := now.Add(-time.Hour * 24); return &t }(),
			CompletedAt: func() *time.Time { t := now.Add(-time.Hour*24 + time.Hour*5); return &t }(),
		},
		{
			OrderNo:     "ORD" + now.Add(-time.Hour*48).Format("20060102150405") + "004",
			UserID:      4,
			RoomID:      4,
			StartTime:   now.Add(-time.Hour * 48),
			EndTime:     now.Add(-time.Hour*48 + time.Hour*3),
			Duration:    180,
			Status:      3,
			TotalAmount: 240.00,
			PaidAmount:  0,
			CancelTime:  func() *time.Time { t := now.Add(-time.Hour*48 + time.Hour); return &t }(),
		},
	}

	for _, order := range orders {
		if err := db.Create(&order).Error; err != nil {
			log.Logger.Error("Failed to create test order:", err)
		}
	}

	log.Logger.Info(fmt.Sprintf("Test orders created: %d", len(orders)))
}

func initTestMemberships(db *gorm.DB) {
	var count int
	db.Model(&Membership{}).Count(&count)
	if count > 0 {
		return
	}

	now := time.Now()
	memberships := []Membership{
		{
			UserID:           1,
			Level:            2,
			Points:           2500,
			Balance:          500.00,
			TotalConsumed:    1500.00,
			TotalRecharged:   2000.00,
			MembershipStatus: 1,
			JoinedAt:         now.Add(-time.Hour * 720),
		},
		{
			UserID:           2,
			Level:            1,
			Points:           800,
			Balance:          200.00,
			TotalConsumed:    400.00,
			TotalRecharged:   500.00,
			MembershipStatus: 1,
			JoinedAt:         now.Add(-time.Hour * 360),
		},
		{
			UserID:           3,
			Level:            3,
			Points:           6000,
			Balance:          1200.00,
			TotalConsumed:    3800.00,
			TotalRecharged:   5000.00,
			MembershipStatus: 1,
			JoinedAt:         now.Add(-time.Hour * 1440),
		},
		{
			UserID:           4,
			Level:            0,
			Points:           100,
			Balance:          0,
			TotalConsumed:    100.00,
			TotalRecharged:   0,
			MembershipStatus: 1,
			JoinedAt:         now.Add(-time.Hour * 24),
		},
	}

	for _, membership := range memberships {
		if err := db.Create(&membership).Error; err != nil {
			log.Logger.Error("Failed to create test membership:", err)
		}
	}

	log.Logger.Info(fmt.Sprintf("Test memberships created: %d", len(memberships)))
}

func initTestPayments(db *gorm.DB) {
	var count int
	db.Model(&Payment{}).Count(&count)
	if count > 0 {
		return
	}

	now := time.Now()
	payments := []Payment{
		{
			OrderID:       2,
			UserID:        2,
			Amount:        250.00,
			PaymentType:   PaymentTypeWechat,
			Status:        1,
			TransactionNo: "TXN" + now.Format("20060102150405") + "001",
			PaidAt:        &now,
		},
		{
			OrderID:       3,
			UserID:        3,
			Amount:        400.00,
			PaymentType:   PaymentTypeWechat,
			Status:        1,
			TransactionNo: "TXN" + now.Add(-time.Hour*24).Format("20060102150405") + "002",
			PaidAt:        func() *time.Time { t := now.Add(-time.Hour * 24); return &t }(),
		},
	}

	for _, payment := range payments {
		if err := db.Create(&payment).Error; err != nil {
			log.Logger.Error("Failed to create test payment:", err)
		}
	}

	log.Logger.Info(fmt.Sprintf("Test payments created: %d", len(payments)))
}

func initTestRechargeRecords(db *gorm.DB) {
	var count int
	db.Model(&RechargeRecord{}).Count(&count)
	if count > 0 {
		return
	}

	now := time.Now()
	records := []RechargeRecord{
		{
			UserID:     1,
			Amount:     1000.00,
			GiftAmount: 200.00,
			Status:     1,
			CreatedAt:  now.Add(-time.Hour * 720),
		},
		{
			UserID:     1,
			Amount:     500.00,
			GiftAmount: 80.00,
			Status:     1,
			CreatedAt:  now.Add(-time.Hour * 360),
		},
		{
			UserID:     2,
			Amount:     500.00,
			GiftAmount: 80.00,
			Status:     1,
			CreatedAt:  now.Add(-time.Hour * 360),
		},
		{
			UserID:     3,
			Amount:     2000.00,
			GiftAmount: 400.00,
			Status:     1,
			CreatedAt:  now.Add(-time.Hour * 1440),
		},
		{
			UserID:     3,
			Amount:     1000.00,
			GiftAmount: 200.00,
			Status:     1,
			CreatedAt:  now.Add(-time.Hour * 720),
		},
	}

	for _, record := range records {
		if err := db.Create(&record).Error; err != nil {
			log.Logger.Error("Failed to create test recharge record:", err)
		}
	}

	log.Logger.Info(fmt.Sprintf("Test recharge records created: %d", len(records)))
}
