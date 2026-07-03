package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	// 用户状态
	UserStatusNormal = 1
	// 用户状态：禁用
	UserStatusDisabled = 0
	// 管理员状态
	AdminStatusNormal = 1
	// 管理员状态：禁用
	AdminStatusDisabled = 0
	// 房间类型状态
	RoomTypeStatusNormal = 1
	// 房间类型状态：隐藏类型
	RoomTypeStatusHidden = 0
	// 房间状态
	RoomStatusMaintenance = 0
	// 房间状态：启用
	RoomStatusAvailable = 1
	// 房间状态：已使用
	RoomStatusInUse = 2
	// 房间状态：已预约
	RoomStatusReserved = 3
	// 订单状态
	OrderStatusPending = 0
	// 订单状态：已激活
	OrderStatusActive = 1
	// 订单状态：已完成
	OrderStatusCompleted = 2
	// 订单状态：已取消
	OrderStatusCancelled = 3
	// 订单状态：退款中
	OrderStatusRefunding = 4
	// 订单状态：已退款
	OrderStatusRefunded = 5
	// 支付状态
	PaymentStatusPending = 0
	// 支付状态：已支付
	PaymentStatusSuccess = 1
	// 支付状态：支付失败
	PaymentStatusFailed = 2
	// 会员状态
	MembershipStatusNormal = 1
	// 会员状态：已过期
	MembershipStatusExpired = 2
	// 支付方式
	PaymentTypeWechat = 1
	// 支付方式：支付宝
	PaymentTypeAlipay = 2
	// 支付方式：余额支付
	PaymentTypeBalance = 3
)

// 用户表
type User struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	OpenID    string    `gorm:"column:openid;size:100;unique_index" json:"open_id" comment:"微信openid"`
	Phone     string    `gorm:"column:phone;size:20;unique_index" json:"phone" comment:"手机号"`
	Password  string    `gorm:"column:password;size:255" json:"-" comment:"密码"`
	Nickname  string    `gorm:"column:nickname;size:100" json:"nickname" comment:"昵称"`
	Realname  string    `gorm:"column:realname;size:100" json:"realname" comment:"真实姓名"`
	Avatar    string    `gorm:"column:avatar;size:500" json:"avatar" comment:"头像"`
	Gender    int       `gorm:"column:gender;default:0" json:"gender" comment:"性别"`
	Status    int       `gorm:"column:status;default:1" json:"status" comment:"状态"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
}

// 用户表
func (User) TableName() string {
	return "users"
}

// 管理员表
type Admin struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	Username  string    `gorm:"column:username;size:50;unique_index" json:"username" comment:"用户名"`
	Password  string    `gorm:"column:password;size:255" json:"-" comment:"密码"`
	Realname  string    `gorm:"column:realname;size:100" json:"realname" comment:"真实姓名"`
	Phone     string    `gorm:"column:phone;size:20;unique_index" json:"phone" comment:"手机号"`
	Email     string    `gorm:"column:email;size:50;unique_index" json:"email" comment:"邮箱"`
	RoleType  int       `gorm:"column:role;default:1" json:"role" comment:"管理员类型"`
	RoleID    int64     `gorm:"column:role_id;default:1" json:"role_id" comment:"角色ID"`
	Role      AdminRole `gorm:"foreignkey:RoleID;association_autoupdate:false;association_autocreate:false" json:"role" comment:"角色信息"`
	Status    int       `gorm:"column:status;default:1" json:"status" comment:"状态"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
}

// 管理员表
func (Admin) TableName() string {
	return "admins"
}

