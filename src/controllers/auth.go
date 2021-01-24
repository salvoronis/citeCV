package controllers

import (
	"net/http"
	"utils"
)

func Test(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(200, "OK", "all right")
	utils.Respond(w, resp)
}
