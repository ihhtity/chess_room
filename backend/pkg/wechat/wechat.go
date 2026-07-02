// 定义包名为 wechat
package wechat

// 导入所需的依赖包
import (
	// 导入项目配置包
	"chess-room-backend/pkg/config"
	// 导入 JSON 编码解码包
	"encoding/json"
	// 导入格式化包
	"fmt"
	// 导入 IO 工具包
	"io"
	// 导入 HTTP 客户端包
	"net/http"
)

// 定义微信会话响应结构体
type SessionResponse struct {
	// 用户唯一标识
	OpenID string `json:"openid"`
	// 会话密钥
	SessionKey string `json:"session_key"`
	// 用户在开放平台的唯一标识符
	UnionID string `json:"unionid"`
	// 错误码
	ErrCode int `json:"errcode"`
	// 错误信息
	ErrMsg string `json:"errmsg"`
}

// 获取微信用户会话信息
func GetSession(code string) (*SessionResponse, error) {
	// 构建请求 URL，包含小程序 appid、密钥和授权码
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		config.Cfg.Wechat.AppID, config.Cfg.Wechat.AppSecret, code,
	)

	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	// 如果请求失败，返回错误
	if err != nil {
		return nil, err
	}
	// 延迟关闭响应体
	defer resp.Body.Close()

	// 读取响应体内容
	body, err := io.ReadAll(resp.Body)
	// 如果读取失败，返回错误
	if err != nil {
		return nil, err
	}

	// 定义结果变量
	var result SessionResponse
	// 解析 JSON 响应
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	// 检查微信返回的错误码
	if result.ErrCode != 0 {
		// 返回微信错误信息
		return nil, fmt.Errorf("微信错误: %d - %s", result.ErrCode, result.ErrMsg)
	}

	// 返回解析后的结果
	return &result, nil
}

// 定义微信用户信息响应结构体
type UserInfoResponse struct {
	// 用户唯一标识
	OpenID string `json:"openid"`
	// 用户昵称
	Nickname string `json:"nickname"`
	// 用户性别
	Sex int `json:"sex"`
	// 用户所在省份
	Province string `json:"province"`
	// 用户所在城市
	City string `json:"city"`
	// 用户所在国家
	Country string `json:"country"`
	// 用户头像 URL
	AvatarURL string `json:"headimgurl"`
	// 用户在开放平台的唯一标识符
	UnionID string `json:"unionid"`
	// 错误码
	ErrCode int `json:"errcode"`
	// 错误信息
	ErrMsg string `json:"errmsg"`
}

// 获取微信用户信息
func GetUserInfo(accessToken, openID string) (*UserInfoResponse, error) {
	// 构建请求 URL，包含访问令牌和用户 openid
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s",
		accessToken, openID,
	)

	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	// 如果请求失败，返回错误
	if err != nil {
		return nil, err
	}
	// 延迟关闭响应体
	defer resp.Body.Close()

	// 读取响应体内容
	body, err := io.ReadAll(resp.Body)
	// 如果读取失败，返回错误
	if err != nil {
		return nil, err
	}

	// 定义结果变量
	var result UserInfoResponse
	// 解析 JSON 响应
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	// 检查微信返回的错误码
	if result.ErrCode != 0 {
		// 返回微信错误信息
		return nil, fmt.Errorf("微信错误: %d - %s", result.ErrCode, result.ErrMsg)
	}

	// 返回解析后的结果
	return &result, nil
}

// 生成支付签名（模拟实现）
func GeneratePaySign(params map[string]string) string {
	// 返回模拟签名
	return "模拟签名"
}

// 创建微信支付订单（模拟实现）
func CreateOrder(tradeNO, body, totalFee, openID string) (map[string]interface{}, error) {
	// 返回模拟的预支付订单信息
	return map[string]interface{}{
		"prepay_id": "模拟预支付ID_" + tradeNO,
	}, nil
}
