package ioc

import (
	"fmt"
	"gin_boot/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Dbname,
		cfg.Database.Charset,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold 慢 SQL 阈值
			LogLevel:                  logger.Info, // Log level 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略记录器的ErrRecordNotFound错误
			ParameterizedQueries:      false,       // 不要在SQL日志中包含参数
			Colorful:                  true,        // Disable color
		},
	)

	// Globally mode 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: cfg.Database.Prefix, // 数据库的表前缀
		},
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true, // 在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
	})
	if err != nil {
		log.Panic("mysql 连接失败", err)
	}

	return db
}
