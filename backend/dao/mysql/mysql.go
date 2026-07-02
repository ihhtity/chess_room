// 定义mysql包
package mysql

// 导入fmt包用于格式化
import (
	// 导入fmt包用于格式化输出
	"fmt"
	// 导入time包用于时间相关操作
	"time"

	// 导入gorm包用于ORM操作
	"github.com/jinzhu/gorm"
	// 导入mysql驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// 导入viper包用于配置管理
	"github.com/spf13/viper"

	// 导入自定义log包
	"chess-room-backend/pkg/log"
)

// 声明全局数据库连接变量
var DB *gorm.DB

// 初始化数据库连接
func Init() {
	// 从配置文件中读取数据库主机地址
	host := viper.GetString("database.host")
	// 从配置文件中读取数据库端口
	port := viper.GetInt("database.port")
	// 从配置文件中读取数据库用户名
	username := viper.GetString("database.username")
	// 从配置文件中读取数据库密码
	password := viper.GetString("database.password")
	// 从配置文件中读取数据库名称
	database := viper.GetString("database.database")
	// 从配置文件中读取字符集
	charset := viper.GetString("database.charset")
	// 从配置文件中读取最大空闲连接数
	maxIdleConns := viper.GetInt("database.max_idle_conns")
	// 从配置文件中读取最大打开连接数
	maxOpenConns := viper.GetInt("database.max_open_conns")
	// 从配置文件中读取连接最大生命周期
	connMaxLifetime := viper.GetInt("database.conn_max_lifetime")

	// 检查并创建数据库（如果不存在）
	if err := createDatabaseIfNotExists(host, port, username, password, database); err != nil {
		// 创建数据库失败，记录致命错误并退出
		log.Logger.Fatal(fmt.Sprintf("创建数据库失败: %v", err))
	}

	// 构建数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, database, charset)

	// 声明错误变量
	var err error
	// 打开数据库连接
	DB, err = gorm.Open("mysql", dsn)
	// 检查连接是否成功
	if err != nil {
		// 连接失败，记录致命错误并退出
		log.Logger.Fatal(fmt.Sprintf("连接数据库失败: %v", err))
	}

	// 设置最大空闲连接数
	DB.DB().SetMaxIdleConns(maxIdleConns)
	// 设置最大打开连接数
	DB.DB().SetMaxOpenConns(maxOpenConns)
	// 设置连接最大生命周期
	DB.DB().SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	// 检查是否为调试模式
	if viper.GetString("server.mode") == "debug" {
		// 开启调试模式，打印SQL日志
		DB.LogMode(true)
	}

	// 记录数据库连接成功日志
	log.Logger.Info(fmt.Sprintf("数据库 %s 成功连接成功", database))
}

// 检查并创建数据库（如果不存在）
func createDatabaseIfNotExists(host string, port int, username string, password string, database string) error {
	// 构建连接MySQL服务器的DSN（不指定数据库）
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port)

	// 连接到MySQL服务器
	db, err := gorm.Open("mysql", dsn)
	// 检查连接是否成功
	if err != nil {
		// 返回连接错误
		return fmt.Errorf("failed to connect to MySQL server: %v", err)
	}
	// 延迟关闭临时连接
	defer db.Close()

	// 定义查询结果结构体
	var result struct {
		// 数据库数量
		Count int
	}
	// 查询指定名称的数据库是否存在
	err = db.Raw("SELECT COUNT(*) as count FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", database).Scan(&result).Error
	// 检查查询是否出错
	if err != nil {
		// 返回查询错误
		return fmt.Errorf("failed to check database existence: %v", err)
	}
	// 获取查询结果
	exists := result.Count

	// 判断数据库是否存在
	if exists == 0 {
		// 数据库不存在，记录日志
		log.Logger.Info(fmt.Sprintf("数据库 %s 不存在，正在创建...", database))
		// 执行创建数据库的SQL语句
		err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", database)).Error
		// 检查创建是否成功
		if err != nil {
			// 返回创建错误
			return fmt.Errorf("创建数据库失败: %v", err)
		}
		// 记录数据库创建成功日志
		log.Logger.Info(fmt.Sprintf("数据库 %s 创建成功", database))
	} else {
		// 数据库已存在，记录日志
		log.Logger.Info(fmt.Sprintf("数据库 %s 已存在", database))
	}

	// 返回nil表示成功
	return nil
}

// 关闭数据库连接
func Close() {
	// 检查数据库连接是否不为空
	if DB != nil {
		// 关闭数据库连接，忽略错误
		_ = DB.Close()
	}
}
