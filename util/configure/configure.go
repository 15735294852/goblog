/**
	加载配置模块 因为写不了,就先用第三方的包封装下
	选择这个包的原因是: 支持自由设置你的配置文件
 */
package configure

import (
	"github.com/jinzhu/configor"
	"os"
	"path/filepath"
)

//配置名
const CONFDIRNAME = "conf"
//默认配置文件
const DEFAULTCONFFILENAME = "baseconf.json"

/**
	读取配置文件
	#TODO 未引进日志模块,后续要加一个日志记录
 */
func ReadConf(config interface{}, fileName string) error {

	var err error
	execpath, err := os.Executable()//获取程序路径

	if err != nil {
		return err
	}

	var fileFullPath = filepath.Join(filepath.Dir(execpath), "./" + CONFDIRNAME + "/" + fileName)//格式化处理文件路径

	configor.Load(config,fileFullPath)
	if err != nil {
		return err
	}
	return err

}