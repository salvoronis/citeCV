package auth

import (
	"fmt"
	"net/http"
	"databaseutils"
)

func Login(res http.ResponseWriter, req *http.Request) {
	_ = databaseutils.GetDB()
	fmt.Fprintf(res, "It is fucking fantastic, when you make a well archetected programm")
}
