package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goblog/data"
	"goblog/util/configure"
	"goblog/util/logger"
)

// 数据库配置
type DbConfig struct {
	Goblog      *SingleConfig `json:"goblog"`
	Goblog_test *SingleConfig `json:"goblog_test"`
}

type SingleConfig struct {
	Ip       string `json:"ip"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
	Database string `json:"database"`
}


func main() {
	article := new(data.Articles)
	article.Insert()
	//+++++++++获取配置相关 start +++++++++++++++++++//
	conf := DbConfig{}
	configure.ReadConf(&conf, "databases.json")
	fmt.Printf("%s", conf.Goblog.Password)
	//+++++++++获取配置相关 end ++++++++++++++++++++//

	//+++++++++日志相关 start +++++++++++++++++++//

	logger.LogVar.Error().Msg("error 日志级别测试")
	logger.LogVar.Debug().Msg("debug 日志级别测试")
	logger.LogVar.Fatal().Msg("fatal 日志级别测试")

	//+++++++++日志相关 end ++++++++++++++++++++//

}
