package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func DbConnect() {
	dsn := "root:123456@tcp(localhost:3306)/pro?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 自定义日志输出器
			logger.Config{
				SlowThreshold:             200 * time.Millisecond, // 慢查询阈值
				LogLevel:                  logger.Info,            // 日志级别
				IgnoreRecordNotFoundError: true,                   // 忽略记录未找到错误
				Colorful:                  true,                   // 彩色日志
			},
		),
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	sqlDB, err := Db.DB()
	if err != nil {
		fmt.Println(err.Error())
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}
