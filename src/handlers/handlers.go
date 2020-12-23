package handlers

import (
	"net/http"

	"auth"
	"root"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func GetServerRouters(store *sessions.CookieStore) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", root.NewRoot(store))
	router.HandleFunc("/login", auth.NewLogin(store))
	router.HandleFunc("/register", auth.NewRegister(store))
	return router
}
