// 定义响应包，用于处理HTTP响应
package response

// 导入所需的包
import (
	// HTTP状态码常量
	"net/http"

	// 自定义错误码包
	"chess-room-backend/pkg/errno"

	// Gin Web框架
	"github.com/gin-gonic/gin"
)

// 定义统一的响应结构体
type Response struct {
	// 响应状态码
	Code int `json:"code"`
	// 响应消息
	Message string `json:"message"`
	// 响应数据，omitempty表示为空时不序列化
	Data interface{} `json:"data,omitempty"`
}

// 成功响应，仅返回数据
func Success(c *gin.Context, data interface{}) {
	// 返回200状态码的成功响应
	c.JSON(http.StatusOK, Response{
		// 状态码200
		Code: 200,
		// 消息"成功"
		Message: "成功",
		// 响应数据
		Data: data,
	})
}

// 成功响应，自定义消息和数据
func SuccessWithMsg(c *gin.Context, message string, data interface{}) {
	// 返回200状态码的成功响应
	c.JSON(http.StatusOK, Response{
		// 状态码200
		Code: 200,
		// 自定义消息
		Message: message,
		// 响应数据
		Data: data,
	})
}

// 失败响应，自定义状态码和消息
func Fail(c *gin.Context, code int, message string) {
	// 返回200 HTTP状态码，但包含错误业务码
	c.JSON(http.StatusOK, Response{
		// 业务错误码
		Code: code,
		// 错误消息
		Message: message,
	})
}

// 服务器错误响应
func Error(c *gin.Context, message string) {
	// 返回200 HTTP状态码，但包含500业务错误码
	c.JSON(http.StatusOK, Response{
		// 服务器内部错误码500
		Code: 500,
		// 错误消息
		Message: message,
	})
}

// 未授权响应
func Unauthorized(c *gin.Context) {
	// 返回401 HTTP状态码
	c.JSON(http.StatusUnauthorized, Response{
		// 未授权状态码401
		Code: 401,
		// 消息"未授权"
		Message: "未授权",
	})
}

// 统一错误处理
func HandleError(c *gin.Context, err error) {
	// 判断是否为自定义错误类型
	if e, ok := err.(*errno.Error); ok {
		// 是自定义错误，返回错误码和消息
		c.JSON(http.StatusOK, Response{
			// 转换错误码为int
			Code: int(e.Code),
			// 错误消息
			Message: e.Message,
		})
		// 结束处理
		return
	}
	// 非自定义错误，返回500错误码和错误信息
	c.JSON(http.StatusOK, Response{
		// 服务器内部错误码500
		Code: 500,
		// 错误信息
		Message: err.Error(),
	})
}
