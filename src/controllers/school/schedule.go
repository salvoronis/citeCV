package school

import (
	"databaseutils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetSchedule(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	classIdStr := params["class_id"]

	classId, err := strconv.Atoi(classIdStr)
	if err != nil {
		log.Printf("Can not convert string to int %v\n", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(databaseutils.GetSchedule(classId))
}
