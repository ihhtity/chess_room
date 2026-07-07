// 包声明：中间件包
package middleware

// 导入必要的依赖包
import (
	// 业务逻辑层
	"chess-room-backend/logic"
	// 数据模型层
	"chess-room-backend/model"
	// JSON编码
	"encoding/json"
	// IO工具（已弃用，建议使用io）
	"io/ioutil"
	// 字符串处理
	"strings"

	// Gin Web框架
	"github.com/gin-gonic/gin"
)

// 操作日志中间件构造函数
func OperationLogMiddleware() gin.HandlerFunc {
	// 返回Gin处理函数
	return func(c *gin.Context) {
		// 从上下文中获取管理员ID
		adminID, exists := c.Get("admin_id")
		// 如果不存在则跳过
		if !exists {
			c.Next()
			return
		}

		// 类型断言转换为int64
		adminIDInt, ok := adminID.(int64)
		// 类型转换失败则跳过
		if !ok {
			c.Next()
			return
		}

		// 获取请求方法
		method := c.Request.Method
		// 获取请求路径
		path := c.Request.URL.Path
		// 获取客户端IP
		ip := c.ClientIP()

		// 定义请求体数据变量
		var bodyData []byte
		// 只对POST和PUT请求读取body
		if method == "POST" || method == "PUT" {
			// 读取请求体
			body, err := ioutil.ReadAll(c.Request.Body)
			// 读取成功
			if err == nil {
				// 保存body数据
				bodyData = body
				// 关闭原始body
				c.Request.Body.Close()
				// 重新设置body以便后续读取
				c.Request.Body = ioutil.NopCloser(strings.NewReader(string(body)))
			}
		}

		// 执行后续中间件和处理函数
		c.Next()

		// 获取响应状态码
		statusCode := c.Writer.Status()
		// 错误状态码不记录日志
		if statusCode >= 400 {
			return
		}

		// 提取模块名称
		module := extractModule(path)
		// 提取操作类型
		action := extractAction(method, path)
		// 提取目标ID
		targetID := extractTargetID(path)

		// 定义内容变量
		var content string
		// 如果有body数据
		if len(bodyData) > 0 {
			// 转换为字符串
			content = string(bodyData)
			// 截断过长内容
			if len(content) > 500 {
				content = content[:500] + "..."
			}
		}

		// 模块和操作都有效才记录
		if module != "" && action != "" {
			// 创建日志对象
			log := &model.OperationLog{
				AdminID:  adminIDInt,
				Action:   action,
				Module:   module,
				TargetID: targetID,
				Content:  content,
				IP:       ip,
			}
			// 异步创建操作日志
			go logic.CreateOperationLog(log)
		}
	}
}

// 从路径提取模块名称
func extractModule(path string) string {
	// 按斜杠分割路径
	pathParts := strings.Split(path, "/")
	// 检查路径深度
	if len(pathParts) >= 3 {
		// 获取模块名
		module := pathParts[2]
		// 特殊处理admin模块
		if module == "admin" && len(pathParts) >= 4 {
			// 根据子路径判断
			switch pathParts[3] {
			case "login":
				return "admin"
			case "profile":
				return "admin"
			case "reset-password":
				return "admin"
			case "change-password":
				return "admin"
			default:
				return module
			}
		}
		return module
	}
	return ""
}

// 从方法和路径提取操作类型
func extractAction(method, path string) string {
	// 根据HTTP方法判断
	switch method {
	case "POST":
		// 登录操作
		if strings.HasSuffix(path, "/login") {
			return "login"
		}
		// 充值操作
		if strings.Contains(path, "/recharge") {
			return "recharge"
		}
		// 分配操作
		if strings.Contains(path, "/assign") {
			return "assign"
		}
		// 默认创建
		return "create"
	case "PUT":
		// 重置密码
		if strings.Contains(path, "/reset-password") {
			return "reset_password"
		}
		// 修改密码
		if strings.Contains(path, "/change-password") {
			return "change_password"
		}
		// 更新状态
		if strings.Contains(path, "/status") {
			return "update_status"
		}
		// 确认操作
		if strings.Contains(path, "/confirm") {
			return "confirm"
		}
		// 完成操作
		if strings.Contains(path, "/complete") {
			return "complete"
		}
		// 取消操作
		if strings.Contains(path, "/cancel") {
			return "cancel"
		}
		// 分配操作
		if strings.Contains(path, "/assign") {
			return "assign"
		}
		// 默认更新
		return "update"
	case "DELETE":
		// 删除操作
		return "delete"
	case "GET":
		// 查看资料
		if strings.Contains(path, "/profile") {
			return "view_profile"
		}
		// 查看权限
		if strings.Contains(path, "/mine") {
			return "view_permissions"
		}
		// GET请求默认不记录
		return ""
	default:
		// 其他方法不记录
		return ""
	}
}

// 从路径提取目标ID
func extractTargetID(path string) int64 {
	// 按斜杠分割路径
	pathParts := strings.Split(path, "/")
	// 遍历路径各部分
	for i, part := range pathParts {
		// 跳过admin相关路径，查找数字ID
		if i > 0 && pathParts[i-1] != "admin" && pathParts[i-1] != "admins" {
			// 尝试解析为数字
			if id, err := json.Number(part).Int64(); err == nil {
				return id
			}
		}
	}
	// 未找到返回0
	return 0
}
