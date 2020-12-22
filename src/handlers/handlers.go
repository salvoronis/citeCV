package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"root"
)

func GetServerRouters() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", root.Root)
	return router
}
