package blog

import (
	"github.com/beego/orm"
)

//文章基本信息
type Articles struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Title string `json:"title"`
}


func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Articles))
}

//func (article *Articles)Insert(artcle)(int64) {
//
//	o := orm.NewOrm()
//	o.Using("online")
//	articleId,err := o.Insert(artcle)
//	if err != nil {
//		log.Fatal("insert article err")
//	}
//
//	return articleId
//}
