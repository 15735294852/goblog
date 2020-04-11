package blog

import (
	"github.com/beego/orm"
	"goblog/model"
	"log"
)

//文章基本信息
type Articles struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Articles))
	model.InitDB()
}

//插入博客
func (article *Articles) Insert() int64 {
	articleId, err := model.DB.Insert(article)
	if err != nil {
		log.Fatal("insert article err")
	}

	return articleId
}


//博客列表
func (article *Articles) List(page int,size int) (articles []Articles){
	offset := 2 * ( page - 1 )
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id","user_id","title").
		From("articles").
		OrderBy("id").Desc().
		Limit(size).Offset(offset)

	// 导出SQL语句
	sql := qb.String()

	// 执行SQL语句
	model.DB.Raw(sql).QueryRows(&articles)

	return articles
}
