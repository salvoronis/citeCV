package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"utils"
	"config"
	"controllers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/auth/test", controllers.Test).Methods("GET")

	router.Use(utils.JwtAuth)

	log.Printf("Listening to %s\n", config.GetRoot())

	err := http.ListenAndServe(config.GetRoot(), router)
	if err != nil {
		log.Println(err)
	}
}
