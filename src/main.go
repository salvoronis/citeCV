package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"utils"
	"config"
	"controllers/auth"
	"controllers/school"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/auth/login", auth.Login).Methods("POST")
	router.HandleFunc("/auth/register", auth.Register).Methods("POST")
	router.HandleFunc("/school/classes", school.GetClasses).Methods("GET")

	router.Use(utils.JwtAuth)

	log.Printf("Listening to %s\n", config.GetRoot())

	err := http.ListenAndServe(config.GetRoot(), router)
	if err != nil {
		log.Println(err)
	}
}
