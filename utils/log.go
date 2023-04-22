package utils

import (
	"gingo/config"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"path"
)

var Log = logrus.New()

func InitLog() {
	logFileName := path.Join(config.Config.Log.Path, config.Config.Log.Name+".log")

	split(logFileName)

	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	Log.SetLevel(logrus.DebugLevel)
}

// 日志分割
func split(logFile string) {
	Log.SetOutput(&lumberjack.Logger{
		Filename:   logFile,                            // 日志文件位置
		LocalTime:  true,                               // 日志文件分割时使用本地时间
		MaxSize:    config.Config.Log.Split.MaxSize,    // 单文件最大容量,单位是MB
		MaxAge:     config.Config.Log.Split.MaxAge,     // 保留过期文件的最大时间间隔,单位是天
		MaxBackups: config.Config.Log.Split.MaxBackups, // 最大保留过期文件个数
		Compress:   config.Config.Log.Split.Compress,   // 是否需要压缩分割后的日志,使用的gzip压缩
	})
}
