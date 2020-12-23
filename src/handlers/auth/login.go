package auth

import (
	"databaseutils"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	storeL *sessions.CookieStore
)

func login(res http.ResponseWriter, req *http.Request) {
	_ = databaseutils.GetDB()
	fmt.Fprintf(res, "It is fucking fantastic, when you make a well archetected programm")
}

func NewLogin(storeS *sessions.CookieStore) func(http.ResponseWriter, *http.Request) {
	storeL = storeS
	return login
}
