package utils

import (
	"encoding/json"
	"net/http"
	"crypto/sha256"
	"encoding/hex"
	"time"
	"math/rand"
)

func Message(status int, mesType string, message string) (map[string]interface{}) {
	return map[string]interface{}{"code": status, "type": mesType, "message": message}
}

func Respond(w http.ResponseWriter, data map[string] interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetSHA256(secret string) string {
	sha256Bytes := sha256.Sum256([]byte(secret))
	return hex.EncodeToString(sha256Bytes[:])
}

func RandomStr(length int) string {
	var password string
	for i := 0;i < length;i++ {
		rand := randomInt()
		Char := string('a' + byte(rand))
		password += Char
	}
	return password
}

func randomInt() int{
	rand.Seed(time.Now().UTC().UnixNano())
	bytes := rand.Intn(26)
	return bytes
}
