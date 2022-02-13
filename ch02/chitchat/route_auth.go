package main

import (
	"net/http"
)

func authenticate(write http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	user, _ := data.UserByEmail(request.PostFormValue("email"))
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(write, &cookie) // ブラウザにクッキーをセット
		http.Redirect(write, request, "/", 302)
	} else {
		http.Redirect(write, request, "/login", 302)
	}
}
