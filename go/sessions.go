package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("auth-key")
	store = sessions.NewCookieStore(key)
)

func secret(writer http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "auth-cookie")

	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(writer, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(writer, "You're successfully connected :)")
}

func login(writer http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "auth-cookie")
	// Check password match VS Database
	session.Values["authenticated"] = true
	session.Save(req, writer)
}

func logout(writer http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "auth-cookie")

	session.Values["authenticated"] = false
	session.Save(req, writer)
}
