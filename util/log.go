package util

import (
	"github.com/happy3014/happybase/config"
	"go.uber.org/zap"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLog(conf config.LogConfig) {

}

func Logger() *zap.Logger {
	return logger
}

func SugarLogger() *zap.SugaredLogger {
	return sugarLogger
}