// 管理员角色表
type AdminRole struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	Name        string    `gorm:"column:name;size:50;unique_index" json:"name" comment:"角色名称"`
	Level       int       `gorm:"column:level;default:1" json:"level" comment:"角色层级（数字越小层级越高）"`
	Description string    `gorm:"column:description;size:200" json:"description" comment:"角色描述"`
	Status      int       `gorm:"column:status;default:1" json:"status" comment:"状态"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
}

// 管理员角色表
func (AdminRole) TableName() string {
	return "admin_roles"
}

// 权限表
type Permission struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	Code        string    `gorm:"column:code;size:100;unique_index" json:"code" comment:"权限标识"`
	Name        string    `gorm:"column:name;size:100" json:"name" comment:"权限名称"`
	Group       string    `gorm:"column:group;size:50" json:"group" comment:"权限分组"`
	Module      string    `gorm:"column:module;size:50" json:"module" comment:"所属模块"`
	Description string    `gorm:"column:description;size:200" json:"description" comment:"描述"`
	SortOrder   int       `gorm:"column:sort_order;default:0" json:"sort_order" comment:"排序"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
}

// 权限表
func (Permission) TableName() string {
	return "permissions"
}

// 角色权限关联表
type RolePermission struct {
	RoleID       int64 `gorm:"column:role_id;index" json:"role_id" comment:"角色ID"`
	PermissionID int64 `gorm:"column:permission_id;index" json:"permission_id" comment:"权限ID"`
}

// 角色权限关联表
func (RolePermission) TableName() string {
	return "role_permissions"
}

// 房间类型表
type RoomType struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	Name        string    `gorm:"column:name;size:100;not null" json:"name" comment:"名称"`
	Description string    `gorm:"column:description;size:500" json:"description" comment:"描述"`
	BasePrice   float64   `gorm:"column:base_price;type:decimal(10,2);default:0" json:"base_price" comment:"基础价格"`
	MaxPeople   int       `gorm:"column:max_people;default:4" json:"max_people" comment:"最大人数"`
	SortOrder   int       `gorm:"column:sort_order;default:0" json:"sort_order" comment:"排序顺序"`
	Status      int       `gorm:"column:status;default:1" json:"status" comment:"状态"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
}

// 房间类型表
func (RoomType) TableName() string {
	return "room_types"
}

// 房间表
type Room struct {
	ID          int64      `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	Name        string     `gorm:"column:name;size:100;not null" json:"name" comment:"名称"`
	TypeID      int64      `gorm:"column:type_id;index" json:"type_id" comment:"类型类型ID"`
	Type        RoomType   `gorm:"foreignkey:TypeID" json:"type" comment:"类型类型"`
	Floor       string     `gorm:"column:floor;size:20" json:"floor" comment:"楼层"`
	Capacity    int        `gorm:"column:capacity;default:0" json:"capacity" comment:"容量人数"`
	Equipment   string     `gorm:"column:equipment;size:500" json:"equipment" comment:"设备"`
	Images      string     `gorm:"column:images;size:1000" json:"images" comment:"图片"`
	Description string     `gorm:"column:description;size:500" json:"description" comment:"描述"`
	Status      int        `gorm:"column:status;default:1" json:"status" comment:"状态"`
	SortOrder   int        `gorm:"column:sort_order;default:0" json:"sort_order" comment:"排序顺序"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"-" comment:"删除时间"`
}

// 房间表
func (Room) TableName() string {
	return "rooms"
}

// 时间槽表
type TimeSlot struct {
	ID           int64     `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	TypeID       int64     `gorm:"column:type_id;index" json:"type_id" comment:"类型类型ID"`
	Name         string    `gorm:"column:name;size:50;not null" json:"name" comment:"名称"`
	StartTime    string    `gorm:"column:start_time;size:10" json:"start_time" comment:"开始时间"`
	EndTime      string    `gorm:"column:end_time;size:10" json:"end_time" comment:"结束时间"`
	Price        float64   `gorm:"column:price;type:decimal(10,2);default:0" json:"price" comment:"价格"`
	WeekdayPrice float64   `gorm:"column:weekday_price;type:decimal(10,2);default:0" json:"weekday_price" comment:"工作日价格"`
	WeekendPrice float64   `gorm:"column:weekend_price;type:decimal(10,2);default:0" json:"weekend_price" comment:"周末价格"`
	HolidayPrice float64   `gorm:"column:holiday_price;type:decimal(10,2);default:0" json:"holiday_price" comment:"节假日价格"`
	SortOrder    int       `gorm:"column:sort_order;default:0" json:"sort_order" comment:"排序顺序"`
	Status       int       `gorm:"column:status;default:1" json:"status" comment:"状态"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
}

