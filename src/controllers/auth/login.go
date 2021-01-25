package auth

import (
	"databaseutils"
	"encoding/json"
	"log"
	"models"
	"net/http"
	"utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.LoginForm{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Println(err)
		return
	}

	response := make(map[string]interface{})

	if !user.Validate() {
		w.WriteHeader(http.StatusForbidden)
		utils.Respond(w,utils.Message(403, "Forbidden", "Invalid login or password"))
		return
	}

	pass, id := databaseutils.CheckUser(user.Login, user.Password)
	if !pass {
		w.WriteHeader(http.StatusForbidden)
		utils.Respond(w, utils.Message(403, "Forbidden", "Incorrect login or password"))
		return
	}

	response["token"] = utils.CreateJwtToken(utils.Token{UserId: id, Login: user.Login})
	response["status"] = utils.Message(200, "OK", "You're in")

	utils.Respond(w, response)
}
