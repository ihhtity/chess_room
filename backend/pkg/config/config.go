// 定义包名为 config
package config

// 导入 viper 配置库和日志库
import (
	"log"

	"github.com/spf13/viper"
)

// 定义总配置结构体，包含所有模块配置
type Config struct {
	Server   ServerConfig   // 服务器配置
	Database DatabaseConfig // 数据库配置
	Redis    RedisConfig    // Redis配置
	JWT      JWTConfig      // JWT配置
	Wechat   WechatConfig   // 微信配置
	Log      LogConfig      // 日志配置
	CORS     CORSConfig     // CORS配置
}

// 定义服务器配置结构体
type ServerConfig struct {
	Port   string // 服务器端口
	Mode   string // 运行模式
	Domain string // 服务域名
}

// 定义CORS配置结构体
type CORSConfig struct {
	AllowedOrigins []string // 允许的源列表
}

// 定义数据库配置结构体
type DatabaseConfig struct {
	Host            string // 数据库主机地址
	Port            int    // 数据库端口
	Username        string // 数据库用户名
	Password        string // 数据库密码
	Database        string // 数据库名称
	Charset         string // 字符集
	MaxIdleConns    int    // 最大空闲连接数
	MaxOpenConns    int    // 最大打开连接数
	ConnMaxLifetime int    // 连接最大生命周期
}

// 定义Redis配置结构体
type RedisConfig struct {
	Host     string // Redis主机地址
	Port     int    // Redis端口
	Password string // Redis密码
	DB       int    // 数据库编号
	PoolSize int    // 连接池大小
}

// 定义JWT配置结构体
type JWTConfig struct {
	Secret     string // JWT密钥
	ExpireTime int64  // 过期时间
}

// 定义微信配置结构体
type WechatConfig struct {
	AppID        string // 应用ID
	AppSecret    string // 应用密钥
	MchID        string // 商户ID
	APIKey       string // API密钥
	PayNotifyURL string // 支付回调URL
}

// 定义日志配置结构体
type LogConfig struct {
	Level  string // 日志级别
	Format string // 日志格式
	Output string // 日志输出
}

// 定义全局配置变量
var Cfg *Config

// 初始化配置函数，读取并解析配置文件
func Init(configPath string) {
	// 设置配置文件路径
	viper.SetConfigFile(configPath)
	// 设置配置文件类型为yaml
	viper.SetConfigType("yaml")
	// 自动读取环境变量
	viper.AutomaticEnv()

	// 读取配置文件，失败则退出程序
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 初始化全局配置对象
	Cfg = &Config{}
	// 将配置解析到结构体，失败则退出程序
	if err := viper.Unmarshal(Cfg); err != nil {
		log.Fatalf("解析配置失败: %v", err)
	}
}