// 时间槽表
func (TimeSlot) TableName() string {
	return "time_slots"
}

// 订单表
type Order struct {
	ID          int64      `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	OrderNo     string     `gorm:"column:order_no;size:50;unique_index" json:"order_no" comment:"订单号"`
	UserID      int64      `gorm:"column:user_id;index" json:"user_id" comment:"用户ID"`
	User        User       `gorm:"foreignkey:UserID" json:"user" comment:"用户"`
	RoomID      int64      `gorm:"column:room_id;index" json:"room_id" comment:"房间ID"`
	Room        Room       `gorm:"foreignkey:RoomID" json:"room" comment:"房间"`
	StartTime   time.Time  `gorm:"column:start_time" json:"start_time" comment:"开始时间"`
	EndTime     time.Time  `gorm:"column:end_time" json:"end_time" comment:"结束时间"`
	Duration    int        `gorm:"column:duration;default:0" json:"duration" comment:"持续时间"`
	Status      int        `gorm:"column:status;default:0" json:"status" comment:"状态"`
	TotalAmount float64    `gorm:"column:total_amount;type:decimal(10,2);default:0" json:"total_amount" comment:"总金额"`
	PaidAmount  float64    `gorm:"column:paid_amount;type:decimal(10,2);default:0" json:"paid_amount" comment:"已支付金额"`
	Remark      string     `gorm:"column:remark;size:500" json:"remark" comment:"备注"`
	PaidAt      *time.Time `gorm:"column:paid_at" json:"paid_at" comment:"已支付时间"`
	CancelTime  *time.Time `gorm:"column:cancel_time" json:"cancel_time" comment:"取消时间"`
	CompletedAt *time.Time `gorm:"column:completed_at" json:"completed_at" comment:"完成时间"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
}

// 订单表
func (Order) TableName() string {
	return "orders"
}

// 支付表
type Payment struct {
	ID            int64      `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	OrderID       int64      `gorm:"column:order_id;index" json:"order_id" comment:"订单ID"`
	UserID        int64      `gorm:"column:user_id;index" json:"user_id" comment:"用户ID"`
	Amount        float64    `gorm:"column:amount;type:decimal(10,2);default:0" json:"amount" comment:"金额"`
	PaymentType   int        `gorm:"column:payment_type;default:0" json:"payment_type" comment:"支付类型"`
	Status        int        `gorm:"column:status;default:0" json:"status" comment:"状态"`
	TransactionNo string     `gorm:"column:transaction_no;size:100" json:"transaction_no" comment:"交易号"`
	PaidAt        *time.Time `gorm:"column:paid_at" json:"paid_at" comment:"已支付时间"`
	RefundedAt    *time.Time `gorm:"column:refunded_at" json:"refunded_at" comment:"已退款时间"`
	CreatedAt     time.Time  `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt     time.Time  `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
}

// 支付表
func (Payment) TableName() string {
	return "payments"
}

// 会员表
type Membership struct {
	ID               int64      `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	UserID           int64      `gorm:"column:user_id;unique_index" json:"user_id" comment:"用户ID"`
	User             User       `gorm:"foreignkey:UserID" json:"user" comment:"用户"`
	Level            int        `gorm:"column:level;default:0" json:"level" comment:"等级"`
	Points           int        `gorm:"column:points;default:0" json:"points" comment:"积分"`
	Balance          float64    `gorm:"column:balance;type:decimal(10,2);default:0" json:"balance" comment:"余额"`
	TotalConsumed    float64    `gorm:"column:total_consumed;type:decimal(10,2);default:0" json:"total_consumed" comment:"总消费金额"`
	TotalRecharged   float64    `gorm:"column:total_recharged;type:decimal(10,2);default:0" json:"total_recharged" comment:"总充值金额"`
	MembershipStatus int        `gorm:"column:membership_status;default:1" json:"membership_status" comment:"会员状态"`
	JoinedAt         time.Time  `gorm:"column:joined_at" json:"joined_at" comment:"加入时间"`
	ExpiredAt        *time.Time `gorm:"column:expired_at" json:"expired_at" comment:"过期时间"`
	CreatedAt        time.Time  `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt        time.Time  `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
}

