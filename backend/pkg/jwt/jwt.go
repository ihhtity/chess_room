// 定义 JWT 包
package jwt

// 导入必要的包
import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// 定义自定义声明结构体
type Claims struct {
	UserID               int64  `json:"user_id"`  // 用户 ID
	Username             string `json:"username"` // 用户名
	Role                 int    `json:"role"`     // 用户角色
	RoleID               int64  `json:"role_id"`  // 角色ID
	jwt.RegisteredClaims        // 内嵌标准声明
}

// 生成 JWT 令牌
func GenerateToken(userID int64, username string, role int, roleID int64) (string, error) {
	secret := viper.GetString("jwt.secret")
	expireTime := viper.GetInt64("jwt.expire_time")

	claims := Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RoleID:   roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expireTime))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// 解析 JWT 令牌
func ParseToken(tokenString string) (*Claims, error) {
	// 从配置中获取密钥
	secret := viper.GetString("jwt.secret")

	// 解析令牌字符串
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 返回密钥用于验证签名
		return []byte(secret), nil
	})

	// 解析出错时返回错误
	if err != nil {
		return nil, err
	}

	// 验证声明并检查令牌是否有效
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	// 令牌无效时返回错误
	return nil, err
}
