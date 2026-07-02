package router

import (
	"chess-room-backend/controller"
	"chess-room-backend/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 设置路由函数，返回 Gin 引擎实例
func SetupRouter() *gin.Engine {
	// 创建默认的 Gin 引擎
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.LoggerMiddleware())

	r.GET("/health", controller.HealthCheck)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	// 创建用户路由组
	user := api.Group("/user")
	{
		// 用户登录接口
		user.POST("/login", controller.UserLogin)
		// 获取用户信息接口（需要认证）
		user.GET("/info", middleware.AuthMiddleware(), controller.GetUserProfile)
		// 更新用户信息接口（需要认证）
		user.PUT("/info", middleware.AuthMiddleware(), controller.UpdateUserProfile)
		// 修改密码接口（需要认证）
		user.POST("/change-password", middleware.AuthMiddleware(), controller.ChangePassword)
		// 获取用户列表接口（需要管理员权限）
		user.GET("/", middleware.AdminMiddleware(), controller.GetUserList)
		// 更新用户状态（需要管理员权限）
		user.PUT("/:id/status", middleware.AdminMiddleware(), controller.UpdateUserStatus)
	}

	// 创建房间列表路由组（公开）
	rooms := api.Group("/rooms")
	{
		// 获取房间列表
		rooms.GET("/", controller.GetRoomList)
		// 获取房间详情
		rooms.GET("/:id", controller.GetRoomDetail)
	}

	// 创建房间管理路由组（需要管理员权限）
	room := api.Group("/room")
	{
		// 获取房间列表
		room.GET("/", controller.GetRoomList)
		// 获取房间详情
		room.GET("/:id", controller.GetRoomDetail)
		// 创建房间
		room.POST("/", middleware.AdminMiddleware(), controller.CreateRoom)
		// 更新房间
		room.PUT("/:id", middleware.AdminMiddleware(), controller.UpdateRoom)
		// 删除房间
		room.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteRoom)
	}

	// 创建房间类型路由组
	roomType := api.Group("/room-type")
	{
		// 获取房间类型列表
		roomType.GET("/", controller.GetRoomTypeList)
		// 获取房间类型详情
		roomType.GET("/:id", controller.GetRoomTypeDetail)
		// 创建房间类型（需要管理员权限）
		roomType.POST("/", middleware.AdminMiddleware(), controller.CreateRoomType)
		// 更新房间类型（需要管理员权限）
		roomType.PUT("/:id", middleware.AdminMiddleware(), controller.UpdateRoomType)
		// 删除房间类型（需要管理员权限）
		roomType.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteRoomType)
	}

	// 创建订单路由组（用户端）
	orders := api.Group("/orders")
	{
		// 获取订单列表（需要认证）
		orders.GET("/", middleware.AuthMiddleware(), controller.GetOrderList)
		// 获取订单详情（需要认证）
		orders.GET("/:id", middleware.AuthMiddleware(), controller.GetOrderDetail)
		// 创建订单（需要认证）
		orders.POST("/", middleware.AuthMiddleware(), controller.CreateOrder)
		// 取消订单（需要认证）
		orders.PUT("/:id/cancel", middleware.AuthMiddleware(), controller.CancelOrder)
		// 确认订单（需要认证）
		orders.PUT("/:id/confirm", middleware.AuthMiddleware(), controller.ConfirmOrder)
		// 完成订单（需要认证）
		orders.PUT("/:id/complete", middleware.AuthMiddleware(), controller.CompleteOrder)
	}

	// 创建订单管理路由组（管理员端）
	order := api.Group("/order")
	{
		// 获取订单列表（需要管理员权限）
		order.GET("/", middleware.AdminMiddleware(), controller.GetOrderList)
		// 获取订单详情（需要管理员权限）
		order.GET("/:id", middleware.AdminMiddleware(), controller.GetOrderDetail)
		// 更新订单（需要管理员权限）
		order.PUT("/:id", middleware.AdminMiddleware(), controller.UpdateOrder)
		// 确认订单（需要管理员权限）
		order.PUT("/:id/confirm", middleware.AdminMiddleware(), controller.ConfirmOrder)
		// 完成订单（需要管理员权限）
		order.PUT("/:id/complete", middleware.AdminMiddleware(), controller.CompleteOrder)
		// 删除订单（需要管理员权限）
		order.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteOrder)
	}

	// 创建微信相关路由组
	wechat := api.Group("/wechat")
	{
		// 微信登录接口
		wechat.POST("/login", controller.WechatLogin)
	}

	// 创建会员路由组
	membership := api.Group("/membership")
	{
		// 获取会员信息（需要认证）
		membership.GET("/", middleware.AuthMiddleware(), controller.GetMembership)
		// 会员充值（需要认证）
		membership.POST("/recharge", middleware.AuthMiddleware(), controller.Recharge)
		// 获取会员列表（需要管理员权限）
		membership.GET("/list", middleware.AdminMiddleware(), controller.GetMembershipList)
		// 获取充值记录（需要认证）
		membership.GET("/recharges", middleware.AuthMiddleware(), controller.GetRechargeRecords)
	}

	// 创建会员管理路由组（管理员端）
	memberships := api.Group("/memberships")
	{
		// 获取会员列表（需要管理员权限）
		memberships.GET("/", middleware.AdminMiddleware(), controller.GetMembershipList)
		// 创建会员（需要管理员权限）
		memberships.POST("/", middleware.AdminMiddleware(), controller.CreateMembership)
		// 获取会员详情（需要管理员权限）
		memberships.GET("/:id", middleware.AdminMiddleware(), controller.GetMembershipDetail)
		// 更新会员信息（需要管理员权限）
		memberships.PUT("/:id", middleware.AdminMiddleware(), controller.UpdateMembership)
		// 删除会员（需要管理员权限）
		memberships.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteMembership)
	}

	// 创建充值路由组
	recharge := api.Group("/recharge")
	{
		// 发起充值（需要认证）
		recharge.POST("/", middleware.AuthMiddleware(), controller.Recharge)
	}

	// 创建管理员路由组
	admin := api.Group("/admin")
	{
		// 管理员登录
		admin.POST("/login", controller.AdminLogin)
		// 获取管理员信息（需要管理员权限）
		admin.GET("/profile", middleware.AdminMiddleware(), controller.GetAdminProfile)
		// 更新管理员信息（需要管理员权限）
		admin.PUT("/profile", middleware.AdminMiddleware(), controller.UpdateAdminProfile)
		// 修改管理员密码（需要管理员权限）
		admin.POST("/change-password", middleware.AdminMiddleware(), controller.AdminChangePassword)
	}

	// 创建活动路由组
	activities := api.Group("/activities")
	{
		// 获取活动列表（公开）
		activities.GET("/", controller.GetActivityList)
		// 获取活动详情（公开）
		activities.GET("/:id", controller.GetActivityDetail)
		// 创建活动（需要管理员权限）
		activities.POST("/", middleware.AdminMiddleware(), controller.CreateActivity)
		// 更新活动（需要管理员权限）
		activities.PUT("/:id", middleware.AdminMiddleware(), controller.UpdateActivity)
		// 删除活动（需要管理员权限）
		activities.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteActivity)
	}

	// 创建公告路由组
	announcements := api.Group("/announcements")
	{
		// 获取公告列表（公开）
		announcements.GET("/", controller.GetAnnouncementList)
		// 获取公告详情（公开）
		announcements.GET("/:id", controller.GetAnnouncementDetail)
		// 创建公告（需要管理员权限）
		announcements.POST("/", middleware.AdminMiddleware(), controller.CreateAnnouncement)
		// 更新公告（需要管理员权限）
		announcements.PUT("/:id", middleware.AdminMiddleware(), controller.UpdateAnnouncement)
		// 删除公告（需要管理员权限）
		announcements.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteAnnouncement)
	}

	// 创建充值套餐路由组
	rechargePackages := api.Group("/recharge-packages")
	{
		// 获取充值套餐列表（公开）
		rechargePackages.GET("/", controller.GetRechargePackageList)
		// 获取充值套餐详情（公开）
		rechargePackages.GET("/:id", controller.GetRechargePackageDetail)
		// 创建充值套餐（需要管理员权限）
		rechargePackages.POST("/", middleware.AdminMiddleware(), controller.CreateRechargePackage)
		// 更新充值套餐（需要管理员权限）
		rechargePackages.PUT("/:id", middleware.AdminMiddleware(), controller.UpdateRechargePackage)
		// 删除充值套餐（需要管理员权限）
		rechargePackages.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteRechargePackage)
	}

	// 创建时间槽路由组
	timeSlots := api.Group("/time-slots")
	{
		// 获取时间槽列表（公开）
		timeSlots.GET("/", controller.GetTimeSlotList)
		// 获取时间槽详情（公开）
		timeSlots.GET("/:id", controller.GetTimeSlotDetail)
		// 创建时间槽（需要管理员权限）
		timeSlots.POST("/", middleware.AdminMiddleware(), controller.CreateTimeSlot)
		// 更新时间槽（需要管理员权限）
		timeSlots.PUT("/:id", middleware.AdminMiddleware(), controller.UpdateTimeSlot)
		// 删除时间槽（需要管理员权限）
		timeSlots.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteTimeSlot)
	}

	// 创建评价路由组
	reviews := api.Group("/reviews")
	{
		// 获取评价列表（公开）
		reviews.GET("/", controller.GetReviewList)
		// 获取评价详情（公开）
		reviews.GET("/:id", controller.GetReviewDetail)
		// 创建评价（需要认证，管理员可指定用户）
		reviews.POST("/", middleware.AuthMiddleware(), controller.CreateReview)
		// 更新评价（需要管理员权限）
		reviews.PUT("/:id", middleware.AdminMiddleware(), controller.UpdateReview)
		// 删除评价（需要管理员权限）
		reviews.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteReview)
	}

	// 创建节假日路由组
	holidays := api.Group("/holidays")
	{
		// 获取节假日列表（公开）
		holidays.GET("/", controller.GetHolidayList)
		// 获取节假日详情（公开）
		holidays.GET("/:id", controller.GetHolidayDetail)
		// 创建节假日（需要管理员权限）
		holidays.POST("/", middleware.AdminMiddleware(), controller.CreateHoliday)
		// 更新节假日（需要管理员权限）
		holidays.PUT("/:id", middleware.AdminMiddleware(), controller.UpdateHoliday)
		// 删除节假日（需要管理员权限）
		holidays.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteHoliday)
	}

	// 创建支付路由组
	paymentsApi := api.Group("/payments")
	{
		// 获取支付列表（公开）
		paymentsApi.GET("/", controller.GetPaymentList)
		// 获取支付详情（公开）
		paymentsApi.GET("/:id", controller.GetPaymentDetail)
		// 更新支付记录（需要管理员权限）
		paymentsApi.PUT("/:id", middleware.AdminMiddleware(), controller.UpdatePayment)
		// 删除支付记录（需要管理员权限）
		paymentsApi.DELETE("/:id", middleware.AdminMiddleware(), controller.DeletePayment)
	}

	// 创建充值记录路由组
	rechargeRecords := api.Group("/recharge-records")
	{
		// 获取充值记录列表（公开）
		rechargeRecords.GET("/", controller.GetRechargeRecordList)
		// 获取充值记录详情（公开）
		rechargeRecords.GET("/:id", controller.GetRechargeRecordDetail)
		// 创建充值记录（需要管理员权限）
		rechargeRecords.POST("/", middleware.AdminMiddleware(), controller.CreateRechargeRecord)
		// 更新充值记录（需要管理员权限）
		rechargeRecords.PUT("/:id", middleware.AdminMiddleware(), controller.UpdateRechargeRecord)
		// 删除充值记录（需要管理员权限）
		rechargeRecords.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteRechargeRecord)
	}

	// 创建通知路由组
	notifications := api.Group("/notifications")
	{
		// 获取通知列表（公开）
		notifications.GET("/", controller.GetNotificationList)
		// 获取通知详情（公开）
		notifications.GET("/:id", controller.GetNotificationDetail)
		// 创建通知（需要管理员权限）
		notifications.POST("/", middleware.AdminMiddleware(), controller.CreateNotification)
		// 更新通知（需要管理员权限）
		notifications.PUT("/:id", middleware.AdminMiddleware(), controller.UpdateNotification)
		// 删除通知（需要管理员权限）
		notifications.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteNotification)
		// 标记全部已读（需要认证）
		notifications.POST("/mark-all-read", middleware.AuthMiddleware(), controller.MarkAllNotificationAsRead)
	}

	// 创建操作日志路由组
	operationLogs := api.Group("/operation-logs")
	{
		// 获取操作日志列表（公开）
		operationLogs.GET("/", controller.GetOperationLogList)
		// 获取操作日志详情（公开）
		operationLogs.GET("/:id", controller.GetOperationLogDetail)
		// 创建操作日志（需要管理员权限）
		operationLogs.POST("/", middleware.AdminMiddleware(), controller.CreateOperationLog)
		// 删除操作日志（需要管理员权限）
		operationLogs.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteOperationLog)
	}

	// 返回配置好的路由引擎
	return r
}