// 会员表
func (Membership) TableName() string {
	return "memberships"
}

// 充值记录表
type RechargeRecord struct {
	ID         int64     `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	UserID     int64     `gorm:"column:user_id;index" json:"user_id" comment:"用户ID"`
	Amount     float64   `gorm:"column:amount;type:decimal(10,2);default:0" json:"amount" comment:"金额"`
	GiftAmount float64   `gorm:"column:gift_amount;type:decimal(10,2);default:0" json:"gift_amount" comment:"赠送金额"`
	PaymentID  int64     `gorm:"column:payment_id;index" json:"payment_id" comment:"支付ID"`
	Status     int       `gorm:"column:status;default:0" json:"status" comment:"状态"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
}

// 充值记录表
func (RechargeRecord) TableName() string {
	return "recharge_records"
}

// 通知表
type Notification struct {
	ID         int64      `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	UserID     int64      `gorm:"column:user_id;index" json:"user_id" comment:"用户ID"`
	Type       int        `gorm:"column:type;default:0" json:"type" comment:"类型"`
	Title      string     `gorm:"column:title;size:100" json:"title" comment:"标题"`
	Content    string     `gorm:"column:content;size:500" json:"content" comment:"内容"`
	ReadStatus int        `gorm:"column:read_status;default:0" json:"read_status" comment:"读取状态"`
	Link       string     `gorm:"column:link;size:500" json:"link" comment:"链接"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"-" comment:"删除时间"`
}

// 通知表
func (Notification) TableName() string {
	return "notifications"
}

// 操作日志表
type OperationLog struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	AdminID   int64     `gorm:"column:admin_id;index" json:"admin_id" comment:"管理员ID"`
	Action    string    `gorm:"column:action;size:100" json:"action" comment:"操作类型"`
	Module    string    `gorm:"column:module;size:50" json:"module" comment:"模块"`
	TargetID  int64     `gorm:"column:target_id" json:"target_id" comment:"目标ID"`
	Content   string    `gorm:"column:content;size:500" json:"content" comment:"操作内容"`
	IP        string    `gorm:"column:ip;size:50" json:"ip" comment:"IP"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
}

// 操作日志表
func (OperationLog) TableName() string {
	return "operation_logs"
}

// 节假日表
type Holiday struct {
	ID          int64      `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	Name        string     `gorm:"column:name;size:100" json:"name" comment:"名称"`
	Date        string     `gorm:"column:date;size:20" json:"date" comment:"日期"`
	IsHoliday   int        `gorm:"column:is_holiday;default:1" json:"is_holiday" comment:"是否节假日"`
	Description string     `gorm:"column:description;size:200" json:"description" comment:"描述"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"-" comment:"删除时间"`
}

// 节假日表
func (Holiday) TableName() string {
	return "holidays"
}

// 活动表
type Activity struct {
	ID          int64      `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	Name        string     `gorm:"column:name;size:100;not null" json:"name" comment:"活动名称"`
	Description string     `gorm:"column:description;size:500" json:"description" comment:"活动描述"`
	Image       string     `gorm:"column:image;size:500" json:"image" comment:"活动图片"`
	Discount    float64    `gorm:"column:discount;type:decimal(3,2);default:1" json:"discount" comment:"折扣率"`
	ValidFrom   time.Time  `gorm:"column:valid_from" json:"valid_from" comment:"开始时间"`
	ValidTo     time.Time  `gorm:"column:valid_to" json:"valid_to" comment:"结束时间"`
	Status      int        `gorm:"column:status;default:1" json:"status" comment:"状态"`
	SortOrder   int        `gorm:"column:sort_order;default:0" json:"sort_order" comment:"排序"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"-" comment:"删除时间"`
}

