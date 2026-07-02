// @title 棋室管理系统 API
// @version 1.0
// @description 棋室管理系统后端 API 文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// 定义 main 包
package main

// 导入所需依赖包
import (
	// MySQL 数据库操作包
	"chess-room-backend/dao/mysql"
	// Redis 缓存操作包
	"chess-room-backend/dao/redis"
	// 引入 Swagger 文档（下划线表示仅执行 init 函数）
	_ "chess-room-backend/docs"
	// 数据模型包
	"chess-room-backend/model"
	// 配置管理包
	"chess-room-backend/pkg/config"
	// 日志管理包
	"chess-room-backend/pkg/log"
	// 路由配置包
	"chess-room-backend/router"
	// 上下文包，用于超时控制
	"context"
	// 格式化包
	"fmt"
	// HTTP 服务包
	"net/http"
	// 操作系统功能包
	"os"
	// 信号处理包
	"os/signal"
	// 系统调用包
	"syscall"
	// 时间处理包
	"time"

	// Viper 配置管理库
	"github.com/spf13/viper"
)

// 初始化数据库：检查表是否存在，不存在则自动迁移并初始化默认数据
func initDatabase() {
	// 定义变量存储表数量
	var tableCount int
	// 检查数据库表是否存在
	err := mysql.DB.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = ?", viper.GetString("database.database")).Count(&tableCount).Error
	// 判断查询是否出错
	if err != nil {
		// 记录错误日志
		log.Logger.Error("检查数据库表是否存在失败:", err)
		// 出错则直接返回
		return
	}

	// 判断表数量是否为 0
	if tableCount == 0 {
		// 数据库为空，执行自动迁移
		log.Logger.Info("未发现数据表，正在使用 AutoMigrate 初始化数据库...")
		// 执行数据库自动迁移
		if err := model.AutoMigrate(mysql.DB); err != nil {
			// 迁移失败则记录致命错误并退出
			log.Logger.Fatal("数据库自动迁移失败:", err)
		}
		// 初始化默认数据
		// model.InitDefaultData(mysql.DB)
		// 记录初始化成功日志
		log.Logger.Info("数据库初始化成功")
	} else {
		// 数据库已有表，更新表结构
		log.Logger.Info(fmt.Sprintf("发现 %d 张数据表，正在执行 AutoMigrate 更新表结构...", tableCount))
		// 执行数据库表结构更新
		if err := model.AutoMigrate(mysql.DB); err != nil {
			// 更新失败则记录致命错误并退出
			log.Logger.Fatal("更新数据库表结构失败:", err)
		}
		// 记录更新成功日志
		log.Logger.Info("数据库表结构更新成功")
	}
}

// 主函数：程序入口
func main() {
	// 初始化配置文件
	config.Init("./config/config.yaml")
	// 初始化日志系统
	log.Init()

	// 初始化 MySQL 连接
	mysql.Init()
	// 延迟关闭 MySQL 连接
	defer mysql.Close()

	// 初始化 Redis 连接
	redis.Init()
	// 延迟关闭 Redis 连接
	defer redis.Close()

	// 初始化数据库表结构
	initDatabase()

	// 设置 HTTP 路由
	r := router.SetupRouter()

	// 创建 HTTP 服务器实例
	server := &http.Server{
		// 服务器监听地址
		Addr: ":" + viper.GetString("server.port"),
		// 请求处理器
		Handler: r,
		// 读取超时时间
		ReadTimeout: 10 * time.Second,
		// 写入超时时间
		WriteTimeout: 10 * time.Second,
		// 最大请求头大小（1MB）
		MaxHeaderBytes: 1 << 20,
	}

	// 启动 HTTP 服务（协程中运行）
	go func() {
		// 记录服务启动日志
		log.Logger.Info("服务正在启动，端口：" + viper.GetString("server.port"))
		// 启动服务并监听端口
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// 服务启动失败则记录致命错误
			log.Logger.Fatal("服务启动失败:", err)
		}
	}()

	// 创建信号通道，缓冲大小为 1
	quit := make(chan os.Signal, 1)
	// 注册信号监听：中断信号和终止信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞等待信号
	<-quit

	// 记录服务关闭日志
	log.Logger.Info("服务正在关闭...")

	// 创建 5 秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 延迟取消上下文
	defer cancel()

	// 优雅关闭服务
	if err := server.Shutdown(ctx); err != nil {
		// 关闭失败则记录致命错误
		log.Logger.Fatal("服务关闭失败:", err)
	}

	// 记录服务关闭成功日志
	log.Logger.Info("服务已优雅关闭")
}
