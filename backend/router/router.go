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
		// 获取用户详情（需要管理员权限）
		user.GET("/:id", middleware.AdminMiddleware(), controller.GetUserDetail)
		// 创建用户（需要管理员权限）
		user.POST("/", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.CreateUser)
		// 更新用户（需要管理员权限）
		user.PUT("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdateUser)
		// 更新用户状态（需要管理员权限）
		user.PUT("/:id/status", middleware.AdminMiddleware(), controller.UpdateUserStatus)
		// 批量更新用户（需要管理员权限）
		user.PUT("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchUpdateUser)
		// 删除用户（需要管理员权限）
		user.DELETE("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.DeleteUser)
		// 批量删除用户（需要管理员权限）
		user.DELETE("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchDeleteUser)
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
		room.POST("/", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.CreateRoom)
		// 更新房间
		room.PUT("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdateRoom)
		// 批量更新房间（需要管理员权限）
		room.PUT("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchUpdateRoom)
		// 删除房间（需要管理员权限）
		room.DELETE("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.DeleteRoom)
		// 批量删除房间（需要管理员权限）
		room.DELETE("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchDeleteRoom)
	}

	// 创建房间类型路由组
	roomType := api.Group("/room-type")
	{
		// 获取房间类型列表
		roomType.GET("/", controller.GetRoomTypeList)
		// 获取房间类型详情
		roomType.GET("/:id", controller.GetRoomTypeDetail)
		// 创建房间类型（需要管理员权限）
		roomType.POST("/", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.CreateRoomType)
		// 更新房间类型（需要管理员权限）
		roomType.PUT("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdateRoomType)
		// 批量更新房间类型（需要管理员权限）
		roomType.PUT("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchUpdateRoomType)
		// 删除房间类型（需要管理员权限）
		roomType.DELETE("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.DeleteRoomType)
		// 批量删除房间类型（需要管理员权限）
		roomType.DELETE("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchDeleteRoomType)
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
		order.PUT("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdateOrder)
		// 批量更新订单（需要管理员权限）
		order.PUT("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchUpdateOrder)
		// 确认订单（需要管理员权限）
		order.PUT("/:id/confirm", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.ConfirmOrder)
		// 完成订单（需要管理员权限）
		order.PUT("/:id/complete", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.CompleteOrder)
		// 删除订单（需要管理员权限）
		order.DELETE("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.DeleteOrder)
		// 批量删除订单（需要管理员权限）
		order.DELETE("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchDeleteOrder)
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
		memberships.POST("/", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.CreateMembership)
		// 获取会员详情（需要管理员权限）
		memberships.GET("/:id", middleware.AdminMiddleware(), controller.GetMembershipDetail)
		// 更新会员信息（需要管理员权限）
		memberships.PUT("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdateMembership)
		// 批量更新会员（需要管理员权限）
		memberships.PUT("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchUpdateMembership)
		// 删除会员（需要管理员权限）
		memberships.DELETE("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.DeleteMembership)
		// 批量删除会员（需要管理员权限）
		memberships.DELETE("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchDeleteMembership)
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
		admin.POST("/login", controller.AdminLogin)
		admin.GET("/profile", middleware.AdminMiddleware(), controller.GetAdminProfile)
		admin.PUT("/profile", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdateAdminProfile)
		admin.POST("/change-password", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.AdminChangePassword)

		roles := admin.Group("/roles", middleware.AdminMiddleware(), middleware.OperationLogMiddleware())
		{
			roles.GET("/", controller.GetAdminRoleList)
			roles.GET("/:id", controller.GetAdminRoleDetail)
			roles.GET("/available", controller.GetAvailableRoles)
			roles.POST("/", controller.CreateAdminRole)
			roles.PUT("/:id", controller.UpdateAdminRole)
			roles.PUT("/batch", controller.BatchUpdateAdminRole)
			roles.DELETE("/:id", controller.DeleteAdminRole)
		}

		permissions := admin.Group("/permissions", middleware.AdminMiddleware(), middleware.OperationLogMiddleware())
		{
			permissions.GET("/", controller.GetPermissionList)
			permissions.GET("/grouped", controller.GetPermissionListByGroup)
			permissions.GET("/:id", controller.GetPermissionDetail)
			permissions.GET("/role", controller.GetRolePermissions)
			permissions.GET("/mine", controller.GetMyPermissions)
			permissions.POST("/", controller.CreatePermission)
			permissions.PUT("/:id", controller.UpdatePermission)
			permissions.PUT("/batch", controller.BatchUpdatePermission)
			permissions.DELETE("/:id", controller.DeletePermission)
			permissions.POST("/role/:role_id", controller.AssignPermissions)
		}

		admins := admin.Group("/admins", middleware.AdminMiddleware(), middleware.OperationLogMiddleware())
		{
			admins.GET("/", controller.GetAdminList)
			admins.POST("/", controller.CreateAdmin)
			admins.PUT("/:id", controller.UpdateAdmin)
			admins.PUT("/batch", controller.BatchUpdateAdmin)
			admins.DELETE("/:id", controller.DeleteAdmin)
			admins.POST("/:id/reset-password", controller.ResetAdminPassword)
		}
	}

	// 创建活动路由组
	activities := api.Group("/activities")
	{
		// 获取活动列表（公开）
		activities.GET("/", controller.GetActivityList)
		// 获取活动详情（公开）
		activities.GET("/:id", controller.GetActivityDetail)
		// 创建活动（需要管理员权限）
		activities.POST("/", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.CreateActivity)
		// 更新活动（需要管理员权限）
		activities.PUT("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdateActivity)
		// 批量更新活动（需要管理员权限）
		activities.PUT("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchUpdateActivity)
		// 删除活动（需要管理员权限）
		activities.DELETE("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.DeleteActivity)
		// 批量删除活动（需要管理员权限）
		activities.DELETE("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchDeleteActivity)
	}

	// 创建公告路由组
	announcements := api.Group("/announcements")
	{
		// 获取公告列表（公开）
		announcements.GET("/", controller.GetAnnouncementList)
		// 获取公告详情（公开）
		announcements.GET("/:id", controller.GetAnnouncementDetail)
		// 创建公告（需要管理员权限）
		announcements.POST("/", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.CreateAnnouncement)
		// 更新公告（需要管理员权限）
		announcements.PUT("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdateAnnouncement)
		// 批量更新公告（需要管理员权限）
		announcements.PUT("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchUpdateAnnouncement)
		// 删除公告（需要管理员权限）
		announcements.DELETE("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.DeleteAnnouncement)
		// 批量删除公告（需要管理员权限）
		announcements.DELETE("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchDeleteAnnouncement)
	}

	// 创建充值套餐路由组
	rechargePackages := api.Group("/recharge-packages")
	{
		// 获取充值套餐列表（公开）
		rechargePackages.GET("/", controller.GetRechargePackageList)
		// 获取充值套餐详情（公开）
		rechargePackages.GET("/:id", controller.GetRechargePackageDetail)
		// 创建充值套餐（需要管理员权限）
		rechargePackages.POST("/", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.CreateRechargePackage)
		// 更新充值套餐（需要管理员权限）
		rechargePackages.PUT("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdateRechargePackage)
		// 批量更新充值套餐（需要管理员权限）
		rechargePackages.PUT("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchUpdateRechargePackage)
		// 删除充值套餐（需要管理员权限）
		rechargePackages.DELETE("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.DeleteRechargePackage)
		// 批量删除充值套餐（需要管理员权限）
		rechargePackages.DELETE("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchDeleteRechargePackage)
	}

	// 创建时间槽路由组
	timeSlots := api.Group("/time-slots")
	{
		// 获取时间槽列表（公开）
		timeSlots.GET("/", controller.GetTimeSlotList)
		// 获取时间槽详情（公开）
		timeSlots.GET("/:id", controller.GetTimeSlotDetail)
		// 创建时间槽（需要管理员权限）
		timeSlots.POST("/", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.CreateTimeSlot)
		// 更新时间槽（需要管理员权限）
		timeSlots.PUT("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdateTimeSlot)
		// 批量更新时间槽（需要管理员权限）
		timeSlots.PUT("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchUpdateTimeSlot)
		// 删除时间槽（需要管理员权限）
		timeSlots.DELETE("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.DeleteTimeSlot)
		// 批量删除时间槽（需要管理员权限）
		timeSlots.DELETE("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchDeleteTimeSlot)
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
		reviews.PUT("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdateReview)
		// 批量更新评价（需要管理员权限）
		reviews.PUT("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchUpdateReview)
		// 删除评价（需要管理员权限）
		reviews.DELETE("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.DeleteReview)
		// 批量删除评价（需要管理员权限）
		reviews.DELETE("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchDeleteReview)
	}

	// 创建节假日路由组
	holidays := api.Group("/holidays")
	{
		// 获取节假日列表（公开）
		holidays.GET("/", controller.GetHolidayList)
		// 获取节假日详情（公开）
		holidays.GET("/:id", controller.GetHolidayDetail)
		// 创建节假日（需要管理员权限）
		holidays.POST("/", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.CreateHoliday)
		// 更新节假日（需要管理员权限）
		holidays.PUT("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdateHoliday)
		// 批量更新节假日（需要管理员权限）
		holidays.PUT("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchUpdateHoliday)
		// 删除节假日（需要管理员权限）
		holidays.DELETE("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.DeleteHoliday)
		// 批量删除节假日（需要管理员权限）
		holidays.DELETE("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchDeleteHoliday)
	}

	// 创建支付路由组
	paymentsApi := api.Group("/payments")
	{
		// 获取支付列表（公开）
		paymentsApi.GET("/", controller.GetPaymentList)
		// 获取支付详情（公开）
		paymentsApi.GET("/:id", controller.GetPaymentDetail)
		// 更新支付记录（需要管理员权限）
		paymentsApi.PUT("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdatePayment)
		// 批量更新支付记录（需要管理员权限）
		paymentsApi.PUT("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchUpdatePayment)
		// 删除支付记录（需要管理员权限）
		paymentsApi.DELETE("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.DeletePayment)
		// 批量删除支付记录（需要管理员权限）
		paymentsApi.DELETE("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchDeletePayment)
	}

	// 创建充值记录路由组
	rechargeRecords := api.Group("/recharge-records")
	{
		// 获取充值记录列表（公开）
		rechargeRecords.GET("/", controller.GetRechargeRecordList)
		// 获取充值记录详情（公开）
		rechargeRecords.GET("/:id", controller.GetRechargeRecordDetail)
		// 创建充值记录（需要管理员权限）
		rechargeRecords.POST("/", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.CreateRechargeRecord)
		// 更新充值记录（需要管理员权限）
		rechargeRecords.PUT("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdateRechargeRecord)
		// 批量更新充值记录（需要管理员权限）
		rechargeRecords.PUT("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchUpdateRechargeRecord)
		// 删除充值记录（需要管理员权限）
		rechargeRecords.DELETE("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.DeleteRechargeRecord)
		// 批量删除充值记录（需要管理员权限）
		rechargeRecords.DELETE("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchDeleteRechargeRecord)
	}

	// 创建通知路由组
	notifications := api.Group("/notifications")
	{
		// 获取通知列表（公开）
		notifications.GET("/", controller.GetNotificationList)
		// 获取通知详情（公开）
		notifications.GET("/:id", controller.GetNotificationDetail)
		// 创建通知（需要管理员权限）
		notifications.POST("/", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.CreateNotification)
		// 更新通知（需要管理员权限）
		notifications.PUT("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.UpdateNotification)
		// 批量更新通知（需要管理员权限）
		notifications.PUT("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchUpdateNotification)
		// 删除通知（需要管理员权限）
		notifications.DELETE("/:id", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.DeleteNotification)
		// 批量删除通知（需要管理员权限）
		notifications.DELETE("/batch", middleware.AdminMiddleware(), middleware.OperationLogMiddleware(), controller.BatchDeleteNotification)
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
		// 批量删除操作日志（需要管理员权限）
		operationLogs.DELETE("/batch", middleware.AdminMiddleware(), controller.BatchDeleteOperationLog)
		// 批量更新操作日志（需要管理员权限）
		operationLogs.PUT("/batch", middleware.AdminMiddleware(), controller.BatchUpdateOperationLog)
	}

	// 创建定时任务路由组（管理员端）
	cronJobs := api.Group("/admin/cron-jobs", middleware.AdminMiddleware(), middleware.OperationLogMiddleware())
	{
		cronJobs.GET("/", controller.GetCronJobList)
		cronJobs.GET("/:id", controller.GetCronJobDetail)
		cronJobs.POST("/", controller.CreateCronJob)
		cronJobs.PUT("/:id", controller.UpdateCronJob)
		cronJobs.PUT("/batch", controller.BatchUpdateCronJob)
		cronJobs.DELETE("/:id", controller.DeleteCronJob)
		cronJobs.DELETE("/batch", controller.BatchDeleteCronJob)
	}

	// 创建导出路由组
	export := api.Group("/export", middleware.AdminMiddleware())
	{
		export.GET("/rooms", controller.ExportRoomList)
		export.GET("/activities", controller.ExportActivityList)
		export.GET("/orders", controller.ExportOrderList)
		export.GET("/admins", controller.ExportAdminList)
		export.GET("/members", controller.ExportMemberList)
		export.GET("/reviews", controller.ExportReviewList)
		export.GET("/announcements", controller.ExportAnnouncementList)
		export.GET("/recharge-packages", controller.ExportRechargePackageList)
		export.GET("/time-slots", controller.ExportTimeSlotList)
		export.GET("/holidays", controller.ExportHolidayList)
		export.GET("/users", controller.ExportUserList)
		export.GET("/notifications", controller.ExportNotificationList)
		export.GET("/payments", controller.ExportPaymentList)
		export.GET("/recharge-records", controller.ExportRechargeRecordList)
		export.GET("/operation-logs", controller.ExportOperationLogList)
		export.GET("/roles", controller.ExportRoleList)
		export.GET("/permissions", controller.ExportPermissionList)
		export.GET("/all", controller.ExportAllData)
	}

	// 返回配置好的路由引擎
	return r
}
