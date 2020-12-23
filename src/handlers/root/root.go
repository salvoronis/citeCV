package root

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	store *sessions.CookieStore
)

func root (res http.ResponseWriter, req *http.Request) {
	session, err := store.Get(req, "user-data")
	if err != nil {
		log.Printf("Can't get cookies %v\n", err)
	}
	if session.Values["authenticated"] == true {
		//person := session.Values["user"].(*user)
	}
}

func NewRoot(storeS *sessions.CookieStore) func(http.ResponseWriter, *http.Request) {
	store = storeS
	return root
}
