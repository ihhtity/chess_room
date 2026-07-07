// 声明包名为 middleware
package middleware

// 导入必要的包
import (
	// HTTP 状态码常量
	"net/http"
	// 字符串处理工具

	// 项目配置包
	"chess-room-backend/pkg/config"
	// Gin Web 框架
	"github.com/gin-gonic/gin"
)

// CORS 中间件函数，返回 Gin 处理函数
func CORSMiddleware() gin.HandlerFunc {
	// 返回处理函数闭包
	return func(c *gin.Context) {
		// 获取请求头中的 Origin 字段
		origin := c.Request.Header.Get("Origin")
		// 初始化允许标志为 false
		isAllowed := false

		// 设置允许的 HTTP 方法
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 设置允许的请求头
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept")
		// 设置暴露给客户端的响应头
		c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type, Authorization")

		// 如果 Origin 为空，允许请求（同源请求或通过代理访问）
		if origin == "" {
			isAllowed = true
		} else {
			// 遍历配置中允许的源列表
			for _, allowedOrigin := range config.Cfg.CORS.AllowedOrigins {
				// 检查 Origin 是否与允许的源完全匹配
				if origin == allowedOrigin {
					isAllowed = true
					break
				}
			}
		}

		// 如果是预检请求（OPTIONS 方法）
		if c.Request.Method == "OPTIONS" {
			// 对于预检请求，如果源被允许则设置 CORS 头
			if isAllowed && origin != "" {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Access-Control-Allow-Credentials", "true")
			}
			// 返回 204 No Content 状态码并终止请求
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// 如果 Origin 被允许，设置 CORS 头
		if isAllowed && origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		// 继续执行后续中间件或处理器
		c.Next()
	}
}
