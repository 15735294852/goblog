package main

import (
	"goblog/route"
)


func main() {
	//artcle := new(blog.Articles)
	//
	//artcle.UserId = 0
	//artcle.Title = "第一篇文章"
	//articleId, _ := model.DB.Insert(artcle)
	//
	//fmt.Println(articleId)

	//article := new(controller.Articles)
	//article.Insert()
	//+++++++++获取配置相关 start +++++++++++++++++++//

	//+++++++++获取配置相关 end ++++++++++++++++++++//

	//+++++++++日志相关 start +++++++++++++++++++//

	//logger.LogVar.Error().Msg("error 日志级别测试")
	//logger.LogVar.Debug().Msg("debug 日志级别测试")
	//logger.LogVar.Fatal().Msg("fatal 日志级别测试")

	//+++++++++日志相关 end ++++++++++++++++++++//
	//路由
	route.Route()

}

