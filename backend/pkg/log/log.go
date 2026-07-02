// 声明当前包名为 log
package log

// 导入所需的依赖包
import (
	// 操作系统功能包，用于文件操作
	"os"

	// 结构化日志库
	"github.com/sirupsen/logrus"
	// 配置管理库
	"github.com/spf13/viper"
)

// 定义全局 Logger 变量，供外部使用
var Logger *logrus.Logger

// 初始化日志配置
func Init() {
	// 创建新的 logrus 实例
	Logger = logrus.New()
	// 设置日志输出到标准输出
	Logger.SetOutput(os.Stdout)
	// 设置默认日志级别为 Info
	Logger.SetLevel(logrus.InfoLevel)

	// 从配置中读取日志级别
	level := viper.GetString("log.level")
	// 根据配置值切换日志级别
	switch level {
	// 调试级别
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
	// 警告级别
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
	// 错误级别
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
	// 致命级别
	case "fatal":
		Logger.SetLevel(logrus.FatalLevel)
	// 默认使用信息级别
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}

	// 从配置中读取日志格式
	format := viper.GetString("log.format")
	// 判断是否使用 JSON 格式
	if format == "json" {
		// 设置 JSON 格式化器
		Logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		// 设置文本格式化器，启用完整时间戳
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	// 从配置中读取日志输出路径
	output := viper.GetString("log.output")
	// 如果配置了输出路径
	if output != "" {
		// 打开或创建日志文件
		file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		// 如果文件打开成功
		if err == nil {
			// 将日志输出重定向到文件
			Logger.SetOutput(file)
		}
	}
}
