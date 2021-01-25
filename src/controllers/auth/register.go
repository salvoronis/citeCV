package auth

import (
	"databaseutils"
	"encoding/json"
	"log"
	"models"
	"net/http"
	"utils"
)

func Register(w http.ResponseWriter, r *http.Request){
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Println(err)
		return
	}
	if !user.Validate() {
		resp := utils.Message(403, "Forbidden", "Invalid data")
		w.WriteHeader(http.StatusForbidden)
		utils.Respond(w, resp)
		return
	}

	databaseutils.SaveUser(*user)

	tmp, err := databaseutils.GetUserByLogin(user.Login)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		utils.Respond(w, utils.Message(403, "Forbidden", "Can't find ID"))
	}

	response := make(map[string]interface{})
	response["token"] = utils.CreateJwtToken(utils.Token{UserId: tmp.Id, Login: user.Login})
	response["status"] = utils.Message(200, "OK", "OK")

	utils.Respond(w, response)
}
