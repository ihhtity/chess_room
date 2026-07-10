// 声明包名为 middleware
package middleware

// 导入必要的包
import (
	"net/http"

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

		isAllowed := false
		if origin == "" {
			isAllowed = true
		} else {
			for _, allowedOrigin := range config.Cfg.CORS.AllowedOrigins {
				if origin == allowedOrigin {
					isAllowed = true
					break
				}
			}
		}

		if isAllowed {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type, Authorization")
		}

		if c.Request.Method == "OPTIONS" {
			if isAllowed {
				c.AbortWithStatus(http.StatusNoContent)
			} else {
				c.AbortWithStatus(http.StatusForbidden)
			}
			return
		}

		c.Next()
	}
}
