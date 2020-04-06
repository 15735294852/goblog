package main

import (
	"fmt"
	"goblog/model/blog"
)

func main() {
	artcle := new(blog.Articles)

	artcle.UserId = 0
	artcle.Title = "第一篇文章"
	articleId := artcle.Insert()

	fmt.Println(articleId)

	//article := new(data.Articles)
	//article.Insert()
	//+++++++++获取配置相关 start +++++++++++++++++++//
	//conf := DbConfig{}
	//configure.ReadConf(&conf, "databases.json")
	//fmt.Printf("%s", conf.Goblog.Password)
	//+++++++++获取配置相关 end ++++++++++++++++++++//

	//+++++++++日志相关 start +++++++++++++++++++//

	//logger.LogVar.Error().Msg("error 日志级别测试")
	//logger.LogVar.Debug().Msg("debug 日志级别测试")
	//logger.LogVar.Fatal().Msg("fatal 日志级别测试")

	//+++++++++日志相关 end ++++++++++++++++++++//

}
