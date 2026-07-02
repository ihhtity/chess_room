// 声明 redis 包
package redis

// 导入 context 包用于上下文管理
import (
	// 导入 context 包
	"context"
	"log"

	// 导入 fmt 包用于格式化输出
	"fmt"
	// 导入 time 包用于时间操作
	"time"

	// 导入 go-redis 客户端库
	"github.com/go-redis/redis/v8"
	// 导入 viper 配置库
	"github.com/spf13/viper"
)

// 声明全局 Redis 客户端变量
var RDB *redis.Client

// 声明全局上下文变量
var ctx = context.Background()

// 初始化 Redis 连接
func Init() {
	// 从配置中获取 Redis 主机地址
	host := viper.GetString("redis.host")
	// 从配置中获取 Redis 端口
	port := viper.GetInt("redis.port")
	// 从配置中获取 Redis 密码
	password := viper.GetString("redis.password")
	// 从配置中获取 Redis 数据库编号
	db := viper.GetInt("redis.db")

	// 创建新的 Redis 客户端
	RDB = redis.NewClient(&redis.Options{
		// 设置服务器地址
		Addr: fmt.Sprintf("%s:%d", host, port),
		// 设置连接密码
		Password: password,
		// 设置数据库编号
		DB: db,
	})

	// 测试连接是否成功
	_, err := RDB.Ping(ctx).Result()
	// 如果连接失败则抛出异常
	if err != nil {
		panic(fmt.Sprintf("Redis 连接失败: %v", err))
	}

	// 如果连接成功，记录日志
	log.Println("Redis 连接成功")
}

// 关闭 Redis 连接
func Close() {
	// 检查客户端是否已初始化
	if RDB != nil {
		// 关闭连接并忽略错误
		_ = RDB.Close()
	}
}

// 设置键值对，带过期时间
func Set(key string, value interface{}, expiration int) error {
	// 执行 SET 命令，将过期时间转换为秒
	return RDB.Set(ctx, key, value, time.Duration(expiration)*time.Second).Err()
}

// 获取指定键的值
func Get(key string) (string, error) {
	// 执行 GET 命令
	return RDB.Get(ctx, key).Result()
}

// 删除指定键
func Del(key string) error {
	// 执行 DEL 命令
	return RDB.Del(ctx, key).Err()
}

// 检查键是否存在
func Exists(key string) (bool, error) {
	// 执行 EXISTS 命令获取键数量
	exists, err := RDB.Exists(ctx, key).Result()
	// 返回是否存在（数量大于0）
	return exists > 0, err
}

// 将键的值加1
func Incr(key string) (int64, error) {
	// 执行 INCR 命令
	return RDB.Incr(ctx, key).Result()
}

// 将键的值减1
func Decr(key string) (int64, error) {
	// 执行 DECR 命令
	return RDB.Decr(ctx, key).Result()
}

// 设置哈希表字段
func HSet(key string, fields ...interface{}) error {
	// 执行 HSET 命令
	return RDB.HSet(ctx, key, fields...).Err()
}

// 获取哈希表指定字段的值
func HGet(key string, field string) (string, error) {
	// 执行 HGET 命令
	return RDB.HGet(ctx, key, field).Result()
}

// 获取哈希表所有字段和值
func HGetAll(key string) (map[string]string, error) {
	// 执行 HGETALL 命令
	return RDB.HGetAll(ctx, key).Result()
}
