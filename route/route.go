package route

import (
	"goblog/controller"
	"net/http"
)


func Route()  {

	//路由
	mux := http.NewServeMux()

	//前端页面
	files := http.FileServer(http.Dir("static")) // 创建一个能够为指定目录中的静态文件服务的处理器
	mux.Handle("/", http.StripPrefix("/", files))//将上面的处理器给多路复用器handle函数

	//前端首页
	mux.HandleFunc("/Index",controller.Index)
	//后端首页
	mux.HandleFunc("/Admin/Index",controller.Index)

	//博文列表
	mux.HandleFunc("/Blog/Index",controller.Index)
	//博文添加
	mux.HandleFunc("/Blog/Add",controller.BlogIndex)
	//博文修改
	mux.HandleFunc("/Blog/Update",controller.Index)
	//博文删除
	mux.HandleFunc("/Blog/Delete",controller.Index)

	//主题列表
	mux.HandleFunc("/Subject/Index",controller.Index)
	//主题添加
	mux.HandleFunc("/Subject/Add",controller.Index)
	//主题修改
	mux.HandleFunc("/Subject/Update",controller.Index)
	//主题修改
	mux.HandleFunc("/Subject/Delete",controller.Index)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	// 设置服务器监听请求端口
	server.ListenAndServe()
}

