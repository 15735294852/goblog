package route

import (
	"fmt"
	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"goblog"
	"net/http"
)

// GET /threads/new
// Show the new thread form page
func newThread(writer http.ResponseWriter, request *http.Request) {
	_, err := main.session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		main.generateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /signup
// Create the user account
func createThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := main.session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			main.danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			main.danger(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			main.danger(err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func readThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		main.error_message(writer, request, "Cannot read thread")
	} else {
		_, err := main.session(writer, request)
		if err != nil {
			main.generateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			main.generateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// POST /thread/post
// Create the post
func postThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := main.session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			main.danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			main.danger(err, "Cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			main.error_message(writer, request, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			main.danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
