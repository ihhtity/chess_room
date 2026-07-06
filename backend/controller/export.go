package controller

// 导入业务逻辑包
import (
	// 导入房间相关业务逻辑
	"chess-room-backend/logic"
	// 导入数据模型
	"chess-room-backend/model"
	// 导入响应处理包
	"chess-room-backend/pkg/response"
	// 导入JSON编码包
	"encoding/json"
	// 导入HTTP协议包
	"net/http"
	// 导入字符串转换包
	"strconv"
	// 导入时间处理包
	"time"

	// 导入Gin框架
	"github.com/gin-gonic/gin"
)

// 导出房间列表
func ExportRoomList(c *gin.Context) {
	// 获取类型ID查询参数
	typeID := c.Query("type_id")
	// 获取楼层查询参数
	floor := c.Query("floor")
	// 获取状态查询参数
	status := c.Query("status")
	// 获取名称查询参数
	name := c.Query("name")

	// 初始化类型ID整数值
	typeIDInt := 0
	// 初始化状态整数值
	statusInt := 0
	// 声明错误变量
	var err error

	// 如果类型ID不为空
	if typeID != "" {
		// 将类型ID转换为整数
		typeIDInt, err = strconv.Atoi(typeID)
		// 如果转换失败
		if err != nil {
			// 返回参数错误响应
			response.Fail(c, 400, "类型ID格式错误")
			// 结束函数执行
			return
		}
	}

	// 如果状态不为空
	if status != "" {
		// 将状态转换为整数
		statusInt, err = strconv.Atoi(status)
		// 如果转换失败
		if err != nil {
			// 返回参数错误响应
			response.Fail(c, 400, "状态格式错误")
			// 结束函数执行
			return
		}
	}

	// 调用业务逻辑获取过滤后的房间列表
	rooms, _, err := logic.GetRoomListFiltered(typeIDInt, floor, statusInt, name, 0, 0)
	// 如果获取失败
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=rooms.json")
	// 返回JSON格式的房间数据
	c.JSON(http.StatusOK, rooms)
}

// 导出活动列表
func ExportActivityList(c *gin.Context) {
	// 调用业务逻辑获取活动列表
	activities, _, err := logic.GetActivityListAdminFiltered("", 0, 0, 0)
	// 如果获取失败
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=activities.json")
	// 返回JSON格式的活动数据
	c.JSON(http.StatusOK, activities)
}

// 导出订单列表
func ExportOrderList(c *gin.Context) {
	// 调用业务逻辑获取所有订单列表
	orders, _, err := logic.GetOrderList(0, 0, 0, "", time.Time{}, time.Time{}, 0, 0)
	// 如果获取失败
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=orders.json")
	// 返回JSON格式的订单数据
	c.JSON(http.StatusOK, orders)
}

// 导出管理员列表
func ExportAdminList(c *gin.Context) {
	// 调用业务逻辑获取所有管理员列表
	admins, _, err := logic.GetAdminList(0, "", "", 0, 0, 0, 0)
	// 如果获取失败
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=admins.json")
	// 返回JSON格式的管理员数据
	c.JSON(http.StatusOK, admins)
}

// 导出会员列表
func ExportMemberList(c *gin.Context) {
	// 调用业务逻辑获取所有会员列表
	members, _, err := logic.GetMembershipList(0, 0, 0, 0, 0)
	// 如果获取失败
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}
	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=members.json")
	// 返回JSON格式的会员数据
	c.JSON(http.StatusOK, members)
}

// 导出评论列表
func ExportReviewList(c *gin.Context) {
	// 调用业务逻辑获取所有评论列表
	reviews, _, err := logic.GetReviewList(0, 0, 0)
	// 如果获取失败
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=reviews.json")
	// 返回JSON格式的评论数据
	c.JSON(http.StatusOK, reviews)
}

// 导出公告列表
func ExportAnnouncementList(c *gin.Context) {
	// 调用业务逻辑获取所有公告列表
	announcements, _, err := logic.GetAnnouncementList(0, 0)
	// 如果获取失败
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=announcements.json")
	// 返回JSON格式的公告数据
	c.JSON(http.StatusOK, announcements)
}

// 导出储值套餐列表
func ExportRechargePackageList(c *gin.Context) {
	// 调用业务逻辑获取所有储值套餐列表
	packages, _, err := logic.GetRechargePackageList(0, 0)
	// 如果获取失败
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=recharge_packages.json")
	// 返回JSON格式的储值套餐数据
	c.JSON(http.StatusOK, packages)
}

// 导出时间槽列表
func ExportTimeSlotList(c *gin.Context) {
	// 调用业务逻辑获取所有时间槽列表
	timeSlots, _, err := logic.GetTimeSlotList("", 0, 0)
	// 如果获取失败
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=time_slots.json")
	// 返回JSON格式的时间槽数据
	c.JSON(http.StatusOK, timeSlots)
}

// 导出节假日列表
func ExportHolidayList(c *gin.Context) {
	holidays, _, err := logic.GetHolidayList("", 0, 0)
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=holidays.json")
	// 返回JSON格式的节假日数据
	c.JSON(http.StatusOK, holidays)
}

// 导出用户列表
func ExportUserList(c *gin.Context) {
	// 声明用户列表变量
	var users []model.User
	// 调用业务逻辑获取所有用户列表
	if err := logic.GetUserList(&users); err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=users.json")
	// 返回JSON格式的用户数据
	c.JSON(http.StatusOK, users)
}

