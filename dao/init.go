package dao

import (
	"Termbin/config"
	"Termbin/model"
	"context"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	// 数据库连接信息
	dsn := config.Conf.MySQL.User + ":" + config.Conf.MySQL.Password +
		"@tcp(" + config.Conf.MySQL.Host + ":" + config.Conf.MySQL.Port + ")/" +
		config.Conf.MySQL.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	_db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn, // data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
	}), &gorm.Config{
		// Logger: ormLogger, // 日志配置
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，不加s
		},
	})
	if err != nil {
		log.Println("gorm: failed to open MySQL.")
		panic(err)
	}

	// 设置连接池
	sqlDB, _ := _db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(200) // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	db = _db

	migration()

}

// NewDBClient 创建数据库连接对象
//
// 用于在业务逻辑中获取数据库连接对象。
func NewDBClient(ctx context.Context) *gorm.DB {
	_db := db
	return _db.WithContext(ctx)
}

// migration 用于迁移表
func migration() {
	// AutoMigrate 用于自动迁移 schema，保持 schema 是最新的。
	err := db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{}, &model.Clipboard{})
	if err != nil {
		log.Println("gorm: failed to migrate table.")
		panic(err)
	}
}
