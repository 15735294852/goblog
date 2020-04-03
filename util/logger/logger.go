/**
	日志的封装, 依赖第三方包地址:https://github.com/rs/zerolog
	日志级别:
		panic (zerolog.PanicLevel, 5)
		fatal (zerolog.FatalLevel, 4)
		error (zerolog.ErrorLevel, 3)
		warn (zerolog.WarnLevel, 2)
		info (zerolog.InfoLevel, 1)
		debug (zerolog.DebugLevel, 0)
		trace (zerolog.TraceLevel, -1)
*/
package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/rs/zerolog"
)

const LOGDIRNAME = "log"
const LOGFILENAME = "service.log"

var LogVar zerolog.Logger

func init() {

	execpath, err := os.Executable() //获取脚本路径
	if err != nil {
		fmt.Println("logger Init ERROR, please check conf")
		os.Exit(0)
	}
	fmt.Printf("%s",execpath)
	//TODO 1 这边差个自动创建log目录的代码,需要手动创建log,否则报参数错误 2 生成不同目录的日志,暂时不考虑
	fileName 			:= filepath.Join(filepath.Dir(execpath), "./" + LOGDIRNAME + "/" + LOGFILENAME)
	filePointer, err 	:= os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	zerolog.CallerSkipFrameCount = 2 //日志会带什么文件的多少行报错
	LogVar = zerolog.New(filePointer).With().Caller().Timestamp().Logger()
}