// 导出通知列表
func ExportNotificationList(c *gin.Context) {
	// 调用业务逻辑获取所有通知列表
	notifications, _, err := logic.GetNotificationList("", "", "", 0, 0)
	// 如果获取失败
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=notifications.json")
	// 返回JSON格式的通知数据
	c.JSON(http.StatusOK, notifications)
}

// 导出支付记录列表
func ExportPaymentList(c *gin.Context) {
	// 调用业务逻辑获取所有支付记录列表
	payments, _, err := logic.GetPaymentList("", "", "", 0, 0)
	// 如果获取失败
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=payments.json")
	// 返回JSON格式的支付记录数据
	c.JSON(http.StatusOK, payments)
}

// 导出储值记录列表
func ExportRechargeRecordList(c *gin.Context) {
	// 调用业务逻辑获取所有储值记录列表
	records, _, err := logic.GetRechargeRecordList("", "", 0, 0)
	// 如果获取失败
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=recharge_records.json")
	// 返回JSON格式的储值记录数据
	c.JSON(http.StatusOK, records)
}

// 导出操作日志列表
func ExportOperationLogList(c *gin.Context) {
	logs, _, err := logic.GetOperationLogList("", "", 0, 0)
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=operation_logs.json")
	// 返回JSON格式的操作日志数据
	c.JSON(http.StatusOK, logs)
}

// 导出角色列表
func ExportRoleList(c *gin.Context) {
	// 获取名称查询参数
	name := c.Query("name")
	// 获取状态查询参数
	status := c.Query("status")

	// 初始化状态整数值
	statusInt := 0
	// 声明错误变量
	var err error

	// 如果状态不为空
	if status != "" {
		// 将状态转换为整数
		statusInt, err = strconv.Atoi(status)
		// 如果转换失败
		if err != nil {
			// 返回参数错误响应
			response.Fail(c, 400, "状态格式错误")
			// 结束函数执行
			return
		}
	}

	// 调用业务逻辑获取角色列表
	roles, _, err := logic.GetAdminRoleList(name, statusInt, 0, 0)
	// 如果获取失败
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=roles.json")
	// 返回JSON格式的角色数据
	c.JSON(http.StatusOK, roles)
}

// 导出权限列表
func ExportPermissionList(c *gin.Context) {
	// 获取名称查询参数
	name := c.Query("name")
	// 获取分组查询参数
	group := c.Query("group")

	// 调用业务逻辑获取权限列表
	permissions, _, err := logic.GetPermissionList(name, group, 0, 0)
	// 如果获取失败
	if err != nil {
		// 处理错误响应
		response.HandleError(c, err)
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=permissions.json")
	// 返回JSON格式的权限数据
	c.JSON(http.StatusOK, permissions)
}

// 导出所有数据
func ExportAllData(c *gin.Context) {
	// 获取类型ID查询参数
	typeID := c.Query("type_id")
	// 获取楼层查询参数
	floor := c.Query("floor")
	// 获取状态查询参数
	status := c.Query("status")
	// 获取名称查询参数
	name := c.Query("name")

	// 初始化类型ID整数值
	typeIDInt := 0
	// 初始化状态整数值
	statusInt := 0
	// 声明错误变量
	var err error

	// 如果类型ID不为空
	if typeID != "" {
		// 将类型ID转换为整数
		typeIDInt, err = strconv.Atoi(typeID)
		// 如果转换失败
		if err != nil {
			// 返回参数错误响应
			response.Fail(c, 400, "类型ID格式错误")
			// 结束函数执行
			return
		}
	}

	// 如果状态不为空
	if status != "" {
		// 将状态转换为整数
		statusInt, err = strconv.Atoi(status)
		// 如果转换失败
		if err != nil {
			// 返回参数错误响应
			response.Fail(c, 400, "状态格式错误")
			// 结束函数执行
			return
		}
	}

	// 调用业务逻辑获取房间列表（过滤）
	rooms, _, _ := logic.GetRoomListFiltered(typeIDInt, floor, statusInt, name, 0, 0)
	// 调用业务逻辑获取活动列表
	activities, _, _ := logic.GetActivityListAdminFiltered("", 0, 0, 0)
	// 调用业务逻辑获取订单列表
	orders, _, _ := logic.GetOrderList(0, 0, 0, "", time.Time{}, time.Time{}, 0, 0)
	// 调用业务逻辑获取会员列表
	members, _, _ := logic.GetMembershipList(0, 0, 0, 0, 0)
	// 调用业务逻辑获取评论列表
	reviews, _, _ := logic.GetReviewList(0, 0, 0)
	// 调用业务逻辑获取用户列表
	var users []model.User
	// 调用业务逻辑获取用户列表
	logic.GetUserList(&users)

	// 组合所有数据
	allData := map[string]interface{}{
		"rooms":      rooms,
		"activities": activities,
		"orders":     orders,
		"members":    members,
		"reviews":    reviews,
		"users":      users,
	}

	// 转换为JSON字符串
	data, err := json.MarshalIndent(allData, "", "  ")
	// 如果转换失败
	if err != nil {
		// 处理错误响应
		response.Fail(c, 500, err.Error())
		// 结束函数执行
		return
	}

	// 设置响应内容类型为JSON
	c.Header("Content-Type", "application/json")
	// 设置响应头，指示浏览器下载文件
	c.Header("Content-Disposition", "attachment; filename=all_data.json")
	// 返回JSON格式的所有数据
	c.String(http.StatusOK, string(data))
}
