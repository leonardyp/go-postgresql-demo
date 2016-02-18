package logger

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var logDebugFile *logs.BeeLogger
var logErrorFile *logs.BeeLogger
var logStd *logs.BeeLogger

func init() {
	logDebugFile = logs.NewLogger(1000)
	logDebugFile.SetLogger("file", `
	{
		"filename":"`+beego.AppConfig.String("appname")+`_debug.log",
		"maxsize":`+fmt.Sprint(5<<20)+`
	}`) //最大~5M
	logDebugFile.EnableFuncCallDepth(true)
	logDebugFile.SetLogFuncCallDepth(3)

	logErrorFile = logs.NewLogger(1000)
	logErrorFile.SetLogger("file", `{
		"filename":"`+beego.AppConfig.String("appname")+`_error.log",
		"maxsize":`+fmt.Sprint(5<<20)+`
	}`)
	logErrorFile.EnableFuncCallDepth(true)
	logErrorFile.SetLogFuncCallDepth(3)

	logStd = logs.NewLogger(10)
	logStd.SetLogger("console", "")
	logStd.EnableFuncCallDepth(true)
	logStd.SetLogFuncCallDepth(3)
}

func Debug(f string, v ...interface{}) {
	logDebugFile.Debug(f, v...)
}

func Error(f string, v ...interface{}) {
	logErrorFile.Error(f, v...)
}
func ErrorStd(f string, v ...interface{}) {
	logStd.Error(f, v...)
}
func DebugStd(f string, v ...interface{}) {
	logStd.Debug(f, v...)
}
