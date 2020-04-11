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
	//mux.HandleFunc("/Index",safeHander(controller.BlogController{}.BlogIndex))
	////后端首页
	//mux.HandleFunc("/Admin/Index",safeHander(controller.BlogController{}.BlogIndex))
	//
	////博文列表
	mux.HandleFunc("/Blog/Index",safeHander((&controller.BlogController{}).Index))
	//博文添加
	mux.HandleFunc("/Blog/Add",safeHander((&controller.BlogController{}).Add))
	//博文修改
	//mux.HandleFunc("/Blog/Update",safeHander(controller.BlogController{}.BlogIndex))
	////博文删除
	//mux.HandleFunc("/Blog/Delete",safeHander(controller.BlogController{}.BlogIndex))
	//
	////主题列表
	//mux.HandleFunc("/Subject/Index",safeHander(controller.BlogController{}.BlogIndex))
	////主题添加
	//mux.HandleFunc("/Subject/Add",safeHander(controller.BlogController{}.BlogIndex))
	////主题修改
	//mux.HandleFunc("/Subject/Update",safeHander(controller.BlogController{}.BlogIndex))
	////主题修改
	//mux.HandleFunc("/Subject/Delete",safeHander(controller.BlogController{}.BlogIndex))

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	// 设置服务器监听请求端口
	server.ListenAndServe()
}

func safeHander(fn http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter,r *http.Request) () {
		if r.Method=="OPTION" {
			return
		}
		w.Header().Set("Content-type", "application/json;charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")             //返回数据格式是json
		fn(w,r)
	}
}

