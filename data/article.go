package data

import (
	"github.com/beego/orm"
	"log"
)

//文章基本信息
type Articles struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Title string `json:"title"`
}

//文章内容
type ArticleContent struct {
	Id int `json:"id"`
	ArticleId int64 `json:"article_id"`
	Content string `json:"content"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Articles), new(ArticleContent))
}

func (article *Articles)Insert()(int64,int64) {
	artcle := new(Articles)
	artcle.UserId = 0
	artcle.Title = "第一篇文章"

	articleId,err := DB.Insert(artcle)
	if err != nil {
		log.Fatal("insert article err")
	}

	artcleContent := new(ArticleContent)
	artcleContent.ArticleId = articleId
	artcleContent.Content = "第一篇文章的内容"
	contentId,err := DB.Insert(artcleContent)
	if err != nil {
		log.Fatal("insert artcleContent err")
	}

	return articleId,contentId
}
