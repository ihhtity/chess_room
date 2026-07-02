// 定义工具包
package utils

// 导入加密随机数生成器
import (
	// 加密安全的随机数生成
	"crypto/rand"
	// SHA256 哈希算法
	"crypto/sha256"
	// Base64 编码
	"encoding/base64"
	// 格式化输入输出
	"fmt"
	// 大整数运算
	"math/big"
	// 时间处理
	"time"

	// bcrypt 密码哈希
	"golang.org/x/crypto/bcrypt"
)

// 生成订单号
func GenerateOrderNo() string {
	// 获取当前时间戳，格式为年月日时分秒
	timestamp := time.Now().Format("20060102150405")
	// 生成6位随机字符串
	random := generateRandomString(6)
	// 返回格式化的订单号，前缀为ORD
	return fmt.Sprintf("ORD%s%s", timestamp, random)
}

// 生成指定长度的随机字符串
func generateRandomString(length int) string {
	// 定义字符集：数字和大写字母
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 创建字节切片存储结果
	b := make([]byte, length)
	// 遍历每个位置生成随机字符
	for i := range b {
		// 生成随机索引
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		// 根据随机索引选取字符
		b[i] = charset[n.Int64()]
	}
	// 返回生成的随机字符串
	return string(b)
}

// 对密码进行哈希加密
func HashPassword(password string) (string, error) {
	// 使用 bcrypt 生成密码哈希
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 返回哈希字符串和可能的错误
	return string(bytes), err
}

// 验证密码与哈希是否匹配
func CheckPasswordHash(password, hash string) bool {
	// 比较密码和哈希值
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// 返回验证结果，无错误则匹配成功
	return err == nil
}

// 计算字符串的 SHA256 哈希值并返回 Base64 编码
func SHA256(data string) string {
	// 创建 SHA256 哈希器
	h := sha256.New()
	// 写入数据
	h.Write([]byte(data))
	// 返回 Base64 编码的哈希结果
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// 获取今天开始时间（00:00:00）
func GetTodayStart() time.Time {
	// 获取当前时间
	now := time.Now()
	// 返回今天零点时间
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// 获取今天结束时间（23:59:59.999999999）
func GetTodayEnd() time.Time {
	// 获取当前时间
	now := time.Now()
	// 返回今天最后一纳秒的时间
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())
}

// 计算两个时间点之间的分钟数
func CalculateDurationMinutes(start, end time.Time) int {
	// 计算时间差并转换为分钟
	return int(end.Sub(start).Minutes())
}
