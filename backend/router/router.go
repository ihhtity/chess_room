// 定义包名为 router
package router

// 导入控制器包
import (
	// 导入控制器包
	"chess-room-backend/controller"
	// 导入中间件包
	"chess-room-backend/middleware"

	// 导入 Gin 框架
	"github.com/gin-gonic/gin"
)

// 设置路由函数，返回 Gin 引擎实例
func SetupRouter() *gin.Engine {
	// 创建默认的 Gin 引擎
	r := gin.Default()

	// 使用 CORS 中间件
	r.Use(middleware.CORSMiddleware())
	// 使用日志中间件
	r.Use(middleware.LoggerMiddleware())

	// 注册健康检查路由
	r.GET("/health", controller.HealthCheck)

	// 创建 API 路由组
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
	}

	// 创建订单管理路由组（管理员端）
	order := api.Group("/order")
	{
		// 获取订单列表（需要管理员权限）
		order.GET("/", middleware.AdminMiddleware(), controller.GetOrderList)
		// 获取订单详情（需要管理员权限）
		order.GET("/:id", middleware.AdminMiddleware(), controller.GetOrderDetail)
		// 确认订单（需要管理员权限）
		order.PUT("/:id/confirm", middleware.AdminMiddleware(), controller.ConfirmOrder)
		// 完成订单（需要管理员权限）
		order.PUT("/:id/complete", middleware.AdminMiddleware(), controller.CompleteOrder)
		// 删除订单（需要管理员权限）
		order.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteOrder)
	}

	// 创建支付路由组（复数形式）
	payments := api.Group("/payments")
	{
		// 创建支付（需要认证）
		payments.POST("/", middleware.AuthMiddleware(), controller.CreatePayment)
		// 根据订单号获取支付信息（需要认证）
		payments.GET("/:orderNo", middleware.AuthMiddleware(), controller.GetPaymentByOrderNo)
	}

	// 创建支付路由组（单数形式）
	payment := api.Group("/payment")
	{
		// 创建支付（需要认证）
		payment.POST("/", middleware.AuthMiddleware(), controller.CreatePayment)
		// 获取支付详情（需要认证）
		payment.GET("/:id", middleware.AuthMiddleware(), controller.GetPaymentDetail)
		// 获取支付列表（需要认证）
		payment.GET("/", middleware.AuthMiddleware(), controller.GetPaymentList)
		// 微信支付回调通知（公开）
		payment.POST("/wechat-notify", controller.WechatPayNotify)
		// 发起微信支付（需要认证）
		payment.POST("/wechat", middleware.AuthMiddleware(), controller.WechatPay)
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

	// 创建评价路由组
	reviews := api.Group("/reviews")
	{
		// 获取评价列表（公开）
		reviews.GET("/", controller.GetReviewList)
		// 获取评价详情（公开）
		reviews.GET("/:id", controller.GetReviewDetail)
		// 创建评价（需要认证）
		reviews.POST("/", middleware.AuthMiddleware(), controller.CreateReview)
		// 删除评价（需要管理员权限）
		reviews.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteReview)
	}

	// 返回配置好的路由引擎
	return r
}
