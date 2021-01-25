package school

import (
	"databaseutils"
	"encoding/json"
	"net/http"
)

func GetClasses(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(databaseutils.GetClasses())
}
