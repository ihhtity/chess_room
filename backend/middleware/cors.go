// 定义中间件包
package middleware

// 导入必要的标准库和第三方库
import (
	// HTTP状态码常量
	"net/http"
	// 字符串处理工具
	"strings"

	// Gin Web框架
	"github.com/gin-gonic/gin"
)

// 创建CORS中间件函数，返回Gin处理函数
func CORSMiddleware() gin.HandlerFunc {
	// 返回匿名处理函数
	return func(c *gin.Context) {
		// 获取请求头中的Origin字段
		origin := c.Request.Header.Get("Origin")

		// 初始化允许标志为false
		isAllowed := false
		// 判断Origin是否为空
		if origin == "" {
			// 空Origin允许访问
			isAllowed = true
		} else {
			// 检查是否为本地开发环境地址
			isAllowed = strings.HasPrefix(origin, "http://localhost:") ||
				// 检查是否为127.0.0.1本地地址
				strings.HasPrefix(origin, "http://127.0.0.1:") ||
				// 检查是否为0.0.0.0地址
				strings.HasPrefix(origin, "http://0.0.0.0:")
		}

		// 如果允许访问，设置CORS响应头
		if isAllowed {
			// 设置允许的Origin
			c.Header("Access-Control-Allow-Origin", origin)
			// 允许携带凭证（如Cookie）
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		// 设置允许的HTTP方法
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 设置允许的请求头
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept")
		// 设置允许暴露的响应头
		c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type, Authorization")

		// 处理预检请求（OPTIONS方法）
		if c.Request.Method == "OPTIONS" {
			// 返回204无内容状态码并终止请求
			c.AbortWithStatus(http.StatusNoContent)
			// 直接返回，不执行后续中间件
			return
		}

		// 继续执行后续中间件或处理函数
		c.Next()
	}
}
