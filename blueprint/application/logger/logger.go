package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type AppLogger struct {
	Logger *zap.Logger
	Level  zapcore.LevelEnabler
}

func InitLogger(level zapcore.LevelEnabler) *AppLogger {
	writerSyncer := GetLogWriter()

	encoder := GetEncoder()

	core := zapcore.NewCore(encoder, writerSyncer, level)

	appLogger := new(AppLogger)
	appLogger.Logger = zap.New(core, zap.AddCaller())

	defer appLogger.Logger.Sync()

	return appLogger
}

func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func GetLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./calendar_slave/log/logger.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}

	return zapcore.AddSync(lumberJackLogger)
}
