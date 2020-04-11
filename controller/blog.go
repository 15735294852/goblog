package controller

import (
	"fmt"
	"goblog/model/blog"
	"net/http"
	"strconv"
)

type BlogController struct {
	BaseController
}

func (b *BlogController) Add(w http.ResponseWriter, request *http.Request) {
	//参数
	cont := request.PostFormValue("cont")
	title := request.PostFormValue("title")
	//cont := request.PostForm["cont"]
	//title := request.PostForm["title"]
	fmt.Printf("%s",cont)
	fmt.Printf("%s",title)

	artcle := new(blog.Articles)

	artcle.UserId = 0
	//artcle.Title = "第一篇文章"
}

//博客列表 https://www.cnblogs.com/liuhe688/p/11063945.html
func (b *BlogController) Index(w http.ResponseWriter, request *http.Request) {
	page := request.PostFormValue("page")
	size := request.PostFormValue("size")

	//string转为int
	pageInt,_ := strconv.Atoi(page)
	sizeInt,_ := strconv.Atoi(size)

	article := new(blog.Articles)
	data := article.List(pageInt,sizeInt)

	b.Response.Data = data
	b.Success(w)
}