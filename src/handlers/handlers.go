package handlers

import (
	"net/http"

	"root"
	"auth"

	"github.com/gorilla/mux"
)

func GetServerRouters() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", root.Root)
	router.HandleFunc("/login", auth.Login)
	return router
}
