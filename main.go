package main

import (
	"fmt"
	"goblog/util/configure"
)


// 数据库配置
type DbConfig struct {
	Goblog *SingleConfig `json:"goblog"`
	Goblog_test     *SingleConfig `json:"goblog_test"`
}
type SingleConfig struct {
	Ip       string `json:"ip"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

func init()  {
	fmt.Println("test")
}

func main() {

	//+++++++++获取配置相关 start +++++++++++++++++++//
	conf := DbConfig{}
	configure.ReadConf(&conf, "databases.json")
	fmt.Printf("%s",conf.Goblog.Password)
	//+++++++++获取配置相关 end ++++++++++++++++++++//

	//+++++++++日志相关 start +++++++++++++++++++//

	//+++++++++日志相关 end ++++++++++++++++++++//




}