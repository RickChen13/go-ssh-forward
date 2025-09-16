package flog

import (
	"go-ssh-forward/app/common/Const"
	"log"

	"github.com/natefinch/lumberjack"
)

func Debug(data string) {
	Log("", "log", data)
}

func Error(data string) {
	Log("", "error", data)
}

func Log(appendDir, name string, data string) {
	filename := Const.BASE_PATH + "/log" + appendDir
	filename += "/" + name + ".log"
	log.SetOutput(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,   // 每个日志文件的最大大小（MB）
		MaxBackups: 10,   // 保留旧文件的最大数量
		MaxAge:     28,   // 保留旧文件的最大天数
		Compress:   true, // 是否压缩旧文件
	})
	log.Println(data)
}
