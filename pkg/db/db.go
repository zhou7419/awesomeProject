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

	logFile, err := os.Create("log/gorm.log")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(logFile, "\r\n", log.LstdFlags), // 自定义日志输出器
			logger.Config{
				SlowThreshold: 200 * time.Millisecond, // 慢查询阈值
				LogLevel:      logger.Info,            // 日志级别
			},
		),
		//Logger: logger.Default.LogMode(logger.Info),
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
