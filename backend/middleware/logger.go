// 定义中间件包
package middleware

// 导入所需的依赖包
import (
	// 用于记录时间
	"time"

	// 项目内部的日志包
	"chess-room-backend/pkg/log"

	// Gin Web框架
	"github.com/gin-gonic/gin"
)

// 创建日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	// 返回Gin中间件处理函数
	return func(c *gin.Context) {
		// 记录请求开始时间
		startTime := time.Now()

		// 执行后续中间件和处理函数
		c.Next()

		// 计算请求处理耗时
		duration := time.Since(startTime)
		// 使用结构化日志记录请求信息
		log.Logger.Infof("[GIN] %s %s %d %v",
			// HTTP请求方法
			c.Request.Method,
			// 请求URL路径
			c.Request.URL.Path,
			// HTTP响应状态码
			c.Writer.Status(),
			// 请求处理耗时
			duration,
		)
	}
}
