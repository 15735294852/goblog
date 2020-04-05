package data

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"github.com/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"goblog/util/configure"
	"log"
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

func init() {
	conf := DbConfig{}
	configure.ReadConf(&conf, "databases.json")

	host := conf.Goblog_test.Host
	port := conf.Goblog_test.Port
	user := conf.Goblog_test.User
	password := conf.Goblog_test.Password
	database := conf.Goblog_test.Database
	character := conf.Goblog_test.Character

	dataSource := user+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset="+character
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// set default database
	orm.RegisterDataBase("default", "mysql", dataSource, 30)
	DB = orm.NewOrm()
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
