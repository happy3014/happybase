package log

import (
	"fmt"
	"github.com/happy3014/happybase/config"
	"github.com/happy3014/happybase/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"path/filepath"
	"strings"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLog(conf config.LogConfig) error {
	// 初始化目录
	err := utils.CreateDirIfNotExists(conf.LogDir)
	if err != nil {
		return err
	}
	logFilePath := filepath.Join(conf.LogDir, conf.LogFilename)
	writer := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    conf.MaxFileSize,
		MaxBackups: conf.MaxBackupNum,
		LocalTime:  false,
		Compress:   false,
	}
	zapConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,                               // 修改日志级别的编码方式，当前值: 全部大写
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000000"), // 修改时间编码格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	zapEncoder := zapcore.NewJSONEncoder(zapConfig)
	zapLogLevel := getZapLoggerLevel(conf.LogLevel)
	zapCore := zapcore.NewCore(zapEncoder, zapcore.AddSync(writer), zapLogLevel)
	logger = zap.New(zapCore, zap.AddCaller())
	sugarLogger = logger.Sugar()

	return nil
}

func DestroyLog() {
	if logger != nil {
		err := logger.Sync()
		if err != nil {
			fmt.Printf("failed to sync logger: %v\n", err)
		}
	}
	if sugarLogger != nil {
		err := sugarLogger.Sync()
		if err != nil {
			fmt.Printf("failed to sync sugar logger: %v\n", err)
		}
	}
}

func Logger() *zap.Logger {
	return logger
}

func SugarLogger() *zap.SugaredLogger {
	return sugarLogger
}

func getZapLoggerLevel(level string) zapcore.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	case "panic":
		return zap.PanicLevel
	default:
		return zap.InfoLevel
	}
}
