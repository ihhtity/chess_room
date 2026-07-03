// 定义中间件包
package middleware

// 导入字符串处理包
import (
	// 字符串操作库
	"strings"

	// JWT工具包
	"chess-room-backend/pkg/jwt"
	// 响应工具包
	"chess-room-backend/pkg/response"

	// Gin Web框架
	"github.com/gin-gonic/gin"
)

// 认证中间件，验证用户Token
func AuthMiddleware() gin.HandlerFunc {
	// 返回Gin中间件处理函数
	return func(c *gin.Context) {
		// 从请求头获取Authorization字段
		token := c.GetHeader("Authorization")
		// 判断Token是否为空
		if token == "" {
			// 返回未授权响应
			response.Unauthorized(c)
			// 终止后续处理
			c.Abort()
			// 结束当前函数
			return
		}

		// 移除Token前缀"Bearer "
		token = strings.Replace(token, "Bearer ", "", 1)
		// 解析JWT Token获取声明信息
		claims, err := jwt.ParseToken(token)
		// 判断解析是否出错
		if err != nil {
			// 返回未授权响应
			response.Unauthorized(c)
			// 终止后续处理
			c.Abort()
			// 结束当前函数
			return
		}

		// 将用户ID存入上下文
		c.Set("user_id", claims.UserID)
		// 将用户名存入上下文
		c.Set("username", claims.Username)
		// 将用户角色存入上下文
		c.Set("role", claims.Role)
		// 继续执行后续中间件
		c.Next()
	}
}

// 管理员权限中间件，验证管理员Token
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)
		claims, err := jwt.ParseToken(token)
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		if claims.Role != 0 && claims.Role != 1 {
			response.Fail(c, 403, "没有权限")
			c.Abort()
			return
		}

		c.Set("admin_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("role_id", claims.RoleID)
		c.Next()
	}
}
