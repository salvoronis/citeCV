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

	router.HandleFunc("/auth/login", controllers.Login).Methods("POST")
	router.HandleFunc("/auth/register", controllers.Register).Methods("POST")

	router.Use(utils.JwtAuth)

	log.Printf("Listening to %s\n", config.GetRoot())

	err := http.ListenAndServe(config.GetRoot(), router)
	if err != nil {
		log.Println(err)
	}
}
