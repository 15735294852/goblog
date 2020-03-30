package route

import (
	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"goblog"
	"net/http"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := main.session(writer, request)
	if err != nil {
		main.generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		main.generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		main.error_message(writer, request, "Cannot get threads")
	} else {
		_, err := main.session(writer, request)
		if err != nil {
			main.generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			main.generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
