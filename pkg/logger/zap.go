package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var GormLogger *zap.Logger

func InitLogger() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"log/gin.log"}
	var err error
	Logger, err = config.Build()
	if err != nil {
		panic(err)
	}

	defer Logger.Sync()
}

func InitGormLogger() {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config.OutputPaths = []string{"log/gorm.log"}
	var err error
	GormLogger, err = config.Build()
	if err != nil {
		panic(err)
	}

	defer GormLogger.Sync()
}
