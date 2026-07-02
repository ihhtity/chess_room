// 声明 main 包，程序入口包
package main

// 导入项目所需的依赖包
import (
	// 导入 MySQL 数据库操作包
	"chess-room-backend/dao/mysql"
	// 导入 Redis 缓存操作包
	"chess-room-backend/dao/redis"
	// 导入数据模型包
	"chess-room-backend/model"
	// 导入配置管理包
	"chess-room-backend/pkg/config"
	// 导入日志管理包
	"chess-room-backend/pkg/log"
	// 导入路由配置包
	"chess-room-backend/router"
	// 导入格式化 I/O 包
	"fmt"
	// 导入 HTTP 网络服务包
	"net/http"
	// 导入时间处理包
	"time"

	// 导入 Viper 配置管理库
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
		model.InitDefaultData(mysql.DB)
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

// 程序入口：加载配置、初始化组件、启动服务
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

	// 调用数据库初始化函数
	initDatabase()

	// 设置 HTTP 路由
	r := router.SetupRouter()

	// 创建 HTTP 服务器配置
	server := &http.Server{
		// 设置服务器监听地址和端口
		Addr: ":" + viper.GetString("server.port"),
		// 设置请求处理器
		Handler: r,
		// 设置读取超时时间为 10 秒
		ReadTimeout: 10 * time.Second,
		// 设置写入超时时间为 10 秒
		WriteTimeout: 10 * time.Second,
		// 设置请求头最大字节数为 1MB
		MaxHeaderBytes: 1 << 20,
	}

	// 启动 HTTP 服务
	log.Logger.Info("服务正在启动，端口：" + viper.GetString("server.port"))
	// 启动服务器并监听请求
	if err := server.ListenAndServe(); err != nil {
		// 启动失败则记录致命错误并退出
		log.Logger.Fatal("服务启动失败:", err)
	}
}
