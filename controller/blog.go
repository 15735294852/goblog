package controller

import (
	"fmt"
	"goblog/model/blog"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request)  {
	// 先写成这样,为了解决跨域
	Cors(writer)

	fmt.Println("2323")

}

func BlogIndex(w http.ResponseWriter, request *http.Request) {
	// 先写成这样,为了解决跨域
	Cors(w)
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