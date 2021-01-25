package utils

import (
	"config"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserId uint
	Login string
	jwt.StandardClaims
}

var JwtAuth = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/auth/login", "/auth/register"}
		reqPath := r.URL.Path

		for _, val := range notAuth {
			if val == reqPath {
				next.ServeHTTP(w,r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHead := r.Header.Get("Authorization")

		if tokenHead == "" {
			response = Message(403, "Forbidden", "Missing token")
			w.WriteHeader(http.StatusForbidden)
			Respond(w, response)
			return
		}

		tk := &Token{}

		token, err := jwt.ParseWithClaims(tokenHead, tk, func(token *jwt.Token)(interface{}, error) {
			return []byte(config.GetSecret()), nil
		})

		if err != nil {
			response = Message(403, "Forbidden", "Can't decode token")
			w.WriteHeader(http.StatusForbidden)
			Respond(w, response)
			log.Println(err)
			return
		}

		if !token.Valid {
			response = Message(403, "Forbidden", "Token not valid")
			w.WriteHeader(http.StatusForbidden)
			Respond(w, response)
			return
		}

		log.Printf("User %d, %s\n", tk.UserId, tk.Login)
		//ctx := context.WithValue(r.Context(), "user", )
		next.ServeHTTP(w,r)
	})
}

func CreateJwtToken(tk Token) string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, _ := token.SignedString([]byte(config.GetSecret()))
	return tokenStr
}