// 活动表
func (Activity) TableName() string {
	return "activities"
}

// 公告表
type Announcement struct {
	ID        int64      `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	Title     string     `gorm:"column:title;size:200;not null" json:"title" comment:"标题"`
	Content   string     `gorm:"column:content;size:2000" json:"content" comment:"内容"`
	Type      int        `gorm:"column:type;default:0" json:"type" comment:"类型"`
	Status    int        `gorm:"column:status;default:1" json:"status" comment:"状态"`
	SortOrder int        `gorm:"column:sort_order;default:0" json:"sort_order" comment:"排序"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-" comment:"删除时间"`
}

// 公告表
func (Announcement) TableName() string {
	return "announcements"
}

// 充值套餐表
type RechargePackage struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	Name        string    `gorm:"column:name;size:100;not null" json:"name" comment:"套餐名称"`
	Amount      float64   `gorm:"column:amount;type:decimal(10,2);not null" json:"amount" comment:"充值金额"`
	GiftAmount  float64   `gorm:"column:gift_amount;type:decimal(10,2);default:0" json:"gift_amount" comment:"赠送金额"`
	GiftPoints  int       `gorm:"column:gift_points;default:0" json:"gift_points" comment:"赠送积分"`
	Description string    `gorm:"column:description;size:500" json:"description" comment:"描述"`
	SortOrder   int       `gorm:"column:sort_order;default:0" json:"sort_order" comment:"排序"`
	Status      int       `gorm:"column:status;default:1" json:"status" comment:"状态"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
}

// 充值套餐表
func (RechargePackage) TableName() string {
	return "recharge_packages"
}

// 评价表
type Review struct {
	ID        int64      `gorm:"primary_key;auto_increment" json:"id" comment:"主键"`
	OrderID   int64      `gorm:"column:order_id;unique_index" json:"order_id" comment:"订单ID"`
	UserID    int64      `gorm:"column:user_id;index" json:"user_id" comment:"用户ID"`
	User      User       `gorm:"foreignkey:UserID" json:"user" comment:"用户"`
	RoomID    int64      `gorm:"column:room_id;index" json:"room_id" comment:"房间ID"`
	Rating    int        `gorm:"column:rating;default:5" json:"rating" comment:"评分(1-5)"`
	Content   string     `gorm:"column:content;size:1000" json:"content" comment:"评价内容"`
	Images    string     `gorm:"column:images;size:1000" json:"images" comment:"图片"`
	Status    int        `gorm:"column:status;default:1" json:"status" comment:"状态"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-" comment:"删除时间"`
}

// 评价表
func (Review) TableName() string {
	return "reviews"
}

// 自动迁移数据库表
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},            // 用户表
		&Admin{},           // 管理员表
		&AdminRole{},       // 管理员角色表
		&Permission{},      // 权限表
		&RolePermission{},  // 角色权限关联表
		&RoomType{},        // 房间类型表
		&Room{},            // 房间表
		&TimeSlot{},        // 时间槽表
		&Order{},           // 订单表
		&Payment{},         // 支付表
		&Membership{},      // 会员表
		&RechargeRecord{},  // 充值记录表
		&Notification{},    // 通知表
		&OperationLog{},    // 操作日志表
		&Holiday{},         // 节假日表
		&Activity{},        // 活动表
		&Announcement{},    // 公告表
		&RechargePackage{}, // 充值套餐表
		&Review{},          // 评价表
	).Error
}
