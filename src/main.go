package main

import (
	"config"
	"handlers"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

func main() {
	store := sessions.NewCookieStore([]byte("I'll change it later"))

	handleDirs()
	http.Handle("/", handlers.GetServerRouters(store))

	log.Printf("Listening to %s\n",config.GetRoot())
	http.ListenAndServe(config.GetRoot(), nil)
}

func handleDirs(){
	http.Handle(
		"/js/",
		http.StripPrefix(
			"/js/",
			http.FileServer(http.Dir("./scripts"))))
	http.Handle(
		"/css/",
		http.StripPrefix(
			"/css/",
			http.FileServer(http.Dir("./css"))))
	http.Handle(
		"/img/",
		http.StripPrefix(
			"/img/",
			http.FileServer(http.Dir("./images"))))
}
