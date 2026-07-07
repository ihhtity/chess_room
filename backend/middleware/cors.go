// 声明包名为 middleware
package middleware

// 导入必要的包
import (
	// HTTP 状态码常量
	"net/http"
	// 字符串处理工具

	// 项目配置包

	// Gin Web 框架
	"github.com/gin-gonic/gin"
)

// CORS 中间件函数，返回 Gin 处理函数
func CORSMiddleware() gin.HandlerFunc {
	// 返回处理函数闭包
	return func(c *gin.Context) {
		// 获取请求头中的 Origin 字段
		origin := c.Request.Header.Get("Origin")

		// 设置允许的 HTTP 方法
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 设置允许的请求头
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept")
		// 设置暴露给客户端的响应头
		c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type, Authorization")

		// 如果 Origin 不为空，设置允许跨域的源
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		// 如果是预检请求（OPTIONS 方法）
		if c.Request.Method == "OPTIONS" {
			// 返回 204 No Content 状态码并终止请求
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// 继续执行后续中间件或处理器
		c.Next()
	}
}
