package main

import (
  "fmt"
  "net/http"
  "database/sql"
  _ "github.com/lib/pq"
  "crypto/md5"
  "encoding/hex"

  "github.com/gorilla/sessions"
)

type pupil struct{
  username string
  mail string
  password string
  index string
}

var (
  key = []byte("super-secret-key")
  store = sessions.NewCookieStore(key)

  connStr = "host=localhost port=5432 user=postgres password=qwerty dbname=postgres sslmode=disable"
  db, err = sql.Open("postgres", connStr)
)

func secret(w http.ResponseWriter, r *http.Request) {
  session, _ := store.Get(r, "cookie-name")

  if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
      http.Error(w, "Forbidden", http.StatusForbidden)
      return
  }

  fmt.Fprintln(w, "The cake is a lie!")
}

func GetMd5(text string) string {
  h := md5.New()
  h.Write([]byte(text))
  return hex.EncodeToString(h.Sum(nil))
}

func main() {
  http.HandleFunc("/profile", profile)
  http.HandleFunc("/login", login)
  http.HandleFunc("/logout", logout)
  http.HandleFunc("/register", register)
  http.HandleFunc("/index", index)
  http.HandleFunc("/", root)
  http.Handle("/js/",http.StripPrefix("/js/", http.FileServer(http.Dir("./scripts"))))
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
  http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./images"))))

  http.ListenAndServe(":8080", nil)
}
