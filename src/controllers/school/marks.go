package school

import (
	"databaseutils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetMarks(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	dateStr := params["date"]
	studIdStr := params["student_id"]

	studId, err := strconv.Atoi(studIdStr)
	if err != nil {
		log.Printf("Can not convert string to int %v\n", err)
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Printf("Can not convert strig to time %v\n", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(databaseutils.GetMarks(studId, date, date.AddDate(0,0,7)))
}
