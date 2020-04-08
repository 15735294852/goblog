package model

import (
	"fmt"
	"github.com/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"goblog/util/configure"
	"os"
)

// 数据库配置
type DbConfig struct {
	Goblog      *SingleConfig `json:"goblog"`
	Goblog_test *SingleConfig `json:"goblog_test"`
}

type SingleConfig struct {
	Host       string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Character string `json:"character"`
}

var DB orm.Ormer

func InitDB() {
	conf := DbConfig{}
	configure.ReadConf(&conf, "databases.json")

	//host := conf.Goblog_test.Host
	//port := conf.Goblog_test.Port
	//user := conf.Goblog_test.User
	//password := conf.Goblog_test.Password
	//database := conf.Goblog_test.Database
	//character := conf.Goblog_test.Character

	host2 := conf.Goblog.Host
	port2 := conf.Goblog.Port
	user2 := conf.Goblog.User
	password2 := conf.Goblog.Password
	database2 := conf.Goblog.Database
	character2 := conf.Goblog.Character

	//dataSource := user+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset="+character
	dataSource2 := user2+":"+password2+"@tcp("+host2+":"+port2+")/"+database2+"?charset="+character2
	fmt.Printf("%s",dataSource2)
	os.Exit(1)
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// set default database
	//orm.RegisterDataBase("default", "mysql", dataSource, 30)
	orm.RegisterDataBase("default", "mysql", dataSource2, 30)
	//注册第二个数据库
	//orm.RegisterDataBase("online", "mysql", dataSource2, 30)
	DB = orm.NewOrm()
}
